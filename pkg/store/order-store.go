package store

import (
	"fmt"
	"sync"

	"github.com/iboware/location-history/pkg/model"
)

// OrderStore is an in-memory database of orders with history
type OrderStore struct {
	sync.Mutex
	orders map[string][]model.Location
}

func New() *OrderStore {
	s := &OrderStore{}
	s.orders = make(map[string][]model.Location)
	return s
}

// CreateTask creates a new task in the store.
func (s *OrderStore) AppendLocation(orderId string, location model.Location) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.orders[orderId]; !ok {
		s.orders[orderId] = make([]model.Location, 0)
	}
	s.orders[orderId] = append(s.orders[orderId], location)
}

// GetTask retrieves a task from the store, by id. If no such id exists, an
// error is returned.
func (s *OrderStore) GetHistory(orderId string, max *int) (model.OrderHistory, error) {
	s.Lock()
	defer s.Unlock()
	history, ok := s.orders[orderId]
	if !ok {
		return model.OrderHistory{}, fmt.Errorf("order with id=%v not found", orderId)
	}

	if max != nil && *max <= len(history) {
		history = history[len(history)-*max:]
	}

	return model.OrderHistory{
		OrderId: &orderId,
		History: &history,
	}, nil
}

// DeleteTask deletes the task with the given id. If no such id exists, an error
// is returned.
func (s *OrderStore) DeleteHistory(orderId string) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.orders[orderId]; !ok {
		return fmt.Errorf("order with id=%v not found", orderId)
	}

	delete(s.orders, orderId)
	return nil
}
