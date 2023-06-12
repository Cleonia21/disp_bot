package telegram

import (
	proc "disp_bot/processing"
	"errors"
	"fmt"
	"github.com/mymmrac/telego"
	"os"
)

type Bot struct {
	telegram  *telego.Bot
	messPacks map[int64]*messagesPack

	tgGetChan  <-chan telego.Update
	tgSendChan chan<- *telego.SendMessageParams

	procGetChan  chan<- proc.UnProcData
	procSendChan <-chan proc.ProcData
}

func Init(procGetChan chan<- proc.UnProcData, procSendChan <-chan proc.ProcData) *Bot {
	b := &Bot{}
	botToken := ""
	b.messPacks = make(map[int64]*messagesPack)

	// Create Bot with debug on
	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	err := errors.New("")
	b.telegram, err = telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get tgGetChan channel
	b.tgGetChan, _ = b.telegram.UpdatesViaLongPolling(nil)

	// Stop reviving tgGetChan from update channel
	defer b.telegram.StopLongPolling()

	b.procGetChan = procGetChan
	b.procSendChan = procSendChan
	return b
}

func (b *Bot) Start() {
	go b.tgHandler()
	//go b.procHandler()
}

//func (b *Bot) procHandler() {
//	for {
//		pack, ok := <-b.responseChan
//		if ok {
//			break
//		}
//		b.forwardPack(pack)
//	}
//}
//
//func (b *Bot) forwardPack(pack ResponsePack) {
//	reply := tu.Message(tu.ID(pack.chatID), pack.reply)
//	reply.WithReplyToMessageID(pack.message.id)
//	_, _ = b.telegram.SendMessage(reply)
//}
