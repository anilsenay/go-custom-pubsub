package pubsub

import (
	"fmt"

	"github.com/fatih/color"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var cyan = color.New(color.FgCyan).SprintFunc()

type PubSubManager struct {
	buffer_size int
	topics      map[string]map[string]chan string
}

func NewPubSubManager(buffer_size int) *PubSubManager {
	return &PubSubManager{
		buffer_size: buffer_size,
		topics:      make(map[string]map[string]chan string),
	}
}

func (m *PubSubManager) Subscribe(topic, subscriber string) error {
	if _, ok := m.topics[topic]; !ok {
		return fmt.Errorf("topic %s does not exist", topic)
	}
	if _, ok := m.topics[topic][subscriber]; ok {
		return fmt.Errorf("subscriber %s already exists", subscriber)
	}
	m.topics[topic][subscriber] = make(chan string, m.buffer_size)
	fmt.Printf("%s subscriber: %s -> topic: %s\n", cyan("SUBSCRIBED"), subscriber, topic)
	return nil
}

func (m *PubSubManager) Enqueue(topic, msg string) error {
	if _, ok := m.topics[topic]; !ok {
		return fmt.Errorf("topic %s does not exist", topic)
	}

	for _, subscriber := range m.topics[topic] {
		subscriber <- msg
	}
	fmt.Printf("%s message: %s\n", green("ENQUEUED"), msg)
	return nil
}

func (m *PubSubManager) Dequeue(topic, subscriber string) (string, error) {
	if _, ok := m.topics[topic]; !ok {
		return "", fmt.Errorf("topic %s does not exist", topic)
	}

	if _, ok := m.topics[topic][subscriber]; !ok {
		return "", fmt.Errorf("subscriber %s does not exist", subscriber)
	}
	msg := <-m.topics[topic][subscriber]
	fmt.Printf("%s message: %s\n", yellow("DEQUEUED"), msg)
	return msg, nil
}

func (m *PubSubManager) CreateTopic(topic string) {
	m.topics[topic] = make(map[string]chan string, m.buffer_size)
}

func (m *PubSubManager) DeleteTopic(topic string) {
	for _, subscriber := range m.topics[topic] {
		close(subscriber)
	}
	delete(m.topics, topic)
}

func (m *PubSubManager) Close() {
	for _, topic := range m.topics {
		for _, subscriber := range topic {
			close(subscriber)
		}
	}
}
