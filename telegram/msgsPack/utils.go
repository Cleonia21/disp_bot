package msgsPack

func (p *MsgsPack) mailButtonPrefixText() string {
	if p.conf.Mail == true {
		return "не обрабатывать " + idMail
	}
	return "обрабатывать " + idMail
}
