package user_svc

import (
	"errors"
	"fmt"
	"forum/internal/model"
	"forum/internal/repository/user_repo"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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

type Userer interface {
	Create(user *model.UserDto) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type userService struct {
	repo user_repo.Userer
}

func NewUser(repo user_repo.Userer) Userer {
	return &userService{
		repo: repo,
	}
}

func generateHashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error("error while generation hash: %v", err)
	}

	return string(hash)
}

func (s *userService) Create(user *model.UserDto) (int, error) {
	user.Password = generateHashPassword(user.Password)
	id, err := s.repo.Create(model.UserDtoToEntity(user))
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *userService) GenerateToken(email string, password string) (string, error) {
	user, err := s.repo.GetBySignIn(email, generateHashPassword(password))
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

	userToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", fmt.Errorf("service: generate token: get string - %w", err)
	}

	return userToken, nil
}

func (s *userService) ParseToken(accessToken string) (int, error) {
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
