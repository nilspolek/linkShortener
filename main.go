package main

import "linkShortener/db"

func main() {
	store := db.NewLinkStore("link.db")
	store.AddLink("HalloWelt")
}
