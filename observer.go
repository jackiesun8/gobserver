package observer

import (
	"github.com/satori/go.uuid"
)

type Observer interface {
	Notify(value interface{})
}

type Publisher struct {
	Name      string
	Observers map[string]Observer
}

func (publisher *Publisher) Subscribe(observer Observer) (subscriptionId string, err error) {
	var b []byte
	b, err = uuid.NewV4().MarshalText()
	if err != nil {
		return
	}
	subscriptionId = string(b)
	publisher.Observers[subscriptionId] = observer
	return
}

func (publisher *Publisher) Unsubscribe(subscriptionId string) {
	delete(publisher.Observers, subscriptionId)
}

func (publisher *Publisher) Publish(value interface{}) {
	for _, obs := range publisher.Observers {
		obs.Notify(value)
	}
}

func NewPublisher(name string) *Publisher {
	publisher := &Publisher{Name: name}
	publisher.Observers = map[string]Observer{}
	return publisher
}
