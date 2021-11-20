package main

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/schetininl/availability_parser/configs"
	"github.com/schetininl/availability_parser/internal/app/parser"
	"github.com/schetininl/availability_parser/internal/pkg/bot"
)

func main() {
	log.Info("Starting ...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &configs.Config{}
	if errParse := env.Parse(cfg); errParse != nil {
		log.Panic(errParse)
	}

	bot, err := bot.New(cfg.TgBotToken, cfg.TgChatID)
	if err != nil {
		log.Fatal(err)
	}

	service := parser.New(
		parser.WithBot(bot),
		parser.WithTarget(cfg.TargetURL, cfg.TargetText),
	)

	service.StartWorker(cfg.MinDuration, cfg.MaxDuration)
}
