package parser

import (
	proc "disp_bot/processing"
	"disp_bot/telegram"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

type StateRegMark string
type Location string

type Resource map[StateRegMark]Location

type Parser struct {
	locationsRegExp map[StateRegMark]string
	stateMarkRegExp string
}

func Init() *Parser {
	p := &Parser{}
	p.locationsRegExp = make(map[StateRegMark]string)
	patterns := p.locationsPatterns()
	p.fillLocationsRegExp(patterns)
	return p
}

func (p *Parser) Parse(data proc.UnParsedData) proc.ParsedData {
	parsedData := proc.ParsedData{
		Chat47:        p.AnyChat(data.Chat47, ""),
		ChatFlower:    p.AnyChat(data.ChatFlower, ""),
		ChatStretches: p.stretchesChat(data.ChatStretches),
		OneC:          p.oneC(data.OneC),
		Mail:          p.Mail(),
	}
	return parsedData
}

func (p *Parser) fillStateMarkRegExp() {
	p.stateMarkRegExp = `(?i)[авекмнорстух]\d{3}[авекмнорстух]{2}\d{2,3}`
}

func (p *Parser) fillLocationsRegExp(patterns map[string][]string) {
	for regMark, locations := range patterns {
		var regStrs []string
		for _, location := range locations {
			regStrs = append(regStrs, "("+location+")")
		}
		p.locationsRegExp[StateRegMark(regMark)] = strings.Join(regStrs, "|")
	}
}

func (p *Parser) locationsPatterns() (patterns map[string][]string) {
	patterns = make(map[string][]string)
	file, err := os.Open("locations.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteArray, _ := io.ReadAll(file)

	err = json.Unmarshal(byteArray, &patterns)
	if err != nil {
		log.Fatal(err)
	}
	return patterns
}

func (p *Parser) Mail() (res Resource) {
	return p.mail()
}

// Перегоны
func (p *Parser) StretchesChat(mess []telegram.Message) (res Resource) {
	return p.stretchesChat(mess)
}

// 1C
func (p *Parser) OneC(mess []telegram.Message) (res Resource) {
	return p.oneC(mess)
}

// любой телеграмм чат
func (p *Parser) AnyChat(mess []telegram.Message, loc Location) (res Resource) {
	return p.anyChat(mess, loc)
}
