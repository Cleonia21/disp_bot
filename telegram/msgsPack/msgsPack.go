package msgsPack

import (
	"disp_bot/utils"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"time"
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

func (p *MsgsPack) GetUpdate(update <-chan *telego.Update) {
	sendMenuAfterSecond := make(chan int)
	sendMenuNow := make(chan int)
	go p.reSendMenu(sendMenuAfterSecond, sendMenuNow)
	for u := range update {
		if u.CallbackQuery != nil {
			p.handleCallbackQuery(u)
			sendMenuNow <- 0
		} else if u.Message != nil {
			p.saveMsg(u.Message)
			sendMenuAfterSecond <- 0
		}
	}
}

func (p *MsgsPack) reSendMenu(renewTimer <-chan int, doNow <-chan int) {
	var t time.Time
	var dellFlag bool
	for {
		select {
		case <-renewTimer:
			t = time.Now()
			dellFlag = true
		case <-doNow:
			dellFlag = false
			p.deleteInlineMenu()
			p.sendInlineMenu()
		default:
			if dellFlag &&
				time.Since(t) > time.Since(time.Now())+time.Second {
				p.deleteInlineMenu()
				p.sendInlineMenu()
				dellFlag = false
			}
		}
	}
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
