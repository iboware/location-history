package server

import (
	"net/http"

	"github.com/iboware/location-history/pkg/model"
	"github.com/iboware/location-history/pkg/store"

	"github.com/labstack/echo/v4"
)

type LocationServer struct {
	store *store.OrderStore
}

func NewLocationServer() *LocationServer {
	store := store.New()
	return &LocationServer{store: store}
}

// Delete history
// (DELETE /location/{order_id})
func (s *LocationServer) DeleteHistory(ctx echo.Context, orderId string) error {
	err := s.store.DeleteHistory(orderId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return ctx.NoContent(http.StatusOK)

}

// Returns order history
// (GET /location/{order_id})
func (s *LocationServer) GetHistory(ctx echo.Context, orderId string, params model.GetHistoryParams) error {
	history, err := s.store.GetHistory(orderId, params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, history)
}

// Appends a location
// (PUT /location/{order_id})
func (s *LocationServer) AppendHistory(ctx echo.Context, orderId string) error {
	var location model.Location
	err := ctx.Bind(&location)
	if err != nil {
		return err
	}

	if location.Lat == nil || location.Lng == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Lat and Lng are required!")
	}

	s.store.AppendLocation(orderId, location)
	return ctx.NoContent(http.StatusOK)
}
