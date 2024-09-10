package db

import (
	"github.com/nilspolek/goLog"
	"testing"
)

func TestNewLink(t *testing.T) {
	goLog.Debug("TestNewLink:\t%s", NewLink(6))
	if len(NewLink(10)) != 10 {
		t.Fatalf("Error wrong length of NewLink:\t%d", len(NewLink(10)))
	}
}

func TestNewLinkStore(t *testing.T) {
	store := NewLinkStore("link.db")
	store.AddLink("HalloWelt")
}
