package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const DEBUG = true

func main() {

	db, err := sql.Open("pgx", "")

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/employerLogin", employerLogin(db))
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", http.DefaultServeMux))
}

func employerLogin(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
