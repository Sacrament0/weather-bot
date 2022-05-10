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
	// Создание конфига
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err, "Не удалось создать конфиг")
	}

	//---------------------------------------------------------

	// Инициализация запроса погоды через OWM
	owm, err := owm.NewCurrent("C", "ru", cfg.OWMApiKey)
	if err != nil {
		log.Fatal(err, "Не удалось подключиться к сервису")
	}

	// Инициализация сервиса запроса погоды
	service := openWM.NewService(cfg, owm)

	//------------------------------------------------------------

	// Создание нового бота
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	// Флаг для вывода логов в консоли
	bot.Debug = true

	// Оборачивание бота в структуру Bot
	telegramBot := telegram.NewBot(bot, cfg, service)

	// Запускаем бота
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
