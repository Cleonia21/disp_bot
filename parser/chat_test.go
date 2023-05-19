package parser

import (
	"reflect"
	"testing"
)

func TestChat(t *testing.T) {
	tests := []struct {
		name      string
		mess      []string
		wantAutos map[string]struct{}
	}{
		{
			"",
			[]string{
				"у543сх790\n" +
					"е869со799\n" +
					"к242рн799\n" +
					"к242рн799",
				"пригнали а670ме797\n" +
					"а670ме798\n" +
					"а670ме799",
				"+",
				"к242рн700 пригнали",
				"lerjlevefkv",
				"11123132312332csdjcwdwwc",
				"",
			},
			map[string]struct{}{
				"у543сх790": {},
				"е869со799": {},
				"к242рн799": {},
				"а670ме797": {},
				"а670ме798": {},
				"а670ме799": {},
				"к242рн700": {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAutos := Chat(tt.mess); !reflect.DeepEqual(gotAutos, tt.wantAutos) {
				t.Errorf("Chat() = %v, want %v", gotAutos, tt.wantAutos)
			}
		})
	}
}
