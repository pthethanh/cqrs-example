package event

import (
	"time"
)

type Message interface {
	Key() string
}

type BookCreatedMessage struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

func (m *BookCreatedMessage) Key() string {
	return "book.created"
}
