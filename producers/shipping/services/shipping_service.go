package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anilsenay/go-basic-pubsub/producers/shipping/models"
)

type ShippingService struct {
	pubsub_url string
	topic      string
}

func NewShippingService(pubsub_url, topic string) *ShippingService {
	return &ShippingService{
		pubsub_url: pubsub_url,
		topic:      topic,
	}
}

func (s *ShippingService) UpdateStatus(status models.ShippingStatus) error {
	body := models.Message{
		Topic: s.topic,
		Body:  status,
	}

	j, err := json.Marshal(body)
	if err != nil {
		return err
	}

	fmt.Printf("Sending message to pubsub with order id: %d\n", status.OrderID)
	resp, err := http.DefaultClient.Post(s.pubsub_url, "application/json", bytes.NewBuffer(j))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
