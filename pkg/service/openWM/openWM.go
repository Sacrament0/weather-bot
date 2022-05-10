package openWM

import (
	"fmt"

	"github.com/Sacrament0/weather-bot/pkg/config"
	owm "github.com/briandowns/openweathermap"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Структура сервиса
type Service struct {
	cfg *config.Config
	owm *owm.CurrentWeatherData
}

// конструктор для сервиса
func NewService(cfg *config.Config, owm *owm.CurrentWeatherData) *Service {
	return &Service{cfg: cfg, owm: owm}
}

// Получение погоды из OWM
func (s *Service) GetWeather(location *tgbotapi.Location) (response string, err error) {

	err = s.owm.CurrentByCoordinates(
		&owm.Coordinates{
			Longitude: location.Longitude,
			Latitude:  location.Latitude,
		},
	)
	if err != nil {
		return "", err
	}

	// формируем ответ
	message := fmt.Sprintf("Сейчас %v °C, ощущается как %v °C",
		s.owm.Main.Temp, s.owm.Main.FeelsLike)

	return message, nil
}
