package service

import (
	"forum/internal/repository"
)

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
