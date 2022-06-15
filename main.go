package main

import (
	"log"
	"net/http"
	"time"
)

const DEBUG = true

func main() {

	db, err := connect()
	sessionTicker := time.NewTicker(SESSION_DURATION)
	expireSession(sessionTicker)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/employerLogin", login(db))
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", http.DefaultServeMux))
}
