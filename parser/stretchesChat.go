package parser

import (
	"github.com/mymmrac/telego"
	"strings"
)

func (p *Parser) stretchesChat(updates []telego.Update) (res Resource) {
	res = make(Resource)

	for _, update := range updates {
		text := update.Message.Text
		strs := strings.Split(text, "\n")
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
