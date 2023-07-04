package parser

import (
	"disp_bot/utils"
	"reflect"
	"testing"
)

func TestParser_anyChat(t *testing.T) {
	type args struct {
		messages []utils.Message
		location string
	}
	tests := []struct {
		name     string
		args     args
		wantMsgs map[string]utils.Message
	}{
		{
			name: "",
			args: args{
				messages: []utils.Message{
					{
						ID:   1,
						Text: "пригнали\nх543мо797\nу342ее790",
					},
					{
						ID:   2,
						Text: "Добрый вечер!\nпригнали\nо642нс790",
					},
					{
						ID:   3,
						Text: "р285нр790 готова износ передних тормозных колодок и дисков 100%(1.1мм /19.8мм)",
					},
				},
				location: "47",
			},
			wantMsgs: map[string]utils.Message{
				"х543мо797": {
					ID:   1,
					Text: "пригнали\nх543мо797\nу342ее790",
					Loc:  "47",
					Mark: "х543мо797",
				},
				"у342ее790": {
					ID:   1,
					Text: "пригнали\nх543мо797\nу342ее790",
					Loc:  "47",
					Mark: "у342ее790",
				},
				"о642нс790": {
					ID:   2,
					Text: "Добрый вечер!\nпригнали\nо642нс790",
					Loc:  "47",
					Mark: "о642нс790",
				},
				"р285нр790": {
					ID:   3,
					Text: "р285нр790 готова износ передних тормозных колодок и дисков 100%(1.1мм /19.8мм)",
					Loc:  "47",
					Mark: "р285нр790",
				},
			},
		},
	}

	p := Init()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMsgs := p.anyChat(tt.args.messages, tt.args.location); !reflect.DeepEqual(gotMsgs, tt.wantMsgs) {
				t.Errorf("anyChat() = %v, want %v", gotMsgs, tt.wantMsgs)
			}
		})
	}
}
