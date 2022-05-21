package commands

import (
	"bot/internal/database"
	"bot/internal/startapp"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func addTask(app *startapp.App, task string) {
	app.Conn.AddTask(app.Update.Message.From.UserName, task)
	app.Bot.Send(tgbotapi.NewMessage(app.Update.Message.Chat.ID, "Задача добавлена!"))
}

//todo отдельно выполненные, отдельно невыполненные
func showAllTasks(app *startapp.App) {
	var modArr []database.Model
	modArr, err := app.Conn.GetAllTasks(app.Update.Message.From.UserName)
	res, _ := database.BuildAllTasks(modArr)
	fmt.Println(res)
	if err != nil {
		log.Println(err)
	}
	app.Bot.Send(tgbotapi.NewMessage(app.Update.Message.Chat.ID, "тут будут задачи"))
}

func doneTask(app *startapp.App, task string) {
	taskNbr, err := strconv.Atoi(task)
	if err != nil {
		app.Bot.Send(tgbotapi.NewMessage(app.Update.Message.Chat.ID, "Неверно указан номер команды"))
		return
	}
	err = app.Conn.DoneTask(app.Update.Message.From.UserName, taskNbr)
	if err != nil {
		log.Println(err)
		return
	}
	app.Bot.Send(tgbotapi.NewMessage(app.Update.Message.Chat.ID, "Задача выполнена!"))
}
