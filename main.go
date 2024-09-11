package main

import (
	"linkShortener/db"
	"net/http"
)

const (
	defaultAddr = ":8080"
	defaultDest = "https://google.com/"
)

func main() {
	li := &LinkHandler{
		ls: *db.NewLinkStore("link.db"),
	}
	defer li.ls.Close()
	setupHandlers(li)
	http.ListenAndServe(defaultAddr, nil)
}

func setupHandlers(li *LinkHandler) {
	http.Handle("/", li)
}
