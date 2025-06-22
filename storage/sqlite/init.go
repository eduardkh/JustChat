package sqlite

import (
	"database/sql"
)

// Init runs schema migrations.
func Init(db *sql.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT UNIQUE, password_hash TEXT)`
	_, err := db.Exec(schema)
	return err
}
