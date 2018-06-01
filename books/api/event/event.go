package event

import "github.com/golovers/cqrs-example/books/api/schema"

type EventStore interface {
	Close()
	PublishBookCreated(moew schema.Book) error
}

var impl EventStore

func SetEventStore(es EventStore) {
	impl = es
}

func Close() {
	impl.Close()
}

func PublishBookCreated(book schema.Book) error {
	return impl.PublishBookCreated(book)
}
