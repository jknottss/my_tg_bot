package main

import (
	"bot/internal/commands"
	"bot/internal/startapp"
	"time"
)

//TODO написать unit - тесты
func main() {
	time.Sleep(3 * time.Second)
	app := startapp.InitApp()
	commands.Handler(app)
}
