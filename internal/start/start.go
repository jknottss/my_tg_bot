package start

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type App struct {
	Bot    *tgbotapi.BotAPI
	Update tgbotapi.Update
	//сюда же коннекшн в БД закинуть
}

func Start() *App {
	app := &App{}
	bot, err := tgbotapi.NewBotAPI(getToken())
	if err != nil {
		log.Panic(err)
	}
	app.Bot = bot
	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return app
}

func getToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv("TOKEN")
}
