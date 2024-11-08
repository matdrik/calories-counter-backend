package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	calories_counter_backend "server"
	"server/internal/repository"
	"time"
)

const (
	salt       = "some random salt"
	signingKey = ",senfjqnwefo22211"
	tokenTTL   = time.Hour * 12
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.ID,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	return claims.UserId, nil
}

func (s *AuthService) CreateUser(user calories_counter_backend.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// TODO - заменить на bcrypt?
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
