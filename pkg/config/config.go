package config

import (
	"github.com/spf13/viper"
)

// Структура конфига
type Config struct {
	TelegramToken string
	Errors        Errors
	OWMApiKey     string
}

// Структура для хранения ошибок
type Errors struct {
	Default         string `mapstructure:"default"`
	UnknownCommand  string `mapstructure:"unknownCommand"`
	UnknownMessage  string `mapstructure:"unknownMessage"`
	UnableToGetData string `mapstructure:"unableToGetData"`
}

// Инициализация конфига
func Init() (*Config, error) {
	// Путь к директории конфига (имя папки)
	viper.AddConfigPath("configs")
	// Имя конфига в папке
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// переменная для хранения конфига
	var cfg Config

	// парсим конфиг из файла
	if err := viper.UnmarshalKey("errors", &cfg.Errors); err != nil {
		return nil, err
	}

	// парсим переменные окружения
	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Парсит переменные окружения
func parseEnv(cfg *Config) error {

	// Загружаем переменные из .env в систему
	// локальный вариант --------------------------------
	// if err := godotenv.Load(); err != nil {
	// 	log.Print("No .env file found")
	// }

	// cfg.TelegramToken, _ = os.LookupEnv("TOKEN")
	// cfg.OWMApiKey, _ = os.LookupEnv("OWM_API_KEY")

	// серверный вариант ------------------------------
	// парсим токен и даём ему ключ "token"
	if err := viper.BindEnv("token"); err != nil {
		return err
	}

	if err := viper.BindEnv("owm_api_key"); err != nil {
		return err
	}

	// загружаем в кофиг токен по ключу
	cfg.TelegramToken = viper.GetString("token")
	cfg.OWMApiKey = viper.GetString("owm_api_key")

	return nil
}
