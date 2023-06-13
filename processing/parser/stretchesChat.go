package parser

import (
	"disp_bot/utils"
	"strings"
)

func newResource(mark string, location string, message utils.Message) utils.Resource {
	var res utils.Resource
	if mark != "" && location != "" {
		res = utils.Resource{
			StRegMark: mark,
			Loc:       location,
			Analyzed:  true,
			Mess:      message,
		}
	} else if location != "" {
		message.AddReply("Не распознан ГРЗ")
		res = utils.Resource{
			Analyzed: false,
			Mess:     message,
		}
	} else if mark != "" {
		message.AddReply("Не распознан сервис")
		res = utils.Resource{
			Analyzed: false,
			Mess:     message,
		}
	}
	return res
}

func (p *Parser) stretchesChat(messages []utils.Message) (res []utils.Resource) {
	res = make([]utils.Resource, 10)

	for _, mess := range messages {
		strs := strings.Split(mess.Text, "\n")
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
