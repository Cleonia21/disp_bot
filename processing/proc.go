package processing

import (
	"disp_bot/processing/analyzer"
	"disp_bot/processing/parser"
	"disp_bot/utils"
	"log"
)

type Proc struct {
	parser   *parser.Parser
	analyzer *analyzer.Analyzer

	loggingUndef   bool
	mailTestFlag   bool
	mailParsedData func() (service map[string]utils.Message, kuzov map[string]utils.Message)
}

func Init(loggingUndef bool) *Proc {
	p := &Proc{}

	p.parser = parser.Init()
	p.analyzer = analyzer.Init()

	p.loggingUndef = loggingUndef

	return p
}

func (p *Proc) Processing(msgs []utils.Message, conf utils.Conf) (respMsgs []utils.Message) {
	if p.mailTestFlag {
		conf.Mail = false
	}

	parsedData, undef := p.parser.Parse(msgs, conf)

	if p.loggingUndef {
		log.Printf("undefined messages\n")
		for _, u := range undef {
			u.Print()
		}
	}

	if p.mailTestFlag {
		parsedData.MailService, parsedData.MailKuzov = p.mailParsedData()
		conf.Mail = true
	}

	analyzedMsgs := p.analyzer.Analyze(parsedData, conf)
	respMsgs = append(respMsgs, analyzedMsgs...)
	return
}
