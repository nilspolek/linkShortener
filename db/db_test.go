package db

import (
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
}
