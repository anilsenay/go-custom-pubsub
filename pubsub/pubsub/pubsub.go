package pubsub

import (
	"fmt"

	"github.com/fatih/color"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

type PubSubManager struct {
	buffer_size int
	topics      map[string]chan string
}

func NewPubSubManager(buffer_size int) *PubSubManager {
	return &PubSubManager{
		buffer_size: buffer_size,
		topics:      make(map[string]chan string),
	}
}

func (m *PubSubManager) Enqueue(topic, msg string) error {
	if _, ok := m.topics[topic]; !ok {
		return fmt.Errorf("topic %s does not exist", topic)
	}
	m.topics[topic] <- msg
	fmt.Printf("%s message: %s\n", green("ENQUEUED"), msg)
	return nil
}

func (m *PubSubManager) Dequeue(topic string) (string, error) {
	if _, ok := m.topics[topic]; !ok {
		return "", fmt.Errorf("topic %s does not exist", topic)
	}

	msg := <-m.topics[topic]
	fmt.Printf("%s message: %s\n", yellow("DEQUEUED"), msg)
	return msg, nil
}

func (m *PubSubManager) CreateTopic(topic string) {
	m.topics[topic] = make(chan string, m.buffer_size)
}

func (m *PubSubManager) DeleteTopic(topic string) {
	delete(m.topics, topic)
}

func (m *PubSubManager) Close() {
	for _, topic := range m.topics {
		close(topic)
	}
}
