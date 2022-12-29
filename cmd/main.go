package main

import (
	"log"
	"os"
	"time"

	"github.com/Ksenofontovas/garantex-notify-bot/internal/UI/tgbot"
	"github.com/Ksenofontovas/garantex-notify-bot/internal/repository"
	"github.com/Ksenofontovas/garantex-notify-bot/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	botDebug := false
	updateTimeout := 60

	garantexHost := "https://garantex.io"

	db, err := repository.NewPudgeDB("./DB", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repos := repository.NewRepository(garantexHost, 5*time.Second, db)
	service := service.NewService(repos)
	bot, err := tgbot.NewTgBot(os.Getenv("BOT_API_KEY"), botDebug, updateTimeout, service)
	if err != nil {
		log.Panic(err)
	}

	bot.GetUpdates()
}
