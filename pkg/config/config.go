package config

import (
	"github.com/spf13/viper"
)

// Config structure
type Config struct {
	TelegramToken string
	Errors        Errors
	OWMApiKey     string
}

// Error structure
type Errors struct {
	Default         string `mapstructure:"default"`
	UnknownCommand  string `mapstructure:"unknownCommand"`
	UnknownMessage  string `mapstructure:"unknownMessage"`
	UnableToGetData string `mapstructure:"unableToGetData"`
}

// Config init
func Init() (*Config, error) {

	// config folder path (folder name)
	viper.AddConfigPath("configs")

	// Config name in folder
	viper.SetConfigName("main")

	// Read config according to the path
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Config variable
	var cfg Config

	// Parse config to Config variable
	if err := viper.UnmarshalKey("errors", &cfg.Errors); err != nil {
		return nil, err
	}

	// Parse env var
	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Parses env var
func parseEnv(cfg *Config) error {

	// Download env from .env to system
	// case for local server --------------------------------
	// if err := godotenv.Load(); err != nil {
	// 	log.Print("No .env file found")
	// }

	// cfg.TelegramToken, _ = os.LookupEnv("TOKEN")
	// cfg.OWMApiKey, _ = os.LookupEnv("OWM_API_KEY")

	// case for remote server -------------------------------

	// parse token with key "token"
	if err := viper.BindEnv("token"); err != nil {
		return err
	}
	// parse token with key "api key"
	if err := viper.BindEnv("owm_api_key"); err != nil {
		return err
	}

	// put parsed config to Config var
	cfg.TelegramToken = viper.GetString("token")
	cfg.OWMApiKey = viper.GetString("owm_api_key")

	return nil
}
