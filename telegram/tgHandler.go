package telegram

import (
	proc "disp_bot/processing"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// inputIDs
const (
	ID_start = iota
	ID_count
)

type Message struct {
	id        int
	text      string
	replyText string
}

func NewMessage(text string) Message {
	return Message{text: text}
}

func (m *Message) Text() string {
	return m.text
}

func (m *Message) AddReply(text string) {
	m.replyText += text
}

type messagesPack struct {
	chatID     int64
	procStatus int64
	messages   map[int64][]Message
}

func (mp *messagesPack) write(mess *telego.Message) {
	messages := mp.messages[mp.procStatus]
	messages = append(messages, Message{id: mess.MessageID, text: mess.Text})
	mp.messages[mp.procStatus] = messages
}

func (b *Bot) tgHandler() {
	// Loop through all tgGetChan when they came
	for update := range b.tgGetChan {
		if update.Message == nil {
			continue
		}
		//update.Message.From.ID
		b.messageHandler(update.Message)
	}
}

func (b *Bot) messageHandler(mess *telego.Message) {
	messPack := b.findMessPack(mess)
	procStatus, tgStatus := b.getMessStatus(mess)

	if procStatus != proc.ID_default {
		messPack.procStatus = procStatus
	}

	if tgStatus == ID_start {
		if messPack.procStatus == proc.ID_default {
			_, _ = b.telegram.SendMessage(tu.Message(tu.ID(messPack.chatID), "ввод не распознан"))
		} else {
			messPack.write(mess)
		}
	}
	if tgStatus == ID_count {
		b.sendPack(messPack)
		messPack.procStatus = proc.ID_default
	}
}

func (b *Bot) findMessPack(mess *telego.Message) *messagesPack {
	chatID := mess.Chat.ID
	if pack, ok := b.messPacks[chatID]; ok {
		return pack
	}
	pack := &messagesPack{chatID: chatID, procStatus: proc.ID_default}
	b.messPacks[chatID] = pack
	return pack
}

func (b *Bot) getMessStatus(mess *telego.Message) (procStatus int64, tgStatus int64) {
	text := mess.Text
	switch text {
	case "***47***":
		return proc.ID_47, -1
	case "***цветок***":
		return proc.ID_flower, -1
	case "***1C***":
		return proc.ID_oneC, -1
	case "***перегоны***":
		return proc.ID_stretches, -1
	case "***анализ***":
		return -1, ID_count
	default:
		return -1, ID_start
	}
}

func (b *Bot) sendPack(pack *messagesPack) {
	unProcPack := proc.UnProcData{
		ID:        pack.chatID,
		MessPacks: pack.messages,
	}
	b.procGetChan <- unProcPack
}
