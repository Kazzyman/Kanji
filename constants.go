package main

var fileOfCardsK = []charSetStructK{
	{"西", "west", "West", "sai, sei　さい, せい", "nishi　にし", "西 (west)"},
	{"若者", "youth", "Young person", "しゃく, にゃく", "わかい",
		"わかもの、若者 "},
	{"者", "person", "Person", "しゃ, じゃ", "-",
		"わかもの、若者 - young person"},
	{"東", "east", "East", "tou　どう", "higashi　ひがし", "東京 (Tokyo)"},
	{"南", "south", "South", "nan　なん", "minami　みなみ", "南 (south)"},
	{"四", "four", "Four", "shi　し", "yo(tsu), yo, yon　よん", "四日 (4th day of the month)"},
	{"分", "minute", "Minute, part, to understand, to divide", "bun　ぶん, bu, fun　ふん", "wa(karu)",
		"三十分 (thirty minutes), 自分 (oneself)"},
	{"千", "thousand", "Thousand", "sen　せん", "chi　ち", "千万円 (10 million Yen)"},
	{"円", "yen", "Yen, circle, and round", "en　えん", "maru(i)　まるい", "円い (round)"},
	{"日", "day", "Day, sun", "nichi　にち, jitsu　じつ", "hi　ひ, ka　か", "明日 (tomorrow)"},
	{"週", "week", "Week", "shuu　しゅう", "—", "毎週 (every week)"},
	{"月", "month", "Month, moon", "getsu　げつ, gatsu　がつ", "tsuki　つき", "月曜日 (Monday)"},
	{"年", "year", "Year", "nen　ねん", "toshi　とし", "今年 (this year), 去年 (last year)"},

	{"間", "time", "Time frame, span of time", "kan　かん, ken　けん", "aida　あいだ",
		"時間 (time, hours)"},
	{"時", "time", "Time, hour", "ji　じ", "toki　とき", "時計 (clock, watch)"},

	{"午", "noon", "Noon", "go　ご", "—", "午前 (morning, A.M.)"},

	{"先", "before", "Before, ahead, future", "sen　せん", "saki　さき",
		"先週 (last week), 先生 (teacher, master)"},
	{"前", "before", "Before", "zen　ぜん", "mae　まえ", "名前 (name)"},

	{"後", "after", "After, later, behind", "go　ご, kou　こう", "ato　あと", "午後 (afternoon, P.M.)"},
	{"今", "now", "Now", "kon　こん, kin　きん", "ima　いま", "今晩 (this evening), 今朝 (this morning)"},

	{"来", "to-come", "To come", "rai　らい", "ku(ru)　く(る)", "来月 (next month), 来る (to come)"},
	{"半", "half", "Half, middle", "han　まん", "naka(ba)　なか(ば)", "半分 (half)"},
	{"毎", "every", "Every, each", "mai　まい", "—", "毎日(every day)"},
	{"男", "man", "Male", "だん, なん", "おとこ",
		"おとこのこ、男の子 – boy; おとこのひと、男の人 – man; おとこ、男 – male"},

	{"水", "water", "Water", "sui　すい", "mizu　みづ", "水曜日 (Wednesday)"},
	{"金", "money", "Money, gold", "kin　きん, kon　こん", "kane　かね", "金曜日 (Friday)"},
	{"花", "flower", "Flower", "ka　か", "hana　はな", "花火 (fireworks)"},
	{"気", "spirit", "Spirit", "ki　き, ke　け", "—", "元気 (healthy, spirit, fine)"},
	{"生", "life", "Life, to live, to be born, to grow", "sei　せい, shou　しょう",
		"i(kiru), u(mareru), ha(yasu)",
		"生徒 (pupil)"},
	{"車", "car", "Car, vehicle", "sha　しゃ", "kuruma　くるま", "電車 (electric train)"},
	{"語", "language", "Language, word, to chat", "go　ご", "kata(ru)　かた(る)", "英語 (English)"},
	{"耳", "ear", "Ear", "ji　じ", "mimi　みみ", "耳 (ear)"},
	{"目", "eye", "Eye", "moku　もく", "me　め", "目 (eye)"},
	{"名", "name", "Name", "mei　めい, myou　みょう", "na　な", "名前 (name)"},
	{"見", "to-see", "To see, to be visible, to show", "ken　けん", "mi(ru)　み(る)", "見せる (to show)"},
	{"聞", "to-hear", "To hear, to listen, to ask", "mon　もん, bun　ぶん", "ki(ku)　き(く)",
		"聞く (to listen, to hear)"},
	{"書", "to-write", "To write", "sho　しょ", "ka(ku)　か(く)", "辞書 (dictionary)"},
	{"読", "to-read", "To read", "doku　どく", "yo(mu)　よ(む)", "読む (to read)"},

	{"言", "word", "Word, To talk", "gen　げん, gon　ごん", "i(u)　い(う)", "言う (to say)"},
	{"話", "to-talk", "To talk, conversation", "wa　わ", "hanashi　はなし, hana(su)　はな(す)",
		"電話 (telephone)"},

	{"買", "to-buy", "To buy", "bai ばい", "ka か(uう)", "買い物 (shopping)"},
	{"行", "to-go", "To go, to carry out", "kou こう", "i い(ku く), okona おこな(u う)", "銀行 (bank)"},
	{"食", "to-eat", "To eat, food", "shokuしょく", "taた (beru べる　)", "食堂 (dining room)"},
	{"会", "to-meet", "To meet, society", "kaiかい, eえ", "aあ(uう)", "会社 (company)"},
	{"少", "few", "A little, few", "shouしょう", "sukoすこ(shiし), sukuすく(naiない)", "少ない (few)"},
	{"新", "new", "New", "shinしん", "ataraあたら(shiiしい)", "新しい (new), 新聞 (newspaper)"},
	{"小", "little", "Little, small", "shouしょう", "chiiちい(saiさい), koこ", "小さい (little)"},
	{"安", "cheap", "Cheap, safety, peace", "anあん", "yasuやす(iい)", "安い (cheap)"},
	{"白", "white", "White", "hakuはく, byakuびゃく", "shiroしろ, しろ(い)",
		"白い (white), 面白い (interesting)"},
}

