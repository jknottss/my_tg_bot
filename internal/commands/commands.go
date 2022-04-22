package commands

import (
	"bot/internal/start"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func addTask(app *start.App) {
	app.Bot.Send(tgbotapi.NewMessage(app.Update.Message.Chat.ID, "Задача добавлена!"))
	showAllTasks(app)

}

func showAllTasks(app *start.App) {
	app.Bot.Send(tgbotapi.NewMessage(app.Update.Message.Chat.ID, "Список задач:"))
}

func taskDone(app *start.App) {
	app.Bot.Send(tgbotapi.NewMessage(app.Update.Message.Chat.ID, "Задача выполнена!"))
}
