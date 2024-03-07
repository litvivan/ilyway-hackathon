package item

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"

	"github.com/litvivan/ilyway/app/models"
	"github.com/litvivan/ilyway/infra/db/storage"
	"github.com/litvivan/ilyway/pkg/repofilter"
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

type ListFilter struct {
	MinParticipantCount *int      `json:"min_participant_count"`
	MaxParticipantCount *int      `json:"max_participant_count"`
	ActivityType        *string   `json:"activity_type"`
	MinAuthorRating     *float64  `json:"min_author_rating"`
	City                *string   `json:"city"`
	MinStartAt          time.Time `json:"min_start_at"`
	MaxStartAt          time.Time `json:"max_start_at"`
}

func (lf ListFilter) FromFilter(f repofilter.Filter) (ListFilter, error) {
	lf = ListFilter{}
	fromRepofilter, err := f.Make(&lf)
	if err != nil {
		return ListFilter{}, err
	}

	listFilter, ok := fromRepofilter.(*ListFilter)
	if !ok {
		return ListFilter{}, errors.Errorf("invalid list filter %T", fromRepofilter)
	}

	return *listFilter, nil
}

func (r *Repo) List(ctx context.Context, filter repofilter.Filter) ([]*models.Item, error) {
	listFilter, err := ListFilter{}.FromFilter(filter)
	if err != nil {
		return nil, err
	}

	filterParams := storage.ListItemsParams{}

	if listFilter.MinParticipantCount != nil {
		filterParams.MinParticipantCount = sql.NullInt32{Int32: int32(*listFilter.MinParticipantCount), Valid: true}
	}

	if listFilter.MaxParticipantCount != nil {
		filterParams.MaxParticipantCount = sql.NullInt32{Int32: int32(*listFilter.MaxParticipantCount), Valid: true}
	}

	if listFilter.ActivityType != nil {
		filterParams.ActivityType = sql.NullString{String: *listFilter.ActivityType, Valid: true}
	}

	if listFilter.MinAuthorRating != nil {
		filterParams.MinAuthorRating = sql.NullFloat64{Float64: *listFilter.MinAuthorRating, Valid: true}
	}

	if !listFilter.MinStartAt.IsZero() {
		filterParams.MinStartAt = sql.NullTime{Time: listFilter.MinStartAt.UTC(), Valid: true}
	}

	if !listFilter.MaxStartAt.IsZero() {
		filterParams.MaxStartAt = sql.NullTime{Time: listFilter.MaxStartAt.UTC(), Valid: true}
	}

	dbModels, err := r.q.ListItems(ctx, filterParams)
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
