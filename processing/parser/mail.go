package parser

import (
	mailReader "disp_bot/processing/parser/mail"
	"disp_bot/telegram"
	"encoding/json"
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

func (p *Parser) mail() (res []Resource) {
	var Mail mailReader.Mail

	param := getMailParam()
	Mail.Connect(param.Addr, param.UserName, param.Password)
	defer Mail.Close()

	msgs := Mail.GetEmailMessages()

	res = make([]Resource, 10)
	for _, msg := range msgs {
		stateRegMark := p.findRegMark(msg.Subject)
		location := p.findLocation(msg.Body)
		if stateRegMark != "" && location != "" {
			res = append(res, Resource{
				StRegMark: stateRegMark,
				Loc:       location,
				analyzed:  true,
				mess:      telegram.NewMessage(msg.Subject + "\n" + msg.Body),
			})
		} else if stateRegMark == "" {
			res = append(res, Resource{
				analyzed: false,
				mess: telegram.NewMessage("Тема письма: " + msg.Subject + "\n" +
					"Тело письма: " + msg.Body + "\n" +
					"Не распознан ГРЗ\n"),
			})
		} else if location == "" {
			res = append(res, Resource{
				analyzed: false,
				mess: telegram.NewMessage("Тема письма: " + msg.Subject + "\n" +
					"Тело письма: " + msg.Body + "\n" +
					"Не распознан сервис\n"),
			})
		}
	}
	return res
}
