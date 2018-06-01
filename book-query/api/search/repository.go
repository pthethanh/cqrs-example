package search

import (
	"context"

	"github.com/golovers/cqrs-example/book-query/api/schema"
)

type Repository interface {
	Close()
	InsertBook(ctx context.Context, book schema.Book) error
	SearchBooks(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Book, error)
}

var impl Repository

func SetRepository(repo Repository) {
	impl = repo
}

func Close() {
	impl.Close()
}

func InsertBook(ctx context.Context, book schema.Book) error {
	return impl.InsertBook(ctx, book)
}

func SearchBooks(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Book, error) {
	return impl.SearchBooks(ctx, query, skip, take)
}
