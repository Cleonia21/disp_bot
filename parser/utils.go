package parser

import "regexp"

func findMark(s string) string {
	var re = regexp.MustCompile(`(?i)[авекмнорстух]\d{3}[авекмнорстух]{2}\d{2,3}`)
	matches := re.FindStringSubmatch(s)
	if matches != nil {
		return matches[0]
	}
	return ""
}

func findService(s string) string {
	var re = regexp.MustCompile(`(?i)(цкр)|(авторусь)`)
	matches := re.FindStringSubmatch(s)
	if matches != nil {
		return matches[0]
	}
	return ""
}

func findTypeTM(s string) bool {
	var re = regexp.MustCompile(`(?i)(Плановое ТО)`)
	matches := re.FindStringSubmatch(s)
	if matches != nil {
		return true
	}
	return false
}
