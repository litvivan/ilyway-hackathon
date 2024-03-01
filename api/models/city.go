package models

import (
	domainModels "github.com/litvivan/ilyway/app/models"
)

type City struct {
	Name string `json:"name"`
}

func MapCities(cities []domainModels.City) []City {
	mappedItems := make([]City, len(cities))
	for i, city := range cities {
		mappedItems[i] = City{
			Name: city.Name,
		}
	}

	return mappedItems
}
