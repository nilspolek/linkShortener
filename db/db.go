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

func (store *LinkStore) AddLink(dest string) string {
	slink := NewLink(6)
	insertLink := `
	INSERT INTO links (sLink, dLink) VALUES (?, ?);
	`
	goLog.LogOnError(store.db.Exec(insertLink, slink, dest))
	return slink
}

func (store *LinkStore) GetLink(dlink string) string {
	var slink string
	getLink := `
	SELECT sLink FROM links WHERE dLink = ?;
	`
	store.db.QueryRow(getLink, dlink).Scan(&slink)
	return slink
}
func (store *LinkStore) GetDest(slink string) string {
	var dlink string
	getLink := `
	SELECT dLink FROM links WHERE sLink = ?;
	`
	store.db.QueryRow(getLink, slink).Scan(&dlink)
	return dlink
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
