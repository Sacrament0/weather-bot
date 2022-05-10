package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	// ошибка неизвестная команда
	errUnknownCommand  = errors.New("unknown command")
	errUnknownMessage  = errors.New("unknown message")
	errUnableToGetData = errors.New("unable to get data from server")
)

// Обработка ошибок
func (b *Bot) HandleError(chatID int64, err error) {

	// дефолтное сообщение о неизвестной ошибке
	msg := tgbotapi.NewMessage(chatID, b.cfg.Errors.Default)

	switch err {

	case errUnknownCommand:
		msg.Text = b.cfg.Errors.UnknownCommand
		b.bot.Send(msg)
	case errUnknownMessage:
		msg.Text = b.cfg.Errors.UnknownMessage
		b.bot.Send(msg)
	case errUnableToGetData:
		msg.Text = b.cfg.Errors.UnableToGetData
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)

	}
}
