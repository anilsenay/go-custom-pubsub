package services

import (
	"encoding/json"
	"fmt"

	"github.com/anilsenay/go-custom-pubsub/consumers/notification/models"
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

func (s *NotificationService) Consume() error {
	if !s.client.IsSubscribed() {
		return fmt.Errorf("consumer is not subscribed to pubsub")
	}

	resp, err := s.client.Consume()
	if err != nil {
		return err
	}

	order := models.Order{}
	err = json.Unmarshal(resp, &order)
	if err != nil {
		return err
	}

	err = s.SendNotification(order)
	if err != nil {
		return err
	}

	return nil
}
