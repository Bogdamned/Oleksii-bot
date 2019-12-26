package auth

import (
	"BotLeha/models"
	"context"
)

//UserRepository is an abstraction to work with DBs
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username, password string) (*models.User, error)
}
