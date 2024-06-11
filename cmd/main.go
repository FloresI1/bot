package main

import (
	"log"

	"telegram-bot/config"
	"telegram-bot/handler"
	"telegram-bot/scheduler"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Инициализация базы данных и загрузка конфигурации
	db := config.InitDB("bot.db")
	defer db.Close()

	token, chatID := config.LoadTokenAndChatID(db)

	// Инициализация бота
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Запуск задачи по расписанию, если токен и Chat ID загружены
	if token != "" && chatID != "" {
		go scheduler.ScheduleTask(bot, chatID)
	}

	for update := range updates {
		if update.Message != nil {
			handler.HandleUpdate(bot, update.Message, db)
		}
	}
}
