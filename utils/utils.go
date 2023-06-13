package utils

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
}

func NewMessage(text string) Message {
	return Message{Text: text}
}

func (m *Message) AddReply(text string) {
	m.ReplyText += text
}

type ProcData struct {
	ID        int64
	MessPacks []Message
}

type UnProcData struct {
	ID        int64
	MessPacks map[int64][]Message
}

type Resource struct {
	StRegMark string
	Loc       string
	Analyzed  bool
	Mess      Message
}

type ParsedData struct {
	Chat47        []Resource
	ChatFlower    []Resource
	OneC          []Resource
	Mail          []Resource
	ChatStretches []Resource
}

type UnParsedData struct {
	Chat47        []Message
	ChatFlower    []Message
	OneC          []Message
	ChatStretches []Message
}
