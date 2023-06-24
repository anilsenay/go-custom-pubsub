package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anilsenay/go-basic-pubsub/consumer/models"
)

func TestNotificationService_Consume(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		order := models.Order{
			OrderID:    123,
			CustomerID: 1234,
		}
		json.NewEncoder(w).Encode(order)
	}))

	s := &NotificationService{
		pubsub_url: server.URL,
		topic:      "TEST_TOPIC",
	}

	order, err := s.Consume()
	if err != nil {
		t.Error(err)
	}

	if order.OrderID != 123 {
		t.Errorf("Expected order id to be 123, got %d", order.OrderID)
	}
}
