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

func getMailParam() (mailParam mailParam) {
	file, err := os.Open("../../conf/mail.json")
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

func mailParse(mark string, location string, mess mailReader.EmailMessage) (
	res utils.Resource,
	unident utils.Message,
	err error) {

	if mark != "" && location != "" {
		res = utils.Resource{
			Loc:  location,
			Mess: utils.NewMessage(mess.Subject + "\n" + mess.Body),
		}
	} else if location != "" {
		unident = utils.NewMessage("Тема письма: " + mess.Subject + "\n" +
			"Тело письма: " + mess.Body + "\n" +
			"Не распознан ГРЗ\n")
		err = errors.New("unident message")
	} else if mark != "" {
		unident = utils.NewMessage("Тема письма: " + mess.Subject + "\n" +
			"Тело письма: " + mess.Body + "\n" +
			"Не распознан сервис\n")
		err = errors.New("unident message")
	}
	return res, unident, err
}

func (p *Parser) mail() (resces map[string]utils.Resource, unidents []utils.Message) {
	var Mail mailReader.Mail

	param := getMailParam()
	Mail.Connect(param.Addr, param.UserName, param.Password)
	defer Mail.Close()
	messages := Mail.GetEmailMessages()

	resces = make(map[string]utils.Resource, 10)
	for _, mess := range messages {
		mark := p.findMark(mess.Subject)
		location := p.findLocation(mess.Body)

		res, unident, err := mailParse(mark, location, mess)
		if err == nil {
			resces[mark] = res
		} else {
			unidents = append(unidents, unident)
		}
	}
	return resces, unidents
}
