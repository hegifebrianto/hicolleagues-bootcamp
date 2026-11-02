package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// db adalah koneksi database yang dipakai di seluruh handler.
// Package database/sql sebenarnya mengelola SEBUAH POOL koneksi, jadi
// variabel global ini aman dipakai bersamaan oleh banyak request.
var db *sql.DB

func initDB() {
	// Sebaiknya DSN ini dibaca dari environment variable di aplikasi
	// nyata (lihat catatan "Tugas mandiri" di slide penutup), bukan
	// ditulis langsung di kode seperti di sini.
	dsn := "root:@tcp(localhost:3306)/hicolleagues?parseTime=true"

	var err error

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// sql.Open TIDAK langsung membuka koneksi -- cuma menyiapkan
	// strukturnya. Ping() yang benar-benar mencoba konek, supaya kita
	// tahu dari awal kalau DSN/server-nya salah.
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Terhubung ke MySQL")
}
