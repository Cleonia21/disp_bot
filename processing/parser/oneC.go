package parser

import (
	"disp_bot/telegram"
	"strings"
)

func (p *Parser) rewriteLocation(str string, location string) string {
	newLocation := p.findLocation(str)
	if newLocation != "" {
		location = newLocation
	}
	return location
}

func (p *Parser) oneC(messages []telegram.Message) (res []Resource) {
	res = make([]Resource, 10)

	var location string
	for _, mess := range messages {
		strs := strings.Split(mess.Text(), "\n")
		for _, str := range strs {
			location = p.rewriteLocation(str, location)
			mark := p.findRegMark(str)
			if location != "" && mark != "" {
				res = append(res, Resource{
					StRegMark: mark,
					Loc:       location,
					analyzed:  true,
					mess:      mess,
				})
			}
		}
	}
	return res
}
