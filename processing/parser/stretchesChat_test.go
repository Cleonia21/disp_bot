package parser

import (
	"disp_bot/utils"
	"reflect"
	"testing"
)

func TestParser_stretchesChat(t *testing.T) {
	p := Init()
	type args struct {
		messages []utils.Message
	}
	tests := []struct {
		name         string
		args         args
		wantResces   map[string]utils.Message
		wantUnIdents []utils.Message
	}{
		{
			name: "",
			args: args{[]utils.Message{
				{
					Text: "https://ticket.belkacar.ru/ticket/NWOPA9\nс688но790\nКазаков Моторс",
					ID:   1,
				},
				{
					Text: "https://ticket.belkacar.ru/ticket/TOBUQC\nе620сн790\nДженерал",
					ID:   2,
				},
				{
					Text: "https://ticket.belkacar.ru/ticket/12198802\nр353мх797\nЦкр-авто",
					ID:   3,
				},
				{
					Text: "https://ticket.belkacar.ru/ticket/12205928\nр835мм797\nВ офис (после 0:00)\nНовогиреево",
					ID:   4,
				},
				{
					Text: "https://ticket.belkacar.ru/ticket/AJZNQP\nв412нс790\nПоречная\nНа линии",
					ID:   5,
				},
				{
					Text: "===",
					ID:   6,
				},
			}},
			wantResces: map[string]utils.Message{
				"с688но790": {
					Loc:  "казаков",
					Text: "https://ticket.belkacar.ru/ticket/NWOPA9\nс688но790\nКазаков Моторс",
					ID:   1,
					Mark: "с688но790",
				},
				"е620сн790": {
					Loc:  "дженерал",
					Mark: "е620сн790",
					Text: "https://ticket.belkacar.ru/ticket/TOBUQC\nе620сн790\nДженерал",
					ID:   2,
				},
				"р353мх797": {
					Loc:  "цкр",
					Mark: "р353мх797",
					Text: "https://ticket.belkacar.ru/ticket/12198802\nр353мх797\nЦкр-авто",
					ID:   3,
				},
				"р835мм797": {
					Loc:  "офис",
					Mark: "р835мм797",
					Text: "https://ticket.belkacar.ru/ticket/12205928\nр835мм797\nВ офис (после 0:00)\nНовогиреево",
					ID:   4,
				},
				"в412нс790": {
					Loc:  "поречная",
					Mark: "в412нс790",
					Text: "https://ticket.belkacar.ru/ticket/AJZNQP\nв412нс790\nПоречная\nНа линии",
					ID:   5,
				},
			},
			wantUnIdents: []utils.Message{
				{
					Text:      "===",
					ID:        6,
					ReplyText: "Сообщение не распознано",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResces, gotUnIdents := p.stretchesChat(tt.args.messages)
			if !reflect.DeepEqual(gotResces, tt.wantResces) {
				t.Errorf("stretchesChat() gotResces = %v, want %v", gotResces, tt.wantResces)
			}
			if !reflect.DeepEqual(gotUnIdents, tt.wantUnIdents) {
				t.Errorf("stretchesChat() gotUnIdents = %v, want %v", gotUnIdents, tt.wantUnIdents)
			}
		})
	}
}

func Test_removeUnprocPart(t *testing.T) {
	tests := []struct {
		name         string
		str          string
		wantProcPart string
	}{
		{
			name: "",
			str: "https://ticket.belkacar.ru/ticket/NCPH7A\n" +
				"е666кх777\n" +
				"на цветочный\n" +
				"из 47го",
			wantProcPart: "https://ticket.belkacar.ru/ticket/NCPH7A\n" +
				"е666кх777\n" +
				"на цветочный",
		},
		{
			name: "",
			str: "https://ticket.belkacar.ru/ticket/NCPH7A\n" +
				"е666кх777\n" +
				"на цветочный\n" +
				"из 47го" +
				"",
			wantProcPart: "https://ticket.belkacar.ru/ticket/NCPH7A\n" +
				"е666кх777\n" +
				"на цветочный",
		},
		{
			name: "",
			str: "https://ticket.belkacar.ru/ticket/NCPH7A\n" +
				"е666кх777\n" +
				"на цветочный\n" +
				"из 47го\n" +
				"лдвоидлв\n" +
				"ваолмраво\n" +
				"вамоавлд\n" +
				"8234832свы\n" +
				"ысвыасываэаэжывмсю.счмчсм*(*?:*\n" +
				"амвмва\n",
			wantProcPart: "https://ticket.belkacar.ru/ticket/NCPH7A\n" +
				"е666кх777\n" +
				"на цветочный",
		},
		{
			name: "",
			str: "https://ticket.belkacar.ru/ticket/NCPH7A\n" +
				"е666кх777\n",
			wantProcPart: "https://ticket.belkacar.ru/ticket/NCPH7A\n" +
				"е666кх777\n",
		},
		{
			name:         "",
			str:          "+",
			wantProcPart: "+",
		},
		{
			name:         "",
			str:          "",
			wantProcPart: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotProcPart := removeUnprocPart(tt.str); gotProcPart != tt.wantProcPart {
				t.Errorf("removeUnprocPart() = %v, want %v", gotProcPart, tt.wantProcPart)
			}
		})
	}
}

func TestParser_removeURL(t *testing.T) {
	p := Init()
	tests := []struct {
		name          string
		str           string
		wantEditedStr string
	}{
		{
			name: "",
			str: "1\n" +
				"sjvlksjlkcjklasdjlc43789r4398cjskdcs\n" +
				"http\n" +
				"https://ticket.belkacar.ru/ticket/NCPH7A",
			wantEditedStr: "1\n" +
				"sjvlksjlkcjklasdjlc43789r4398cjskdcs\n" +
				"http", // not \n
		},
		{
			name: "",
			str: "1\n" +
				"sjvlksjlkcjklasdjlc43789r4398cjskdcs\n" +
				"http\n" +
				"ehrgjkherkjgkjergjkerhttps://ticket.belkacar.ru/ticket/NCPH7A\n" +
				"jhnrfvgkjldnkjvkdf\n",
			wantEditedStr: "1\n" +
				"sjvlksjlkcjklasdjlc43789r4398cjskdcs\n" +
				"http\n" +
				"jhnrfvgkjldnkjvkdf\n",
		},
		{
			name: "",
			str: "ehrgjkherkjgkjergjkerhttps://ticket.belkacar.ru/ticket/NCPH7A\n" +
				"1\n" +
				"sjvlksjlkcjklasdjlc43789r4398cjskdcs\n" +
				"http\n" +
				"jhnrfvgkjldnkjvkdf\n",
			wantEditedStr: "1\n" +
				"sjvlksjlkcjklasdjlc43789r4398cjskdcs\n" +
				"http\n" +
				"jhnrfvgkjldnkjvkdf\n",
		},
		{
			name:          "",
			str:           "+",
			wantEditedStr: "+",
		},
		{
			name:          "",
			str:           "",
			wantEditedStr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEditedStr := p.removeURL(tt.str); gotEditedStr != tt.wantEditedStr {
				t.Errorf("removeURL() = %v, want %v", gotEditedStr, tt.wantEditedStr)
			}
		})
	}
}
