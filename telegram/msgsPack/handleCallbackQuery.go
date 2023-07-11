package msgsPack

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (p *MsgsPack) handleCallbackQuery(update *telego.Update) {
	_, ok := p.messages[p.status]
	if ok && p.oldCallbackQueryData != "" {
		p.modConf()
		p.modButtons()
	}
	p.modStatus(update.CallbackQuery.Data)
	p.modInlineReplyText()
	if p.processingMsgs() {
		return
	}
	p.clearMsgs()
	p.oldCallbackQueryData = update.CallbackQuery.Data
	err := p.telegram.AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{CallbackQueryID: update.CallbackQuery.ID})
	if err != nil {
		p.telegram.Logger().Errorf(err.Error())
	}
}

func (p *MsgsPack) modConf() {
	switch p.oldCallbackQueryData {
	case id1C:
		p.conf.OneC = true
	case idFlower:
		p.conf.ChatFlower = true
	case id47:
		p.conf.Chat47 = true
	case idStretches:
		p.conf.ChatStretches = true
	}
}

func (p *MsgsPack) modButtons() {
	switch p.oldCallbackQueryData {
	case id1C:
		p.buttons[id1C] = tu.InlineKeyboardButton(id1C + "✅").
			WithCallbackData(id1C)
	case idFlower:
		p.buttons[idFlower] = tu.InlineKeyboardButton(idFlower + "✅").
			WithCallbackData(idFlower)
	case id47:
		p.buttons[id47] = tu.InlineKeyboardButton(id47 + "✅").
			WithCallbackData(id47)
	case idStretches:
		p.buttons[idStretches] = tu.InlineKeyboardButton(idStretches + "✅").
			WithCallbackData(idStretches)
	}
}

func (p *MsgsPack) modStatus(data string) {
	p.status = data
}

func (p *MsgsPack) clearMsgs() {
	delete(p.messages, p.status)
}

func (p *MsgsPack) modInlineReplyText() {
	p.inlineText = "жду " + p.status + "..."
}
