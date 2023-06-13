package parser

import (
	"disp_bot/utils"
)

func (p *Parser) anyChat(messages []utils.Message, location string) (res []utils.Resource) {
	res = make([]utils.Resource, 10)
	for _, mess := range messages {
		marks := p.findRegMarks(mess.Text)
		for _, mark := range marks {
			res = append(res, utils.Resource{
				StRegMark: mark,
				Loc:       location,
				Analyzed:  true,
				Mess:      mess,
			})
		}
	}
	return res
}
