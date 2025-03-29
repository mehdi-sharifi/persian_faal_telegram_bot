package services

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"persian_faal_bot/internal/models"
)

func StartBot(token string, poems []models.Poem) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Bot started: @%s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		chatID := update.Message.Chat.ID
		text := update.Message.Text

		if text == "/start" {
			msg := tgbotapi.NewMessage(chatID, "به ربات فال حافط خوش آمدید")
			bot.Send(msg)
		} else if text == "/faal" {
			poem := GetRandomPoem(poems)
			response := fmt.Sprintf("📜 *فال:*\n\n_%s_\n\n💡 *تفصیر:* %s", poem.Text, poem.Interpretation)
			msg := tgbotapi.NewMessage(chatID, response)
			msg.ParseMode = tgbotapi.ModeMarkdown
			bot.Send(msg)

		} else {
			msg := tgbotapi.NewMessage(chatID, "لطفا برای فال گرفتن /faal را بزنید")
			msg.ParseMode = tgbotapi.ModeMarkdown
			bot.Send(msg)
		}
	}
}
