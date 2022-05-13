package commands

import (
	"bot/internal/startapp"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

const add = "добавь"
const list = "список"
const done = "сделал"

func Handler(app *startapp.App) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := app.Bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil { // If we got a message
			app.Update = update
			doCommand(app)
		}
	}
}

func doCommand(app *startapp.App) {
	input := strings.Split(app.Update.Message.Text, " ")
	command := strings.ToLower(input[0])
	if command == add || command == done {
		if len(input) < 2 {
			app.Bot.Send(tgbotapi.NewMessage(app.Update.Message.Chat.ID, "Укажите задачу!"))
			return
		}
	}
	//строку отправлять в базу
	str := strings.Join(input[1:], " ")
	fmt.Println(str)
	switch {
	case command == add:
		addTask(app)
	case command == list:
		showAllTasks(app)
	case command == done:
		taskDone(app)
	default:
		log.Printf("[%s] %s", app.Update.Message.From.UserName, app.Update.Message.Text)
		app.Bot.Send(tgbotapi.NewMessage(app.Update.Message.Chat.ID, "Неверная команда"))
	}
}
