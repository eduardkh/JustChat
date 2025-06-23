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

// GetUserByUsername retrieves a user by username.
func (s *Store) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	row := s.DB.QueryRowContext(ctx, `SELECT id, username, password_hash FROM users WHERE username = ?`, username)
	user := &models.User{}
	if err := row.Scan(&user.ID, &user.Username, &user.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
