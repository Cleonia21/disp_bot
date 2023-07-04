package processing

import (
	"disp_bot/utils"
	"reflect"
	"testing"
)

func TestProc_Processing2(t *testing.T) {
	type args struct {
		msgs []utils.Message
		conf utils.Conf
	}

	type testStruct struct {
		name         string
		args         args
		wantRespMsgs []utils.Message
	}

	textFlower := []string{
		"# х554мо797 закрыт на линии",
		"# о145вх790 закрыт на линии",
		"# о642нс790 закрыт на линии",
		"# а164мн797 закрыт на линии",
		"пригнали\nр092св790",
		"Рано пригнали,от ТО проехала 5 тыс км",
		"пригнали\nк348рн799\nо309мс797",
		"пригнали\nх232но797",
		"пригнали\nн998мх797",
		"пригнали\nе640нт790",
		"пригнали\nх543мо797\nу342ее790",
		"пригнали\nк095рн799",
		"пригнали\nх592ем790",
		"ошибочно вам отправили",
		"272 точно к нам, показывает что она на 47 км",
		"пригнали\nв590се790",
		"пригнали\nх554мо797",
		"пригнали\nа164мн797\nо145вх790",
		"Добрый вечер!\nпригнали\nо642нс790",
	}
	text47 := []string{
		"пригнали\nк083хс750\nт208ем790",
		"О116ВХ790 готова",
		"т172нх790 готова",
		"пригнали\nк409рн799",
		"пригнали\nу346ее790",
		"р345рм799 готова",
		"р285нр790 готова износ передних тормозных колодок и дисков 100%(1.1мм /19.8мм)",
		"е968со799 готова",
		"е415ео790 готова",
		"пригнали\nх370но790\nо116вх790",
		"пригнали\nр345рм799\nт172нх790",
		"к272рн799 готова",
		"у742еа790 готова",
		"м665еа790 готова",
		"е930нс797 готова",
		"пригнали\nе968со799\nр285нр790",
		"р123ст799 готова",
		"т994хв750 готова",
		"пригнали\nе415ео790",
		"в785рн799 готова",
		"к302еу790 готова",
		"+",
		"е286нт790 готова",
		"Нет заявки",
		"пригнали\nк272рн799",
		"пригнали\nм665еа790",
		"пригнали\nу742еа790",
		"Спасибо",
		"выполнять то-8",
		"мин",
		"Коллеги подскажите пожалуйста",
		"е930нс797 Подскажите какое То выполнять?",
		"пригнали\nт994хв750\nе930нс797\nр123ст799",
		"пригнали\nв785рн799",
		"е979ре799 готова",
		"а607ме797 готова",
		"пригнали\nк302еу790\nе286нт790\nх515нн790",
		"пригнали\nа607ме797",
		"Добрый вечер.",
		"Добрый вечер!\nпригнали\nе979ре799",
	}
	textStretches := []string{
		"https://ticket.belkacar.ru/ticket/PV1YAH\nе979ре799\nАвторемонт Плюс",
		"https://ticket.belkacar.ru/ticket/CIJ8PL\nм438ек790\nКазаков Моторс",
		"https://ticket.belkacar.ru/ticket/S5HBAX\nр133ст799\nШикана",
		"https://ticket.belkacar.ru/ticket/12204884\nв162мв797\nКазаков Моторс",
		"https://ticket.belkacar.ru/ticket/11929465\nт271мо797\nКазаков Моторс",
		"https://ticket.belkacar.ru/ticket/WSS8DF\nу401нн797\n1-я Магистральная",
		"https://ticket.belkacar.ru/ticket/12138044\nр092св790\nЦветочный ТО",
		"BelkaCar\nhttps://ticket.belkacar.ru/ticket/12142308\nх232но797\nТо цветочный",
		"https://ticket.belkacar.ru/ticket/0OJRZN\nе640нт790\nТО Цветочный",
		"https://ticket.belkacar.ru/ticket/11808863\nх543мо797\nТо цветочный",
		"https://ticket.belkacar.ru/ticket/XUCPYD\nу342ее790\nТО Цветочный",
		"https://ticket.belkacar.ru/ticket/GKASDO\nу789ев790\nКазаков Моторс",
		"https://ticket.belkacar.ru/ticket/QVDMRX\nт190мк797\nКазаков Моторс",
		"https://ticket.belkacar.ru/ticket/12171124\nк242рн799\nШикана",
		"https://ticket.belkacar.ru/ticket/ZXZG11\nк030хс750\nШикана",
		"https://ticket.belkacar.ru/ticket/12148653\nв785рн799\n47км",
		"https://ticket.belkacar.ru/ticket/12195079\nх344ев790\nЦкр-авто",
		"https://ticket.belkacar.ru/ticket/12187213\nа341са790\nКазак",
		"https://ticket.belkacar.ru/ticket/NWOPA9\nс688но790\nКазаков Моторс",
		"+",
		"https://ticket.belkacar.ru/ticket/12202405\nк945нт790\nКулак",
		"https://ticket.belkacar.ru/ticket/12141209\nе412нн797\n26191\nДубровка",
		"https://ticket.belkacar.ru/ticket/12207230\nт208ем790\n127012\nТо47",
		"https://ticket.belkacar.ru/ticket/12207957\nк083хс750\n152466\nТо 47",
		"https://ticket.belkacar.ru/ticket/LLY3TQ\nх634нн797\n21757\nДубровка",
		"https://ticket.belkacar.ru/ticket/12198262\nе792но797\n26053\nТо Дубровка",
		"Вернула",
		"Едет обратно",
		"https://ticket.belkacar.ru/ticket/12085538\nм060мн797\n7560\nТо Дубровка",
		"https://ticket.belkacar.ru/ticket/FFXIXG\nр525ме797\n37366\nДубровка",
		"https://ticket.belkacar.ru/ticket/12207217\nк409рн799\n143417\nТо47",
		"https://ticket.belkacar.ru/ticket/12208002\nу346ее790\n112280\nТо 47",
		"https://ticket.belkacar.ru/ticket/12093653\nс345му797\n29517\nДубровка",
		"https://ticket.belkacar.ru/ticket/12093652\nс303му797\n29329\nТо Дубровка",
		"[ Альбом ]\nх210нн797 колесо левое заднее пробито @SP_carsh_msc_night",
		"https://ticket.belkacar.ru/ticket/12206306\nх370но790\n100102\nТо 47",
		"https://ticket.belkacar.ru/ticket/12092217\nо116вх790\n131726\nТо47",
		"https://ticket.belkacar.ru/ticket/12031821\nк348рн799\n129008\nЦветочный",
		"https://ticket.belkacar.ru/ticket/ZCY3O9\nо309мс797\n30825\nТо цветочный",
		"https://ticket.belkacar.ru/ticket/12142298\nн998мх797\n25764\nЦветочный",
		"https://ticket.belkacar.ru/ticket/11978180\nр345рм799\n143344\nТо47",
		"https://ticket.belkacar.ru/ticket/12206295\nт172нх790\n98362\nТо 47",
		"https://ticket.belkacar.ru/ticket/12092476\nе968со799\n146308\nТо47",
		"https://ticket.belkacar.ru/ticket/12144093\nр285нр790\n97373\nТо47",
		"https://ticket.belkacar.ru/ticket/12207206\nк095рн799\n149560\nЦветочный",
		"https://ticket.belkacar.ru/ticket/11978190\nх592ем790\n120541\nТо цветочный",
		"https://ticket.belkacar.ru/ticket/12092473\nе415ео790\n149163\nТо47",
		"https://ticket.belkacar.ru/ticket/12149898\nв590се790\n64690\nЦветочный",
		"https://ticket.belkacar.ru/ticket/11918189\nк272рн799\n161133\nТо47",
		"https://ticket.belkacar.ru/ticket/12092487\nм665еа790\n121926\nТо47",
		"https://ticket.belkacar.ru/ticket/12208004\nу742еа790\n130215\nТо47",
		"https://ticket.belkacar.ru/ticket/11971692\nх554мо797\n30713\nТо цветочный",
		"https://ticket.belkacar.ru/ticket/11918194\nо145вх790\n123152\nТо цветочный",
		"https://ticket.belkacar.ru/ticket/KTUGBU\nт994хв750\n146138\nТо47",
		"https://ticket.belkacar.ru/ticket/BB1DO0\nе930нс797\n128687\nТо 47",
		"https://ticket.belkacar.ru/ticket/12207982\nр123ст799\n157876\nТо47",
		"https://ticket.belkacar.ru/ticket/11863123\nа164мн797\n41179\nТо цветочный",
		"https://ticket.belkacar.ru/ticket/12206270\nе286нт790\n95908\nТо47",
		"Только если это киа",
		"https://ticket.belkacar.ru/ticket/12200809\nо642нс790\n80954\nТо цветочный",
		"https://ticket.belkacar.ru/ticket/12150749\nх515нн790\n98079\nТо47",
		"https://ticket.belkacar.ru/ticket/12200610\nк302еу790\n112173\nТо 47",
		"https://ticket.belkacar.ru/ticket/12027423\nа607ме797\n34762\nТо47",
		"https://ticket.belkacar.ru/ticket/12207952\nе979ре799\n159665\nТо 47",
	}
	text1C := []string{
		"кулак\n" +
			"1  ;к945нт790/Z94C251ABMR147949/Rio X/86 692/Пользователь забыл телефон в авто, и разбил окно чтобы его достать);\n" +
			"АВТОЛАЙТ.Полбина 29,с1  \n" +
			"АВТОСТЕКЛОУСТАНОВКА.Поречная д 10  \n" +
			"АСЦ Дубровка.Geely Exeed Chery.2яМашиностроения.д6  \n" +
			"2  ;е412нн797/Y4K8622Z6NB912431/Coolray/26 191/Плановое ТО;\n" +
			"3  ;е792но797/Y4K8622Z9NB912780/Coolray/26 053/Плановое ТО;\n" +
			"4  ;м060мн797/LVTDB21B3ND334867/LX/7 536/Плановое ТО;\n" +
			"7  ;р525ме797/Y4K8622Z2NB908926/Coolray/37 366/Плановое ТО;\n" +
			"8  ;с303му797/LVTDD21B8ND097254/TXL/29 329/Плановое ТО;\n" +
			"9  ;с345му797/LVTDD21BXND089110/TXL/29 517/Плановое ТО;\n" +
			"11  ;х634нн797/Y4K8622Z8NB912799/Coolray/21 757/Плановое ТО;\n" +
			"Казаков Моторс.Костомаровская наб.д.29А  \n" +
			"1  ;а341са790/LVVDB21B9MD292811/Tiggo 4/64 207/Замена колодок (Остаток ПТК = 4,7, Остаток ЗТК = 0);\n" +
			"2  ;в162мв797/XZGEE04A5NA818563/Jolion/31 559/Замена колодок (Остаток ПТК = 4,7, Остаток ЗТК = 0);\n" +
			"4  ;м438ек790/Z94C251ABLR101525/Rio X-Line/105 864/Замена колодок (Остаток ПТК = 2, Остаток ЗТК = 0);\n" +
			"6  ;с688но790/Z94C251ABMR152995/Rio X/92 470/;\n" +
			"7  ;т190мк797/XZGEE04A9NA817500/Jolion/35 150/Замена колодок (Остаток ПТК = 5, Остаток ЗТК = 0);\n" +
			"8  ;т271мо797/XZGEE04A4NA822247/Jolion/37 799/;\n" +
			"9  ;у789ев790/Z94C251ABLR102361/Rio X-Line/133 293/Замена колодок (Остаток ПТК = 2, Остаток ЗТК = 0);\n" +
			"МЭЙДЖОР CHERRY. Цветочный прзд д 17  \n" +
			"1  ;а164мн797/XZGEE04A7NA818886/Jolion/41 178/Плановое ТО;\n" +
			"2  ;в590се790/LVVDB21B7MD293133/Tiggo 4/64 690/Плановое ТО;\n" +
			"3  ;е640нт790/Z94C251ABMR148183/Rio X/95 026/Плановое ТО;\n" +
			"4  ;к095рн799/Z94C251ABKR063276/Rio X-Line/149 560/Плановое ТО;\n" +
			"5  ;к348рн799/Z94C251ABKR063110/Rio X-Line/129 008/Плановое ТО;\n" +
			"6  ;н998мх797/XZGEE04A3NA819274/Jolion/25 764/Плановое ТО;\n" +
			"7  ;о145вх790/Z94C251ABLR118285/Rio X-Line/123 129/Плановое ТО;\n" +
			"8  ;о309мс797/LVVDB21B1ND227713/Tiggo 7 Pro/30 824/Плановое ТО;\n" +
			"9  ;о642нс790/Z94C251ABMR161017/Rio X/80 917/Плановое ТО;\n" +
			"10  ;р092св790/LVVDB21B1MD293127/Tiggo 4/68 755/Плановое ТО;\n" +
			"11  ;у342ее790/Z94C251ABLR118291/Rio X-Line/125 069/Плановое ТО;\n" +
			"12  ;х232но797/XZGEE04A4NA822443/Jolion/27 695/Плановое ТО;\n" +
			"13  ;х543мо797/XZGEE04A9NA823734/Jolion/33 563/Плановое ТО;\n" +
			"14  ;х554мо797/XZGEE04A3NA823728/Jolion/30 713/Плановое ТО;\n" +
			"15  ;х592ем790/Z94C251ABLR117000/Rio X-Line/120 541/Плановое ТО;\n" +
			"МЭЙДЖОР KIA VW HAVAL .47 км МКАД  \n" +
			"1  ;а607ме797/XZGEE04A4NA821373/Jolion/34 726/Плановое ТО;\n" +
			"2  ;в785рн799/Z94C251ABKR063143/Rio X-Line/161 547/Плановое ТО;\n" +
			"3  ;е286нт790/Z94C251ABMR151817/Rio X/95 852/Плановое ТО;\n" +
			"4  ;е415ео790/Z94C251ABLR067204/Rio X-Line/149 163/Плановое ТО;\n" +
			"5  ;е930нс797/XW8ZZZ61ZLG007257/Polo/128 654/Плановое ТО;\n" +
			"6  ;е968со799/Z94C251ABLR067205/Rio X-Line/146 308/Плановое ТО;\n" +
			"7  ;е979ре799/Z94C251ABLR069865/Rio X-Line/159 663/Плановое ТО;\n" +
			"8  ;к083хс750/Z94C251ABKR064635/Rio X-Line/152 466/Плановое ТО;\n" +
			"9  ;к272рн799/Z94C251ABKR063098/Rio X-Line/161 133/Плановое ТО;\n" +
			"10  ;к302еу790/XW8ZZZ61ZLG026503/Polo/112 162/Плановое ТО;\n" +
			"11  ;к409рн799/Z94C251ABKR062963/Rio X-Line/143 426/Плановое ТО;\n" +
			"13  ;м665еа790/Z94C251ABLR100454/Rio X-Line/121 926/Плановое ТО;\n" +
			"14  ;о116вх790/Z94C251ABLR115406/Rio X-Line/131 726/Плановое ТО;\n" +
			"15  ;р123ст799/Z94C251ABLR069143/Rio X-Line/157 808/Плановое ТО;\n" +
			"16  ;р285нр790/Z94C251ABMR149796/Rio X/97 373/Плановое ТО;",

		"17  ;р345рм799/Z94C251ABKR064408/Rio X-Line/143 344/Плановое ТО;\n" +
			"19  ;т172нх790/Z94C251ABMR150019/Rio X/98 362/Плановое ТО;\n" +
			"20  ;т208ем790/Z94C251ABLR117961/Rio X-Line/127 012/Плановое ТО;\n" +
			"22  ;т994хв750/Z94C251ABKR064529/Rio X-Line/146 136/Плановое ТО;\n" +
			"23  ;у346ее790/Z94C251ABLR115010/Rio X-Line/112 280/Плановое ТО;\n" +
			"24  ;у742еа790/Z94C251ABLR101870/Rio X-Line/130 215/Плановое ТО;\n" +
			"25  ;х370но790/Z94C251ABMR154224/Rio X/100 102/Плановое ТО;\n" +
			"26  ;х515нн790/XW8ZZZ5NZNG002795/Tiguan/98 022/Плановое ТО;\n" +
			"СТО ШИКАНА.Хлебозаводский проезд 7А  \n" +
			"2  ;к030хс750/Z94C251ABKR064546/Rio X-Line/164 649/Не работает кондиционер+греется;\n" +
			"3  ;к242рн799/Z94C251ABKR063232/Rio X-Line/147 075/диагностика подвески  + пинается коробка;\n" +
			"8  ;р133ст799/Z94C251ABLR069463/Rio X-Line/155 191/сломана шпилька на переднем правом;\n" +
			"ЦКР-АВТО. Перерва д19,с2  \n" +
			"1  ;х344ев790/Z94C251ABLR101307/Rio X-Line/129 033/антенна, спойлер переднего бампера;\n",
	}

	var msgs []utils.Message
	for _, text := range textFlower {
		msgs = append(msgs, utils.Message{Text: text, From: utils.ID_flower})
	}
	for _, text := range text47 {
		msgs = append(msgs, utils.Message{Text: text, From: utils.ID_47})
	}
	for _, text := range textStretches {
		msgs = append(msgs, utils.Message{Text: text, From: utils.ID_stretches})
	}
	for _, text := range text1C {
		msgs = append(msgs, utils.Message{Text: text, From: utils.ID_oneC})
	}

	tests := []testStruct{
		{
			name: "большая куча говна",
			args: args{
				msgs: msgs,
				conf: utils.Conf{
					Chat47:        true,
					ChatFlower:    true,
					OneC:          true,
					ChatStretches: true,
					Mail:          true,
				},
			},
			wantRespMsgs: []utils.Message{
				{
					Text:      "https://ticket.belkacar.ru/ticket/S5HBAX\nр133ст799\nШикана",
					ReplyText: "не отправлен на service",
					From:      utils.ID_stretches,
					Loc:       "шикана",
					Mark:      "р133ст799",
				},
				{
					Text:      "https://ticket.belkacar.ru/ticket/WSS8DF\nу401нн797\n1-я Магистральная",
					ReplyText: "не отправлен на kuzov",
					From:      utils.ID_stretches,
					Loc:       "магистральная",
					Mark:      "у401нн797",
				},

				{
					Text:      "р133ст799шикана",
					ReplyText: "р133ст799 не найден в чате перегонов или нахождение в источнике не подразумевается",
					From:      0,
					Loc:       "шикана",
					Mark:      "р133ст799",
				},
			},
		},
	}

	p := Init(false)
	p.mailTestFlag = true
	p.mailParsedData = func() (service map[string]utils.Message, kuzov map[string]utils.Message) {
		service = map[string]utils.Message{
			"х344ев790": {
				Text: "х344ев790цкр",
				Loc:  "цкр",
				Mark: "х344ев790",
			},
			"к242рн799": {
				Text: "к242рн799шикана",
				Loc:  "шикана",
				Mark: "к242рн799",
			},
			"к030хс750": {
				Text: "к030хс750шикана",
				Loc:  "шикана",
				Mark: "к030хс750",
			},
			/*
				"": {
					Text: "",
					Loc:  "",
					Mark: "",
				},
			*/
		}
		kuzov = map[string]utils.Message{
			"к945нт790": {
				Text: "к945нт790\nна кулак",
				Loc:  "кулак",
				Mark: "к945нт790",
			},
			"р133ст799": {
				Text: "р133ст799шикана",
				Loc:  "шикана",
				Mark: "р133ст799",
			},
		}
		return
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRespMsgs := p.Processing(tt.args.msgs, tt.args.conf); !reflect.DeepEqual(gotRespMsgs, tt.wantRespMsgs) {
				t.Errorf("Processing():")
				for _, msg := range gotRespMsgs {
					msg.Print()
				}
				t.Errorf("want:")
				for _, msg := range tt.wantRespMsgs {
					msg.Print()
				}
			}
		})
	}
}

