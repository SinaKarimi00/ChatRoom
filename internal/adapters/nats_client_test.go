package adapters

import (
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

func TestNATSClient(t *testing.T) {
	client, err := NewNATSClient("nats://localhost:4222")
	if err != nil {
		t.Fatalf("Failed to connect to NATS: %v", err)
	}

	received := make(chan string, 1)

	_, err = client.Subscribe("test", func(msg *nats.Msg) {
		received <- string(msg.Data)
	})
	if err != nil {
		t.Fatalf("Failed to subscribe: %v", err)
	}

	err = client.Publish("test", []byte("Hello, NATS!"))
	if err != nil {
		t.Fatalf("Failed to publish: %v", err)
	}

	select {
	case msg := <-received:
		if msg != "Hello, NATS!" {
			t.Errorf("Expected 'Hello, NATS!', got '%s'", msg)
		}
	case <-time.After(2 * time.Second):
		t.Error("Did not receive a message in time")
	}
}
