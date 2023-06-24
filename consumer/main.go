package main

import (
	"flag"
	"time"

	"github.com/anilsenay/go-basic-pubsub/consumer/services"
)

var pubsub_url = flag.String("h", "http://localhost:8080/consume", "URL to consume messages from")
var topic = flag.String("t", "ORDER", "Topic to consume messages from")
var consumer_count = flag.Int("c", 5, "Number of consumers to spawn")
var delay = flag.Int("d", 5000, "Delay between messages in milliseconds")

func main() {
	flag.Parse()

	consumers := make([]*services.NotificationService, *consumer_count)
	for i := 0; i < *consumer_count; i++ {
		consumers[i] = services.NewNotificationService(*pubsub_url, *topic)
	}

	for {
		for _, consumer := range consumers {
			go func(c *services.NotificationService) {
				c.Consume()
			}(consumer)
		}
		time.Sleep(time.Duration(*delay) * time.Millisecond)
	}
}
