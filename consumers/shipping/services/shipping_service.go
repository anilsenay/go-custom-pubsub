package services

import (
	"encoding/json"
	"fmt"

	"github.com/anilsenay/go-custom-pubsub/consumers/shipping/models"
)

type pubSubClient interface {
	Consume() ([]byte, error)
	Subscribe() error
	IsSubscribed() bool
}

type ShippingService struct {
	client pubSubClient
}

func NewShippingService(client pubSubClient) *ShippingService {
	return &ShippingService{
		client: client,
	}
}

func (s *ShippingService) RegisterShipping(order models.Order) error {
	fmt.Printf("Registering shipping for customer: %d with order: %d \n", order.CustomerID, order.OrderID)
	return nil
}

func (s *ShippingService) Consume() (*models.Order, error) {
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
