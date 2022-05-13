package main

import (
	"bot/internal/commands"
	"bot/internal/startapp"
)

//TODO написать unit - тесты
func main() {
	app := startapp.InitApp()
	commands.Handler(app)
}
