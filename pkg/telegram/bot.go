package telegram

import (
	"log"

	"github.com/Sacrament0/weather-bot/pkg/config"
	"github.com/Sacrament0/weather-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Структура описывающая бота
type Bot struct {
	bot *tgbotapi.BotAPI
	cfg *config.Config
	service service.Servicer
}

// Конструктор для бота. Создает переменную структуру типа Bot и помещает туда созданного бота типа *tgbotapi.BotAPI
// Почему принимает ссылку - Чтобы не копировать всего бота в метод и не выделять память
// а также чтобы была возможность изменять данные
func NewBot(bot *tgbotapi.BotAPI, cfg *config.Config, service service.Servicer) *Bot {
	return &Bot{bot: bot, cfg: cfg, service: service}
}

// Создает канал для отправки сообщений и обрабатывает входящие сообщения
func (b *Bot) Start() error {

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	// КАНАЛ ДЛЯ ПОЛУЧЕНИЯ СООБЩЕНИЙ
	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	// ОБРАБОТКА СООБЩЕНИЙ
	b.handleUpdates(updates)

	return nil

}

// initUpdatesChannel инициализирует канал для получения-передачи сообщений
func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {

	// Создание конфигурации для получения обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Возвращает канал для получения значеий от API
	return b.bot.GetUpdatesChan(u)
}

// handleUpdates получает и отправляет сообщения пользователю
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {

	// Итерация по каналу
	for update := range updates {
		// если сообщение пустое, скипаем цикл
		if update.Message == nil {
			continue
		}
		// проверка является ли сообщение командой
		if update.Message.IsCommand() {

			//ОБРАБОТКА КОМАНДЫ -------------------------------------------
			if err := b.handleCommand(update.Message); err != nil {
				// обработка ошибки
				b.HandleError(update.Message.Chat.ID, err)
			}
			continue
		}
		// ОБРАБОТКА СООБЩЕНИЯ ---------------------------------------------
		if err := b.handleMessage(update.Message); err != nil {
			// обработка ошибки
			b.HandleError(update.Message.Chat.ID, err)
		}

	}

}
