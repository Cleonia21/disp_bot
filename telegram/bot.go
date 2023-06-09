package telegram

import (
	"disp_bot/conf"
	"disp_bot/processing"
	"errors"
	"fmt"
	"github.com/mymmrac/telego"
	"os"
)

type Bot struct {
	telegram  *telego.Bot
	messPacks map[telego.ChatID]chan *telego.Update

	proc *processing.Proc

	tgGetChan  <-chan telego.Update
	tgSendChan chan<- *telego.SendMessageParams
}

func Init() *Bot {
	b := &Bot{} //
	botToken := conf.TOKEN
	b.messPacks = make(map[telego.ChatID]chan *telego.Update)

	// Create Bot with debug on
	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	err := errors.New("")
	b.telegram, err = telego.NewBot(botToken) //, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b.proc = processing.Init(true)
	return b
}

func (b *Bot) Start() {
	// Get tgGetChan channel
	b.tgGetChan, _ = b.telegram.UpdatesViaLongPolling(nil)

	// Stop reviving tgGetChan from update channel
	defer b.telegram.StopLongPolling()

	// Loop through all tgGetChan when they came
	for update := range b.tgGetChan {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}
		//update.Message.From.ID
		b.msgsProc(&update)
	}
}
