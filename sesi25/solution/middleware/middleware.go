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

// Auth: tanpa token valid -> 401. Token valid -> claims disimpan di context, lanjut ke next.
func (m *Middleware) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		token, ok := strings.CutPrefix(header, "Bearer ")
		if !ok || strings.TrimSpace(token) == "" {
			respon.Error(w, http.StatusUnauthorized, "token tidak ada")
			return
		}
		claims, err := m.Tokens.Cek(strings.TrimSpace(token))
		if err != nil {
			respon.Error(w, http.StatusUnauthorized, err.Error())
			return
		}
		ctx := context.WithValue(r.Context(), userKey, claims)
		next(w, r.WithContext(ctx))
	}
}

// RequireRole: harus dipasang DI DALAM Auth. Role tidak cocok -> 403.
func RequireRole(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := UserDariContext(r.Context())
		if c == nil || c.Role != role {
			respon.Error(w, http.StatusForbidden, "akses ditolak: butuh role "+role)
			return
		}
		next(w, r)
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
