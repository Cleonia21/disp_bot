package parser

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func (p *Parser) findRegMark(s string) string {
	var re = regexp.MustCompile(p.stateMarkRegExp)
	matches := re.FindStringSubmatch(s)
	if matches != nil {
		return matches[0]
	}
	return ""
}

func (p *Parser) findRegMarks(s string) (marks []string) {
	strs := strings.Split(s, "\n")
	for _, str := range strs {
		if mark := p.findRegMark(str); mark != "" {
			marks = append(marks, mark)
		}
	}
	return marks
}

func (p *Parser) findTO(s string) string {
	var re = regexp.MustCompile(p.techServiceRegExp)
	matches := re.FindStringSubmatch(s)
	if matches != nil {
		return matches[0]
	}
	return ""
}

func (p *Parser) findLocation(s string) string {
	for location, regExpStr := range p.locationsRegExp {
		re := regexp.MustCompile(regExpStr)
		matches := re.FindStringSubmatch(s)
		if matches != nil {
			return string(location)
		}
	}
	return ""
}

func stateMarkRegExp() string {
	return `(?i)[авекмнорстух]\d{3}[авекмнорстух]{2}\d{2,3}`
}

func techServiceRegExp() string {
	return "Плановое ТО"
}

func locationsRegExp() (locationsRegExp map[string]string) {
	locationsRegExp = make(map[string]string)
	patterns := locationsPatterns()

	for regMark, locations := range patterns {
		var regStrs []string
		for _, location := range locations {
			regStrs = append(regStrs, "("+location+")")
		}
		locationsRegExp[regMark] = strings.Join(regStrs, "|")
	}
	return
}

func locationsPatterns() (patterns map[string][]string) {
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
