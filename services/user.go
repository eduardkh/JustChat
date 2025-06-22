package services

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/justchatapp/justchat/models"
	"github.com/justchatapp/justchat/storage"
)

// UserServiceType defines user related operations.
type UserServiceType struct {
	store storage.UserStore
}

// UserService is the default service instance.
var UserService *UserServiceType

// NewUserService creates a new UserServiceType.
func NewUserService(store storage.UserStore) *UserServiceType {
	return &UserServiceType{store: store}
}

// Register registers a new user with hashed password.
func (s *UserServiceType) Register(ctx context.Context, username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &models.User{Username: username, PasswordHash: string(hash)}
	return s.store.CreateUser(ctx, user)
}
