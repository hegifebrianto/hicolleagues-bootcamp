package handler

import (
	"encoding/json"
	"net/http"

	"project3/middleware"
	"project3/respon"
	"project3/service"
)

type AuthHandler struct{ Svc *service.AuthService }

type daftarReq struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Daftar(w http.ResponseWriter, r *http.Request) {
	var req daftarReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respon.Error(w, http.StatusBadRequest, "JSON tidak valid")
		return
	}
	u, err := h.Svc.Daftar(req.Nama, req.Email, req.Password)
	if err != nil {
		respon.ErrorAPI(w, err)
		return
	}
	respon.Data(w, http.StatusCreated, u) // PasswordHash tidak ikut (json:"-")
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respon.Error(w, http.StatusBadRequest, "JSON tidak valid")
		return
	}
	token, err := h.Svc.Login(req.Email, req.Password)
	if err != nil {
		respon.ErrorAPI(w, err)
		return
	}
	respon.Data(w, http.StatusOK, map[string]string{"token": token})
}

// Saya: contoh endpoint yang membaca identitas dari context (diisi middleware Auth).
func (h *AuthHandler) Saya(w http.ResponseWriter, r *http.Request) {
	c := middleware.UserDariContext(r.Context())
	respon.Data(w, http.StatusOK, map[string]string{"id": c.Sub, "role": c.Role})
}
