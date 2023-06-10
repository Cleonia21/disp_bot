package parser

import (
	"encoding/json"
	"github.com/mymmrac/telego"
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

func (p *Parser) Init() {
	p.locationsRegExp = make(map[StateRegMark]string)
	patterns := p.locationsPatterns()
	p.fillLocationsRegExp(patterns)
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
func (p *Parser) StretchesChat(updates []telego.Update) (res Resource) {
	return p.stretchesChat(updates)
}

// 1C
func (p *Parser) OneC(updates []telego.Update) (res Resource) {
	return p.oneC(updates)
}

// любой телеграмм чат
func (p *Parser) AnyChat(updates []telego.Update) (res Resource) {
	return p.anyChat(updates, "")
}
