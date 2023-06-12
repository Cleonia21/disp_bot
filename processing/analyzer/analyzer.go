package analyzer

import proc "disp_bot/processing"

type Analyzer struct {
}

func Init() *Analyzer {
	return &Analyzer{}
}

func Analyze(data proc.ParsedData) proc.ProcData {
	return proc.ProcData{}
}
