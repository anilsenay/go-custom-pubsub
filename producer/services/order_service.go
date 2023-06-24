package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/anilsenay/go-basic-pubsub/producer/models"
)

type Message struct {
	Topic string
	Body  models.Order
}

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

func (s *OrderService) CreateOrder() error {
	randomOrderId := rand.Intn(100000)

	body := Message{
		Topic: s.topic,
		Body: models.Order{
			OrderID:    randomOrderId,
			CustomerID: 1,
			ItemID:     1,
			Quantity:   1,
			Price:      1.0,
			Total:      1.0,
		},
	}

	j, err := json.Marshal(body)
	if err != nil {
		return err
	}

	fmt.Printf("Sending message to pubsub with order id: %d\n", randomOrderId)
	_, err = http.DefaultClient.Post(s.pubsub_url, "application/json", bytes.NewBuffer(j))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
