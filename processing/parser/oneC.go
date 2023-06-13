package parser

import (
	"disp_bot/utils"
	"strings"
)

func (p *Parser) rewriteLocation(str string, location string) string {
	newLocation := p.findLocation(str)
	if newLocation != "" {
		location = newLocation
	}
	return location
}

func (p *Parser) oneC(messages []utils.Message) (res []utils.Resource) {
	res = make([]utils.Resource, 10)

	var location string
	for _, mess := range messages {
		strs := strings.Split(mess.Text, "\n")
		for _, str := range strs {
			location = p.rewriteLocation(str, location)
			mark := p.findRegMark(str)
			if location != "" && mark != "" {
				res = append(res, utils.Resource{
					StRegMark: mark,
					Loc:       location,
					Analyzed:  true,
					Mess:      mess,
				})
			}
		}
	}
	return res
}
