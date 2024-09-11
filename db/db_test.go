package db

import (
	"os"
	"testing"
)

func TestNewLink(t *testing.T) {
	if len(NewLink(10)) != 10 {
		t.Fatalf("Error wrong length of NewLink:\t%d", len(NewLink(10)))
	}
}

func TestNewLinkStore(t *testing.T) {
	store := NewLinkStore("link.db")
	defer store.Close()
	sLink := store.AddLink("Hallo Welt")
	got := store.GetDest(sLink)
	if got != "Hallo Welt" {
		t.Fatalf("Error wrong destination:\t%s", got)
	}
	os.Remove("link.db")
}

func TestEmptyLink(t *testing.T) {
	store := NewLinkStore("link.db")
	defer store.Close()
	sLink := store.AddLink("")
	got := store.GetDest(sLink)
	if got != "" {
		t.Fatalf("Error wrong destination:\t%s", got)
	}
	os.Remove("link.db")
}
