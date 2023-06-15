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

func (p *Parser) oneC(messages []utils.Message) (to, repair map[string]utils.Resource) {
	to = make(map[string]utils.Resource, 10)
	repair = make(map[string]utils.Resource, 10)

	var location string
	for _, mess := range messages {
		strs := strings.Split(mess.Text, "\n")
		for _, str := range strs {
			location = p.rewriteLocation(str, location)
			mark := p.findRegMark(str)
			if location != "" && mark != "" {
				if p.findTO(str) != "" {
					to[mark] = utils.Resource{
						Loc:  location,
						Mess: mess,
					}
				} else {
					repair[mark] = utils.Resource{
						Loc:  location,
						Mess: mess,
					}
				}
			}
		}
	}
	return
}
