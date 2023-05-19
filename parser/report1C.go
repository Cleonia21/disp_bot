package parser

import (
	"strings"
)

type ServiceType struct {
	Service ServiceID
	// technical maintenance(true), others(false)
	TypeTM bool
}

/*
Report1C обрабытывает массив сообщений из отчета 1С,
сообщения могут перекрывать друг друга

	сервис\n
	грз...выполняемые работы

возвращает map[грз]{сервис, ТО}
нераспознаные строки пропускаются
*/
func Report1C(mess []string) (cars map[string]ServiceType) {
	cars = make(map[string]ServiceType)

	var strs []string
	for _, val := range mess {
		strs = append(strs, strings.Split(val, "\n")...)
	}

	var serv ServiceID
	for _, val := range strs {
		_, ok := gServices[val]
		if ok {
			serv = ServiceID(val)
			continue
		}
		carNum := findMark(val)
		if carNum == "" {
			continue
		}
		typeTM := findTypeTM(val)
		cars[carNum] = ServiceType{serv, typeTM}
	}
	return
}
