package telegram

import (
	"disp_bot/utils"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// inputIDs
const (
	ID_start = iota
	ID_count
)

type messagesPack struct {
	chatID     int64
	procStatus int64
	messages   map[int64][]utils.Message
}

func (mp *messagesPack) write(mess *telego.Message) {
	messages := mp.messages[mp.procStatus]
	messages = append(messages, utils.Message{ID: mess.MessageID, Text: mess.Text})
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

	if procStatus != utils.ID_default {
		messPack.procStatus = procStatus
	}

	if tgStatus == ID_start {
		if messPack.procStatus == utils.ID_default {
			_, _ = b.telegram.SendMessage(tu.Message(tu.ID(messPack.chatID), "ввод не распознан"))
		} else {
			messPack.write(mess)
		}
	}
	if tgStatus == ID_count {
		b.sendPack(messPack)
		messPack.procStatus = utils.ID_default
	}
}

func (b *Bot) findMessPack(mess *telego.Message) *messagesPack {
	chatID := mess.Chat.ID
	if pack, ok := b.messPacks[chatID]; ok {
		return pack
	}
	pack := &messagesPack{chatID: chatID, procStatus: utils.ID_default}
	b.messPacks[chatID] = pack
	return pack
}

func (b *Bot) getMessStatus(mess *telego.Message) (procStatus int64, tgStatus int64) {
	text := mess.Text
	switch text {
	case "***47***":
		return utils.ID_47, -1
	case "***цветок***":
		return utils.ID_flower, -1
	case "***1C***":
		return utils.ID_oneC, -1
	case "***перегоны***":
		return utils.ID_stretches, -1
	case "***анализ***":
		return -1, ID_count
	default:
		return -1, ID_start
	}
}

func (b *Bot) sendPack(pack *messagesPack) {
	unProcPack := utils.UnProcData{
		ID:        pack.chatID,
		MessPacks: pack.messages,
	}
	b.procGetChan <- unProcPack
}
