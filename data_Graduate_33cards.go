package main

var fileOfCardsGraduate = []charSetStructKanji{

	{"安", "cheap", "peace",
		"anあん", "yasuやす(iい)",
		"安い (cheap)",
		" - or; safety"},

	{"耳", "ear", "hearing",
		"ony: ジ (ji)",
		"kun: みみ (mimi)",
		"耳鳴り(miminari; ringing in ears), 耳栓(mimisen; earplugs), 耳飾り(mimi kazari; earring)",
		"聴覚(choukaku; hearing), 聞こえる(kikoeru; audible), 聞く(kiku; listen)"},

	{"東", "east", "eastern, eastern direction",
		"ony: トウ (tō)(tou)",
		"kun: ひがし (higashi)",
		"東京(Tōkyō; Tokyo), 東北(Tōhoku; northeast region), 東海道(Tōkaidō; Tokaido region), 東日(toujitsu; rising sun)",
		"東口(Higashiguchi; east exit), 東側(higashigawa; eastern side), 東向き(higashimuki; facing east), 東洋(touyou; the East)\n" +
			"compare: gather 集:東"},

	{"八", "eight", "octuple",
		"ony: ハチ (hachi)",
		"kun: や (ya), やつ",
		"八日(youka; eight days, 8th day of the month), 八人(hachinin; eight people), 八回(hakkai; eight times)",
		"八月(hachigatsu; August), 八十(hachijuu; eighty), 八百(happyaku; eight hundred)\n" +
			"compare: 入 to Enter, 六 six"},

	{"毎", "every", "each",
		"mai　まい", "—",
		"毎日(every day)",
		" "},

	{"目", "eye", "vision, ocular",
		"ony: モク (moku); ボク (boku)",
		"kun: め (me)",
		"目力(mejikara; eyesight), 目的(mokuteki; purpose), 目玉(medama; eyeball)",
		"視力(shiryoku; eyesight), 見る(miru; see), 見える(mieru; visible)"},

	{"父", "father", "dad, papa, parent",
		"ony: フ (fu); ブ(bu)",
		"kun: ちち (chichi)",
		"父親(fubo; father), 父上(chichiue; father), 義父(gifu; foster father)",
		"お父さん(otosan; father), 父ちゃん(tochan; dad), 爸(papa; dad)"},

	{"半", "half", "middle",
		"han　まん", "naka(ba)　なか(ば)",
		"半分 (half)",
		" "},

	{"手", "hand", "pyka",
		"shu", "te",
		"手紙 (letter)",
		" Russian: pyka"},

	{"生", "life", "to grow",
		"せい, しょう", "i(kiru), u(mareru), ha(yasu)",
		"生徒 (pupil)",
		"to live, to be born, to grow\n" +
			"compare: 午 noon, "},

	{"小", "small", "little",
		"ony: ショウ(shou); コ(ko)",
		"kun: ちい(chiisai)",
		"小さい(chiisai; small), 小学校(shougakkou; elementary school), 少女(shoujo; girl)",
		"最小(saishou; smallest), 少量(shouryou; small amount), 細かい(komakai; tiny)"},

	{"男", "man", "male",
		"ony: ダン (dan); ナン (nan)",
		"kun: おとこ (otoko)",
		"男性(dansei; male), 男子(danshi; boy, dannoshi 男の子), 男っ気(otokogokoro; manliness)",
		"男の人(otoko no hito; man), 男らしい(otokorashii; manly), 男子学生(danshigakusei; male student)"},

	{"金", "gold", "money",
		"ony: キン (kin)",
		"kun: かね (kane)",
		"金曜日(kin'youbi; Friday), 金メダル(kinmedaru; gold medal), お金(okane; money)",
		"金庫(kinko; safe), 金魚(kingyo; goldfish), 銀行(ginkou; bank)"},

	{"名", "name", "moniker",
		"mei　めい, myou　みょう", "na　な",
		"名前 (name)",
		"compare 右 right "},

	{"九", "nine", "ninth",
		"ony: キュウ (kyuu)",
		"kun: ここの (kokono)",
		"ここのつ(kokonotsu; nine), 九日(kokonoka; nine days), 九人(kyuunin; nine people), 九回(kyuukai; nine times)",
		"九月(kugatsu; September), 九十(kyuujuu; ninety), 九州(kyuushuu; Kyushu)"},

	{"午", "noon", "midday",
		"ご", "—",
		"午前 (morning, A.M.)",
		"compare: 生 l i f e,  "},

	{"今", "now", "the present",
		"kon　こん, kin　きん", "ima　いま",
		"今晩 (this evening), 今朝 (this morning)",
		"compare: 令 (order, rule, decree)"},

	{"気", "feeling", "spirit, ghost",
		"ony: キ (ki)",
		"kun: -",
		"元気(healthy, spirit, fine), 気分(kibun; feeling), 気持ち(kimochi; feeling), 気象(kishou; weather)",
		"大気(taiki; atmosphere), 気圧(kiatsu; atmospheric pressure), 気候(kikou; climate)"},

	{"買", "buy", "purchase",
		"ony: バイ (bai)",
		"kun: か (ka)",
		"購入(kounyuu; purchase), 買い物(kaimono; shopping), 買取(kaitori; buying)",
		"買う(kau; buy), 売買(baibai; buying and selling), 代金(daikin; payment)"},

	{"来", "to come", "approach",
		"rai　らい", "ku(ru)　く(る)",
		"来月 (next month), 来る (to come)",
		" "},

	{"食", "eat", "food",
		"ony: ショク (shoku)",
		"kun: た (as in taberu)",
		"食事(shokuji; meal), 食品(shokuhin; food), 食堂(shokudou; dining room/hall)",
		"食べる(taberu; eat), 朝食(choushoku; breakfast), 献立(kondate; menu)"},

	{"行", "to go", "carry out",
		"kou : こう", "い(く), okona : おこな(う)",
		"銀行 (bank)",
		" "},

	{"聞", "to hear", "to listen",
		"mon　もん, bun　ぶん", "ki(ku)　き(く)",
		"聞く (to listen, to hear)",
		"to listen, to ask"},

	{"会", "to meet", "society",
		"kaiかい, eえ", "aあ(uう)",
		"会社 (company)",
		" "},

	{"見", "see", "outlook, view, to show, to be visible",
		"ony: ケン (ken)",
		"kun: み (mi)",
		"見解(kenkai; view), 山頂(sanchou; mountain top), 考え方(kangaekata; way of thinking)",
		"見る(miru; see), 見せる (to show), 見える(mieru; be visible), 一見(ikken; glance)"},

	{"立", "stand", "establish, consist, to stand",
		"ony: リツ (ritsu); リュウ(ryuu)",
		"kun: た (ta)",
		"立候補(rikouho; to stand for election), 立証(risshou; proof), 独立(dokuritsu; independence)",
		"立つ(tatsu; stand-up, to-stand), 建てる(tateru; build up), 存在(sonzai; existence)"},

	{"西", "west", "western",
		"ony: セイ (sei); サイ (sai)",
		"kun: にし (nishi)",
		"西日本(Nishi-Nihon; western Japan), 西洋(Seiyō; the West), 西南(Seinan; southwest)",
		"西口(Nishiguchi; west exit), 西側(nishigawa; western side), 西向き(nishimuki; facing west)" +
			"西新宿(Nishi-Shinjuku), 西方(Seihō), 西日(nishibi; setting sun), 西暦(seireki; Western calendar),\n" +
			"compare: 百 hundred,  酒 alcohol (man sticks head up to look west, one v-line w 2 fifties under is 100 "},

	{"何", "what", "which",
		"ka　か", "nan　なん, nani　なに",
		"何曜日 (what day of the week)",
		"how many"},

	{"新", "new", "fresh, refresh, renew",
		"ony: シン (shin)",
		"kun: あら (ara), あたら (atara) ????",
		"新しい(new), 新品(shinpin; new article), 新築(shinchiku; new construction), 新設(shinsetsu; new establishment)",
		"新人(shinjin; new employee), 更新(koushin; renewal), 新鮮(shinsen; freshness), 新参者(shinzanza; new entrant)\n" +
			"新聞(newspaper), 新宿(Shinjuku), 新大久保(Shin-Ōkubo), 新南口(Shin-Minamiguchi)"},

	{"外", "outside", "exterior, foreign, to deviate",
		"ony: ガイ (gai)",
		"kun: そと (soto)",
		"外国(gaikoku; foreign country), 外出(gaishutsu; going out), 外観(gaikan; exterior)",
		"屋外(okugai; outdoors), 海外(kaigai; overseas), 表(omote; surface)"},

	{"南", "south", "southern direction",
		"ony: ナン (nan)",
		"kun: みなみ (minami)",
		"南国(Nangoku; southern country), 南極(Nankyoku; South Pole)",
		"南口(Minamiguchi; south exit), 南側(minamigawa; southern side), 南向き(minamimuki; facing south)"},

	{"少", "few", "little, scarce, a little",
		"ony: ショウ (shou)",
		"kun: すく (suku)",
		"少ない(sukunai; few), 少数(shousuu; few/minority), 少量(shouryou; small quantity/amount), 少し(sukoshi; a-little)",
		"年少(nenshou; young age), 少年(shounen; boy), 少女(shoujo; girl), わずか(wazuka; little), 乏しい(toboshii; scarce)"},

	{"花", "flower", "blossom, bloom, beauty",
		"ony: カ (ka)",
		"kun: はな (hana)",
		"花瓶(kabin; vase), 花束(hanataba; bouquet), 花見(hanami; flower viewing), 草花(kusabana), 花言葉(hanakotoba), 花びら(hanabira)",
		"桜の花(sakura no hana; cherry blossoms), 花火(hanabi; fireworks), 花屋(hanaya; florist), 花壇(hanadan), 花道(hanamichi),\n" +
			"compare: 華 (splendor) or \"to flower\" "},
}
