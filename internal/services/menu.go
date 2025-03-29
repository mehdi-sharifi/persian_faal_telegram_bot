package services

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func SendMenu(bot *tgbotapi.BotAPI, chatID int) {
	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("فال جدید"),
			tgbotapi.NewKeyboardButton("درباره ما"),
		),
	)
	msg := tgbotapi.NewMessage(int64(chatID), "یک کزینه را انتخواب کنید")
	msg.ReplyMarkup = buttons
	bot.Send(msg)
}
