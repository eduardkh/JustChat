package storage

import "context"

import "github.com/justchatapp/justchat/models"

// UserStore defines methods to persist users.
type UserStore interface {
	CreateUser(ctx context.Context, u *models.User) error
}
