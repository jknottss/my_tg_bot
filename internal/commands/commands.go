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
	count, err := app.Conn.AddTask(app.Update.Message.From.UserName, task)
	if err != nil {
		log.Print(err)
		return
	}
	app.Bot.Send(tgbotapi.NewMessage(app.Update.Message.Chat.ID, "Задача добавлена! Ее номер: "+strconv.Itoa(count)))
}

func showAllTasks(app *startapp.App) {
	var modArr []database.Model
	modArr, err := app.Conn.GetAllTasks(app.Update.Message.From.UserName)
	res := BuildTasks(modArr)
	fmt.Println(res)
	if err != nil {
		log.Println(err)
		return
	}
	msg := &tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: app.Update.Message.Chat.ID,
		},
		Text:                  res,
		ParseMode:             "html",
		Entities:              nil,
		DisableWebPagePreview: false,
	}
	app.Bot.Send(msg)
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

func BuildTasks(tasks []database.Model) (res string) {
	done := "<ins>Выполненные:</ins>\n"
	actual := "<ins>Актуальные:</ins>\n"
	for _, val := range tasks {
		if val.Done {
			done += fmt.Sprintf("<i>%d. %s</i> \u2705\n", val.Priority, val.Task)
		} else {
			actual += fmt.Sprintf("<i>%d. %s</i>\n", val.Priority, val.Task)
		}
		res = done + "\n" + actual
	}
	return
}
