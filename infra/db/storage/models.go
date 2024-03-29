// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package storage

import (
	"database/sql"
	"time"

	domain "github.com/litvivan/ilyway/app/models"
)

type Item struct {
	ID               int32
	Title            string
	Description      string
	ParticipantCount int32
	ActivityType     domain.ActivityType
	City             string
	AuthorName       string
	AuthorRating     float64
	ImageUrl         string
	FullAddress      string
	HasReward        bool
	Duration         string
	StartAt          time.Time
	CreatedAt        sql.NullTime
	UpdatedAt        sql.NullTime
}
