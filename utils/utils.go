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

type Resource struct {
	Loc  string
	Mess Message
}

type ParsedData struct {
	Chat47        map[string]Resource
	ChatFlower    map[string]Resource
	OneCto        map[string]Resource
	OneCRepair    map[string]Resource
	Mail          map[string]Resource
	ChatStretches map[string]Resource
	Unidentified  []Message
}

type UnParsedData struct {
	Chat47        []Message
	ChatFlower    []Message
	OneC          []Message
	ChatStretches []Message
}
