package scheduler

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const startButton = "🔲 Отправиться в лес"

// ScheduleTask выполняет задачу по сбору дерева каждые 10 минут
func ScheduleTask(bot *tgbotapi.BotAPI, chatID string) {
	for {
		select {
		case <-time.After(10 * time.Minute):
			collectWood(bot, chatID)
		}
	}
}

func collectWood(bot *tgbotapi.BotAPI, chatID string) {
	msg := tgbotapi.NewMessageToChannel(chatID, startButton)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Ошибка при отправке команды на рубку дерева:", err)
	}
}
