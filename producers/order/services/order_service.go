package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anilsenay/go-custom-pubsub/producers/order/models"
)

type OrderService struct {
	pubsub_url string
	topic      string
}

func NewOrderService(pubsub_url, topic string) *OrderService {
	return &OrderService{
		pubsub_url: pubsub_url,
		topic:      topic,
	}
}

func (s *OrderService) CreateOrder(order models.Order) error {
	body := models.Message{
		Topic: s.topic,
		Body:  order,
	}

	j, err := json.Marshal(body)
	if err != nil {
		return err
	}

	fmt.Printf("Sending message to pubsub with order id: %d\n", order.OrderID)
	resp, err := http.DefaultClient.Post(s.pubsub_url, "application/json", bytes.NewBuffer(j))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
