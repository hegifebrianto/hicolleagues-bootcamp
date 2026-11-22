package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

// BukuHandler hanya mengurus HTTP: baca request, panggil service, tulis response.
type BukuHandler struct {
	Service BukuService
}

func NewBukuHandler(s BukuService) BukuHandler { return BukuHandler{Service: s} }

func tulisJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// tulisErrorAPI memetakan error dari lapisan bawah menjadi status HTTP.
// Satu tempat saja -> semua handler konsisten.
func tulisErrorAPI(w http.ResponseWriter, err error) {
	var ev ErrValidasi
	switch {
	case errors.As(err, &ev):
		tulisJSON(w, http.StatusBadRequest, Respons{Error: ev.Error()})
	case errors.Is(err, ErrTidakDitemukan):
		tulisJSON(w, http.StatusNotFound, Respons{Error: err.Error()})
	default:
		log.Printf("error internal: %v", err) // detail hanya di log server
		tulisJSON(w, http.StatusInternalServerError, Respons{Error: "terjadi kesalahan pada server"})
	}
}

func idDariURL(w http.ResponseWriter, r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		tulisJSON(w, http.StatusBadRequest, Respons{Error: "id tidak valid"})
		return 0, false
	}
	return id, true
}

func (h BukuHandler) Semua(w http.ResponseWriter, r *http.Request) {
	hasil, err := h.Service.Semua()
	if err != nil {
		tulisErrorAPI(w, err)
		return
	}
	tulisJSON(w, http.StatusOK, Respons{Data: hasil})
}

func (h BukuHandler) Ambil(w http.ResponseWriter, r *http.Request) {
	id, ok := idDariURL(w, r)
	if !ok {
		return
	}
	b, err := h.Service.Ambil(id)
	if err != nil {
		tulisErrorAPI(w, err)
		return
	}
	tulisJSON(w, http.StatusOK, Respons{Data: b})
}

func (h BukuHandler) Tambah(w http.ResponseWriter, r *http.Request) {
	var b Buku
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		tulisJSON(w, http.StatusBadRequest, Respons{Error: "body tidak valid"})
		return
	}
	hasil, err := h.Service.Tambah(b)
	if err != nil {
		tulisErrorAPI(w, err)
		return
	}
	tulisJSON(w, http.StatusCreated, Respons{Data: hasil})
}

func (h BukuHandler) Hapus(w http.ResponseWriter, r *http.Request) {
	id, ok := idDariURL(w, r)
	if !ok {
		return
	}
	if err := h.Service.Hapus(id); err != nil {
		tulisErrorAPI(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
