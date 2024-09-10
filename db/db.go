package db

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nilspolek/goLog"
)

type LinkStore struct {
	db *sql.DB
}

func NewLinkStore(dbFile string) *LinkStore {
	db := goLog.ExitOnError(sql.Open("sqlite3", dbFile))
	createTable := `
	CREATE TABLE IF NOT EXISTS links (
        sLink TEXT PRIMARY KEY,
        dLink TEXT
    );
	`
	goLog.LogOnError(db.Exec(createTable))
	return &LinkStore{
		db: db,
	}
}

func (store *LinkStore) AddLink(dest string) {
	rows := goLog.LogOnError(store.db.Query("SELECT sLink FROM links"))
	goLog.Debug("%v", rows)
}

func (store *LinkStore) Close() {
	err := store.db.Close()
	if err != nil {
		goLog.Error(err.Error())
	}
}

func NewLink(length int) string {
	bytes := make([]byte, length)
	goLog.LogOnError(rand.Read(bytes))
	return hex.EncodeToString(bytes)[:length]
}
