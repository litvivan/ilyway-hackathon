package city

import (
	"github.com/litvivan/ilyway/app/models"
)

type Repo struct{}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) List() []models.City {
	cities := models.Cities

	list := make([]models.City, 0, len(cities))
	for i := range cities {
		list = append(list, cities[i])
	}

	return list
}
