package parser

import (
	"disp_bot/utils"
	"errors"
	"strings"
)

func strechChatParse(mark string, location string, message utils.Message) (
	res utils.Resource,
	unIdent utils.Message,
	err error) {

	if mark != "" && location != "" {
		res = utils.Resource{
			Loc:  location,
			Mess: message,
		}
	} else if location != "" {
		message.AddReply("Не распознан ГРЗ")
		unIdent = message
		err = errors.New("unident GR mark")
	} else if mark != "" {
		message.AddReply("Не распознан сервис")
		unIdent = message
		err = errors.New("unident service")
	}
	return res, unIdent, err
}

func (p *Parser) removeURL(strs []string) (editedStrs []string) {
	for _, s := range strs {
		if p.findURL(s) == false {
			editedStrs = append(editedStrs, s)
		}
	}
	return editedStrs
}

func (p *Parser) stretchesChat(messages []utils.Message) (
	resces map[string]utils.Resource, unIdents []utils.Message) {

	resces = make(map[string]utils.Resource, 10)
	for _, mess := range messages {
		strs := strings.Split(mess.Text, "\n")
		strs = p.removeURL(strs)
		if len(strs) > 3 {
			strs = strs[:3]
		}
		str := strings.Join(strs, "\n")

		mark := p.findMark(str)
		location := p.findLocation(str)
		res, unident, err := strechChatParse(mark, location, mess)
		if err == nil {
			resces[mark] = res
		} else {
			unIdents = append(unIdents, unident)
		}
	}
	return resces, unIdents
}
