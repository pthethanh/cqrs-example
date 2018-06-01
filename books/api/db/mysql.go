package db

import (
	"context"
	"database/sql"

	"github.com/golovers/cqrs-example/books/api/schema"
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

func (rp *MySQLRepository) InsertBook(ctx context.Context, book schema.Book) error {
	_, err := rp.db.Exec("insert into books(id, name, created_at) values (?,?,?)", book.ID, book.Name, book.CreatedAt)
	return err
}

func (rp *MySQLRepository) Init(ctx context.Context) error {
	_, err := rp.db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			_id varchar(255) NOT NULL,
			name varchar(255) NOT NULL,
			created_at timestamp NOT  NULL,
			PRIMARY KEY (_id));
			`)
	return err
}
