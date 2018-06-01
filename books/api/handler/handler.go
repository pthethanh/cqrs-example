package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golovers/cqrs-example/books/api/db"
	"github.com/golovers/cqrs-example/books/api/event"
	"github.com/golovers/cqrs-example/books/api/schema"
	"github.com/golovers/cqrs-example/common"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &schema.Book{}
	if err := json.NewDecoder(r.Body).Decode(book); err != nil {
		common.ResponseError(w, http.StatusBadRequest, "Invalid input")
		return
	}
	book.CreatedAt = time.Now()
	book.ID = uuid.NewV4().String()
	err := db.InsertBook(context.Background(), *book)
	if err != nil {
		common.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	event.PublishBookCreated(*book)
	logrus.Info("created & published new book: ", book)
	//common.ResponseOK(w, ResponseOK{book.ID})
}

type ResponseOK struct {
	ID string
}
