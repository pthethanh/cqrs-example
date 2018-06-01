package db

import (
	"context"
	"database/sql"

	"github.com/golovers/cqrs-example/book-query/api/schema"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepostory(url string) (*MySQLRepository, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	return &MySQLRepository{
		db: db,
	}, nil
}

func (rp *MySQLRepository) Close() {
	rp.db.Close()
}

func (rp *MySQLRepository) ListBooks(ctx context.Context, skip uint64, take uint64) ([]schema.Book, error) {
	books := make([]schema.Book, 0)
	rows, err := rp.db.Query("select * from books")
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		book := schema.Book{}
		if err := rows.Scan(&book.ID, &book.Name, &book.CreatedAt); err != nil {
			return books, err
		}
		books = append(books, book)
		if err := rows.Err(); err != nil {
			return books, err
		}
	}
	return books, nil
}
