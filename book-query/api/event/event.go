package event

type EventStore interface {
	Close()
	SubsribeBookCreated() (<-chan BookCreatedMessage, error)
	OnBookCreated(f func(BookCreatedMessage)) error
}

var impl EventStore

func SetEventStore(es EventStore) {
	impl = es
}

func Close() {
	impl.Close()
}

func SubsribeBookCreated() (<-chan BookCreatedMessage, error) {
	return impl.SubsribeBookCreated()
}

func OnBookCreated(f func(BookCreatedMessage)) error {
	return impl.OnBookCreated(f)
}
