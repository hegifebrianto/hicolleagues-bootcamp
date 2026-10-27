package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-http-routing-sesi13/internal/tugas"
)

func newTestServer() http.Handler {
	mux := http.NewServeMux()
	New(tugas.NewMemoryRepository()).RegisterRoutes(mux)
	return mux
}

func TestPing(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	newTestServer().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if !strings.Contains(rec.Body.String(), `"pesan":"pong"`) {
		t.Fatalf("body = %s, want JSON pong", rec.Body.String())
	}
}

func TestCreateAndDeleteTask(t *testing.T) {
	server := newTestServer()

	createReq := httptest.NewRequest(http.MethodPost, "/api/tugas", bytes.NewBufferString(`{"judul":"Belajar Go"}`))
	createRec := httptest.NewRecorder()
	server.ServeHTTP(createRec, createReq)
	if createRec.Code != http.StatusCreated {
		t.Fatalf("create status = %d, want %d", createRec.Code, http.StatusCreated)
	}

	deleteReq := httptest.NewRequest(http.MethodDelete, "/api/tugas/1", nil)
	deleteRec := httptest.NewRecorder()
	server.ServeHTTP(deleteRec, deleteReq)
	if deleteRec.Code != http.StatusNoContent {
		t.Fatalf("delete status = %d, want %d", deleteRec.Code, http.StatusNoContent)
	}

	missingReq := httptest.NewRequest(http.MethodDelete, "/api/tugas/1", nil)
	missingRec := httptest.NewRecorder()
	server.ServeHTTP(missingRec, missingReq)
	if missingRec.Code != http.StatusNotFound {
		t.Fatalf("missing status = %d, want %d", missingRec.Code, http.StatusNotFound)
	}
}
