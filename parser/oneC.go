package parser

import (
	"github.com/mymmrac/telego"
)

func (p *Parser) oneC(updates []telego.Update) (res Resource) {
	res = make(Resource)

	for _, update := range updates {
		text := update.Message.Text
		stateRegMark := p.findRegMark(text)
		location := p.findLocation(text)
		if stateRegMark != "" && location != "" {
			res[stateRegMark] = location
		}
	}
	return res
}