var fileOfCardsGuru2 = []charSetStructK{
	{"高", "expensive", "Expensive, high", "kouこう", "takaたか(iい)", "高い (expensive)"},
	{"長", "long", "Long, leader", "chouちょう", "nagaなが(iい)", "長い (long), 部長 (manager)"},
	{"古", "old", "Old", "koこ", "furuふる(iい)", "古い (old)"},
	{"休", "to-rest", "To rest, break, holiday, vacation", "kyuu きゅう", "yasu やす(muむ), yasu やす(mi み)",
		"休む (to take a day off)"},
	{"足", "foot", "Foot, to add", "soku　そく", "ashi　あし, ta(su)　た(す)", "足 (foot)"},
	{"魚", "fish", "Fish", "gyo　ぎょ", "sakana　さかな", "魚 (fish)"},
	{"天", "heaven", "Heaven", "ten　てん", "ame　あめ, ama　あま", "天気 (weather)"},
	{"空", "sky", "Sky, empty", "kuu　くう", "sora　そら, a(keru)　あ(ける)", "空 (sky)"},
	{"雨", "rain", "Rain", "u　う", "ame　あめ", "雨 (rain)"},
	{"母", "mother", "Mother", "bo　ぼ", "haha　はは", "母 (mother)"},
	{"友", "friend", "Friend", "yuu　ゆう", "tomo　とも", "友達 (friend)"},
	{"木", "tree", "Tree, wood", "moku　もく, boku　ぼく", "ki　き, ko　こ", "木曜日 (Thursday)"},
	{"店", "shop", "Store", "ten てん", "mise　みせ", "喫茶店 (coffee shop)"},
	{"駅", "station", "Train station", "eki　えき", "-", "駅前 (in front of the station)"},
	{"道", "street", "Street, path, way", "dou (the way)　どう", "michi　みち", "道具 (tool)"},
	{"社", "shrine", "Shrine, society", "sha　しゃ", "yashiro　やしろ", "社長 (president of a company)"},
	{"北", "north", "North", "hoku　ほく", "kita　きた", "北 (north)"},
	{"右", "right", "Right", "yuu　ゆう", "migi　みぎ", "右 (right)"},
	{"左", "left", "Left", "sa　さ", "hidari　ひだり", "左 (left)"},
	{"三日", "tdom", "3rd day of the month", "san　さん", "mi　みか", "三日(3rd day of the month"},
	{"六", "six", "Six", "roku　ろく", "mut(tsu), mu　む", "六日 (6th day of the month)"},
	{"百", "hundred", "Hundred", "hyaku　ひゃく", "—", "百万円 (1 million Yen)"},
	{"万", "ten-thousand", "Ten thousand", "man　まん, ban　ばん", "—",
		"万年筆 (fountain pen, ten thousand year brush)"},
	{"何", "what", "What, which, how many", "ka　か", "nan　なん, nani　なに",
		"何曜日 (what day of the week)"},
	{"土", "earth", "Earth, ground", "do　ど, to　と", "tsuchi　つち", "土曜日 (Saturday)"},
}

