package analyzer

import (
	"disp_bot/telegram"
)

type Analyzer struct {
}

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

type ProcData struct {
	ID        int64
	MessPacks []telegram.Message
}

func Init() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) Analyze(parsedData ParsedData) ProcData {
	return ProcData{}
}
