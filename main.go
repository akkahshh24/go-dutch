package main

import (
	"context"
	"log"

	"github.com/akkahshh24/go-dutch/token"

	"github.com/akkahshh24/go-dutch/api"
	db "github.com/akkahshh24/go-dutch/db/sqlc"
	"github.com/akkahshh24/go-dutch/util"
	"github.com/jackc/pgx/v5"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := pgx.Connect(context.Background(), config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	defer conn.Close(context.Background())

	store := db.NewStore(conn)
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatalf("cannot create token maker: %v", err)
	}
	server := api.NewServer(config, store, tokenMaker)

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
	log.Println("Server started on", config.ServerAddress)
}
