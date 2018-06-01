package search

import (
	"context"
	"encoding/json"
	"github.com/golovers/cqrs-example/book-query/api/schema"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

type ElasticRepository struct {
	client *elastic.Client
}

func NewElastic(url string) (*ElasticRepository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return &ElasticRepository{client: client}, nil
}

func (e *ElasticRepository) Close() {
	// nothing to close...
}

func (e *ElasticRepository) InsertBook(ctx context.Context, book schema.Book) error {
	_, err := e.client.Index().Index("books").Type("book").Id(book.ID).BodyJson(book).Refresh("wait_for").Do(ctx)
	return err
}

func (e *ElasticRepository) SearchBooks(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Book, error) {
	rs, err := e.client.Search().Index("books").Query(elastic.NewMultiMatchQuery(query, "name").Fuzziness("3").PrefixLength(1)).From(int(skip)).Size(int(take)).Do(ctx)
	if err != nil {
		logrus.Error(err)
		return []schema.Book{}, err
	}
	books := make([]schema.Book, 0)
	for _, hit := range rs.Hits.Hits {
		var book schema.Book
		if err := json.Unmarshal(*hit.Source, &book); err != nil {
			logrus.Error(err)
		}
		books = append(books, book)
	}
	return books, nil
}
