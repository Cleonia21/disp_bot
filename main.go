package main

import (
	proc "disp_bot/processing"
	"disp_bot/processing/parser"
)

//update *telego.Update

func main() {
	p := parser.Init()
	p.Parse(proc.UnParsedData{
		Chat47:        nil,
		ChatFlower:    nil,
		OneC:          nil,
		ChatStretches: nil,
	})
}
