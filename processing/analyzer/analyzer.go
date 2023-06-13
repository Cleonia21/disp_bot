package analyzer

import (
	"disp_bot/utils"
)

type Analyzer struct {
}

func Init() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) Analyze(parsedData utils.ParsedData) utils.ProcData {
	return utils.ProcData{}
}
