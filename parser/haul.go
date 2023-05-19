package parser

import (
	"errors"
	"strings"
)

/*
Haul принимает массив сообщений из чатов перегонов,
где сообщение имеет формат

	url\n
	...ГРЗ...\n
	...куда сервис...
	откуда

возвращает map[грз]сервис и массив не распознаных сообщений
*/
func Haul(mess []string) (cars map[string]ServiceID, notProcessed []string) {
	cars = make(map[string]ServiceID)
	for _, val := range mess {
		// сырая строка, что там на самом деле не известно
		carNum, serv, err := getCarNumAndServ(val)
		if err != nil {
			notProcessed = append(notProcessed, val)
			continue
		}
		// получение валидных данных
		carNum = findMark(carNum)
		serv = findService(serv)
		if carNum == "" || serv == "" {
			notProcessed = append(notProcessed, val)
			continue
		}
		cars[carNum] = ServiceID(serv)
	}
	return
}

func getCarNumAndServ(mes string) (carNum string, serv string, err error) {
	strs := strings.Split(mes, "\n")
	// если сообщение мусорное("+", к примеру)
	if len(strs) < 2 {
		return carNum, serv, errors.New("сообщение не распознано")
	}
	carNum = strs[1]
	// если грз и сервис написали на одной строке("url \n грз сервис")
	if len(strs) == 2 {
		serv = strs[1]
	} else {
		serv = strs[2]
	}
	return
}
