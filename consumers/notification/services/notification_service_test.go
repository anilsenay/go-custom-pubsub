package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anilsenay/go-custom-pubsub/consumers/notification/models"
)

func TestNotificationService_Consume(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		order := models.Order{
			OrderID:    123,
			CustomerID: 1234,
		}
		json.NewEncoder(w).Encode(order)
	}))

	c := &PubSubClient{
		pubsub_url:     server.URL,
		topic:          "TEST_TOPIC",
		subscriptionID: "TEST_SUBSCRIPTION",
	}

	c.Subscribe()

	s := &NotificationService{
		client: c,
	}

	err := s.Consume()
	if err != nil {
		t.Error(err)
	}
}
