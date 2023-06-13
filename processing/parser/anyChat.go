package parser

import (
	"disp_bot/telegram"
)

func (p *Parser) anyChat(messages []telegram.Message, location string) (res []Resource) {
	res = make([]Resource, 10)
	for _, mess := range messages {
		marks := p.findRegMarks(mess.Text())
		for _, mark := range marks {
			res = append(res, Resource{
				StRegMark: mark,
				Loc:       location,
				analyzed:  true,
				mess:      mess,
			})
		}
	}
	return res
}
