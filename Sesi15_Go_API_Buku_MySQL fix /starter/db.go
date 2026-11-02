package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// db adalah koneksi database yang dipakai di seluruh handler.
var db *sql.DB

func initDB() {
	// TODO 1: buat variabel dsn (string) dengan format:
	//   "user:password@tcp(localhost:3306)/tugasdb?parseTime=true"
	// Sesuaikan user/password dengan MySQL di komputermu.

	// TODO 2: panggil sql.Open("mysql", dsn) -- simpan hasilnya ke
	// variabel global db (pakai db, err = sql.Open(...), BUKAN :=,
	// supaya db yang dipakai adalah variabel package-level di atas).
	// Kalau err != nil, panggil log.Fatal(err).

	// TODO 3: panggil db.Ping() untuk benar-benar menguji koneksi.
	// Kalau error, log.Fatal(err) juga.

	log.Println("Terhubung ke MySQL")
}
