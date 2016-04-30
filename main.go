package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kgibilterra/high5-api/muxCustomDB"

	"log"
	"net/http"
)

func main() {
	r := muxCustomDB.NewRouter()
	log.Fatal(http.ListenAndServe(":8787", r))
}
