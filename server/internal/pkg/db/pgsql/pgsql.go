package pgsql

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

var Db *pgx.Conn

func InitDB() {
	uri := "postgres://root:root@ine_pgsql:5432/ine"
	conn, err := pgx.Connect(context.Background(), uri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer conn.Close(context.Background())
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Hore Connect Alhamdulilah")

	_, err = conn.Exec(context.Background(), `set search_path='belivine'`)
	if err != nil {
		log.Fatal(err)
	}
	Db = conn
}
