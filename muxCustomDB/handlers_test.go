package muxCustomDB

import (
	"fmt"
	"io"
	//"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var (
	server      *httptest.Server
	reader      io.Reader
	highFiveUrl string
)

func init() {
	router := NewRouter()
	server = httptest.NewServer(router)
	server.URL = "http://localhost:8787"
	defer server.Close()
}

// TestPostHighFive tests creating a new HighFive.
func TestPostHighFive(t *testing.T) {
	highFive := `{"sender": "Dennis", "receiver": "Kaylyn2", "message": "Thanks"}`
	highFiveUrl = fmt.Sprintf("%s/highFive", server.URL)
	reader = strings.NewReader(highFive)
	request, err := http.NewRequest("POST", highFiveUrl, reader)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	PostHighFive(w, request)
	if w.Code != 201 {
		t.Error("Success expected: %d", w.Code)
	}
}

// TestPostHighFiveError tests posting an incomplete HighFive.
func TestPostHighFiveError(t *testing.T) {
	//highFive := `{"sender": "Dennis", "receiver": "Kaylyn2", "message": "Thanks"}`
	highFiveUrl = fmt.Sprintf("%s/highFive", server.URL)
	reader = strings.NewReader("")
	request, err := http.NewRequest("POST", highFiveUrl, reader)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	PostHighFive(w, request)
	if w.Code != 422 {
		t.Error("Success expected: %d", w.Code)
	}
}

// TestGetHighFives tests the response of all high fives.
func TestGetHighFives(t *testing.T) {
	highFiveUrl = fmt.Sprintf("%s/highFive", server.URL)
	reader = strings.NewReader("")
	request, err := http.NewRequest("GET", highFiveUrl, reader)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	GetHighFives(w, request)
	if w.Code != 200 {
		t.Error("Success expected: %d", w.Code)
	}
}

// TestGetHighFive tests to find a specific high five.
// A new mux route handler is created to save the variables through the
// entire gorilla lifecycle.
func TestGetHighFive(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://localhost:8787/highFive/1", nil)
	w := httptest.NewRecorder()

	m := mux.NewRouter()
	m.HandleFunc("/highFive/{highFiveId}", GetHighFive)
	m.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Error("Success expected: ", w.Code)
	}
}

// TestDeleteHighFive tests to delete a specific high five.
// A new mux route handler is created to save the variables through the
// entire gorilla lifecycle.
func TestDeleteHighFive(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "http://localhost:8787/highFive/1", nil)
	w := httptest.NewRecorder()

	m := mux.NewRouter()
	m.HandleFunc("/highFive/{highFiveId}", DeleteHighFive)
	m.ServeHTTP(w, r)

	if w.Code != 204 {
		t.Error("Success expected: ", w.Code)
	}
}
