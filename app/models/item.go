package models

import "time"

type Item struct {
	ID               int64
	Title            string
	Description      string
	ParticipantCount int
	ActivityType     ActivityType
	AuthorName       string
	AuthorRating     float64
	ImageUrl         string
	HasReward        bool
	Duration         string
	City             string
	FullAddress      string
	StartAt          time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ActivityType string

const (
	ActivityTypeCleaning       ActivityType = "Cleaning"
	ActivityTypeCooking        ActivityType = "Cooking"
	ActivityTypeDelivery       ActivityType = "Delivery"
	ActivityTypeAnimal         ActivityType = "Animal care"
	ActivityTypeRepair         ActivityType = "Repair"
	ActivityTypeElderCompanion ActivityType = "Elder companion"
	ActivityTypeTutoring       ActivityType = "Tutoring"
	ActivityTypePlantcare      ActivityType = "Plant care"
	ActivityTypeShopping       ActivityType = "Shopping"
	ActivityTypePacking        ActivityType = "Packing"
)

var ActivityTypes = []ActivityType{
	ActivityTypeCleaning,
	ActivityTypeCooking,
	ActivityTypeDelivery,
	ActivityTypeAnimal,
	ActivityTypeRepair,
	ActivityTypeElderCompanion,
	ActivityTypeTutoring,
	ActivityTypePlantcare,
	ActivityTypeShopping,
	ActivityTypePacking,
}
