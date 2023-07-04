package parser

import (
	"disp_bot/utils"
	"strings"
)

func (p *Parser) rewriteLocation(str string, location string) string {
	if p.findOneCStretches(str) {
		return location
	}
	newLocation := p.findLocation(str)
	if newLocation != "" {
		location = newLocation
	}
	return location
}

func (p *Parser) oneC(messages []utils.Message) (to, repair map[string]utils.Message) {
	to = make(map[string]utils.Message, 10)
	repair = make(map[string]utils.Message, 10)

	var location string
	for _, msg := range messages {
		strs := strings.Split(msg.Text, "\n")
		for _, str := range strs {
			location = p.rewriteLocation(str, location)
			mark := p.findMark(str)
			if location != "" && mark != "" {
				tmpMsg := msg
				tmpMsg.Text = "text from 1C"
				tmpMsg.Loc = location
				tmpMsg.Mark = mark
				if p.findTO(str) {
					to[mark] = tmpMsg
				} else {
					repair[mark] = tmpMsg
				}
			}
		}
	}
	return
}
