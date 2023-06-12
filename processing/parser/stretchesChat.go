package parser

import (
	"disp_bot/telegram"
	"strings"
)

func (p *Parser) stretchesChat(messages []telegram.Message) (res Resource) {
	res = make(Resource)

	for _, mess := range messages {
		strs := strings.Split(mess.Text(), "\n")
		if len(strs) > 3 {
			strs = strs[:3]
		}
		str := strings.Join(strs, "\n")

		stateRegMark := p.findRegMark(str)
		location := p.findLocation(str)
		if stateRegMark != "" && location != "" {
			res[stateRegMark] = location
		}
	}
	return res
}
