package main

import (
	"bot/internal/commands"
	"bot/internal/start"
)

//TODO написать unit - тесты
func main() {
	app := start.Start()
	commands.Handler(app)
}
