package adapters

import "github.com/nats-io/nats.go"

type NATSClient struct {
	Conn *nats.Conn
}

func NewNATSClient(url string) (*NATSClient, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NATSClient{Conn: nc}, nil
}

func (n *NATSClient) Publish(subject string, message []byte) error {
	return n.Conn.Publish(subject, []byte(message))
}

func (n *NATSClient) Subscribe(subject string, handler func(msg *nats.Msg)) (*nats.Subscription, error) {
	return n.Conn.Subscribe(subject, handler)
}
