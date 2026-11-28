// Package handler: lapisan HTTP. Hanya baca request, panggil service, tulis response.
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"project3/model"
	"project3/respon"
	"project3/service"
)

type BukuHandler struct{ Svc *service.BukuService }

func (h *BukuHandler) Semua(w http.ResponseWriter, r *http.Request) {
	data, err := h.Svc.Semua()
	if err != nil {
		respon.ErrorAPI(w, err)
		return
	}
	respon.Data(w, http.StatusOK, data)
}

func idDariPath(w http.ResponseWriter, r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		respon.Error(w, http.StatusBadRequest, "id harus angka")
		return 0, false
	}
	return id, true
}

func (h *BukuHandler) Ambil(w http.ResponseWriter, r *http.Request) {
	id, ok := idDariPath(w, r)
	if !ok {
		return
	}
	b, err := h.Svc.Ambil(id)
	if err != nil {
		respon.ErrorAPI(w, err)
		return
	}
	respon.Data(w, http.StatusOK, b)
}

func (h *BukuHandler) Tambah(w http.ResponseWriter, r *http.Request) {
	var b model.Buku
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		respon.Error(w, http.StatusBadRequest, "JSON tidak valid")
		return
	}
	baru, err := h.Svc.Tambah(b)
	if err != nil {
		respon.ErrorAPI(w, err)
		return
	}
	respon.Data(w, http.StatusCreated, baru)
}

func (h *BukuHandler) Hapus(w http.ResponseWriter, r *http.Request) {
	id, ok := idDariPath(w, r)
	if !ok {
		return
	}
	if err := h.Svc.Hapus(id); err != nil {
		respon.ErrorAPI(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
