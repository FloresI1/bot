package handler

import (
	"database/sql"
	"strings"

	"telegram-bot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleUpdate обрабатывает входящие сообщения
func HandleUpdate(bot *tgbotapi.BotAPI, message *tgbotapi.Message, db *sql.DB) {
	if strings.HasPrefix(message.Text, "/token ") {
		token := strings.TrimSpace(strings.TrimPrefix(message.Text, "/token "))
		config.SaveConfig(db, "token", token)
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Token сохранен."))
	}

	if strings.HasPrefix(message.Text, "/chatid ") {
		chatID := strings.TrimSpace(strings.TrimPrefix(message.Text, "/chatid "))
		config.SaveConfig(db, "chat_id", chatID)
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Chat ID сохранен."))
	}
}
