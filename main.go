package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:secret@localhost:5000/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8000"
// )

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot loading config", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
