package telegram

import (
	"disp_bot/telegram/msgsPack"
	"github.com/mymmrac/telego"
)

func (b *Bot) msgsProc(update *telego.Update) {
	pack := b.findMessPack(update)
	pack.GetUpdate(update)
}

func (b *Bot) findMessPack(update *telego.Update) *msgsPack.MsgsPack {
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
	b.messPacks[chatID] = pack
	return pack
}
