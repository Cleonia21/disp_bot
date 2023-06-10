package parser

import (
	"github.com/mymmrac/telego"
)

func (p *Parser) anyChat(updates []telego.Update, location Location) (res Resource) {
	res = make(Resource)
	for _, update := range updates {
		text := update.Message.Text
		if stateRegMark := p.findRegMark(text); stateRegMark != "" {
			res[stateRegMark] = location
		}
	}
	return res
}
