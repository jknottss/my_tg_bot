package startapp

import (
	"bot/internal/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type App struct {
	Bot    *tgbotapi.BotAPI
	Update tgbotapi.Update
	Conn   *database.Connection
}

func InitApp() *App {
	conn, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection with base established")
	bot, err := tgbotapi.NewBotAPI(getToken())
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return &App{
		Bot:    bot,
		Update: tgbotapi.Update{},
		Conn:   conn,
	}
}

func getToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv("TOKEN")
}
