package telegram

import (
	"log"

	"github.com/Sacrament0/weather-bot/pkg/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "start"
	helloMessage = "Привет, чтобы узнать погоду, отправь мне своё местоположение. Для этого в твоей клавиатуре есть кнопка Send Location"
)

// Метод для обработки КОМАНД ---------------------------------------------------
func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	// получения значения команды
	switch message.Command() {

	case commandStart:

		// обработка команды старт
		return b.handleStartCommand(message)

	default:

		// обработка неизвестной команды
		return errUnknownCommand

	}

}

// Метод для обработки СООБЩЕНИЙ -------------------------------------------------
func (b *Bot) handleMessage(message *tgbotapi.Message) error {

	// Если в сообщении есть указание локации
	if message.Location != nil {

		// Создание структуры для ответоного сообщения с указанием:
		// Chat ID куда отправляется сообщение и текст сообщения
		msg := tgbotapi.NewMessage(message.Chat.ID, helloMessage)

		// формируем ответ
		// Функция принимает интерфейс
		response, err := service.CreateWeatherMessage(b.service, message.Location)
		if err != nil {
			b.HandleError(message.Chat.ID, err)
		}

		msg.Text = response

		// Отправка сообщения
		b.bot.Send(msg)

		return nil

	} else {

		return errUnknownMessage

	}
}

// Метод для обработки стартовой команды
func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {

	// приветственное сообщение
	msg := tgbotapi.NewMessage(message.Chat.ID, helloMessage)

	// создам кнопку с запросом геолокации
	locationButton := tgbotapi.NewKeyboardButtonLocation("Send Location")

	// помещаем кнопку в сообщение
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{locationButton})

	// Отправка сообщения
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Println("send hello message failed")
		return err
	}

	return nil
}
