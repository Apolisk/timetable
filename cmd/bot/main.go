package main

import (
	"log"
	"os"

	"github.com/Apolisk/bus"
	"github.com/Apolisk/bus/internal/bot"
	"github.com/Apolisk/bus/internal/database"
	"github.com/pressly/goose/v3"
)

func main() {

	db, err := database.Open(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	boot := bus.Bootstrap{
		DB: db,
	}

	if err := goose.Up(db.DB, "sql"); err != nil {
		log.Fatal(err)
	}

	b, err := bot.New("bot.yml", boot)
	if err != nil {
		log.Fatal(err)
	}

	b.Start()
}
