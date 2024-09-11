package main

import (
	"encoding/json"
	"linkShortener/db"
	"net/http"

	"github.com/nilspolek/goLog"
)

type LongLinkReq struct {
	Destination string `json:"destination"`
}

type ShortLink struct {
	Short string `json:"short"`
}

type LinkHandler struct {
	ls db.LinkStore
}

func (h *LinkHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case http.MethodGet:

	case http.MethodPost:
		// Create a new short link

		// Extract the destination URL from the request body
		var reqestData LongLinkReq
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&reqestData); err != nil {
			goLog.Error(err.Error())
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
		// Return the ShortLink object as JSON
		encoder := json.NewEncoder(w)
		sLink := h.ls.AddLink(reqestData.Destination)
		if err := encoder.Encode(ShortLink{Short: sLink}); err != nil {
			goLog.Error(err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		goLog.Info("Added Linkdest %s with shortend Link %s to Database", reqestData.Destination, sLink)
	}
}
