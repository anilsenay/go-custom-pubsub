package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/anilsenay/go-custom-pubsub/pubsub/models"
	"github.com/anilsenay/go-custom-pubsub/pubsub/pubsub"
)

type MessageBody struct {
	OrderID    int `json:"order_id"`
	CustomerID int `json:"customer_id"`
}

func TestNewPubSubHandler(t *testing.T) {

	ps := pubsub.NewPubSubManager(100)
	ps.CreateTopic("TEST_TOPIC")

	handler := NewPubSubHandler("8080", ps)

	go handler.Listen()
	time.Sleep(1 * time.Second) // wait for server to start

	// CONSUMER - SUBSCRIBE
	res, err := http.Get("http://localhost:8080/subscribe?topic=TEST_TOPIC&subscriber=TEST_SUBSCRIBER")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	// PRODUCER - PRODUCE
	msg := models.Message{
		Topic: "TEST_TOPIC",
		Body: MessageBody{
			OrderID:    123,
			CustomerID: 1234,
		},
	}

	b, _ := json.Marshal(msg)
	body := bytes.NewBuffer(b)

	res, err = http.Post("http://localhost:8080/produce", "application/json", body)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	// CONSUMER - CONSUME
	res, err = http.Get("http://localhost:8080/consume?topic=TEST_TOPIC&subscriber=TEST_SUBSCRIBER")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	b, err = io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	msgBody := MessageBody{}
	json.Unmarshal(b, &msgBody)

	if msgBody.OrderID != 123 || msgBody.CustomerID != 1234 {
		t.Errorf("Expected order id %d and customer id %d, got order id %d and customer id %d", 123, 1234, msgBody.OrderID, msgBody.CustomerID)
	}

	// CLOSE
	defer func() {
		if r := recover(); r != nil {
			handler.Close()
		}
	}()
}
