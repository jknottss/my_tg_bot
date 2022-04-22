package commands

import (
	"bot/internal/start"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func Handler(app *start.App) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := app.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			app.Update = update
			input := strings.Split(update.Message.Text, " ")
			command := strings.ToLower(input[0])
			//строку отправлять в базу
			str := strings.Join(input[1:], " ")
			fmt.Println(str)
			switch {
			case command == "добавь":
				addTask(app)
			case command == "список":
				showAllTasks(app)
			case command == "сделал":
				taskDone(app)
			default:
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				app.Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Неверная команда"))
			}
		}
	}
}
