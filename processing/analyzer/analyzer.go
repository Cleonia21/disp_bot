package analyzer

import (
	"disp_bot/utils"
)

type rule struct {
}

type rules struct {
	Дженерал            []string `json:"дженерал"`
	Магистральная       []string `json:"магистральная"`
	Цкр                 []string `json:"цкр"`
	АвторемонтОгородный []string `json:"авторемонт огородный"`
	Сокольники          []string `json:"сокольники"`
	Кулак               []string `json:"кулак"`
	Остап               []string `json:"остап"`
	Волжский            []string `json:"волжский"`
	Полбина             []string `json:"полбина"`
	Выборгская          []string `json:"выборгская"`
	РольфЯрославка      []string `json:"рольф ярославка"`
	РольфВаршавка       []string `json:"рольф варшавка"`
	Автофорум           []string `json:"автофорум"`
	ПроСто              []string `json:"про-сто"`
	АвтолайтПрокатная   []string `json:"автолайт прокатная"`
	Поречная            []string `json:"поречная"`
	РольфВосток         []string `json:"рольф восток"`
	АвторемонтПлюс      []string `json:"авторемонт плюс"`
	Цветок              []string `json:"цветок"`
	Field20             []string `json:"47"`
	Марксистская        []string `json:"марксистская"`
	Шикана              []string `json:"шикана"`
	Дубровка            []string `json:"дубровка"`
	Автоигл             []string `json:"автоигл"`
	РольфЦентр          []string `json:"рольф центр"`
	Обручева            []string `json:"обручева"`
	Авторусь            []string `json:"авторусь"`
	Азимут              []string `json:"азимут"`
	Авалон              []string `json:"авалон"`
	Казаков             []string `json:"казаков"`
	Офис                []string `json:"офис"`
}

type Analyzer struct {
}

func Init() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) Analyze(parsedData utils.ParsedData) utils.ProcData {

}
