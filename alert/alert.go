package alert

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var bot *tgbotapi.BotAPI
var chatID int64 = 5962807794

func InitBot(token string) error {
	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	return err
}

func SendAlert(message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := bot.Send(msg)
	return err
}
