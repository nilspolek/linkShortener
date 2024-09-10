package main

import (
	"linkShortener/db"
	"net/http"
)

const (
	defaultAddr = ":8080"
	defaultDest = "httpa://google.com/"
)

func main() {
	store := db.NewLinkStore("link.db")
	defer store.Close()
	setupHandlers()
	http.ListenAndServe(defaultAddr, nil)
}

func setupHandlers() {
	http.HandleFunc("/", mainHandler)
}
