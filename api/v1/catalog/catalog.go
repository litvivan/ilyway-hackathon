package catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/litvivan/ilyway/api/models"
	"github.com/litvivan/ilyway/api/renderer"
	"github.com/litvivan/ilyway/app/services/catalog"
	"github.com/litvivan/ilyway/app/services/city"
	"github.com/litvivan/ilyway/pkg/repofilter"
)

type Service struct {
	cityService    city.Service
	catalogService catalog.Service
}

func NewService(
	cityService city.Service,
	catalogService catalog.Service,
) *Service {
	return &Service{
		cityService:    cityService,
		catalogService: catalogService,
	}
}

func (svc *Service) GetItem(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	idVal := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idVal)
	if err != nil {
		renderer.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	item, err := svc.catalogService.GetItem(ctx, id)
	if err != nil {
		renderer.RenderError(w, r, http.StatusNotFound, "Item not found")
		return
	}

	mappedItem := models.MapItem(item)

	renderer.RenderResponse(w, r, http.StatusOK, mappedItem)
}

func (svc *Service) ListItems(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var req models.ListItemRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		renderer.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	filter, err := repofilter.NewFromMap(req.Filter)
	if err != nil {
		renderer.RenderError(w, r, http.StatusNotFound, "Error creating filter")
		return
	}

	items, err := svc.catalogService.ListItems(ctx, filter)
	if err != nil {
		renderer.RenderError(w, r, http.StatusNotFound, fmt.Sprintf("Items not found: %s", err.Error()))
		return
	}

	renderer.RenderResponse(w, r, http.StatusOK, models.MapItems(items))
}

func (svc *Service) SeedItems(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	items, err := svc.catalogService.SeedItems(ctx, 10)
	if err != nil {
		renderer.RenderError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	renderer.RenderResponse(w, r, http.StatusOK, models.MapItems(items))
}
