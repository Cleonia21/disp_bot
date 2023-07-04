package parser

import (
	mailReader "disp_bot/processing/parser/mail"
	"disp_bot/utils"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
)

type mailParam struct {
	Addr     string `json:"Addr"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}

func getMailParam(confPath string) (mailParam mailParam) {
	file, err := os.Open(confPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteArray, _ := io.ReadAll(file)

	err = json.Unmarshal(byteArray, &mailParam)
	if err != nil {
		log.Fatal(err)
	}
	return mailParam
}

func mailParse(mark string, location string, mailMsg mailReader.EmailMessage) (msg utils.Message, err error) {
	msg = utils.Message{
		Text: "Тема письма: " + mailMsg.Subject + "\n" + "Тело письма: " + mailMsg.Body,
		Loc:  location,
		Mark: mark,
	}
	if location != "" && mark != "" {
		return
	}
	msg.AddReply("не распознаны грз или сервис")
	err = errors.New("unident message")
	return
}

func (p *Parser) mail(confPath string) (resces map[string]utils.Message, undef []utils.Message) {
	var Mail mailReader.Mail

	param := getMailParam(confPath)
	Mail.Connect(param.Addr, param.UserName, param.Password)
	defer Mail.Close()
	mailMsgs := Mail.GetEmailMessages()

	resces = make(map[string]utils.Message, 10)
	for _, mailMsg := range mailMsgs {
		mark := p.findMark(mailMsg.Subject)
		location := p.findLocation(mailMsg.Body)

		msg, err := mailParse(mark, location, mailMsg)
		if err == nil {
			resces[mark] = msg
		} else {
			undef = append(undef, msg)
		}
	}
	return resces, undef
}
