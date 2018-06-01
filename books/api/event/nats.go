package event

import (
	"bytes"
	"encoding/gob"

	"github.com/golovers/cqrs-example/books/api/schema"
	"github.com/nats-io/go-nats"
)

type NatsEventStore struct {
	nc                      *nats.Conn
	bookCreatedSubscription *nats.Subscription
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
}
func (n *NatsEventStore) PublishBookCreated(book schema.Book) error {
	m := BookCreatedMessage{
		ID:        book.ID,
		Name:      book.Name,
		CreatedAt: book.CreatedAt,
	}

	data := bytes.Buffer{}
	err := gob.NewEncoder(&data).Encode(m)
	if err != nil {
		return err
	}
	return n.nc.Publish(m.Key(), data.Bytes())
}
