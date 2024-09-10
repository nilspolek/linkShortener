package main

import "net/http"

func mainHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case "GET":
		w.Write([]byte("Hallo Welt"))
	case "POST":
		w.Write([]byte("Hallo Welt but Post "))
	}
}