func TestProc_Processing(t *testing.T) {
	type args struct {
		msgs []utils.Message
		conf utils.Conf
	}

	type testStruct struct {
		name         string
		args         args
		wantRespMsgs []utils.Message
	}

	tests := []testStruct{
		{
			name: "валидный перегон по слесарке",
			args: args{
				msgs: []utils.Message{
					{
						Text: "СТО ШИКАНА.Хлебозаводский проезд 7А  \n" +
							"9  ;у749мх797/XZGEE04A9NA822972/Jolion/32 410/Колодок нет, разгоняется плохо, на повышенную очень долго переключает;\n",
						From: utils.ID_oneC,
					},
					{
						Text: "https://ticket.belkacar.ru/ticket/YHZKZ7\nу749мх797\nШикана",
						From: utils.ID_stretches,
					},
				},
				conf: utils.Conf{
					Chat47:        true,
					ChatFlower:    true,
					OneC:          true,
					ChatStretches: true,
					Mail:          false,
				},
			},
			wantRespMsgs: nil,
		},
		{
			name: "валидные перегоны по слесарке",
			args: args{
				msgs: []utils.Message{
					{
						Text: "СТО ШИКАНА.Хлебозаводский проезд 7А  \n" +
							"9  ;у749мх797/XZGEE04A9NA822972/Jolion/32 410/Колодок нет, разгоняется плохо, на повышенную очень долго переключает;\n" +
							"2  ;к030хс750/Z94C251ABKR064546/Rio X-Line/164 649/Не работает кондиционер+греется;\n" +
							"Казаков Моторс.Костомаровская наб.д.29А  \n" +
							"7  ;т190мк797/XZGEE04A9NA817500/Jolion/35 150/Замена колодок (Остаток ПТК = 5, Остаток ЗТК = 0);\n",
						From: utils.ID_oneC,
					},
					{
						Text: "АСЦ Дубровка.Geely Exeed Chery.2яМашиностроения.д6  \n" +
							"https://ticket.belkacar.ru/ticket/12141209\nе412нн797\n26191\nДубровка",
						From: utils.ID_oneC,
					},
					{
						Text: "https://ticket.belkacar.ru/ticket/YHZKZ7\nу749мх797\nШикана",
						From: utils.ID_stretches,
					},
					{
						Text: "https://ticket.belkacar.ru/ticket/ZXZG11\nк030хс750\nШикана",
						From: utils.ID_stretches,
					},
					{
						Text: "https://ticket.belkacar.ru/ticket/12141209\nе412нн797\n26191\nДубровка",
						From: utils.ID_stretches,
					},
					{
						Text: "https://ticket.belkacar.ru/ticket/QVDMRX\nт190мк797\nКазаков Моторс",
						From: utils.ID_stretches,
					},
				},
				conf: utils.Conf{
					Chat47:        true,
					ChatFlower:    true,
					OneC:          true,
					ChatStretches: true,
					Mail:          false,
				},
			},
			wantRespMsgs: nil,
		},
		{
			name: "валидный перегон по кузову",
			args: args{
				msgs: []utils.Message{
					{
						Text: "https://ticket.belkacar.ru/ticket/WSS8DF\nу401нн797\n1-я Магистральная",
						From: utils.ID_stretches,
					},
				},
				conf: utils.Conf{
					Chat47:        true,
					ChatFlower:    true,
					OneC:          true,
					ChatStretches: true,
					Mail:          false,
				},
			},
			wantRespMsgs: nil, //need real mail
		},
		{
			name: "валидный перегон на ТО",
			args: args{
				msgs: []utils.Message{
					{
						Text: "пригнали\nк083хс750",
						From: utils.ID_47,
					},
					{
						Text: "МЭЙДЖОР KIA VW HAVAL .47 км МКАД  \n" +
							"8  ;к083хс750/Z94C251ABKR064635/Rio X-Line/152 466/Плановое ТО;\n",
						From: utils.ID_oneC,
					},
					{
						Text: "https://ticket.belkacar.ru/ticket/12207957\nк083хс750\n152466\nТо 47",
						From: utils.ID_stretches,
					},
				},
				conf: utils.Conf{
					Chat47:        true,
					ChatFlower:    true,
					OneC:          true,
					ChatStretches: true,
					Mail:          false,
				},
			},
			wantRespMsgs: nil,
		},
		{
			name: "валидные перегоны на ТО",
			args: args{
				msgs: []utils.Message{
					{
						Text: "пригнали\nк083хс750\nт208ем790",
						From: utils.ID_47,
					},
					{
						Text: "МЭЙДЖОР KIA VW HAVAL .47 км МКАД  \n" +
							"8  ;к083хс750/Z94C251ABKR064635/Rio X-Line/152 466/Плановое ТО;\n" +
							"20  ;т208ем790/Z94C251ABLR117961/Rio X-Line/127 012/Плановое ТО;\n",
						From: utils.ID_oneC,
					},
					{
						Text: "https://ticket.belkacar.ru/ticket/12207957\nк083хс750\n152466\nТо 47",
						From: utils.ID_stretches,
					},
					{
						Text: "https://ticket.belkacar.ru/ticket/12207230\nт208ем790\n127012\nТо47",
						From: utils.ID_stretches,
					},
				},
				conf: utils.Conf{
					Chat47:        true,
					ChatFlower:    true,
					OneC:          true,
					ChatStretches: true,
					Mail:          false,
				},
			},
			wantRespMsgs: nil,
		},
		{
			name: "невалидный перегон по слесарке без заявки в 1С",
			args: args{
				msgs: []utils.Message{
					{
						Text: "https://ticket.belkacar.ru/ticket/GILYOS\nх169вх790\nКулак",
						From: utils.ID_stretches,
					},
					{
						Text: "https://ticket.belkacar.ru/ticket/12092476\nе968со799\n146308\nТо47",
						From: utils.ID_stretches,
					},
					{
						Text: "МЭЙДЖОР KIA VW HAVAL .47 км МКАД  \n" +
							"6  ;е968со799/Z94C251ABLR067205/Rio X-Line/146 308/Плановое ТО;\n",
						From: utils.ID_oneC,
					},
					{
						Text: "е968со799 готова",
						From: utils.ID_47,
					},
					{
						Text: "пригнали\nе968со799\n",
						From: utils.ID_47,
					},
				},
				conf: utils.Conf{
					Chat47:        true,
					ChatFlower:    true,
					OneC:          true,
					ChatStretches: true,
					Mail:          false,
				},
			},
			wantRespMsgs: []utils.Message{
				{
					Text: "https://ticket.belkacar.ru/ticket/GILYOS\nх169вх790\nКулак",
					From: utils.ID_stretches,

					ReplyText: "не найден в 1С(как заявка на ремонт)",
					Loc:       "кулак",
					Mark:      "х169вх790",
				},
			},
		},
		{
			name: "лишняя заявка в 1С",
			args: args{
				msgs: []utils.Message{
					{
						Text: "https://ticket.belkacar.ru/ticket/12092476\nе968со799\n146308\nТо47",
						From: utils.ID_stretches,
					},
					{
						Text: "МЭЙДЖОР KIA VW HAVAL .47 км МКАД  \n" +
							"6  ;е968со799/Z94C251ABLR067205/Rio X-Line/146 308/Плановое ТО;\n" +
							"19  ;т172нх790/Z94C251ABMR150019/Rio X/98 362/Плановое ТО;\n",
						From: utils.ID_oneC,
					},
					{
						Text: "е968со799 готова",
						From: utils.ID_47,
					},
					{
						Text: "пригнали\nе968со799\n",
						From: utils.ID_47,
					},
				},
				conf: utils.Conf{
					Chat47:        true,
					ChatFlower:    true,
					OneC:          true,
					ChatStretches: true,
					Mail:          false,
				},
			},
			wantRespMsgs: []utils.Message{
				{
					Text: "text from 1C",
					From: utils.ID_oneC,

					ReplyText: "т172нх790 не найден в чате перегонов или нахождение в источнике не подразумевается",
					Loc:       "47",
					Mark:      "т172нх790",
				},
			},
		},
	}

	p := Init(false)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRespMsgs := p.Processing(tt.args.msgs, tt.args.conf); !reflect.DeepEqual(gotRespMsgs, tt.wantRespMsgs) {
				t.Errorf("Processing():")
				for _, msg := range gotRespMsgs {
					msg.Print()
				}
				t.Errorf("want:")
				for _, msg := range tt.wantRespMsgs {
					msg.Print()
				}
			}
		})
	}
}
