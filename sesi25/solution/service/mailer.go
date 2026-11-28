package service

import (
	"log"
	"time"
)

// Mailer = pengirim email. Berupa interface agar mudah diganti/di-mock di test.
type Mailer interface {
	KirimSelamatDatang(email string)
}

// LogMailer mensimulasikan kirim email yang LAMBAT (Delay) lalu mencatat ke log.
type LogMailer struct{ Delay time.Duration }

func (m LogMailer) KirimSelamatDatang(email string) {
	time.Sleep(m.Delay) // pura-pura menghubungi server SMTP
	log.Printf("[email] selamat datang terkirim ke %s", email)
}
