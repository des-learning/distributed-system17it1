package main

import (
	"fmt"
	"sync"
	"time"
)

type Publisher struct {
	subscribers *[]chan interface{}
}

func (p *Publisher) PublishTo(subscribers *[]chan interface{}) {
	p.subscribers = subscribers
}

func (p *Publisher) Publish(message interface{}) {
	for _, subscriber := range *p.subscribers {
		subscriber <- message
	}
}

type Subscriber interface {
	Consume(<-chan interface{})
}

type MyConsumer struct{}

func (m *MyConsumer) Consume(channel <-chan interface{}) {
	for data := range channel {
		fmt.Println(data)
	}
}

type MyConsumer1 struct{}

func (m *MyConsumer1) Consume(channel <-chan interface{}) {
	for data := range channel {
		fmt.Println("MyConsumer1", data)
	}
}

type MyConsumer2 struct {
}

func (m *MyConsumer2) Consume(channel <-chan interface{}) {
	for data := range channel {
		fmt.Println("MyConsumer2 ", data, len(data.(string)))
	}
}

type broker struct {
	topics map[string]*[]chan interface{}
	mux    sync.Mutex
}

func (b *broker) Subscribe(topic string, subscriber Subscriber) {
	b.mux.Lock()
	channel := make(chan interface{}, 1)
	if existingTopic, found := b.topics[topic]; found {
		*existingTopic = append(*existingTopic, channel)
	} else {
		b.topics[topic] = &[]chan interface{}{channel}
	}
	go subscriber.Consume(channel)
	b.mux.Unlock()
}

func (b *broker) Publish(topic string, publisher *Publisher) {
	b.mux.Lock()
	if _, found := b.topics[topic]; !found {
		b.topics[topic] = &[]chan interface{}{}
	}
	publisher.PublishTo(b.topics[topic])
	b.mux.Unlock()
}

func NewBroker() *broker {
	return &broker{
		make(map[string]*[]chan interface{}),
		sync.Mutex{},
	}
}

func main() {
	b := NewBroker()
	pub := &Publisher{}
	con1 := &MyConsumer{}
	con2 := &MyConsumer1{}

	wg := sync.WaitGroup{}
	wg.Add(2)
	b.Publish("abc", pub)

	time.AfterFunc(time.Second*2, func() {
		fmt.Println("consumer 2 joined")
		b.Subscribe("abc", &MyConsumer2{})
		b.Subscribe("abc", con1)
		b.Subscribe("abc", con2)
	})

	for i := 1; ; i++ {
		pub.Publish(fmt.Sprintf("Hello %d", i))
	}

	wg.Wait()
}
