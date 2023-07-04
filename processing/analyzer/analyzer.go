package analyzer

import (
	"disp_bot/utils"
	"reflect"
)

type Data struct {
	Chat47        map[string]utils.Message
	ChatFlower    map[string]utils.Message
	OneCto        map[string]utils.Message
	OneCRepair    map[string]utils.Message
	MailKuzov     map[string]utils.Message
	MailService   map[string]utils.Message
	ChatStretches map[string]utils.Message
}

func rules(conf utils.Conf) func(key string) []rulePack {
	rulesData := map[string][]rulePack{
		"дженерал":      {{mailService: true, oneCService: true}, {oneCto: true}, {mailKuzov: true}},
		"магистральная": {{mailKuzov: true}},
		"цкр":           {{mailService: true, oneCService: true}, {mailKuzov: true}},
		"авторемонт огородный": {{mailKuzov: true}},
		"сокольники":           {{mailKuzov: true}},
		"кулак":                {{mailService: true, oneCService: true}, {mailKuzov: true}},
		"остап":                {{mailKuzov: true}},
		"волжский":             {{mailService: true, oneCService: true}, {mailKuzov: true}},
		"полбина":              {{mailService: true, oneCService: true}, {oneCto: true}, {mailKuzov: true}},
		"выборгская":           {{mailService: true, oneCService: true}, {mailKuzov: true}},
		"рольф ярославка":      {{mailKuzov: true}},
		"рольф варшавка":       {{mailKuzov: true}, {oneCto: true}},
		"автофорум":            {{mailKuzov: true}, {oneCto: true}},
		"про-сто":              {{mailKuzov: true}},
		"автолайт прокатная":   {{mailKuzov: true}},
		"поречная":             {{mailKuzov: true}},
		"рольф восток":         {{mailKuzov: true}},
		"авторемонт плюс":      {{mailKuzov: true}},

		"цветок":       {{mailService: true, oneCService: true}, {oneCto: true, chatFlower: true}},
		"47":           {{mailService: true, oneCService: true}, {oneCto: true, chat47: true}},
		"марксистская": {{mailService: true, oneCService: true}, {oneCto: true}},
		"шикана":       {{mailService: true, oneCService: true}},
		"дубровка":     {{mailService: true, oneCService: true}, {oneCto: true}},
		"автоигл":      {{mailService: true, oneCService: true}},
		//"рольф центр": ["=========не нашел таких перегонов========="],
		"обручева": {{mailService: true, oneCService: true}, {oneCto: true}},
		"авторусь": {{mailService: true, oneCService: true}},
		"азимут":   {{mailService: true, oneCService: true}},
		"авалон":   {{mailService: true, oneCService: true}},

		"казаков": {{oneCService: true}},
		"офис":    {{mailService: true}},
	}

	return func(key string) []rulePack {
		rulesSet := rulesData[key]
		var modRulesSet []rulePack
		for _, rule := range rulesSet {
			rule.oneCto = rule.oneCto && conf.OneC
			rule.oneCService = rule.oneCService && conf.OneC
			rule.chatFlower = rule.chatFlower && conf.ChatFlower
			rule.chat47 = rule.chat47 && conf.Chat47
			rule.mailService = rule.mailService && conf.Mail
			rule.mailKuzov = rule.mailKuzov && conf.Mail
			//conf.ChatStretches
			modRulesSet = append(modRulesSet, rule)
		}
		return modRulesSet
	}
}

type rulePack struct {
	mailService bool
	mailKuzov   bool
	chat47      bool
	chatFlower  bool
	oneCto      bool
	oneCService bool
}

type Analyzer struct {
}

func Init() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) Analyze(parsedData Data, conf utils.Conf) (msgs []utils.Message) {
	getRules := rules(conf)

	for grz, stretch := range parsedData.ChatStretches {
		var tmpMsgs []utils.Message
		var ok bool
		rules := getRules(stretch.Loc)
		for _, rule := range rules {
			if reflect.DeepEqual(rule, rulePack{}) {
				continue
			}
			var checkMsgs []utils.Message
			checkMsgs, ok = a.checkRule(parsedData, rule, grz)
			if ok {
				break
			}
			tmpMsgs = append(tmpMsgs, checkMsgs...)
		}
		if !ok {
			msgs = append(msgs, tmpMsgs...)
		}
	}
	msgs = append(msgs, a.findUnused(parsedData)...)
	return
}

