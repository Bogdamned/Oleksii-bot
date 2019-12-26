package usecase

import (
	"BotLeha/auth"
	"BotLeha/models"
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//AuthClaims auth method
type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

//AuthUseCase instance
type AuthUseCase struct {
	userRepo       auth.UserRepository
	hashSalt       string
	signinKey      []byte
	expireDuration time.Duration
}

//NewAuthUseCase constructor
func NewAuthUseCase(
	userRepo auth.UserRepository,
	hashSalt string,
	signinKey []byte,
	tokenTTLSeconds time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signinKey:      signinKey,
		expireDuration: time.Second * tokenTTLSeconds,
	}
}

// SignUp: user registration
func (a *AuthUseCase) SignUp(ctx context.Context, username, password string) error {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &models.User{
		Username: username,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	return a.userRepo.CreateUser(ctx, user)
}

// SignIn of a user
func (a *AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.userRepo.GetUser(ctx, username, password)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(a.expireDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signinKey)
}
