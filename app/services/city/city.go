package city

import (
	"github.com/litvivan/ilyway/app/models"
	cityRepo "github.com/litvivan/ilyway/app/repo/city"
)

type Service struct {
	repo *cityRepo.Repo
}

func NewService(repo *cityRepo.Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List() []models.City {
	return s.repo.List()
}
