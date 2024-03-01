package models

import (
	"time"

	domainModels "github.com/litvivan/ilyway/app/models"
)

type Item struct {
	ID               int64   `json:"id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	ParticipantCount int     `json:"participant_count"`
	ActivityType     string  `json:"activity_type"`
	City             string  `json:"city"`
	AuthorName       string  `json:"author_name"`
	AuthorRating     float64 `json:"author_rating"`
	ImageUrl         string  `json:"image_url"`
	FullAddress      string  `json:"full_address"`
	HasReward        bool    `json:"has_reward"`
	Duration         string  `json:"duration"`
	StartAt          string  `json:"start_at"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

func MapItems(items []*domainModels.Item) []Item {
	mappedItems := make([]Item, len(items))
	for i := range items {
		mappedItems[i] = MapItem(items[i])
	}

	return mappedItems
}

func MapItem(item *domainModels.Item) Item {
	return Item{
		ID:               int64(item.ID),
		Title:            item.Title,
		Description:      item.Description,
		ParticipantCount: int(item.ParticipantCount),
		City:             item.City,
		ActivityType:     string(item.ActivityType),
		AuthorName:       item.AuthorName,
		AuthorRating:     item.AuthorRating,
		ImageUrl:         item.ImageUrl,
		FullAddress:      item.FullAddress,
		HasReward:        item.HasReward,
		Duration:         item.Duration,
		StartAt:          item.StartAt.Format(time.RFC3339),
		CreatedAt:        item.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        item.UpdatedAt.Format(time.RFC3339),
	}
}
