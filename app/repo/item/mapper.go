package item

import (
	"github.com/litvivan/ilyway/app/models"
	"github.com/litvivan/ilyway/infra/db/storage"
)

func mapItem(dbModel storage.Item) *models.Item {
	return &models.Item{
		ID:               int64(dbModel.ID),
		Title:            dbModel.Title,
		Description:      dbModel.Description,
		ParticipantCount: int(dbModel.ParticipantCount),
		ActivityType:     dbModel.ActivityType,
		AuthorName:       dbModel.AuthorName,
		AuthorRating:     dbModel.AuthorRating,
		ImageUrl:         dbModel.ImageUrl,
		FullAddress:      dbModel.FullAddress,
		City:             dbModel.City,
		HasReward:        dbModel.HasReward,
		Duration:         dbModel.Duration,
		StartAt:          dbModel.StartAt,
		CreatedAt:        dbModel.CreatedAt.Time,
		UpdatedAt:        dbModel.UpdatedAt.Time,
	}
}

func mapItems(dbModels []storage.Item) []*models.Item {
	items := make([]*models.Item, 0, len(dbModels))
	for i := range dbModels {
		items = append(items, mapItem(dbModels[i]))
	}

	return items
}
