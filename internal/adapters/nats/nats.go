package natsadapter

import (
	"log"

	"github.com/nats-io/nats.go"
)

type NatsBroker struct {
    nc *nats.Conn
}

func NewNatsBroker(url string) (*NatsBroker, error) {
    nc, err := nats.Connect(url)
    if err != nil {
        return nil, err
    }

    return &NatsBroker{nc: nc}, nil
}

func (n *NatsBroker) Publish(subject string, data []byte) error {
    return n.nc.Publish(subject, data)
}

func (n *NatsBroker) Subscribe(subject string, handler func([]byte) error) error {
    _, err := n.nc.Subscribe(subject, func(msg *nats.Msg) {
        if err := handler(msg.Data); err != nil {
			log.Print(err.Error())
        }
    })
    return err
}

func (n *NatsBroker) Close() error {
    n.nc.Close()
    return nil
}
