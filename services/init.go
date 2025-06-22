package services

import (
	"database/sql"
	"log"

	"github.com/justchatapp/justchat/storage/sqlite"
)

// InitServices initializes database and services.
func InitServices() {
	db, err := sql.Open("sqlite3", "justchat.db")
	if err != nil {
		log.Fatal(err)
	}
	if err := sqlite.Init(db); err != nil {
		log.Fatal(err)
	}
	store := sqlite.New(db)
	UserService = NewUserService(store)
}
