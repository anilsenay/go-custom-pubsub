package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anilsenay/go-custom-pubsub/producers/shipping/models"
)

func TestShippingService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := models.Message{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("PUBSUB: Received message from producer with order id: %d\n", body.Body.OrderID)

		if body.Topic != "TEST_TOPIC" {
			t.Errorf("Expected topic to be TEST_TOPIC, got %s", body.Topic)
		}

		if body.Body.OrderID != 123 {
			t.Errorf("Expected order id to be 123, got %d", body.Body.OrderID)
		}
	}))

	s := &ShippingService{
		pubsub_url: server.URL,
		topic:      "TEST_TOPIC",
	}

	status := models.ShippingStatus{
		OrderID:    123,
		CustomerID: 1234,
		Status:     "shipped",
	}

	err := s.UpdateStatus(status)
	if err != nil {
		t.Error(err)
	}
}
