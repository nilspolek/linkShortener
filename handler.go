package main

import (
	"encoding/json"
	"fmt"
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

type Link struct {
	ShortLink []string `json:"short"`
	DestLink  string   `json:"destination"`
}

type LinkHandler struct {
	ls          db.LinkStore
	defaultDest string
}

func (h *LinkHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case http.MethodPost:
		// Create a new short link
		var reqestData LongLinkReq
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&reqestData); err != nil {
			goLog.Error(err.Error())
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
		encoder := json.NewEncoder(w)
		sLink := h.ls.AddLink(reqestData.Destination)
		if err := encoder.Encode(ShortLink{Short: sLink}); err != nil {
			goLog.Error(err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		goLog.Info("Added Linkdest %s with shortend Link %s to Database", reqestData.Destination, sLink)
		break
	case http.MethodGet:
		// Redirect to the destination URL
		sLink := r.URL.Path[1:]
		dLink := h.ls.GetDest(sLink)
		if dLink == "" {
			http.Redirect(w, r, h.defaultDest, http.StatusTemporaryRedirect)
			return
		}
		http.Redirect(w, r, dLink, http.StatusTemporaryRedirect)
		goLog.Info("Redirected %s from %s to %s", r.RemoteAddr, sLink, dLink)
		break
	case http.MethodDelete:
		// Delete a short link:
		var reqestData Link
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&reqestData); err != nil {
			goLog.Error(err.Error())
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
		if reqestData.DestLink != "" {
			err := h.ls.DeleteDest(reqestData.DestLink)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				goLog.Error(err.Error())
				return
			}
			goLog.Info("Deleted Destinationlink %s from Database", reqestData.ShortLink)
		}
		if len(reqestData.ShortLink) > 0 {
			for _, sLink := range reqestData.ShortLink {
				err := h.ls.DeleteLink(sLink)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					goLog.Error(err.Error())
					return
				}
				goLog.Info("Deleted Shortlink %s from Database", sLink)
			}
		}
		fmt.Fprintf(w, "Deleted Link/s %s from Database", reqestData)
		w.WriteHeader(http.StatusOK)
		break
	}
}
