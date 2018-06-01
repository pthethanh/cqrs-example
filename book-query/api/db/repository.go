package db

import (
	"context"

	"github.com/golovers/cqrs-example/book-query/api/schema"
)

type Repository interface {
	Close()
	ListBooks(ctx context.Context, skip uint64, take uint64) ([]schema.Book, error)
}

var impl Repository

func SetRepository(repo Repository) {
	impl = repo
}

func Close() {
	impl.Close()
}

func ListBooks(ctx context.Context, skip uint64, take uint64) ([]schema.Book, error) {
	return impl.ListBooks(ctx, skip, take)
}
