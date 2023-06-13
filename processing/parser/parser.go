package parser

import (
	"disp_bot/telegram"
)

type Resource struct {
	StRegMark string
	Loc       string
	analyzed  bool
	mess      telegram.Message
}

type ParsedData struct {
	Chat47        []Resource
	ChatFlower    []Resource
	OneC          []Resource
	Mail          []Resource
	ChatStretches []Resource
}

type UnParsedData struct {
	Chat47        []telegram.Message
	ChatFlower    []telegram.Message
	OneC          []telegram.Message
	ChatStretches []telegram.Message
}

type Parser struct {
	locationsRegExp map[string]string
	stateMarkRegExp string
}

func Init() *Parser {
	p := &Parser{}
	p.locationsRegExp = locationsRegExp()
	p.stateMarkRegExp = stateMarkRegExp()
	return p
}

func (p *Parser) Parse(data UnParsedData) ParsedData {
	parsedData := ParsedData{
		Chat47:        p.anyChat(data.Chat47, ""),
		ChatFlower:    p.anyChat(data.ChatFlower, ""),
		ChatStretches: p.stretchesChat(data.ChatStretches),
		OneC:          p.oneC(data.OneC),
		Mail:          p.mail(),
	}
	return parsedData
}
