package main

import (
	"linkShortener/db"
	"net/http"
)

const (
	defaultAddr = ":8080"
	defaultDest = "https://www.google.com/"
	defaultDB   = "link.db"
)

func main() {
	li := &LinkHandler{
		ls:          *db.NewLinkStore(defaultDB),
		defaultDest: defaultDest,
	}
	defer li.ls.Close()
	setupHandlers(li)
	http.ListenAndServe(defaultAddr, nil)
}

func setupHandlers(li *LinkHandler) {
	http.Handle("/", li)
}
