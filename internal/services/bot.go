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
		if update.Message != nil {
			chatID := update.Message.Chat.ID
			text := update.Message.Text
			switch text {
			case "/start":
				SendMenu(bot, int(chatID))
			case "فال جدید":
				poem := GetRandomPoem(poems)
				response := fmt.Sprintf("📜 *فال:*\n\n_%s_\n\n💡 *تفصیر:* %s", poem.Text, poem.Interpretation)
				msg := tgbotapi.NewMessage(chatID, response)
				msg.ParseMode = tgbotapi.ModeMarkdown
				bot.Send(msg)
			case "درباره ما":
				msg := tgbotapi.NewMessage(chatID, "about")
				msg.ParseMode = tgbotapi.ModeMarkdown
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(chatID, "I didn't understand that. Please use the menu below.")
				bot.Send(msg)
				SendMenu(bot, int(chatID))

			}
		} else if update.CallbackQuery != nil {
			chatID := update.CallbackQuery.Message.Chat.ID
			data := update.CallbackQuery.Data
			if data == "get_poem" {
				poem := GetRandomPoem(poems)
				response := fmt.Sprintf("📜 *فال:*\n\n_%s_\n\n💡 *تفسیر:* %s", poem.Text, poem.Interpretation)
				msg := tgbotapi.NewMessage(chatID, response)
				msg.ParseMode = "Markdown"
				bot.Send(msg)
			} else if data == "about" {
				msg := tgbotapi.NewMessage(chatID, "about")
				msg.ParseMode = tgbotapi.ModeMarkdown
				bot.Send(msg)
			}
		}
	}
}