var fileOfCardsGuru = []charSetStructK{
	{"大", "big", "Big, a lot", "dai, tai", "oo(kii)", "大きい (big), 大変 (dreadful, immense)"},
	{"多", "many", "A lot, many", "ta", "oo(i)", "多い (many), 多分 (probably)"},
	{"飲", "to-drink", "To drink, a drink", "in", "no(mu)", "飲み物 (beverage)"},
	{"出", "to-leave", "To go out, to leave", "shutsu", "de(ru), da(su)", "出かける (to go out)"},
	{"入", "to-enter", "To enter, to put in", "nyuu", "hai(ru), i(reru", "入口 (entrance)"},
	{"電", "electricity", "Electricity", "den", "—", "電気 (electricity)"},
	{"山", "mountain", "Mountain", "san", "yama", "山 (mountain)"},
	{"川", "river", "River", "sen", "kawa", "川 (river)"},
	{"火", "fire", "Fire", "ka", "hi", "火曜日 (Tuesday)"},
	{"父", "father", "Father", "fu", "chichi", "父 (father)"},
	{"国", "country", "State", "koku", "kuni", "外国人 (foreigner)"},
	{"外", "outside", "External, foreign - to deviate", "gai-ge", "soto-hazu(reru)-hoka",
		"外国 (foreign country)"},
	{"学 or 校 or 学校", "school", "School, learning", "gaku, kou", "mana(bu)",
		"大学 (university)"},
	{"中", "middle", "Middle, center, inner, between", "chuu", "naka",
		"日中 (during the day, midday)"},
	{"上", "above", "Up, above", "shou, jou", "ue, u, a(geru)", "上着 (jacket)"},
	{"下", "down", "Down, below", "ka, ge", "ku(daru), shita", "靴下 (socks)"},
	{"五", "five", "Five", "go", "itsu(tsu), itsu", "五日 (5th day of the month)"},
	{"七", "seven", "Seven", "shichi", "nana(tsu), nana", "七日 (7th day of the month)"},
	{"八", "eight", "Eight", "hachi", "yat(tsu), ya", "八日 (8th day of the month)"},
	{"九", "nine", "Nine", "kyuu, ku", "kokono(tsu)", "九日 (9th day of the month)"},
	{"十", "ten", "Ten", "juu, ji", "tou, to", "十日 (10th day of the month)"},
	{"人", "person", "Person", "じん, にん", "ひと",
		"人々 (people); ひと、人 – person; にんげん、人間 – human; じんるい、人類 – humanity"},
	{"者", "person", "Person", "しゃ, じゃ", "-",
		"わかもの、若者 - young person"},
	{"女", "woman", "Woman, girl, female", "jo じょ, nyo にょ", "onna おんな, me め",
		"おんな、女 – female; おんなのひと、女の人 – woman; おんなのこ、女の子 – girl"},
	{"子", "child", "Child", "shi, su", "ko", "子供 (child)"},
	{"本", "book", "Book, source", "hon", "moto", "日本語 (Japanese)"},
	{"立", "to-stand", "To stand", "ritsu", "ta(tsu)", "立つ (to stand)"},
	{"口", "mouth", "Mouth", "kou, ku", "kuchi", "出口 (exit)"},
	{"手", "hand", "Hand", "shu", "te", "手紙 (letter)"},
}

// The deck will draw cards per this aCard var
var aCard = charSetStructK{}

type charSetStructK struct {
	Kanji        string
	Meaning      string
	Long_Meaning string
	Onyomi       string
	Kunyomi      string
	Vocab        string
}

// Constants:
const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorCyan = "\033[36m"
const colorPurple = "\033[35m"
