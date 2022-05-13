package getconfig

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type App struct {
	Bot    *tgbotapi.BotAPI
	Update tgbotapi.Update
}

func Start() *App {
	app := &App{}
	bot, err := tgbotapi.NewBotAPI(getToken())
	if err != nil {
		log.Panic(err)
	}
	app.Bot = bot
	bot.Debug = false
	if err != nil {
		log.Fatal(err)
	}
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

func GetConfig() (*pgx.ConnPoolConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	PsqlHost := os.Getenv("PSQL_HOST")
	PsqlPort := os.Getenv("PSQL_PORT")
	PsqlUser := os.Getenv("PSQL_USER")
	PsqlPassword := os.Getenv("PSQL_PASSWORD")
	PsqlDb := os.Getenv("PSQL_DB")
	Port, _ := strconv.ParseUint(PsqlPort, 10, 64)
	conf := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     PsqlHost,
			Port:     uint16(Port),
			Password: PsqlPassword,
			Database: PsqlDb,
			User:     PsqlUser,
		}}
	return &conf, nil
}
