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
		name    string
		args    args
		wantRes map[string]utils.Resource
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
			wantRes: map[string]utils.Resource{
				"х543мо797": {
					Loc: "47",
					Mess: utils.Message{
						ID:   1,
						Text: "пригнали\nх543мо797\nу342ее790",
					},
				},
				"у342ее790": {
					Loc: "47",
					Mess: utils.Message{
						ID:   1,
						Text: "пригнали\nх543мо797\nу342ее790",
					},
				},
				"о642нс790": {
					Loc: "47",
					Mess: utils.Message{
						ID:   2,
						Text: "Добрый вечер!\nпригнали\nо642нс790",
					},
				},
				"р285нр790": {
					Loc: "47",
					Mess: utils.Message{
						ID:   3,
						Text: "р285нр790 готова износ передних тормозных колодок и дисков 100%(1.1мм /19.8мм)",
					},
				},
			},
		},
	}

	p := Init()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := p.anyChat(tt.args.messages, tt.args.location); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("anyChat() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
