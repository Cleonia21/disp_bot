package telegram

import (
	"disp_bot/telegram/msgsPack"
	"github.com/mymmrac/telego"
)

func (b *Bot) msgsProc(update *telego.Update) {
	packChan := b.getPackChan(update)
	packChan <- update
}

func (b *Bot) getPackChan(update *telego.Update) chan *telego.Update {
	var chatID telego.ChatID
	if update.Message != nil {
		chatID.ID = update.Message.Chat.ID
	} else if update.CallbackQuery != nil {
		chatID.ID = update.CallbackQuery.Message.Chat.ID
	}
	if pack, ok := b.messPacks[chatID]; ok {
		return pack
	}
	pack := msgsPack.Init(b.telegram, chatID)
	packChan := make(chan *telego.Update)
	b.messPacks[chatID] = packChan
	go pack.GetUpdate(packChan)
	return packChan
}
