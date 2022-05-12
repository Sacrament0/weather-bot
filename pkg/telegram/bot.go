package telegram

import (
	"log"

	"github.com/Sacrament0/weather-bot/pkg/config"
	"github.com/Sacrament0/weather-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot structure
type Bot struct {
	bot     *tgbotapi.BotAPI
	cfg     *config.Config
	service service.Servicer
}

// Bot constructor
func NewBot(bot *tgbotapi.BotAPI, cfg *config.Config, service service.Servicer) *Bot {
	return &Bot{bot: bot, cfg: cfg, service: service}
}

// Creates channel for messages and handle them
func (b *Bot) Start() error {

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	// Channel for messages
	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	// Handling messages
	b.handleUpdates(updates)

	return nil

}

// Creates channel for messages
func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {

	// Configuring channel
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

// Getting messages from users and sending messages to users
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {

	// Loop on channel
	for update := range updates {
		// if message from user is empty, skip iteration
		if update.Message == nil {
			continue
		}
		// check if message is command
		if update.Message.IsCommand() {

			// Handle command
			if err := b.handleCommand(update.Message); err != nil {
				// Handle error
				b.HandleError(update.Message.Chat.ID, err)
			}
			continue
		}
		// Handle message
		if err := b.handleMessage(update.Message); err != nil {
			// Handle error
			b.HandleError(update.Message.Chat.ID, err)
		}

	}

}
