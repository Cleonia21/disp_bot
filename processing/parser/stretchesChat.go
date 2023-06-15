package parser

import (
	"disp_bot/utils"
	"errors"
	"strings"
)

func strechChatParse(mark string, location string, message utils.Message) (
	res utils.Resource,
	unident utils.Message,
	err error) {

	if mark != "" && location != "" {
		res = utils.Resource{
			Loc:  location,
			Mess: message,
		}
	} else if location != "" {
		message.AddReply("Не распознан ГРЗ")
		unident = message
		err = errors.New("unident GR mark")
	} else if mark != "" {
		message.AddReply("Не распознан сервис")
		unident = message
		err = errors.New("unident service")
	}
	return res, unident, err
}

func (p *Parser) stretchesChat(messages []utils.Message) (
	resces map[string]utils.Resource, unidents []utils.Message) {

	resces = make(map[string]utils.Resource, 10)
	for _, mess := range messages {
		strs := strings.Split(mess.Text, "\n")
		if len(strs) > 3 {
			strs = strs[:3]
		}
		str := strings.Join(strs, "\n")

		mark := p.findRegMark(str)
		location := p.findLocation(str)
		res, unident, err := strechChatParse(mark, location, mess)
		if err == nil {
			resces[mark] = res
		} else {
			unidents = append(unidents, unident)
		}
	}
	return resces, unidents
}
