package parser

import (
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
