package analyzer

import (
	"disp_bot/utils"
)

func rules() func(key string) []rulePack {
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

		if find {
			delete(data.OneCRepair, grz)
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

		if find {
			delete(data.OneCto, grz)
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

		if find {
			delete(data.Mail, grz)
		}
	}

	if rule.chatFlower {
		_, find := data.ChatFlower[grz]
		if !find {
			ok = false
			msgs = append(msgs, utils.NewMessage(grz+" не найден в чате цветка"))
		}

		if find {
			delete(data.ChatFlower, grz)
		}
	}

	if rule.chat47 {
		_, find := data.Chat47[grz]
		if !find {
			ok = false
			msgs = append(msgs, utils.NewMessage(grz+" не найден в чате 47го"))
		}

		if find {
			delete(data.Chat47, grz)
		}
	}
	return
}

func (a *Analyzer) findUnused(parsedData utils.ParsedData) []utils.Message {
	findRess := func(ress map[string]utils.Resource) []utils.Message {
		var msgs []utils.Message
		for grz, res := range ress {
			res.Mess.AddReply(grz + " не найден в чате перегонов")
			msgs = append(msgs, res.Mess)
		}
		return msgs
	}

	var msgs []utils.Message

	msgs = append(msgs, findRess(parsedData.ChatFlower)...)
	msgs = append(msgs, findRess(parsedData.Chat47)...)
	msgs = append(msgs, findRess(parsedData.OneCRepair)...)
	msgs = append(msgs, findRess(parsedData.OneCto)...)
	msgs = append(msgs, findRess(parsedData.Mail)...)

	return msgs
}

func (a *Analyzer) Analyze(parsedData utils.ParsedData) []utils.Message {
	messages := parsedData.Unidentified
	getRules := rules()

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

	messages = append(messages, a.findUnused(parsedData)...)
	messages = append(messages, parsedData.Unidentified...)

	return messages
}
