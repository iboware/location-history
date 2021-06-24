package store

import (
	"testing"

	"github.com/iboware/location-history/pkg/helper"
	"github.com/iboware/location-history/pkg/model"
)

func TestOrderStore_AppendLocation(t *testing.T) {
	testStore := New()
	testOrders := []struct {
		orderId string
		lat     float32
		lng     float32
	}{
		{"order1", 1, 1},
		{"order1", 1, 2},
		{"order1", 1, 3},
		{"order2", 2, 1},
		{"order2", 2, 2},
		{"order2", 2, 3},
		{"order2", 2, 4},
		{"order3", 3, 1},
	}
	expected := model.OrderHistory{
		History: &[]model.Location{
			{Lat: helper.CreateFloatP32FromFloat(1), Lng: helper.CreateFloatP32FromFloat(2)},
			{Lat: helper.CreateFloatP32FromFloat(1), Lng: helper.CreateFloatP32FromFloat(3)},
		},
		OrderId: helper.CreateStringPFromString("order1"),
	}

	for _, order := range testOrders {
		lct := model.Location{Lat: helper.CreateFloatP32FromFloat(order.lat), Lng: helper.CreateFloatP32FromFloat(order.lng)}
		testStore.AppendLocation(order.orderId, lct)
	}

	res, err := testStore.GetHistory("order1", helper.CreateIntPFromInt(2))

	if err != nil {
		t.Error("Got error while not expecting!")
	}
	if *res.OrderId != *expected.OrderId {
		t.Errorf("Order history, got: %v, want: %v.", res.OrderId, expected.OrderId)
	}

	expectedHistory := *(expected.History)

	if len(expectedHistory) != len(*res.History) {
		t.Errorf("Order history length, got: %d, want: %d.", len(*res.History), len(expectedHistory))

	}
	for i, h := range *res.History {
		if *h.Lat != *expectedHistory[i].Lat || *h.Lng != *expectedHistory[i].Lng {
			t.Errorf("Order history, got: %+v, want: %+v.", h, expectedHistory[i])
		}
	}

}
func TestOrderStore_DeleteHistory(t *testing.T) {
	//TODO
}

func TestOrderStore_GetHistory(t *testing.T) {
	//TODO
}
