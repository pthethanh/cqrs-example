package db

import (
	"context"

	"github.com/golovers/cqrs-example/books/api/schema"
)

type Repository interface {
	Close()
	Init(ctx context.Context) error
	InsertBook(ctx context.Context, book schema.Book) error
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

func Init(ctx context.Context) error {
	return impl.Init(ctx)
}
