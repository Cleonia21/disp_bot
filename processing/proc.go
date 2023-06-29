package processing

import (
	"disp_bot/processing/analyzer"
	"disp_bot/processing/parser"
	"disp_bot/utils"
)

type MailConf struct {
	UserName string
	Pass     string
}

type Checks struct {
	Chat47        bool
	ChatFlower    bool
	OneC          bool
	ChatStretches bool
	Mail          bool
}

type UnProcData struct {
	ID        int64
	MessPacks map[int64][]Message

	ProcConf struct {
		Checks   Checks
		MailConf MailConf
	}
}

type Proc struct {
	parser   *parser.Parser
	analyzer *analyzer.Analyzer
}

func Init() *Proc {
	p := &Proc{}

	p.parser = parser.Init()
	p.analyzer = analyzer.Init()

	return p
}

func (p *Proc) Processing(data *utils.UnProcData) *utils.ProcData {
	unParsedData := utils.UnParsedData{
		Chat47:        data.MessPacks[utils.ID_47],
		ChatFlower:    data.MessPacks[utils.ID_flower],
		OneC:          data.MessPacks[utils.ID_oneC],
		ChatStretches: data.MessPacks[utils.ID_stretches],
	}
	parsedData := p.parser.Parse(unParsedData)
	msgs := p.analyzer.Analyze(parsedData)
	procData := utils.ProcData{
		ID:        data.ID,
		MessPacks: msgs,
	}
	return &procData
}
