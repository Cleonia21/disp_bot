package parser

import "regexp"

func (p *Parser) findRegMark(s string) StateRegMark {
	var re = regexp.MustCompile(p.stateMarkRegExp)
	matches := re.FindStringSubmatch(s)
	if matches != nil {
		return StateRegMark(matches[0])
	}
	return ""
}

func (p *Parser) findLocation(s string) Location {
	for location, regExpStr := range p.locationsRegExp {
		re := regexp.MustCompile(regExpStr)
		matches := re.FindStringSubmatch(s)
		if matches != nil {
			return Location(location)
		}
	}
	return ""
}
