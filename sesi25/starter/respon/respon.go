// Package respon: satu tempat untuk menulis balasan JSON + memetakan error ke status HTTP.
package respon

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"project3/model"
)

type Respons struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func Data(w http.ResponseWriter, status int, data any) { JSON(w, status, Respons{Data: data}) }

func Error(w http.ResponseWriter, status int, pesan string) {
	JSON(w, status, Respons{Error: pesan})
}

// ErrorAPI memetakan error dari lapisan bawah menjadi status HTTP yang tepat.
func ErrorAPI(w http.ResponseWriter, err error) {
	var ev model.ErrValidasi
	switch {
	case errors.As(err, &ev):
		Error(w, http.StatusBadRequest, ev.Error())
	case errors.Is(err, model.ErrTidakDitemukan):
		Error(w, http.StatusNotFound, err.Error())
	case errors.Is(err, model.ErrEmailTerpakai):
		Error(w, http.StatusConflict, err.Error())
	case errors.Is(err, model.ErrKredensial):
		Error(w, http.StatusUnauthorized, err.Error())
	default:
		log.Printf("error internal: %v", err) // detail hanya di log server
		Error(w, http.StatusInternalServerError, "terjadi kesalahan pada server")
	}
}
