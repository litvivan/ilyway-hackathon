package catalog

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/litvivan/ilyway/app/models"
	itemRepository "github.com/litvivan/ilyway/app/repo/item"
	"github.com/litvivan/ilyway/pkg/repofilter"
)

type Service struct {
	itemRepo *itemRepository.Repo
}

func NewService(itemRepo *itemRepository.Repo) *Service {
	return &Service{
		itemRepo: itemRepo,
	}
}

func (s *Service) GetItem(ctx context.Context, id int) (*models.Item, error) {
	return s.itemRepo.Get(ctx, id)
}

func (s *Service) ListItems(ctx context.Context, filter repofilter.Filter) ([]*models.Item, error) {
	return s.itemRepo.List(ctx, filter)
}

func (s *Service) SeedItems(ctx context.Context, count int) ([]*models.Item, error) {
	cityNames := make([]string, 0, len(models.Cities))
	for i := range models.Cities {
		cityNames = append(cityNames, models.Cities[i].Name)
	}

	createdItems := make([]*models.Item, 0, count)

	for i := 0; i < count; i++ {
		activityType := genRandomActivityType()
		createdItem, err := s.itemRepo.Create(ctx, itemRepository.ItemCreateRequest{
			Title:            gofakeit.HipsterSentence(5),
			Description:      gofakeit.HipsterParagraph(5, 5, 10, ""),
			ParticipantCount: genRandomParticipantsCount(),
			ActivityType:     activityType,
			AuthorName:       gofakeit.Name(),
			AuthorRating:     float32(rand.Intn(5)),
			ImageUrl:         genRandomImageUrl(320, 240),
			FullAddress:      gofakeit.Address().Address,
			HasReward:        rand.Intn(2) == 1,
			Duration:         genRandomDuration(),
			City:             cityNames[rand.Intn(len(cityNames))],
			StartAt:          genRandomStartAt(),
		})
		if err != nil {
			return nil, err
		}
		createdItems = append(createdItems, createdItem)
	}

	return createdItems, nil
}

func genRandomParticipantsCount() int {
	num := rand.Intn(50)
	sub := 0
	if num > 5 {
		sub = num % 5
	} else {
		num = 5
	}

	return num - sub
}

func genRandomActivityType() models.ActivityType {
	num := rand.Intn(len(models.ActivityTypes))
	return models.ActivityTypes[num]
}

func genRandomImageUrl(width, height int) string {
	return fmt.Sprintf("https://loremflickr.com/%d/%d/%d", width, height, rand.Intn(10000))
}

func genRandomStartAt() time.Time {
	n := time.Now()

	days := rand.Intn(20)

	return time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, days)
}

func genRandomDuration() string {
	hours := rand.Intn(5)
	if hours == 0 {
		hours = 1
	}

	return fmt.Sprintf("%d h", hours)
}
