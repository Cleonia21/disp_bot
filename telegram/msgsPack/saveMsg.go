package msgsPack

import (
	"disp_bot/utils"
	"github.com/mymmrac/telego"
)

func (p *MsgsPack) saveMsg(mess *telego.Message) {
	if p.status == idDefault {
		return
	}
	customMsg := utils.Message{
		ID:   mess.MessageID,
		Text: mess.Text,
		From: IDConverter(p.status),
	}
	p.messages[p.status] = append(p.messages[p.status], customMsg)
}

func IDConverter(packID string) (utilsID int) {
	switch packID {
	case id1C:
		return utils.ID_oneC
	case idFlower:
		return utils.ID_flower
	case id47:
		return utils.ID_47
	case idStretches:
		return utils.ID_stretches
	default:
		return utils.ID_default
	}
}
