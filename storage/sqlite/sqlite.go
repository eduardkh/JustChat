package sqlite

import (
	"context"
	"database/sql"
)

import _ "github.com/mattn/go-sqlite3"

import "github.com/justchatapp/justchat/models"

// Store implements storage.UserStore using SQLite.
type Store struct {
	DB *sql.DB
}

// New creates a new SQLite store.
func New(db *sql.DB) *Store {
	return &Store{DB: db}
}

// CreateUser inserts a new user.
func (s *Store) CreateUser(ctx context.Context, u *models.User) error {
	_, err := s.DB.ExecContext(ctx, `INSERT INTO users (username, password_hash) VALUES (?, ?)`, u.Username, u.PasswordHash)
	return err
}
