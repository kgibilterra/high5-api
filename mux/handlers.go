// Package high5 provides rest APIs for interacting with High 5's.
package mux

import (
	"github.com/gorilla/mux"

	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetHigh5 returns all of the High 5's in the database.
func GetHighFives(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(highFives); err != nil {
		panic(err)
	}
}

func GetHighFive(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	highFiveId := vars["highFiveId"]
	i, _ := strconv.Atoi(highFiveId)
	h := RepoFindHighFive(i)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(h); err != nil {
		panic(err)
	}
}

func PostHighFive(w http.ResponseWriter, r *http.Request) {
	var highFive HighFive
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &highFive); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	h := RepoCreateHighFive(highFive)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(h); err != nil {
		panic(err)
	}
}
