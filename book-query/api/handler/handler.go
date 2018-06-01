package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/golovers/cqrs-example/common"
	"github.com/golovers/cqrs-example/book-query/api/db"
	"github.com/golovers/cqrs-example/book-query/api/event"
	"github.com/golovers/cqrs-example/book-query/api/schema"
	"github.com/golovers/cqrs-example/book-query/api/search"
	"github.com/sirupsen/logrus"
)

func OnBookCreated(msg event.BookCreatedMessage) {
	book := schema.Book{
		ID:        msg.ID,
		Name:      msg.Name,
		CreatedAt: msg.CreatedAt,
	}
	err := search.InsertBook(context.Background(), book)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("inserted new book into elastic", book)
}

func ListBooks(w http.ResponseWriter, r *http.Request) {
	skip := uint64(0)
	take := uint64(100)
	skipV := r.FormValue("skip")
	takeV := r.FormValue("take")
	var err error
	if len(skipV) != 0 {
		skip, err = strconv.ParseUint(skipV, 10, 64)
	}
	if len(takeV) != 0 {
		take, err = strconv.ParseUint(takeV, 10, 64)
	}

	books, err := db.ListBooks(context.Background(), skip, take)
	if err != nil {
		common.ResponseError(w, http.StatusInternalServerError, "Could not list books")
		logrus.Error("Failed to list book", err)
		return
	}
	logrus.Info("books returned: ", books)
	common.ResponseOK(w, books)
}

func SearchBooks(w http.ResponseWriter, r *http.Request) {
	skip := uint64(0)
	take := uint64(100)
	skipV := r.FormValue("skip")
	takeV := r.FormValue("take")
	query := r.FormValue("query")
	query = strings.Replace(query, "\"", "", -1)

	var err error
	if len(skipV) != 0 {
		skip, err = strconv.ParseUint(skipV, 10, 64)
	}
	if len(takeV) != 0 {
		take, err = strconv.ParseUint(takeV, 10, 64)
	}

	books, err := search.SearchBooks(context.Background(), query, skip, take)
	if err != nil {
		common.ResponseOK(w, []schema.Book{})
		logrus.Error("Nothing to return", err)
		return
	}
	logrus.Info("books returned", books)
	common.ResponseOK(w, books)
}
