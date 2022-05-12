package main

import (
	"log"

	"github.com/Sacrament0/weather-bot/pkg/config"
	"github.com/Sacrament0/weather-bot/pkg/service/openWM"
	"github.com/Sacrament0/weather-bot/pkg/telegram"
	owm "github.com/briandowns/openweathermap"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Config creating
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err, "Не удалось создать конфиг")
	}

	// Weather request init
	owm, err := owm.NewCurrent("C", "ru", cfg.OWMApiKey)
	if err != nil {
		log.Fatal(err, "Не удалось подключиться к сервису")
	}

	// Weather service init
	service := openWM.NewService(cfg, owm)

	// Bot api creating
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	// Bot logs in console
	bot.Debug = true

	// Bot creating
	telegramBot := telegram.NewBot(bot, cfg, service)

	// Bot starting
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
