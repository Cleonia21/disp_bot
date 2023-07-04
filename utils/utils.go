package utils

import "fmt"

// inputIDs
const (
	ID_default = iota
	ID_47
	ID_flower
	ID_oneC
	ID_stretches
)

type Message struct {
	ID        int
	Text      string
	ReplyText string
	From      int

	Loc  string
	Mark string
}

func (m *Message) Print() {
	fmt.Printf("-----------------------\ntext:\n{%v}\nerror: {%v}\nfrom {%v}\nloc  {%v}\nmark {%v}\n-----------------------\n",
		m.Text, m.ReplyText, m.From, m.Loc, m.Mark)
}

func (m *Message) AddReply(text string) {
	m.ReplyText += text
}

type Conf struct {
	Chat47        bool
	ChatFlower    bool
	OneC          bool
	ChatStretches bool
	Mail          bool
}
