package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/anilsenay/go-basic-pubsub/pubsub/models"
)

type PubsubMock struct {
}

func (p *PubsubMock) Enqueue(topic, msg string) error {
	fmt.Printf("Topic: %s, Message: %s\n", topic, msg)
	return nil
}

func (p *PubsubMock) Dequeue(topic string) (string, error) {
	return "message", nil
}

func TestNewPubSubHandler(t *testing.T) {

	ps := &PubsubMock{}

	handler := NewPubSubHandler("8080", ps)

	go handler.Listen()
	time.Sleep(1 * time.Second) // wait for server to start

	// PRODUCER
	msg := models.Message{
		Topic: "TEST_TOPIC",
		Body:  "TEST_BODY",
	}

	b, _ := json.Marshal(msg)
	body := bytes.NewBuffer(b)

	res, err := http.Post("http://localhost:8080/produce", "application/json", body)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	// CONSUMER
	res, err = http.Get("http://localhost:8080/consume?topic=TEST_TOPIC")
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
	if string(b) != "message" {
		t.Errorf("Expected message %s, got %s", "message", string(b))
	}

	// CLOSE
	defer func() {
		if r := recover(); r != nil {
			handler.Close()
		}
	}()
}
