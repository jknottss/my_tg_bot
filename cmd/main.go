package main

import (
	"bot/internal/commands"
	"bot/internal/getconfig"
)

//TODO написать unit - тесты
func main() {
	app := getconfig.Start()
	commands.Handler(app)
}
