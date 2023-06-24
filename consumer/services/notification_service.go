package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anilsenay/go-basic-pubsub/consumer/models"
)

type NotificationService struct {
	pubsub_url string
	topic      string
}

func NewNotificationService(pubsub_url, topic string) *NotificationService {
	return &NotificationService{
		pubsub_url: pubsub_url,
		topic:      topic,
	}
}

func (s *NotificationService) SendNotification(order models.Order) error {
	fmt.Printf("Sending notification to customer: %d\n for order: %d", order.CustomerID, order.OrderID)
	return nil
}

func (s *NotificationService) Consume() (*models.Order, error) {
	resp, err := http.DefaultClient.Get(s.pubsub_url + "?topic=" + s.topic)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	order := models.Order{}
	err = json.NewDecoder(resp.Body).Decode(&order)
	if err != nil {
		return nil, err
	}

	s.SendNotification(order)

	return &order, nil
}
