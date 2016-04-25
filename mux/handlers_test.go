package mux

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

func TestGetHighFives(t *testing.T) {
	highFiveUrl = fmt.Sprintf("%s/highFives", server.URL)
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
