package processing

import (
	"disp_bot/processing/analyzer"
	"disp_bot/processing/parser"
	"disp_bot/telegram"
	"disp_bot/utils"
)

type Proc struct {
	bot      *telegram.Bot
	parser   *parser.Parser
	analyzer *analyzer.Analyzer

	getChan  chan utils.UnProcData
	sendChan chan utils.ProcData
}

func Init() *Proc {
	p := &Proc{}

	p.sendChan = make(chan utils.ProcData)
	p.getChan = make(chan utils.UnProcData)

	p.bot = telegram.Init(p.getChan, p.sendChan)
	p.parser = parser.Init()
	p.analyzer = analyzer.Init()

	return p
}

func (p *Proc) Start() {
	go p.getData()
}

func (p *Proc) Stop() {

}

func (p *Proc) getData() {
	for data := range p.getChan {
		unParsedData := utils.UnParsedData{
			Chat47:        data.MessPacks[utils.ID_47],
			ChatFlower:    data.MessPacks[utils.ID_flower],
			OneC:          data.MessPacks[utils.ID_oneC],
			ChatStretches: data.MessPacks[utils.ID_stretches],
		}
		parsedData := p.parser.Parse(unParsedData)
		procData := p.analyzer.Analyze(parsedData)
		p.sendChan <- utils.ProcData(procData)
	}
}
