package parser

import "disp_bot/telegram"

func (p *Parser) anyChat(messages []telegram.Message, location Location) (res Resource) {
	res = make(Resource)
	for _, mess := range messages {
		text := mess.Text()
		if stateRegMark := p.findRegMark(text); stateRegMark != "" {
			res[stateRegMark] = location
		}
	}
	return res
}
