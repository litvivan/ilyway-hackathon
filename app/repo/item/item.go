package item

import (
	"context"
	"database/sql"
	"time"

	"github.com/litvivan/ilyway/app/models"
	"github.com/litvivan/ilyway/infra/db/storage"
)

type Repo struct {
	db *sql.DB
	q  storage.Queries
}

func NewRepo(
	db *sql.DB,
	q *storage.Queries,
) *Repo {
	return &Repo{
		db: db,
		q:  *q,
	}
}

func (r *Repo) Get(ctx context.Context, id int) (*models.Item, error) {
	dbModel, err := r.q.GetItem(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return mapItem(dbModel), nil
}

func (r *Repo) List(ctx context.Context) ([]*models.Item, error) {
	dbModels, err := r.q.ListItems(ctx)
	if err != nil {
		return nil, err
	}

	return mapItems(dbModels), nil
}

type ItemCreateRequest struct {
	Title            string
	Description      string
	ParticipantCount int
	ActivityType     models.ActivityType
	City             string
	AuthorName       string
	AuthorRating     float32
	ImageUrl         string
	FullAddress      string
	HasReward        bool
	Duration         string
	StartAt          time.Time
}

func (r *Repo) Create(ctx context.Context, req ItemCreateRequest) (*models.Item, error) {
	dbModel, err := r.q.InsertItem(ctx, storage.InsertItemParams{
		Title:            req.Title,
		Description:      req.Description,
		ParticipantCount: int32(req.ParticipantCount),
		ActivityType:     req.ActivityType,
		City:             req.City,
		AuthorName:       req.AuthorName,
		AuthorRating:     float64(req.AuthorRating),
		ImageUrl:         req.ImageUrl,
		FullAddress:      req.FullAddress,
		HasReward:        req.HasReward,
		Duration:         req.Duration,
		StartAt:          req.StartAt,
	})
	if err != nil {
		return nil, err
	}

	return mapItem(dbModel), nil
}
