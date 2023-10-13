package main

// The deck will draw cards per this aCard var
var aCard = charSetStructK{}

type charSetStructK struct {
	Kanji   string
	Meaning string
	Onyomi  string
	Kunyomi string
	Vocab   string
}

/*
type charSetStructK struct {
	Kata   string
	Hira   string
	Romaji string
	HiraHint string
	KataHint string
}

*/

var fileOfCardsK = []charSetStructK{
	{"店", "shop", "ten", "mise", "喫茶店 (coffee shop)"},
	{"駅", "station", "eki", "-", "駅前 (in front of the station)"},
	{"道", "street", "dou", "michi", "道具 (tool)"},
	{"社", "shrine", "sha", "yashiro", "社長 (president of a company)"},
	{"国", "country", "koku", "kuni", "外国人 (foreigner)"},
	{"外", "outside", "gai-ge", "soto-hazu(reru)-hoka", "外国 (foreign country)"},
	{"学", "school", "gaku", "mana(bu)", "大学 (university)"},
}

/*
	{"道", "street-path-way", "dou", "michi", "道具 (tool)"},
	{"社", "shrine-society", "sha", "yashiro", "社長 (president of a company)"},
	{"学", "school-learning", "gaku", "mana(bu)", "大学 (university)"},
*/

// Constants:
const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorCyan = "\033[36m"
const colorPurple = "\033[35m"
