package event

import (
	"bytes"
	"encoding/gob"

	"github.com/nats-io/go-nats"
)

type NatsEventStore struct {
	nc                      *nats.Conn
	bookCreatedSubscription *nats.Subscription
	bookCreatedChan         chan BookCreatedMessage
}

func NewNats(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{
		nc: nc,
	}, nil
}

func (n *NatsEventStore) Close() {
	if n.nc != nil {
		n.nc.Close()
	}
	if n.bookCreatedSubscription != nil {
		n.bookCreatedSubscription.Unsubscribe()
	}
	close(n.bookCreatedChan)
}

func (n *NatsEventStore) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}

func (n *NatsEventStore) SubsribeBookCreated() (<-chan BookCreatedMessage, error) {
	m := BookCreatedMessage{}
	n.bookCreatedChan = make(chan BookCreatedMessage)
	ch := make(chan *nats.Msg)
	var err error
	n.bookCreatedSubscription, err = n.nc.ChanSubscribe(m.Key(), ch)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case msg := <-ch:
				n.readMessage(msg.Data, &m)
				n.bookCreatedChan <- m
			}
		}
	}()
	return (<-chan BookCreatedMessage)(n.bookCreatedChan), nil
}

func (n *NatsEventStore) OnBookCreated(f func(BookCreatedMessage)) error {
	m := BookCreatedMessage{}
	var err error
	n.bookCreatedSubscription, err = n.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		n.readMessage(msg.Data, &m)
		f(m)
	})
	return err
}
