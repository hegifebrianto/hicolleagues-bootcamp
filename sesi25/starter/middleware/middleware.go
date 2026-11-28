// Package middleware: pembungkus handler (cek token, cek role, log, CORS).
package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"project3/auth"
	"project3/model"
	"project3/respon"
)

type ctxKey string

const userKey ctxKey = "user"

type Middleware struct{ Tokens *auth.Penerbit }

// TODO 2a: Auth. Langkah:
//  1. ambil header "Authorization", wajib berawalan "Bearer " (strings.CutPrefix)
//     gagal -> respon.Error(w, 401, "token tidak ada"); return
//  2. m.Tokens.Cek(token); error -> respon.Error(w, 401, err.Error()); return
//  3. simpan claims: ctx := context.WithValue(r.Context(), userKey, claims)
//  4. next(w, r.WithContext(ctx))
func (m *Middleware) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r) // sementara: semua orang lolos (SALAH). Perbaiki sesuai langkah di atas.
	}
}

// TODO 2b: RequireRole. Ambil claims dengan UserDariContext(r.Context()).
//
//	claims nil atau claims.Role != role -> respon.Error(w, 403, "akses ditolak: butuh role "+role); return
//	selain itu next(w, r)
func RequireRole(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r) // sementara: semua orang lolos (SALAH).
	}
}

// UserDariContext mengambil claims yang disimpan Auth (nil bila tidak ada).
func UserDariContext(ctx context.Context) *auth.Claims {
	c, _ := ctx.Value(userKey).(*auth.Claims)
	return c
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s (%v)", r.Method, r.URL.Path, time.Since(t).Round(time.Microsecond))
	})
}

func CORS(origin string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

var _ = model.RoleAdmin

var _ = []any{context.WithValue, strings.CutPrefix, respon.Error}
