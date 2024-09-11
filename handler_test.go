package main

import (
	"encoding/json"
	"linkShortener/db"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/nilspolek/goLog"
)

func setup() {
	li := &LinkHandler{
		ls:          *db.NewLinkStore(defaultDB),
		defaultDest: defaultDest,
	}
	defer li.ls.Close()
	setupHandlers(li)
	http.ListenAndServe(":12345", nil)
}

func TestMain(t *testing.T) {
	go setup()
	time.Sleep(time.Millisecond * 500)
}

func TestGET(t *testing.T) {
	// Test the GET method
	resp, err := http.Get("http://localhost:12345/amazon.com")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK, got %s", resp.Status)
	}
	if resp.Request.URL.Host != "www.google.com" {
		t.Fatalf("Expected www.google.com, got %s", resp.Request.URL.Host)
	}
}
func TestPOSTandDELETE(t *testing.T) {
	// Create a new short link
	var sb strings.Builder
	encoder := json.NewEncoder(&sb)
	if err := encoder.Encode(LongLinkReq{"https://amazon.com"}); err != nil {
		t.Error(err)
	}
	goLog.Debug(sb.String())
	resp, err := http.Post("http://localhost:12345/", "application/json", strings.NewReader(sb.String()))
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK, got %s", resp.Status)
	}

	// Decode the response
	decoder := json.NewDecoder(resp.Body)
	var sLink ShortLink
	if err := decoder.Decode(&sLink); err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	// Verify that the short link works
	resp, err = http.Get("http://localhost:12345/" + sLink.Short)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK, got %s", resp.Status)
	}
	if resp.Request.URL.Host != "www.amazon.com" {
		t.Fatalf("Expected www.amazon.com, got %s", resp.Request.URL.Host)
	}
	sb.Reset()
	encoder = json.NewEncoder(&sb)
	if err := encoder.Encode(Link{
		ShortLink: []string{sLink.Short},
	}); err != nil {
	}
	// Perform DELETE request
	req, err := http.NewRequest("DELETE", "http://localhost:12345/", strings.NewReader(sb.String()))
	if err != nil {
		t.Error(err)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK for DELETE, got %s", resp.Status)
	}

	// Verify that the short link is deleted
	resp, err = http.Get("http://localhost:12345/" + sLink.Short)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK after DELETE, got %s", resp.Status)
	}
	defaultURL, _ := url.Parse(defaultDest)
	if resp.Request.URL.Host != defaultURL.Host {
		t.Fatalf("Expected %s, got %s", defaultURL.Host, resp.Request.URL.Host)
	}
}
