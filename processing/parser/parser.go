package parser

import (
	"disp_bot/utils"
)

type Parser struct {
	locationsRegExp   map[string]string
	stateMarkRegExp   string
	techServiceRegExp string
}

func Init() *Parser {
	p := &Parser{}
	p.locationsRegExp = locationsRegExp()
	p.stateMarkRegExp = stateMarkRegExp()
	p.techServiceRegExp = techServiceRegExp()
	return p
}

func (p *Parser) Parse(data utils.UnParsedData) utils.ParsedData {
	parsedData := utils.ParsedData{
		Chat47:        p.anyChat(data.Chat47, ""),
		ChatFlower:    p.anyChat(data.ChatFlower, ""),
		ChatStretches: p.stretchesChat(data.ChatStretches),
		OneC:          p.oneC(data.OneC),
		Mail:          p.mail(),
	}
	return parsedData
}
