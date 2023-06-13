package parser

import (
	"disp_bot/telegram"
	"strings"
)

func newResource(mark string, location string, message telegram.Message) Resource {
	var res Resource
	if mark != "" && location != "" {
		res = Resource{
			StRegMark: mark,
			Loc:       location,
			analyzed:  true,
			mess:      message,
		}
	} else if location != "" {
		message.AddReply("Не распознан ГРЗ")
		res = Resource{
			analyzed: false,
			mess:     message,
		}
	} else if mark != "" {
		message.AddReply("Не распознан сервис")
		res = Resource{
			analyzed: false,
			mess:     message,
		}
	}
	return res
}

func (p *Parser) stretchesChat(messages []telegram.Message) (res []Resource) {
	res = make([]Resource, 10)

	for _, mess := range messages {
		strs := strings.Split(mess.Text(), "\n")
		if len(strs) > 3 {
			strs = strs[:3]
		}
		str := strings.Join(strs, "\n")

		mark := p.findRegMark(str)
		location := p.findLocation(str)
		res = append(res, newResource(mark, location, mess))
	}
	return res
}
