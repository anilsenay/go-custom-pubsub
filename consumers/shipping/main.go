package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/anilsenay/go-basic-pubsub/consumers/shipping/services"
)

var pubsub_url = flag.String("h", "http://localhost:8080", "URL to consume messages from")
var topic = flag.String("t", "ORDER", "Topic to consume messages from")
var consumer_count = flag.Int("c", 5, "Number of consumers to spawn")
var delay = flag.Int("d", 5000, "Delay between messages in milliseconds")

func main() {
	flag.Parse()

	pubSubClient := services.NewPubSubClient(*pubsub_url, *topic, "shipping-consumer")
	err := pubSubClient.Subscribe()
	if err != nil {
		fmt.Printf("error while subscribing topic: %v", err)
		return
	}

	consumers := make([]*services.ShippingService, *consumer_count)
	for i := 0; i < *consumer_count; i++ {
		consumers[i] = services.NewShippingService(pubSubClient)
	}

	for {
		for _, consumer := range consumers {
			go func(c *services.ShippingService) {
				order, err := c.Consume()
				if err != nil {
					fmt.Printf("error while consuming message: %v", err)
					return
				}
				err = c.RegisterShipping(*order)
				if err != nil {
					fmt.Printf("error while registering shipping: %v", err)
					return
				}
			}(consumer)
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		}
		time.Sleep(time.Duration(rand.Intn(*delay)) * time.Millisecond)
	}
}
