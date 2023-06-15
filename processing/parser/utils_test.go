package parser

import (
	"reflect"
	"testing"
)

func TestParser_findLocation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"",
			args{"lfndvd кулцк"},
			"",
		},
		{
			"",
			args{"кула"},
			"",
		},
		{
			"",
			args{"Есть же тикет для Казакова."},
			"казаков",
		},
		{
			"",
			args{"Максим Тарасов, [14.06.2023 23:57]\nhttps://ticket.belkacar.ru/ticket/OIIZY8\nт417ру790\nКулак"},
			"кулак",
		},
		{
			"",
			args{"Белугин Денис, [14.06.2023 23:39]\nhttps://ticket.belkacar.ru/ticket/XPRX06\nу286нс790 в ЦКР"},
			"цкр",
		},
		{
			"",
			args{"Станислав Качусов, [14.06.2023 22:56]\nhttps://ticket.belkacar.ru/ticket/12228657\nв972рн799\nВ авторемонт плюс\nодинцово"},
			"авторемонт плюс",
		},
		{
			"",
			args{"Белугин Денис, [14.06.2023 22:49]\nhttps://ticket.belkacar.ru/ticket/FYUWHO\nм442ек790 В Казаков"},
			"казаков",
		},
		{
			"",
			args{"Александр Чмирь, [14.06.2023 22:46]\nhttps://ticket.belkacar.ru/ticket/12229756\nр785нс790\nНа первую магистральную\nнапротив тойоты"},
			"магистральная",
		},
		{
			"in url",
			args{"Денис Коробков, [14.06.2023 22:38]\nhttps://ticket.belkacar.ru/ticket/1152775\nр467рм799\nВ казаков"},
			"казаков",
		},
		{
			"",
			args{"Валерия Кретова🧸, [14.06.2023 22:20]\nhttps://ticket.belkacar.ru/ticket/F0CMIL\nа366св790\nНа шикану\nАренда"},
			"шикана",
		},
		{
			"",
			args{"Руслан Залдя, [14.06.2023 22:12]\nhttps://ticket.belkacar.ru/ticket/GCHWBM\nм363ек790\nна шикану"},
			"шикана",
		},
		{
			"",
			args{"Дмитрий Овчинников, [14.06.2023 21:55]\nhttps://ticket.belkacar.ru/ticket/GXCWVM\nк385нт790\nВ казаков"},
			"казаков",
		},
		{
			"",
			args{"Сергей Яшин, [14.06.2023 21:44]\nhttps://ticket.belkacar.ru/ticket/QLPP3U\nв213нт790\nКулак"},
			"кулак",
		},
		{
			"",
			args{"Кирилл Ураков, [14.06.2023 21:39]\nhttps://ticket.belkacar.ru/ticket/6J3GUE\nк115нн797\nВ казаков"},
			"казаков",
		},
		{
			"",
			args{"Андрей Вьюсов, [14.06.2023 18:15]\n==========Смена Вьюсова Андрея========="},
			"",
		},
		{
			"",
			args{"\nАндрей Вьюсов, [14.06.2023 07:42]\nЧеклист"},
			"",
		},
		{
			"",
			args{"Арсений Зайцев, [14.06.2023 07:34]\nhttps://ticket.belkacar.ru/ticket/D7IS7O\nк866рн799\nТо обруч"},
			"обручева",
		},
		{
			"",
			args{"Андрей Вьюсов, [14.06.2023 07:17]\nЧеклист"},
			"",
		},
		{
			"",
			args{"Виттория Гройс, [14.06.2023 07:06]\nhttps://ticket.belkacar.ru/ticket/4NW6GO\nр984ео790\nОфис"},
			"офис",
		},
		{
			"",
			args{"Виттория Гройс, [14.06.2023 07:06]\nhttps://ticket.belkacar.ru/ticket/4NW6GO\nр984ео790\nОфис"},
			"офис",
		},
		{
			"",
			args{"Александр Чмирь, [14.06.2023 06:46]\nhttps://ticket.belkacar.ru/ticket/12181379\nк897нт790\nНа цкр"},
			"цкр",
		},
		{
			"",
			args{"Кирилл Ураков, [14.06.2023 06:41]\nhttps://ticket.belkacar.ru/ticket/8LWYEC\nр313ме797\nВ Офис"},
			"офис",
		},
		{
			"",
			args{"Дмитрий Овчинников, [14.06.2023 06:36]\nhttps://ticket.belkacar.ru/ticket/B7PIMX\nр509ср799\nНа Волжский"},
			"волжский",
		},
		{
			"",
			args{"Валерия Кретова🧸, [14.06.2023 06:29]\nhttps://ticket.belkacar.ru/ticket/12226759\nа720мх797\nМойка + заправить\nВ офис\nЯмонтово"},
			"офис",
		},
		{
			"",
			args{"Руслан Залдя, [14.06.2023 06:28]\nhttps://ticket.belkacar.ru/ticket/12221868\nт198вх790\nШикана"},
			"шикана",
		},
		{
			"",
			args{"+"},
			"",
		},
	}
	p := Init()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.findLocation(tt.args.s); got != tt.want {
				t.Errorf("findLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_findMark(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"",
			args{""},
			"",
		},
		{
			"",
			args{"У999УУ177"},
			"у999уу177",
		},
		{
			"",
			args{"о000мх17"},
			"о000мх17",
		},
		{
			"",
			args{"klervelrУ999УУ177kevekrve4-30543//dw\\"},
			"у999уу177",
		},
		{
			"",
			args{"Белугин Денис, [14.06.2023 23:39]\nhttps://ticket.belkacar.ru/ticket/XPRX06\nу286нс790 в ЦКР"},
			"у286нс790",
		},
		{
			"",
			args{"435234"},
			"",
		},
		{
			"",
			args{"о000мх1в"},
			"",
		},
		{
			"",
			args{"о000мсх17"},
			"",
		},
		{
			"",
			args{"о0030сх17"},
			"",
		},
		{
			"",
			args{"000сх17"},
			"",
		},
	}
	p := Init()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.findMark(tt.args.s); got != tt.want {
				t.Errorf("findMark() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_findMarks(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		wantMarks []string
	}{
		{
			"",
			args{""},
			nil,
		},
		{
			"",
			args{"У999УУ177"},
			[]string{"у999уу177"},
		},
		{
			"",
			args{"о000мх17"},
			[]string{"о000мх17"},
		},
		{
			"",
			args{"klervelrУ999УУ177kevekrve4-30543//dw\\"},
			[]string{"у999уу177"},
		},
		{
			"",
			args{"Белугин Денис, [14.06.2023 23:39]\nhttps://ticket.belkacar.ru/ticket/XPRX06\nу286нс790 в ЦКР"},
			[]string{"у286нс790"},
		},
		{
			"",
			args{"435234"},
			nil,
		},
		{
			"",
			args{"о000мх1в"},
			nil,
		},
		{
			"",
			args{"о000мсх17"},
			nil,
		},
		{
			"",
			args{"о0030сх17"},
			nil,
		},
		{
			"",
			args{"000сх17"},
			nil,
		},
		{
			"",
			args{"Белугин Ду286нс790енис, [14.06.2023 23:39]\nhttps://tickу286нс791et.belkacar.ru/ticket/XPRX06\nу286нс792 в ЦКР"},
			[]string{"у286нс790", "у286нс791", "у286нс792"},
		},
		{
			"",
			args{"пригнали\nс706нх790\nр727еа790\nх330нс790"},
			[]string{"с706нх790", "р727еа790", "х330нс790"},
		},
	}
	p := Init()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMarks := p.findMarks(tt.args.s); !reflect.DeepEqual(gotMarks, tt.wantMarks) {
				t.Errorf("findMarks() = %v, want %v", gotMarks, tt.wantMarks)
			}
		})
	}
}

func TestParser_findTO(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{";р253нк790/Z94C251ABLR117963/Rio X-Line/123 374/Плановое ТО;"},
			true,
		},
		{
			"",
			args{"Плановое ТО"},
			true,
		},
		{
			"",
			args{";р253нк790/Z94C251ABLR117963/Rio X-Line/123 374/Плановое "},
			false,
		},
		{
			"",
			args{"роиамолулдм"},
			false,
		},
		{
			"",
			args{""},
			false,
		},
	}
	p := Init()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.findTO(tt.args.s); got != tt.want {
				t.Errorf("findTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_findURL(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{"Белугин Денис, [14.06.2023 23:39]\nhttps://ticket.belkacar.ru/ticket/XPRX06\nу286нс790 в ЦКР"},
			true,
		},
		{
			"",
			args{"https://ticket.belkacar.ru/ticket/XPRX06"},
			true,
		},
		{
			"",
			args{""},
			false,
		},
	}
	p := Init()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.findURL(tt.args.s); got != tt.want {
				t.Errorf("findURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