func (a *Analyzer) checkRule(data Data, rule rulePack, grz string) (msgs []utils.Message, ok bool) {
	stretchesMsg := data.ChatStretches[grz]
	ok = true

	if rule.oneCService {
		msg, find := data.OneCRepair[grz]
		if !find {
			ok = false
			tmpMsg := stretchesMsg
			tmpMsg.AddReply("не найден в 1С(как заявка на ремонт)")
			msgs = append(msgs, tmpMsg)
		} else if msg.Loc != stretchesMsg.Loc {
			ok = false
			msg.AddReply(grz + " не совпадает с сервисом из перегоны")
			msgs = append(msgs, msg)
		}

		if find {
			delete(data.OneCRepair, grz)
		}
	}

	if rule.oneCto {
		msg, find := data.OneCto[grz]
		if !find {
			ok = false
			tmpMsg := stretchesMsg
			tmpMsg.AddReply("не найден в 1С(как заявка на ТО)")
			msgs = append(msgs, tmpMsg)
		} else if msg.Loc != stretchesMsg.Loc {
			ok = false
			msg.AddReply(grz + " не совпадает с сервисом из перегоны")
			msgs = append(msgs, msg)
		}

		if find {
			delete(data.OneCto, grz)
		}
	}

	if rule.mailKuzov {
		msg, find := data.MailKuzov[grz]
		if !find {
			ok = false
			tmpMsg := stretchesMsg
			tmpMsg.AddReply("не отправлен на kuzov")
			msgs = append(msgs, tmpMsg)
		} else if msg.Loc != stretchesMsg.Loc {
			ok = false
			msg.AddReply(grz + " не совпадает с сервисом из перегоны")
			msgs = append(msgs, msg)
		}

		if find {
			delete(data.MailKuzov, grz)
		}
	}

	if rule.mailService {
		msg, find := data.MailService[grz]
		if !find {
			ok = false
			tmpMsg := stretchesMsg
			tmpMsg.AddReply("не отправлен на service")
			msgs = append(msgs, tmpMsg)
		} else if msg.Loc != stretchesMsg.Loc {
			ok = false
			msg.AddReply(grz + " не совпадает с сервисом из перегоны")
			msgs = append(msgs, msg)
		}

		if find {
			delete(data.MailService, grz)
		}
	}

	if rule.chatFlower {
		_, find := data.ChatFlower[grz]
		if !find {
			ok = false
			tmpMsg := stretchesMsg
			tmpMsg.AddReply("не найден в чате цветка")
			msgs = append(msgs, tmpMsg)
		}

		if find {
			delete(data.ChatFlower, grz)
		}
	}

	if rule.chat47 {
		_, find := data.Chat47[grz]
		if !find {
			ok = false
			tmpMsg := stretchesMsg
			tmpMsg.AddReply("не найден в чате 47го")
			msgs = append(msgs, tmpMsg)
		}

		if find {
			delete(data.Chat47, grz)
		}
	}
	return
}

func (a *Analyzer) findUnused(parsedData Data) (msgs []utils.Message) {
	modAndReturn := func(msgs map[string]utils.Message) (modMsgs []utils.Message) {
		for _, msg := range msgs {
			msg.AddReply(msg.Mark + " не найден в чате перегонов или нахождение в источнике не подразумевается")
			modMsgs = append(modMsgs, msg)
		}
		return
	}

	msgs = append(msgs, modAndReturn(parsedData.ChatFlower)...)
	msgs = append(msgs, modAndReturn(parsedData.Chat47)...)
	msgs = append(msgs, modAndReturn(parsedData.OneCRepair)...)
	msgs = append(msgs, modAndReturn(parsedData.OneCto)...)
	msgs = append(msgs, modAndReturn(parsedData.MailKuzov)...)
	msgs = append(msgs, modAndReturn(parsedData.MailService)...)

	return msgs
}
