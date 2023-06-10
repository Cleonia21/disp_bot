package mail

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"io/ioutil"
	"log"
	"time"
)

type Mail struct {
	client *client.Client
}

func (m *Mail) Connect() {
	m.connect()
	m.login()
}

func (m *Mail) Close() {
	_ = m.client.Logout()
}

func (m *Mail) connect() {
	log.Println("Connecting to server...")

	var err error
	// Connect to server
	m.client, err = client.DialTLS(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")
}

func (m *Mail) login() {
	if err := m.client.Login(username, password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")
}

type Email struct {
	Whom    string
	Subject string
	Body    string
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

func (m *Mail) GetEmail() (email []Email) {
	// Select INBOX
	_, err := m.client.Select("INBOX", true)
	if err != nil {
		log.Fatal(err)
	}

	criterion := imap.NewSearchCriteria()
	criterion.Since = m.getStartShiftTime(time.Now())

	ids, err := m.client.Search(criterion)
	if err != nil {
		log.Fatal(err)
	}

	if len(ids) > 0 {
		seqSet := new(imap.SeqSet)
		seqSet.AddNum(ids...)
		messages := make(chan *imap.Message, 10)
		done := make(chan error, 1)

		// Get the whole message body
		var section imap.BodySectionName
		items := []imap.FetchItem{section.FetchItem()}

		go func() {
			done <- m.client.Fetch(seqSet, items, messages) //[]imap.FetchItem{imap.FetchEnvelope}
		}()

		log.Println("Messages for date:")
		for msg := range messages {
			r := msg.GetBody(&section)
			if r == nil {
				log.Fatal("Server didn't returned message body")
			}

			// Create a new mail reader
			reader, err := mail.CreateReader(r)
			if err != nil {
				log.Fatal(err)
			}

			newEmail := Email{}

			// Print some info about the message
			header := reader.Header
			if subject, err := header.Subject(); err == nil {
				//log.Println("Subject:", subject)
				newEmail.Subject = subject
			}
			// Process each message's part
			for {
				p, err := reader.NextPart()
				if err == io.EOF {
					break
				} else if err != nil {
					log.Fatal(err)
				}

				switch p.Header.(type) {
				case *mail.InlineHeader:
					// This is the message's text (can be plain-text or HTML)
					b, _ := ioutil.ReadAll(p.Body)
					newEmail.Body = string(b)
					//log.Printf("Got text: %v", string(b))
				}
			}
			email = append(email, newEmail)
		}
		if err := <-done; err != nil {
			log.Fatal(err)
		}
	}
	return email
}
