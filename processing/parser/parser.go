package parser

import (
	"disp_bot/utils"
	"log"
	"regexp"
)

type Parser struct {
	locationsRegExp   map[string]*regexp.Regexp
	stateMarkRegExp   *regexp.Regexp
	techServiceRegExp *regexp.Regexp
	urlRegExp         *regexp.Regexp
}

func Init() *Parser {
	p := &Parser{}
	var err error
	p.locationsRegExp, err = locationsRegExp()
	if err != nil {
		log.Println(err)
	}
	p.stateMarkRegExp, err = stateMarkRegExp()
	if err != nil {
		log.Println(err)
	}
	p.techServiceRegExp, err = techServiceRegExp()
	if err != nil {
		log.Println(err)
	}
	p.urlRegExp, err = urlRegExp()
	if err != nil {
		log.Println(err)
	}
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
