package analyzer

import (
	"disp_bot/utils"
)

func getRules() func(key string) []rulePack {
	rules := map[string][]rulePack{
		"дженерал":      {{mail: true, oneCto: true}, {oneCto: true}},
		"магистральная": {{mail: true, oneCto: true}},
		"цкр":           {{mail: true, oneCto: true}},
		"авторемонт огородный": {{mail: true, oneCto: true}},
		"сокольники":           {{mail: true, oneCRepair: true}},
		"кулак":                {{mail: true, oneCRepair: true}},
		"остап":                {{mail: true, oneCRepair: true}},
		"волжский":             {{mail: true, oneCRepair: true}},
		"полбина":              {{mail: true, oneCRepair: true}, {oneCto: true}},
		"выборгская":           {{mail: true, oneCRepair: true}},
		"рольф ярославка":      {{mail: true, oneCRepair: true}},
		"рольф варшавка":       {{mail: true, oneCRepair: true}, {oneCto: true}},
		"автофорум":            {{mail: true, oneCRepair: true}, {oneCto: true}},
		"про-сто":              {{mail: true, oneCRepair: true}},
		"автолайт прокатная":   {{mail: true, oneCRepair: true}},
		"поречная":             {{mail: true, oneCRepair: true}},
		"рольф восток":         {{mail: true, oneCRepair: true}},
		"авторемонт плюс":      {{mail: true, oneCRepair: true}},

		"цветок":       {{mail: true, oneCRepair: true}, {oneCto: true, chatFlower: true}},
		"47":           {{mail: true, oneCRepair: true}, {oneCto: true, chat47: true}},
		"марксистская": {{mail: true, oneCRepair: true}, {oneCto: true}},
		"шикана":       {{mail: true, oneCRepair: true}},
		"дубровка":     {{mail: true, oneCRepair: true}, {oneCto: true}},
		"автоигл":      {{mail: true, oneCRepair: true}},
		//"рольф центр": ["=========не нашел таких перегонов========="],
		"обручева": {{mail: true, oneCRepair: true}, {oneCto: true}},
		"авторусь": {{mail: true, oneCRepair: true}},
		"азимут":   {{mail: true, oneCRepair: true}},
		"авалон":   {{mail: true, oneCRepair: true}},

		"казаков": {{oneCRepair: true}},
		"офис":    {{mail: true}},
	}

	return func(key string) []rulePack {
		return rules[key]
	}
}

type rulePack struct {
	mail       bool
	chat47     bool
	chatFlower bool
	oneCto     bool
	oneCRepair bool
}

type Analyzer struct {
}

func Init() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) checkRule(data utils.ParsedData, rule rulePack, grz string) (msgs []utils.Message, ok bool) {
	resStrech := data.ChatStretches[grz]
	if rule.oneCRepair {
		resOneCRepair, find := data.OneCRepair[grz]
		if !find {
			ok = false
			msgs = append(msgs, utils.NewMessage(grz+" не найден в 1С(как заявка на ремонт)"))
		} else if resOneCRepair.Loc != resStrech.Loc {
			ok = false
			resOneCRepair.Mess.AddReply(grz + " не совпадает с сервисом из перегоны")
			msgs = append(msgs, resOneCRepair.Mess)
		}
	}

	if rule.oneCto {
		resOneCto, find := data.OneCto[grz]
		if !find {
			ok = false
			msgs = append(msgs, utils.NewMessage(grz+" не найден в 1С(как заявка на ТО)"))
		} else if resOneCto.Loc != resStrech.Loc {
			ok = false
			resOneCto.Mess.AddReply(grz + " не совпадает с сервисом из перегоны")
			msgs = append(msgs, resOneCto.Mess)
		}
	}

	if rule.mail {
		resMail, find := data.Mail[grz]
		if !find {
			ok = false
			msgs = append(msgs, utils.NewMessage(grz+" не найден на почте"))
		} else if resMail.Loc != resStrech.Loc {
			ok = false
			resMail.Mess.AddReply(grz + " не совпадает с сервисом из перегоны")
			msgs = append(msgs, resMail.Mess)
		}
	}

	if rule.chatFlower {
		_, find := data.ChatFlower[grz]
		if !find {
			ok = false
			msgs = append(msgs, utils.NewMessage(grz+" не найден в чате цветка"))
		}
	}

	if rule.chat47 {
		_, find := data.Chat47[grz]
		if !find {
			ok = false
			msgs = append(msgs, utils.NewMessage(grz+" не найден в чате 47го"))
		}
	}
	return
}

func (a *Analyzer) Analyze(parsedData utils.ParsedData) []utils.Message {
	messages := parsedData.Unidentified
	getRules := getRules()

	for grz, stretch := range parsedData.ChatStretches {
		var tmpMsgs []utils.Message
		var ok bool
		rules := getRules(stretch.Loc)
		for _, rule := range rules {
			var checkMsgs []utils.Message
			checkMsgs, ok = a.checkRule(parsedData, rule, grz)
			if ok {
				break
			}
			tmpMsgs = append(tmpMsgs, checkMsgs...)
		}
		if !ok {
			messages = append(messages, tmpMsgs...)
		}
	}
	return messages
}
