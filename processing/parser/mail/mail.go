package mail

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"golang.org/x/net/html"
	"io"
	"log"
	"strings"
	"time"
)

type Mail struct {
	client *client.Client
}

func (m *Mail) Connect(addr string, username string, password string) {
	m.connect(addr)
	m.login(username, password)
}

func (m *Mail) Close() {
	_ = m.client.Logout()
}

func (m *Mail) connect(addr string) {
	log.Println("Connecting to server...")

	var err error
	// Connect to server
	m.client, err = client.DialTLS(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")
}

func (m *Mail) login(username string, password string) {
	if err := m.client.Login(username, password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")
}

type EmailMessage struct {
	Whom    string
	Subject string
	Body    string
}

func (m *Mail) selectedMailBox() {
	// Select INBOX
	_, err := m.client.Select("INBOX", true)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Mail) getEmailIds() (ids []uint32) {
	criterion := imap.NewSearchCriteria()
	criterion.Since = m.getStartShiftTime(time.Now())

	ids, err := m.client.Search(criterion)
	if err != nil {
		log.Fatal(err)
	}
	return ids
}

func (m *Mail) getStartShiftTime(t time.Time) time.Time {
	hour := time.Duration(t.Hour())
	startShift := t.Add((19 - hour) * time.Hour)
	if hour > 19 {
		return startShift
	} else {
		return startShift.AddDate(0, 0, -1)
	}
}

func (m *Mail) getEmailSeqSet(ids []uint32) (seqSet *imap.SeqSet) {
	seqSet = new(imap.SeqSet)
	seqSet.AddNum(ids...)
	return seqSet
}

func (m *Mail) htmlToString(htmlString string) string {
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		panic(err)
	}

	var clearText string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			clearText += n.Data
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
			if n.Type == html.ElementNode && n.Data == "p" {
				clearText += "\n"
			}
		}
	}
	f(doc)
	return clearText
}

func (m *Mail) emailMsgBody(reader *mail.Reader) string {
	msgBody := ""
	for {
		p, err := reader.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		switch p.Header.(type) {
		case *mail.InlineHeader:
			b, _ := io.ReadAll(p.Body)
			msgBody += m.htmlToString(string(b))
		}
	}
	return msgBody
}

func (m *Mail) emailMsgAddresses(reader *mail.Reader) string {
	header := reader.Header
	addresses, err := header.AddressList("Sender")
	if err != nil {
		log.Fatal(err)
	}
	addressesString := ""
	for _, address := range addresses {
		addressesString += address.String()
	}
	return addressesString
}

func (m *Mail) emailMsgSubject(reader *mail.Reader) string {
	header := reader.Header
	subject, err := header.Subject()
	if err != nil {
		log.Fatal(err)
	}
	return subject
}

func (m *Mail) processMsg(msg *imap.Message, section *imap.BodySectionName) (emailMsg EmailMessage) {
	r := msg.GetBody(section)
	if r == nil {
		log.Fatal("Server didn't returned message body")
	}

	reader, err := mail.CreateReader(r)
	if err != nil {
		log.Fatal(err)
	}

	emailMsg.Subject = m.emailMsgSubject(reader)
	emailMsg.Whom = m.emailMsgAddresses(reader)
	emailMsg.Body = m.emailMsgBody(reader)
	return emailMsg
}

func (m *Mail) GetEmailMessages() (emailMsgs []EmailMessage) {
	m.selectedMailBox()
	ids := m.getEmailIds()

	if len(ids) > 0 {
		seqSet := m.getEmailSeqSet(ids)

		messages := make(chan *imap.Message, 10)
		done := make(chan error, 1)

		var section imap.BodySectionName
		items := []imap.FetchItem{section.FetchItem()}

		go func() {
			done <- m.client.Fetch(seqSet, items, messages)
		}()

		//log.Println("Messages for date:")
		for msg := range messages {
			emailMsgs = append(emailMsgs, m.processMsg(msg, &section))
		}
		if err := <-done; err != nil {
			log.Fatal(err)
		}
	}
	return emailMsgs
}
