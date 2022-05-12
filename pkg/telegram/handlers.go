package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "start"
	helloMessage = "Привет, чтобы узнать погоду, отправь мне своё местоположение. Для этого в твоей клавиатуре есть кнопка Send Location"
)

// Handle command
func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	// check what the command is
	switch message.Command() {

	case commandStart:

		// Start command handling
		return b.handleStartCommand(message)

	default:

		// Unknown command handling
		return errUnknownCommand

	}

}

// Handle message
func (b *Bot) handleMessage(message *tgbotapi.Message) error {

	// If there is location in message
	if message.Location != nil {

		// Create response message
		msg := tgbotapi.NewMessage(message.Chat.ID, helloMessage)

		// Request service for weather
		response, err := b.service.GetWeather(message.Location)
		if err != nil {
			b.HandleError(message.Chat.ID, errUnableToGetData)
		}

		msg.Text = response

		// Send message
		b.bot.Send(msg)

		return nil

	} else {

		return errUnknownMessage

	}
}

// Handle start command
func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {

	// Create greeting message
	msg := tgbotapi.NewMessage(message.Chat.ID, helloMessage)

	// Make button for location request
	locationButton := tgbotapi.NewKeyboardButtonLocation("Send Location")

	// Put button to message
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{locationButton})

	// Send message
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
