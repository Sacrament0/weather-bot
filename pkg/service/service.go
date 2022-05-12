package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Weather service interface
type Servicer interface {
	GetWeather(location *tgbotapi.Location) (response string, err error)
}
