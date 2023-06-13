package processing

import (
	"disp_bot/processing/analyzer"
	"disp_bot/processing/parser"
	"disp_bot/telegram"
)

// inputIDs
const (
	ID_default = iota
	ID_47
	ID_flower
	ID_oneC
	ID_stretches
)

type ProcData struct {
	ID        int64
	MessPacks []telegram.Message
}

type UnProcData struct {
	ID        int64
	MessPacks map[int64][]telegram.Message
}

type Proc struct {
	bot      *telegram.Bot
	parser   *parser.Parser
	analyzer *analyzer.Analyzer

	getChan  chan UnProcData
	sendChan chan ProcData
}

func Init() *Proc {
	p := &Proc{}

	p.sendChan = make(chan ProcData)
	p.getChan = make(chan UnProcData)

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
		unParsedData := parser.UnParsedData{
			Chat47:        data.MessPacks[ID_47],
			ChatFlower:    data.MessPacks[ID_flower],
			OneC:          data.MessPacks[ID_oneC],
			ChatStretches: data.MessPacks[ID_stretches],
		}
		parsedData := p.parser.Parse(unParsedData)
		procData := p.analyzer.Analyze(analyzer.ParsedData(parsedData))
		p.sendChan <- ProcData(procData)
	}
}
