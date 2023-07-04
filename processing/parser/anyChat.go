package parser

import (
	"disp_bot/utils"
)

func (p *Parser) anyChat(messages []utils.Message, location string) (res map[string]utils.Message) {
	res = make(map[string]utils.Message, 10)
	for _, msg := range messages {
		marks := p.findMarks(msg.Text)
		for _, mark := range marks {
			tmpMsg := msg
			tmpMsg.Loc = location
			tmpMsg.Mark = mark
			res[mark] = tmpMsg
		}
	}
	return res
}
