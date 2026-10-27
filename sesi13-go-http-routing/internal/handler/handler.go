package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go-http-routing-sesi13/internal/tugas"
)

type API struct {
	repo *tugas.MemoryRepository
}

func New(repo *tugas.MemoryRepository) *API {
	return &API{repo: repo}
}

func (a *API) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /ping", a.ping)
	mux.HandleFunc("GET /waktu", a.waktu)
	mux.HandleFunc("GET /api/salam", a.salam)
	mux.HandleFunc("GET /api/tugas", a.daftarTugas)
	mux.HandleFunc("POST /api/tugas", a.buatTugas)
	mux.HandleFunc("GET /api/tugas/{id}", a.ambilTugas)
	mux.HandleFunc("DELETE /api/tugas/{id}", a.hapusTugas)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		// Response headers may already have been sent; nothing else can be done here.
		return
	}
}

func (a *API) ping(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"pesan": "pong"})
}

func (a *API) waktu(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"waktu": time.Now().Format(time.RFC3339)})
}

func (a *API) salam(w http.ResponseWriter, r *http.Request) {
	nama := strings.TrimSpace(r.URL.Query().Get("nama"))
	if nama == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "nama wajib diisi"})
		return
	}

	bahasa := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("bahasa")))
	salam := "" //default indo

	if bahasa == "en" {
		salam = "Hello, " + nama
	} else {
		salam = "halo, " + nama

	}
	writeJSON(w, http.StatusOK, map[string]string{"salam": salam})
}

func (a *API) daftarTugas(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, a.repo.List())
}

func (a *API) buatTugas(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var input struct {
		Judul string `json:"judul"`
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "JSON tidak valid"})
		return
	}
	if strings.TrimSpace(input.Judul) == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "judul wajib diisi"})
		return
	}

	item := a.repo.Create(strings.TrimSpace(input.Judul))
	writeJSON(w, http.StatusCreated, item)
}

func parseID(r *http.Request) (int, error) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		return 0, errors.New("id tidak valid")
	}
	return id, nil
}

func (a *API) ambilTugas(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id tidak valid"})
		return
	}
	item, err := a.repo.Get(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "tugas tidak ditemukan"})
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func (a *API) hapusTugas(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id tidak valid"})
		return
	}
	if err := a.repo.Delete(id); err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "tugas tidak ditemukan"})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
