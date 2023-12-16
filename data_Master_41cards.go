package main

var fileOfCardsMaster = []charSetStructKanji{

	{"車", "car", "vehicle",
		"ony: シャ (sha)",
		"kun: くるま (kuruma)",
		"自動車(jidousha; automobile), 電車(densha; train), 車体(shatai; car body), 車両(sharyou)",
		"新車(shinchya; new car), 車輪(sharin; wheel), 三輪車(sanrinsha; tricycle), 車軸(shajiku), ハンドル(handoru)"},

	{"日", "day", "sun",
		"ony: ジツ (jitsu:じつ), ニチ (nichi:にち)",
		"kun: ひ, か",
		"明日 (tomorrow)",
		" "},

	{"土", "earth-1", "soil, dirt, ground, mud, wet clay",
		"ony: ド, ト  (ど, と) do, to",
		"kun: つち (tsuchi)",
		"土曜日 - Saturday",
		" "},

	{"地", "earth-2", "land, place, solid ground",
		"ony: チ (chi); ジ (ji)",
		"kun: -",
		"地震(jishin; earthquake), 地域(chiiki; region), 土地(tochi; land)",
		"現地(genchi; site), 出身地(shusshin chi; hometown), 地名(chimei; place name)"},

	{"高", "expensive", "high, tall",
		"ony: コウ (kou); タカ (taka)",
		"kun: たか (taka)",
		"高い(takai; high), 最高(saikou; highest), 高原(kougen; plateau)",
		"高校(koukou; high school), 高層(kousou; high-rise), 高速(kousoku; high speed)\n" +
			"高い (expensive)"},

	{"魚", "fish", "seafood",
		"ony: ギョ (gyo), ぎょ (gyo)",
		"kun: さかな (sakana)",
		"魚 (fish)",
		"\ncompare bird:鳥"},

	{"足", "foot", "add, to-add",
		"ソク (soku), そく (soku)",
		"kun: あし (ashi), ta(su) た(す)",
		"足 (foot)",
		" "},

	{"脚", "leg", "foot; base; foundation",
		" -",
		"kun: あし",
		" ",
		"Variant of Chinese: 腳  -  Foot足 or Leg脚, あし"},

	{"友", "friend", "companion",
		"ony: ユウ (yuu ゆう) ==== , 友人(yuujin; friend)",
		"kun: とも (tomo)",
		", 友好(yuukou; friendship), 親友(shin'yuu; close friend), 友達 (ともだち - activities friend)",
		"仲間(nakama; companion), 取引先(torihikisaki; business connection), 共同(kyoudou; joint)"},

	{"天", "heaven", "heavens or celestial sky 空",
		"ony: テン (ten), てん (ten)",
		"kun: あめ (ame), あま (ama)",
		"天気 (weather)",
		" "},

	{"空", "sky", "empty, emptiness, vacuum, void",
		"ony: クウ (kuu); ソラ (sora そら)",
		"kun: から (kara), くう, あける, あく, あき",
		"空港(kuukou; airport), 宇宙空間(uchuu kukan; outer space), 空腹(kuufuku; hunger)",
		"空っぽ (karappo; empty), 空く (aku; become empty), 空き家 (akai; vacant house), 空白 (kuuhaku; blank space)\n" +
			"空っぽ (kuuppo; empty), 空想 (kuusou; imagination), 空蝉 (kuusagi; empty shell)"},

	{"百", "hundred", "many, percent",
		"ony: ヒャク (hyaku)",
		"kun: -",
		"百日(hyakunichi; hundred days), 百人(hyakunin; hundred people), 百回(hyakkai; hundred times)",
		"百パーセント(hyaku paasento; one hundred percent), 二百(nihyaku; two hundred), 三百(sanbyaku; three hundred)\n" +
			"百万円 (1 million Yen)"},

	{"左", "left", "port, direction",
		"ony: サ (sa)",
		"kun: ひだり (hidari)",
		"左手(hidarite; left hand), 左翼(sayoku; left wing), 左折(sasetsu; turn left)",
		"左側(hidarigawa; left side)"},

	{"長", "long", "leader",
		"ony: チョウ (chou)",
		"kun: なが (naga)",
		"長い(nagai; long), 長さ(nagasa; length), 長期(chouki; long time)",
		"会長(kaichou; president), 部長(buchou; department head, manager), 長寿(chouju; longevity)"},

	{"母", "mother", "parent",
		"ony: ボ (bo)",
		"kun: はは (haha)",
		"母親(hahaoya; mother), 母上(hahaue; mother), 義母(gibou; mother-in-law)",
		"お母さん(okaasan; mother), 母ちゃん(kachan; mom), ママ(mama; mom)"},

	{"北", "north", "northern",
		"ony: ホク (hoku)",
		"kun: きた (kita)",
		"北海道(Hokkaidō; Hokkaido), 北極(Hokkyoku; North Pole), 北東(Higashi; northeast)",
		"北口(Kitaguchi; north exit), 北側(kitagawa; northern side), 北向き(kitamuki; facing north)"},

	{"古", "old-1", "former, antiquity",
		"ony: コ (ko)",
		"kun: ふる (furu-i)",
		"古典(koten; classic), 古代(kodai; ancient times), 古文(kobun; classical literature)",
		"古ぼける(fubokeru; grow old), 古着(furugi; old clothes)"},

	{"旧", "old-2", "former",
		"ony: キュウ (kyuu); コ(ko)",
		"kun: ふる (furui)",
		"旧友(kyuuyuu; old friend), 旧姓(kyuusei; maiden name), 旧弊(kyuuhai; old practice)",
		"昔(mukashi; in the past), 古い(furui; old), 用いられた(mochiirareru; used)"},

	{"老", "old-3", "aged",
		"ony: ロウ (rou)",
		"kun: お (o)いる",
		"老人(roujin), 高齢者(koureisha), 老齢(rouireki)",
		"老若男女(rounyounan), 老体(otai), 老衰(rousui); compare: person(formal):old 者:老"},

	{"雨", "rain", "wet, umbrella",
		"ony: ウ (u う)",
		"kun: あめ (ame)",
		"豪雨(gouu; heavy rain), 雨季(uuki; rainy season), 台風の雨(taifuu no ame; typhoon rain)",
		"傘(kasa; umbrella), 雨具(amagu; rainwear), 雨天(uenten; rainy weather)"},

	{"右", "right", "starboard, direction",
		"ony: ウ (u), ユウ　(yuu)",
		"kun: みぎ (migi)",
		"右手(migite; right hand), 右翼(uyoku; right wing), 右折(usettsu; turn right)",
		"右側(migigawa; right side), 坐る(suwaru; sit)"},

	{"店", "shop", "store",
		"ony: てん (ten)",
		"kun: みせ (mise)",
		"喫茶店 (coffee shop)",
		" "},

	{"社", "shrine", "society",
		"ony: しゃ (sha)",
		"kun: やしろ (yashiro)",
		"社長 (president of a company)",
		" "},

	{"六", "six", "6, sextuple",
		"ony: ロク (roku)",
		"kun: む (mu)",
		"六日(muika; six days), 六人(rokunin; six people), 六回(rokkai; six times)",
		"六角(rokkaku; hexagon), 六月(rokugatsu; June), 六十(rokujuu; sixty), 六日(6th day of the month)"},

	{"駅", "station", "Train-Station",
		"ony: えき (eki)",
		"kun: -",
		"駅前 (in front of the station)",
		" "},

	{"道", "street", "path, way",
		"dou (the way) どう", "michi みち",
		"道具 (tool)",
		" "},

	{"三日", "3rd-day-of-the-month", "3rd day of the month",
		"san さん", "mi みか",
		"三日 3rd day of the month",
		"xxxx"},

	{"万", "ten-thousand", "Ten thousand",
		"man まん, ban ばん", "—",
		"万年筆 (fountain pen, ten thousand year brush)",
		" "},

	{"休", "to-rest", "or; break, holiday, vacation",
		"kyuu きゅう", "yasu やす(muむ), yasu やす(mi み)",
		"休む (to take a day off)",
		" "},

	{"水", "water", "Water",
		"sui すい", "mizu みづ",
		"水曜日 (Wednesday)",
		" "},

	{"白", "white", "White",
		"haku はく, byaku びゃく", "shiro しろ, しろ(い)",
		"白い (white), 面白い (interesting)",
		" "},

	// /*
	{"本", "book", "or, source",
		"hon", "moto",
		"日本語 (Japanese)",
		" "},

	{"火", "fire", "Fire",
		"ka", "hi",
		"火曜日 (Tuesday)",
		" "},

	{"大", "big", "a lot",
		"dai, tai", "oo(kii)",
		"大きい (big), 大変 (dreadful, immense)",
		" "},

	{"多", "many", "a lot",
		"ta", "oo(i)",
		"多い (many), 多分 (probably)",
		" "},

	{"飲", "to-drink", "a drink",
		"in", "no(mu)",
		"飲み物 (beverage)",
		" "},

	{"出", "to-leave", "To go out",
		"shutsu", "de(ru), da(su)",
		"出かける (to go out)",
		" "},

	{"入", "to-enter", "or, to put in",
		"nyuu", "hai(ru), i(reru",
		"入口 (entrance)",
		" "},

	{"電", "electricity", "Electricity",
		"den", "—",
		"電気 (electricity)",
		" "},

	{"七", "seven", "Seven",
		"shichi", "nana(tsu), nana",
		"七日 (7th day of the month)",
		" "},

	{"中", "middle", "center, inner, between",
		"chuu", "naka",
		"日中 (during the day, midday)",
		" "},

	{"上", "above", "Up",
		"shou, jou", "ue, u, a(geru)",
		"上着 (jacket)",
		" "},

	{"下", "down", "below",
		"ka, ge", "ku(daru), shita",
		"靴下 (socks)",
		" "},

	{"五", "five", "Five",
		"go", "itsu(tsu), itsu",
		"五日 (5th day of the month)",
		" "},

	{"女", "woman", "girl, female",
		"jo じょ, nyo にょ", "onna おんな, me め",
		"おんな、女 – female; おんなのひと、女の人 – woman; おんなのこ、女の子 – girl",
		" "},

	// */
	{"口", "mouth", "Mouth",
		"kou, ku", "kuchi",
		"出口 (exit)",
		" "},
}
