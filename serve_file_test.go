package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	// Buat permintaan baru
	request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	// Panggil ServeHTTP dengan request yang benar
	router.ServeHTTP(recorder, request)

	// Baca body dari recorder
	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	// Pastikan respons sesuai dengan yang diharapkan
	assert.Equal(t, "hello httprouter", string(body))
}

func TestServeFileGoodbye(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	// Buat permintaan baru
	request := httptest.NewRequest("GET", "http://localhost:3000/files/goodbye.txt", nil)
	recorder := httptest.NewRecorder()

	// Panggil ServeHTTP dengan request yang benar
	router.ServeHTTP(recorder, request)

	// Baca body dari recorder
	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	// Pastikan respons sesuai dengan yang diharapkan
	assert.Equal(t, "goodbye httprouter", string(body))
}
