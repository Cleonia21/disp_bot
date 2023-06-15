package parser

import (
	"disp_bot/utils"
)

func (p *Parser) anyChat(messages []utils.Message, location string) (res map[string]utils.Resource) {
	res = make(map[string]utils.Resource, 10)
	for _, mess := range messages {
		marks := p.findMarks(mess.Text)
		for _, mark := range marks {
			res[mark] = utils.Resource{
				Loc:  location,
				Mess: mess,
			}
		}
	}
	return res
}
