// Package config: semua pengaturan terpusat; kredensial dibaca dari .env / environment.
package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port       string
	Store      string // "mysql" (bawaan) atau "memory"
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	JWTSecret  string
	TokenTTL   time.Duration
	CORSOrigin string
}

func env(kunci, bawaan string) string {
	if v := os.Getenv(kunci); v != "" {
		return v
	}
	return bawaan
}

// Load membaca .env lalu membangun Config. JWT_SECRET WAJIB ada (tidak ada nilai bawaan).
func Load(pathEnv string) (Config, error) {
	MuatDotEnv(pathEnv)

	menit, err := strconv.Atoi(env("TOKEN_TTL_MENIT", "15"))
	if err != nil || menit <= 0 {
		return Config{}, errors.New("TOKEN_TTL_MENIT harus angka positif")
	}
	c := Config{
		Port:       env("PORT", "8089"),
		Store:      env("STORE", "mysql"),
		DBUser:     env("DB_USER", "root"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     env("DB_HOST", "localhost"),
		DBPort:     env("DB_PORT", "3306"),
		DBName:     env("DB_NAME", "bukudb"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		TokenTTL:   time.Duration(menit) * time.Minute,
		CORSOrigin: env("CORS_ORIGIN", "http://localhost:5173"),
	}
	if len(c.JWTSecret) < 16 {
		return Config{}, errors.New("JWT_SECRET wajib diisi di .env (minimal 16 karakter)")
	}
	return c, nil
}

func (c Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}
