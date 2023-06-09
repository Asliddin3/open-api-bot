package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Asliddin3/open-api-bot/bot"
	"github.com/Asliddin3/open-api-bot/config"
	"github.com/Asliddin3/open-api-bot/storage"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sql.Open("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	strg := storage.NewStoragePg(psqlConn)

	botHandler := bot.New(cfg, strg)

	botHandler.Start()
}
