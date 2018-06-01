package main

import (
	"context"
	"fmt"

	"github.com/golovers/cqrs-example/books/api/db"
	"github.com/golovers/cqrs-example/books/api/event"
	"github.com/golovers/cqrs-example/books/api/handler"
	"github.com/golovers/cqrs-example/common"
	"github.com/gorilla/mux"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Cfg struct {
	HTTPAddress   string `envconfig:"HTTP_ADDRESS"`
	MySQLAddress  string `envconfig:"MYSQL_ADDRESS"`
	MySQLDB       string `envconfig:"MYSQL_DB"`
	MySQLUser     string `envconfig:"MYSQL_USER"`
	MySQLPassword string `envconfig:"MYSQL_PASSWORD"`

	NatsAddress string `envconfig:"NATS_ADDRESS"`
}

func main() {
	var cfg Cfg
	common.LoadEnvConfig(&cfg)
	repo, err := db.NewMySQLRepostory(fmt.Sprintf("%v:%v@tcp(%v)/%v?parseTime=true", cfg.MySQLUser, cfg.MySQLPassword, cfg.MySQLAddress, cfg.MySQLDB))
	if err != nil {
		panic(err)
	}
	db.SetRepository(repo)
	defer db.Close()
	es, err := event.NewNats(fmt.Sprintf("nats://%v", cfg.NatsAddress))
	if err != nil {
		panic(err)
	}
	event.SetEventStore(es)
	defer event.Close()

	db.Init(context.Background())

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/books", handler.CreateBook).Methods("POST")

	if err := http.ListenAndServe(cfg.HTTPAddress, r); err != nil {
		panic(err)
	}
}
