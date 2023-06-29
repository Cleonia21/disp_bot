package processing

import (
	"disp_bot/processing/analyzer"
	"disp_bot/processing/parser"
	"disp_bot/utils"
)

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
		MessPacks: msgs,
	}
	return &procData
}
