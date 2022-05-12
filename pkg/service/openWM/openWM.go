package openWM

import (
	"fmt"

	"github.com/Sacrament0/weather-bot/pkg/config"
	owm "github.com/briandowns/openweathermap"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Weather service structure
type Service struct {
	cfg *config.Config
	owm *owm.CurrentWeatherData
}

// Service constructor
func NewService(cfg *config.Config, owm *owm.CurrentWeatherData) *Service {
	return &Service{cfg: cfg, owm: owm}
}

// Requests weather servise and forms weather message
func (s *Service) GetWeather(location *tgbotapi.Location) (response string, err error) {

	// Weather service request
	err = s.owm.CurrentByCoordinates(
		&owm.Coordinates{
			Longitude: location.Longitude,
			Latitude:  location.Latitude,
		},
	)
	if err != nil {
		return "", err
	}

	// Weather message forming
	message := fmt.Sprintf("Сейчас %v °C, ощущается как %v °C",
		s.owm.Main.Temp, s.owm.Main.FeelsLike)

	return message, nil
}
