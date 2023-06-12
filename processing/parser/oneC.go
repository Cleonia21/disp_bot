package parser

import "disp_bot/telegram"

func (p *Parser) oneC(messages []telegram.Message) (res Resource) {
	res = make(Resource)

	for _, mess := range messages {
		text := mess.Text()
		stateRegMark := p.findRegMark(text)
		location := p.findLocation(text)
		if stateRegMark != "" && location != "" {
			res[stateRegMark] = location
		}
	}
	return res
}
