package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"project3/auth"
	"project3/config"
	"project3/handler"
	"project3/middleware"
	"project3/model"
	"project3/repository"
	"project3/service"
)

func main() {
	cfg, err := config.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	var bukuRepo repository.BukuRepo
	var userRepo repository.UserRepo

	if cfg.Store == "memory" {
		log.Println("STORE=memory (data hilang saat server berhenti)")
		bukuRepo = repository.NewBukuMemori(
			model.Buku{Judul: "Laskar Pelangi", Penulis: "Andrea Hirata", Tahun: 2005},
			model.Buku{Judul: "Bumi Manusia", Penulis: "Pramoedya Ananta Toer", Tahun: 1980},
		)
		userRepo = seedUser()
	} else {
		db, err := sql.Open("mysql", cfg.DSN())
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close() //closes connection to the Database
		if err := db.Ping(); err != nil {
			log.Fatalf("tidak bisa terhubung ke MySQL: %v (cek .env, atau jalankan dengan STORE=memory)", err)
		}
		bukuRepo = repository.NewBukuSQL(db)
		userRepo = repository.NewUserSQL(db)
	}

	tokens := auth.NewPenerbit(cfg.JWTSecret, cfg.TokenTTL)
	mw := &middleware.Middleware{Tokens: tokens}
	authSvc := service.NewAuthService(userRepo, tokens, service.LogMailer{Delay: 2 * time.Second})
	bukuSvc := service.NewBukuService(bukuRepo)

	router := handler.NewRouter(
		&handler.BukuHandler{Svc: bukuSvc},
		&handler.AuthHandler{Svc: authSvc},
		mw,
	)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           middleware.Logging(middleware.CORS(cfg.CORSOrigin, router)),
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("server jalan di http://localhost:%s", cfg.Port)
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown: Ctrl+C -> selesaikan request yang sedang berjalan
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("server berhenti")
}

// seedUser menyiapkan 2 akun demo untuk mode memory.
func seedUser() *repository.UserMemori {
	r := repository.NewUserMemori()
	for _, u := range []struct{ nama, email, role string }{
		{"Admin", "admin@mail.com", model.RoleAdmin}, //admin
		{"Budi", "budi@mail.com", model.RoleUser},    //user
	} {
		h, err := auth.HashPassword("rahasia")
		if err != nil {
			log.Fatal(err)
		}
		r.Simpan(model.User{Nama: u.nama, Email: u.email, PasswordHash: h, Role: u.role})
	}
	return r
}
