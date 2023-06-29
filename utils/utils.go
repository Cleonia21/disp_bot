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
	From      int
}

func NewMessage(text string) Message {
	return Message{Text: text}
}

func (m *Message) AddReply(text string) {
	m.ReplyText += text
}

type ProcData struct {
	MessPacks []Message
}

type UnProcData struct {
	MessPacks map[int64][]Message
	Checks    Checks
}

type Checks struct {
	Chat47        bool
	ChatFlower    bool
	OneC          bool
	ChatStretches bool
	Mail          bool
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
	Checks        Checks
}
