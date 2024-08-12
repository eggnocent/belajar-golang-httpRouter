package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "hello world")
	})

	// Buat permintaan baru
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	// Panggil ServeHTTP dengan request yang benar
	router.ServeHTTP(recorder, request)

	// Baca body dari recorder
	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	// Pastikan respons sesuai dengan yang diharapkan
	assert.Equal(t, "hello world", string(body))
}
