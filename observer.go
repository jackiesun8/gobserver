package observer

import (
	"fmt"
	"math/rand"
	"time"
)

type Observer interface {
	Notify(value interface{})
}

type Publisher struct {
	Name      string
	Observers map[string]Observer
}

func (publisher *Publisher) Subscribe(observer Observer) string {
	// genereate Id
	subscriptionId := GenerateIdFromTimestamp()
	publisher.Observers[subscriptionId] = observer
	return subscriptionId
}

func (publisher *Publisher) Unsubscribe(subscriptionId string) {
	delete(publisher.Observers, subscriptionId)
}

func (publisher *Publisher) Publish(value interface{}) {
	for _, obs := range publisher.Observers {
		obs.Notify(value)
	}
}

func GenerateIdFromTimestamp() string {
	return fmt.Sprintf("%x_%x", int32(time.Now().Unix()), rand.Intn(256))
}

func NewPublisher(name string) *Publisher {
	publisher := &Publisher{Name: name}
	publisher.Observers = map[string]Observer{}
	return publisher
}
