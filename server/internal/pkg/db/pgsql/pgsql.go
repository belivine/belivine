package pgsql

import (
	"context"

	"github.com/go-pg/pg/v10"
)

var Db *pg.DB

func InitDB() {
	db := pg.Connect(&pg.Options{
		Addr:     "ine_pgsql:5432",
		User:     "root",
		Password: "root",
		Database: "ine",
	})

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
	Db = db
}
