package observer

import (
	"fmt"
	"testing"
	"time"
)

var (
	Message = ""
)

type TestObserver struct {
	Observer
	SubscriptionId string
}

func (t TestObserver) Notify(value interface{}) {
	Message = fmt.Sprintf("Received %v", value)
}

func Test(t *testing.T) {

	assertEqual := func(val interface{}, exp interface{}) {
		if val != exp {
			t.Errorf("Expected %v, got %v.", exp, val)
		}
	}

	publisher := NewPublisher("TestPublisher")
	testObserver := TestObserver{}
	testObserver2 := TestObserver{}

	testObserver.SubscriptionId = publisher.Subscribe(testObserver)
	testObserver2.SubscriptionId = publisher.Subscribe(testObserver2)

	publisher.Unsubscribe(testObserver2.SubscriptionId)
	publisher.Publish("test")

	time.Sleep(1 * time.Second)

	assertEqual(Message, "Received test")
}
