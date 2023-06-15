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
	parsedData := utils.ParsedData{}
	unidentified := make([]utils.Message, 10)

	parsedData.Chat47 = p.anyChat(data.Chat47, "")
	parsedData.ChatFlower = p.anyChat(data.ChatFlower, "")
	parsedData.OneCto, parsedData.OneCRepair = p.oneC(data.OneC)

	var unidents []utils.Message
	parsedData.ChatStretches, unidents = p.stretchesChat(data.ChatStretches)
	unidentified = append(unidentified, unidents...)
	parsedData.Mail, unidents = p.mail()
	unidentified = append(unidentified, unidents...)

	parsedData.Unidentified = unidentified

	return parsedData
}
