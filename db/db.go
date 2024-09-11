package db

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"

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
	// Add a new link to the database and return the short link
	slink := NewLink(6)
	insertLink := `
	INSERT INTO links (sLink, dLink) VALUES (?, ?);
	`
	goLog.LogOnError(store.db.Exec(insertLink, slink, dest))
	return slink
}

func (store *LinkStore) GetLink(dlink string) string {
	// Return the short link for a given destination link
	var slink string
	getLink := `
	SELECT sLink FROM links WHERE dLink = ?;
	`
	store.db.QueryRow(getLink, dlink).Scan(&slink)
	return slink
}
func (store *LinkStore) GetDest(slink string) string {
	// Return the destination link for a given short link
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

func (store *LinkStore) DeleteLink(slink string) error {
	// Delete a link from the database
	if store.GetDest(slink) == "" {
		return errors.New("Short link does not exist")
	}
	deleteLink := `
	DELETE FROM links WHERE sLink = ?;
	`
	goLog.LogOnError(store.db.Exec(deleteLink, slink))
	return nil
}

func (store *LinkStore) DeleteDest(dlink string) error {
	// Delete a link from the database
	//
	// Check if the destination link exists in the database
	if store.GetLink(dlink) == "" {
		return errors.New("Destination link does not exist")
	}
	deleteLink := `
	DELETE FROM links WHERE dLink = ?;
	`
	goLog.LogOnError(store.db.Exec(deleteLink, dlink))
	return nil
}

func NewLink(length int) string {
	// Generate a random string of length length
	bytes := make([]byte, length)
	goLog.LogOnError(rand.Read(bytes))
	return hex.EncodeToString(bytes)[:length]
}
