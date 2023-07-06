package main

import "disp_bot/telegram"

func main() {
	bot := telegram.Init()
	bot.Start()
}
