package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Интерфейс сервиса погоды
type Servicer interface {
	GetWeather(location *tgbotapi.Location) (response string, err error)
}

// Функция получения показаний погоды независимо от сервиса
func CreateWeatherMessage(s Servicer, location *tgbotapi.Location) (msg string, err error) {

	msg, err = s.GetWeather(location)
	if err != nil {
		return "", err
	}
	
	return msg, nil
}
