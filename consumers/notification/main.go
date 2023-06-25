package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/anilsenay/go-custom-pubsub/consumers/notification/models"
	"github.com/anilsenay/go-custom-pubsub/consumers/notification/services"
)

var defaultTopics = []string{"ORDER", "SHIPMENT"}
var topics models.ArrayFlags
var pubsub_url = flag.String("h", "http://localhost:8080", "URL to consume messages from")
var consumer_count = flag.Int("c", 10, "Number of consumers to spawn")
var delay = flag.Int("d", 5000, "Delay between messages in milliseconds")

type Consumer interface {
	Consume() error
}

func main() {
	flag.Var(&topics, "t", "Topic to consume messages from")
	flag.Parse()

	if len(topics) == 0 {
		topics = append(topics, defaultTopics...)
	}

	consumers := make([]Consumer, *consumer_count)
	for _, topic := range topics {
		pubSubClient := services.NewPubSubClient(*pubsub_url, topic, "notification-consumer")
		err := pubSubClient.Subscribe()
		if err != nil {
			fmt.Printf("error while subscribing topic: %v", err)
			return
		}
		for i := 0; i < *consumer_count; i++ {
			consumers[i] = services.NewNotificationService(pubSubClient)
		}
	}

	for {
		for _, consumer := range consumers {
			go func(c Consumer) {
				err := c.Consume()
				if err != nil {
					fmt.Printf("error while consuming message: %v", err)
					return
				}
			}(consumer)
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		}
		time.Sleep(time.Duration(rand.Intn(*delay)) * time.Millisecond)
	}
}
