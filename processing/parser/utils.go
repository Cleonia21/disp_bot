package parser

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func (p *Parser) findMark(s string) string {
	matches := p.stateMarkRegExp.FindStringSubmatch(s)
	if matches != nil {
		return strings.ToLower(matches[0])
	}
	return ""
}

func (p *Parser) findMarks(s string) (marks []string) {
	matches := p.stateMarkRegExp.FindAllString(s, -1)
	for i := range matches {
		matches[i] = strings.ToLower(matches[i])
	}
	return matches
}

func (p *Parser) findTO(s string) bool {
	return p.techServiceRegExp.FindStringSubmatch(s) != nil
}

func (p *Parser) findLocation(s string) string {
	for location, regExp := range p.locationsRegExp {
		if regExp.MatchString(s) {
			return location
		}
	}
	return ""
}

func (p *Parser) findURL(s string) bool {
	return p.urlRegExp.MatchString(s)
}

func stateMarkRegExp() (*regexp.Regexp, error) {
	return regexp.Compile(`(?i)[авекмнорстух]\d{3}[авекмнорстух]{2}\d{2,3}`)
}

func techServiceRegExp() (*regexp.Regexp, error) {
	return regexp.Compile(`(?i)плановое то`)
}

func urlRegExp() (*regexp.Regexp, error) {
	return regexp.Compile(`(?i)(?:https?|ftp):\/\/[\n\S]+`)
}

func locationsRegExp() (locationsRegExp map[string]*regexp.Regexp, err error) {
	locationsRegExp = make(map[string]*regexp.Regexp)
	patterns := locationsPatterns()

	for regMark, locations := range patterns {
		body := strings.Join(locations, `|`)
		reg, err := regexp.Compile(`(?i)` + body)
		if err != nil {
			return locationsRegExp, err
		}
		locationsRegExp[regMark] = reg
	}
	return
}

func locationsPatterns() (patterns map[string][]string) {
	patterns = make(map[string][]string)
	file, err := os.Open("conf/locations.json")
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
