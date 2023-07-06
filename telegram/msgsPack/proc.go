package msgsPack

import (
	"disp_bot/processing"
	"disp_bot/utils"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (p *MsgsPack) processingMsgs() (sending bool) {
	if p.status == idSend {
		proc := processing.Init(true)
		var buf []utils.Message
		for _, msgs := range p.messages {
			buf = append(buf, msgs...)
		}
		respMsgs := proc.Processing(buf, p.conf)
		p.msgsSend(respMsgs)
		p.oldCallbackQueryData = ""
		p.status = idDefault
		return true
	}
	return false
}

func (p *MsgsPack) msgsSend(msgs []utils.Message) {
	replyText := p.mergingDuplicates(msgs)
	for id, text := range replyText {
		msg := tu.Message(p.chatID, text).
			WithReplyToMessageID(id)
		_, _ = p.telegram.SendMessage(msg)
	}
}

func (p *MsgsPack) mergingDuplicates(msgs []utils.Message) (uniqueMsgs map[int]string) {
	uniqueMsgs = make(map[int]string)
	for _, msg := range msgs {
		uniqueMsgs[msg.ID] += msg.ReplyText + "\n"
	}
	return
}
