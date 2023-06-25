package services

import (
	"fmt"
	"io"
	"net/http"
)

type PubSubClient struct {
	pubsub_url     string
	topic          string
	subscriptionID string
	subscribed     bool
}

func NewPubSubClient(pubsub_url, topic, subscriptionID string) *PubSubClient {
	return &PubSubClient{
		pubsub_url:     pubsub_url,
		topic:          topic,
		subscriptionID: subscriptionID,
	}
}

func (s *PubSubClient) Consume() ([]byte, error) {
	resp, err := http.DefaultClient.Get(s.pubsub_url + "/consume?topic=" + s.topic + "&subscriber=" + s.subscriptionID)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (s *PubSubClient) Subscribe() error {
	fmt.Printf("Sending subscription request to pubsub for consumer with subscriptionID: %s\n", s.subscriptionID)
	resp, err := http.DefaultClient.Post(s.pubsub_url+"/subscribe?topic="+s.topic+"&subscriber="+s.subscriptionID, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	s.subscribed = true
	return nil
}

func (s *PubSubClient) IsSubscribed() bool {
	return s.subscribed
}
