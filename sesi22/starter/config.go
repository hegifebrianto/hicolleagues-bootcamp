package main

import (
	"fmt"
	"os"
)

// Config: semua pengaturan terpusat di satu struct. Kredensial TIDAK ditulis di kode;
// dibaca dari file .env (atau environment variable di terminal).
type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	Port       string
	Store      string // "mysql" (bawaan) atau "memory" (tanpa database, untuk demo/tes)
}

func getEnvDefault(kunci, bawaan string) string {
	if v := os.Getenv(kunci); v != "" {
		return v
	}
	return bawaan
}

func muatKonfigurasi() Config {
	muatDotEnv(".env") // dibaca dari folder tempat `go run .` dijalankan
	return Config{
		DBUser:     getEnvDefault("DB_USER", "root"),
		DBPassword: os.Getenv("DB_PASSWORD"), // boleh kosong
		DBHost:     getEnvDefault("DB_HOST", "localhost"),
		DBPort:     getEnvDefault("DB_PORT", "3306"),
		DBName:     getEnvDefault("DB_NAME", "bukudb"),
		Port:       getEnvDefault("PORT", "8089"),
		Store:      getEnvDefault("STORE", "mysql"),
	}
}

// DSN merakit connection string MySQL dari bagian-bagiannya.
func (c Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}
