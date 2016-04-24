package mux

import (
	"github.com/gorilla/mux"

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
	router := mux.NewRouter()
	server = httptest.NewServer(router)
	server.URL = "http://localhost:8787"
	highFiveUrl = fmt.Sprintf("%s/highFive", server.URL)
}

func TestPostHighFive(t *testing.T) {
	highFive := `{"sender": "Dennis", "receiver": "Kaylyn", "message": "Thanks"}`

	reader = strings.NewReader(highFive)
	request, err := http.NewRequest("POST", highFiveUrl, reader)
	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Error("Success expected: %d", res.StatusCode)
	}
}
