package telegram

import (
	"disp_bot/utils"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (b *Bot) msgsSend(chatID int64, msgs []utils.Message) {
	replyText := b.mergingDuplicates(msgs)
	for id, text := range replyText {
		msg := tu.Message(tu.ID(chatID), text).
			WithReplyToMessageID(id)
		_, _ = b.telegram.SendMessage(msg)
	}
}

func (b *Bot) mergingDuplicates(msgs []utils.Message) (uniqueMsgs map[int]string) {
	uniqueMsgs = make(map[int]string)
	for _, msg := range msgs {
		uniqueMsgs[msg.ID] += msg.ReplyText + "\n"
	}
	return
}
