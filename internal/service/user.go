package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"forum/internal/model"
	"forum/internal/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%SFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
type User interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type UserService struct {
	repo repository.User
}

func NewUser(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *UserService) CreateUser(user model.User) (int, error) {
	user.Password = generateHashPassword(user.Password)
	id, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *UserService) GenerateToken(email string, password string) (string, error) {
	user, err := s.repo.GetUser(email, generateHashPassword(password))
	if err != nil {
		return "", fmt.Errorf("service: generate token: get user - %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	user.Token, err = token.SignedString([]byte(signingKey))
	if err != nil {
		return "", fmt.Errorf("service: generate token: get string - %w", err)
	}

	return user.Token, nil
}

func (s *UserService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, fmt.Errorf("service: parse token: parse claims - %w", err)
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, fmt.Errorf("service: parse token: claims - %w", err)
	}

	return claims.UserId, nil
}
