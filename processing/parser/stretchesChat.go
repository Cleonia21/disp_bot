package parser

import (
	"disp_bot/utils"
	"strings"
)

func (p *Parser) stretchesChat(messages []utils.Message) (
	resces map[string]utils.Message, undef []utils.Message) {

	resces = make(map[string]utils.Message, 10)
	for _, msg := range messages {
		str := p.removeURL(msg.Text)
		mark := p.findMark(str)
		location := p.findLocation(str)
		replyText := identify(mark, location)
		tmpMsg := msg
		tmpMsg.Loc = location
		tmpMsg.Mark = mark
		if replyText == "" {
			resces[mark] = tmpMsg
		} else {
			tmpMsg.AddReply(replyText)
			undef = append(undef, tmpMsg)
		}
	}
	return resces, undef
}

func removeUnprocPart(str string) (procPart string) {
	strs := strings.Split(str, "\n")
	if len(strs) > 3 {
		strs = strs[:3]
	}
	procPart = strings.Join(strs, "\n")
	return procPart
}

func (p *Parser) removeURL(str string) (editedStr string) {
	strs := strings.Split(str, "\n")
	var bufStrs []string
	for _, s := range strs {
		if p.findURL(s) == false {
			bufStrs = append(bufStrs, s)
		}
	}
	editedStr = strings.Join(bufStrs, "\n")
	return editedStr
}

func identify(mark string, loc string) (replyText string) {
	if loc != "" && mark == "" {
		replyText = "Не распознан ГРЗ"
	} else if loc == "" && mark != "" {
		replyText = "Не распознан сервис"
	} else if loc == "" && mark == "" {
		replyText = "Сообщение не распознано"
	}
	return replyText
}
