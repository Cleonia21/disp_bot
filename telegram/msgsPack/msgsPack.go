package msgsPack

import (
	"disp_bot/utils"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	idDefault   = "нажми на кнопку -> присылай сообщения"
	id47        = "47"
	idFlower    = "цветок"
	id1C        = "1C"
	idStretches = "перегоны"
	idMail      = "mail"
	idSend      = "проверить"
)

type MsgsPack struct {
	telegram *telego.Bot
	chatID   telego.ChatID
	messages map[string][]utils.Message
	conf     utils.Conf
	status   string

	buttons              map[string]telego.InlineKeyboardButton
	inlineText           string
	oldInlineMsg         *telego.Message
	oldCallbackQueryData string
}

func Init(telegram *telego.Bot, chatID telego.ChatID) *MsgsPack {
	pack := MsgsPack{}
	pack.chatID = chatID
	pack.status = idDefault
	pack.telegram = telegram
	pack.messages = make(map[string][]utils.Message, 100)

	pack.buttons = map[string]telego.InlineKeyboardButton{
		id47: tu.InlineKeyboardButton(id47).
			WithCallbackData(id47),
		idFlower: tu.InlineKeyboardButton(idFlower).
			WithCallbackData(idFlower),
		id1C: tu.InlineKeyboardButton(id1C).
			WithCallbackData(id1C),
		idStretches: tu.InlineKeyboardButton(idStretches).
			WithCallbackData(idStretches),
		idSend: tu.InlineKeyboardButton(idSend).
			WithCallbackData(idSend),
	}
	pack.inlineText = "нажми на кнопку\n -> \nприсылай сообщения"
	return &pack
}

func (p *MsgsPack) GetUpdate(update *telego.Update) {
	p.deleteInlineMenu()
	if update.CallbackQuery != nil {
		p.handleCallbackQuery(update)
	} else if update.Message != nil {
		p.saveMsg(update.Message)
	}
	p.sendInlineMenu()
}

func (p *MsgsPack) deleteInlineMenu() {
	if p.oldInlineMsg != nil {
		err := p.telegram.DeleteMessage(tu.Delete(p.chatID, p.oldInlineMsg.MessageID))
		if err != nil {
			p.telegram.Logger().Errorf(err.Error())
		}
	}
}

func (p *MsgsPack) sendInlineMenu() {
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			p.buttons[idFlower],
			p.buttons[id1C],
		),
		tu.InlineKeyboardRow(
			p.buttons[idStretches],
			p.buttons[id47],
		),
		tu.InlineKeyboardRow(
			p.buttons[idSend],
		),
	)
	message := tu.Message(
		p.chatID,
		p.inlineText,
	).WithReplyMarkup(inlineKeyboard)
	msg, err := p.telegram.SendMessage(message)
	if err != nil {
		p.telegram.Logger().Errorf(err.Error())
	}
	p.oldInlineMsg = msg
}
