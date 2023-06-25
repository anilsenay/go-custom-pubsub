package services

import (
	"encoding/json"
	"fmt"

	"github.com/anilsenay/go-basic-pubsub/consumer/models"
)

type pubSubClient interface {
	Consume() ([]byte, error)
	Subscribe() error
	IsSubscribed() bool
}

type NotificationService struct {
	client pubSubClient
}

func NewNotificationService(client pubSubClient) *NotificationService {
	return &NotificationService{
		client: client,
	}
}

func (s *NotificationService) SendNotification(order models.Order) error {
	fmt.Printf("Sending notification to customer: %d for order: %d \n", order.CustomerID, order.OrderID)
	return nil
}

func (s *NotificationService) Consume() (*models.Order, error) {
	if !s.client.IsSubscribed() {
		return nil, fmt.Errorf("consumer is not subscribed to pubsub")
	}

	resp, err := s.client.Consume()
	if err != nil {
		return nil, err
	}

	order := models.Order{}
	err = json.Unmarshal(resp, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
