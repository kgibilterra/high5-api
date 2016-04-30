package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kgibilterra/high5-api/muxCustomDB"

	"log"
	"net/http"
)

func main() {
	// db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/high5")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	r := muxCustomDB.NewRouter()
	log.Fatal(http.ListenAndServe(":8787", r))
}
