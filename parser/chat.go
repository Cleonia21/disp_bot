package parser

import "strings"

/*
Chat обрабатывает массив сообщений из чатов сервисов
возвращая map[грз]пустота
*/
func Chat(mess []string) (autos map[string]struct{}) {
	autos = make(map[string]struct{})
	for _, val := range mess {
		strs := strings.Split(val, "\n")
		for _, str := range strs {
			carNum := findMark(str)
			if carNum != "" {
				autos[carNum] = struct{}{}
			}
		}
	}
	return
}
