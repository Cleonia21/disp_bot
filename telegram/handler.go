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
	procStatus int
	messages   []utils.Message
}

func (b *Bot) handler(mess *telego.Message) {
	messPack := b.findMessPack(mess)
	procStatus, tgStatus := b.getMessStatus(mess)

	if procStatus != utils.ID_default {
		messPack.procStatus = procStatus
	}

	if tgStatus == ID_start {
		if messPack.procStatus == utils.ID_default {
			_, _ = b.telegram.SendMessage(tu.Message(tu.ID(messPack.chatID), "ввод не распознан"))
		} else {
			messPack.save(mess)
		}
	}
	if tgStatus == ID_count {
		go b.processing(messPack)
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

func (b *Bot) getMessStatus(mess *telego.Message) (procStatus int, tgStatus int64) {
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

func (mp *messagesPack) save(mess *telego.Message) {
	customMsg := utils.Message{
		ID:   mess.MessageID,
		Text: mess.Text,
		From: mp.procStatus,
	}
	mp.messages = append(mp.messages, customMsg)
}

func (b *Bot) processing(pack *messagesPack) {
	msgs := b.proc.Processing(pack.messages, utils.Conf{
		Chat47:        true,
		ChatFlower:    true,
		OneC:          true,
		ChatStretches: true,
		Mail:          true,
	})

	for _, msg := range msgs {
		b.sendCustomMsg(pack.chatID, &msg)
	}
}

func (b *Bot) sendCustomMsg(chatID int64, msg *utils.Message) {

}
