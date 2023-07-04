package parser

import (
	"disp_bot/processing/analyzer"
	"disp_bot/utils"
	"log"
	"regexp"
)

type Parser struct {
	locationsRegExp     map[string]*regexp.Regexp
	stateMarkRegExp     *regexp.Regexp
	techServiceRegExp   *regexp.Regexp
	urlRegExp           *regexp.Regexp
	oneCStretchesRegExp *regexp.Regexp
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
	p.oneCStretchesRegExp, err = oneCStretchesRegExp()
	if err != nil {
		log.Println(err)
	}
	return p
}

type data struct {
	chat47        []utils.Message
	ChatFlower    []utils.Message
	OneC          []utils.Message
	ChatStretches []utils.Message
}

func (d *data) fill(msgs []utils.Message) {
	for _, msg := range msgs {
		switch msg.From {
		case utils.ID_47:
			d.chat47 = append(d.chat47, msg)
		case utils.ID_flower:
			d.ChatFlower = append(d.ChatFlower, msg)
		case utils.ID_oneC:
			d.OneC = append(d.OneC, msg)
		case utils.ID_stretches:
			d.ChatStretches = append(d.ChatStretches, msg)
		}
	}
}

func (p *Parser) Parse(msgs []utils.Message, conf utils.Conf) (analyzerData analyzer.Data, undef []utils.Message) {
	var d data
	d.fill(msgs)

	if conf.Chat47 {
		analyzerData.Chat47 = p.anyChat(d.chat47, "47")
	}
	if conf.ChatFlower {
		analyzerData.ChatFlower = p.anyChat(d.ChatFlower, "цветок")
	}
	if conf.OneC {
		analyzerData.OneCto, analyzerData.OneCRepair = p.oneC(d.OneC)
	}
	if conf.ChatStretches {
		var tmpUndef []utils.Message
		analyzerData.ChatStretches, tmpUndef = p.stretchesChat(d.ChatStretches)
		undef = append(undef, tmpUndef...)
	}
	if conf.Mail {
		var tmpUndef []utils.Message
		analyzerData.MailKuzov, tmpUndef = p.mail("/home/cleonia/Desktop/disp_bot/conf/mailKuzov.json")
		undef = append(undef, tmpUndef...)
		analyzerData.MailService, tmpUndef = p.mail("/home/cleonia/Desktop/disp_bot/conf/mailService.json")
		undef = append(undef, tmpUndef...)
	}
	return
}
