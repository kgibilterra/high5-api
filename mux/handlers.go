// Package high5 provides rest APIs for interacting with High 5's.
package mux

import (
	"github.com/gorilla/mux"

	"encoding/json"
	"fmt"
	"net/http"
)

// GetHigh5 returns all of the High 5's in the database.
func GetHighFives(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(highFives); err != nil {
		panic(err)
	}
}

func GetHighFive(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	highFiveId := vars["highFiveId"]
	fmt.Fprintln(w, "High Five: ", highFiveId)
}
