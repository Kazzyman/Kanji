package main

var data_file = []charSetStructKanji{
	{"後", "after", "later, behind",
		"ご, こう", "あと",
		"午後 (afternoon, P.M.)",
		" "},

	{"分", "minute", "part, to understand, to divide",
		"ぶん, ぶ, ふん", "わ(かる)",
		"三十分 (thirty minutes), 自分 (oneself)",
		" "},

	{"月", "month", "moon",
		"げつ, がつ", "つき",
		"月曜日 (Monday)",
		" "},

	{"者", "person", "Person",
		"しゃ, じゃ", " -",
		"わかもの、若者 : young person",
		" "},

	{"学 or 校 or 学校", "school", "learning",
		"がく, こう", "まな(ぶ)",
		"大学 (university)",
		" "},

	{"空", "sky", "empty",
		"くう", "そら, あ(ける)",
		"空 (sky)",
		" "},

	{"十", "ten", "10",
		"じゅう, じ", "とう, と",
		"十日 (10th day of the month)",
		" "},

	{"千", "thousand", "Thousand",
		"せん", "ち",
		"千万円 (10 million Yen)",
		" "},

	{"間", "time", "span of time, Time-frame",
		"かん, けん", "あいだ",
		"時間 (time, hours)",
		" "},

	{"読", "to-read", "To read",
		"どく", "よ(む)",
		"読む (to read)",
		" "},

	{"話", "to-talk", "conversation, to speak",
		"わ", "はなし, はな(す)",
		"電話 (telephone)",
		" "},

	{"書", "to-write", "To write",
		"しょ", "か(く)",
		"辞書 (dictionary)",
		" "},

	{"週", "week", "Week",
		"しゅう", "—",
		"毎週 (every week)",
		" "},

	{"言", "word", "To talk",
		"げん, ごん", "い(う)",
		"言う (to say)",
		" "},

	{"年", "year", "Year",
		"ねん", "とし",
		"今年 (this year), 去年 (last year)",
		" "},

	{"円", "yen", "circle, round",
		"えん", "まるい",
		"円い (round)",
		" "},

	{"車", "car", "or, vehicle",
		"sha しゃ", "kuruma くるま",
		"電車 (electric train)",
		" "},

	{"日", "day", "or, sun",
		"にち, じつ", "ひ, か",
		"明日 (tomorrow)",
		" "},

	{"土", "earth", "ground",
		"do ど, to と", "tsuchi つち",
		"土曜日 (Saturday)",
		" "},

	{"高", "expensive", "high",
		"kou こう", "taka たか(iい)",
		"高い (expensive)",
		" "},

	{"魚", "fish", "Fish",
		"gyo ぎょ", "sakana さかな",
		"魚 (fish)",
		" "},

	{"足", "foot", "or, to add",
		"soku そく", "ashi あし, ta(su) た(す)",
		"足 (foot)",
		" "},

	{"友", "friend", "Friend",
		"yuu ゆう", "tomo とも",
		"友達 (friend)",
		" "},

	{"天", "heaven", "Heaven",
		"ten てん", "ame あめ, ama あま",
		"天気 (weather)",
		" "},

	{"百", "hundred", "Hundred",
		"hyaku ひゃく", "—",
		"百万円 (1 million Yen)",
		" "},

	{"左", "left", "Left",
		"sa さ", "hidari ひだり",
		"左 (left)",
		" "},

	{"長", "long", "or, leader",
		"chou ちょう", "naga なが(iい)",
		"長い (long), 部長 (manager)",
		" "},

	{"母", "mother", "Mother",
		"bo ぼ", "haha はは",
		"母 (mother)",
		" "},

	{"北", "north", "North",
		"hoku ほく", "kita きた",
		"北 (north)",
		" "},

	{"古", "old", "Old",
		"ko こ", "furu ふる(iい)",
		"古い (old)",
		" "},

	{"雨", "rain", "Rain",
		"u う", "ame あめ",
		"雨 (rain)",
		" "},

	{"右", "right", "Right",
		"yuu ゆう", "migi みぎ",
		"右 (right)",
		" "},

	{"店", "shop", "or, Store",
		"ten てん", "mise みせ",
		"喫茶店 (coffee shop)",
		" "},

	{"社", "shrine", "or, society",
		"sha しゃ", "yashiro やしろ",
		"社長 (president of a company)",
		" "},

	{"六", "six", "6",
		"roku ろく", "mut(tsu), mu む",
		"六日 (6th day of the month)",
		" "},

	{"駅", "station", "as in: Train Station",
		"eki えき", " -",
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
		"empty vocabulary field [60 char max]"},

	{"水", "water", "Water",
		"sui すい", "みず、(mizu)",
		"水曜日 (Wednesday)",
		"水兄弟 water brother/siblings"},

	{"白", "white", "White",
		"haku はく, byaku びゃく", "shiro しろ, しろ(い)",
		"白い (white), 面白い (interesting)",
		" "},

	{"建築家", "architect", "designer",
		"ony: -", "けんちくか",
		"empty vocabulary field [60 char max]",
		"empty vocabulary field [60 char max]"},

	{"赤ちゃん", "baby", "infant",
		" -", "あか-ちゃん",
		"私は赤女です (Watashi wa aka onna desu) \"I am a red/intense woman\" ;  あかちゃん、(aka-chan)",
		"赤 (read: \"aka\" or \"seki\") means \"red\" and is used metaphorically for something intense or vivid"},

	{"前", "before", "in-front",
		"ぜん", "まえ",
		"名前 (name)",
		"empty vocabulary field [60 char max]"},

	{"少年", "boy-small", "boy",
		" -", " -",
		" – 少 \"a little\": but can also mean \"few\"",
		"Small in years, boy?男, しょうねん"},

	{"歯医者", "dentist", "orthodontist",
		" -", "しかい、はいしゃ",
		" – vocab",
		"Dentist : (shikai, ha-isha)"},

	{"医者", "doctor", "physician",
		" -", " -",
		" – vocab ",
		"Doctor: いしゃ"},

	{"エンジニア", "engineer", "Engineer",
		" -", " -",
		" – vocab",
		" "},

	{"消防士", "firefighter", "fireman",
		" -", "しょうぼうし",
		" – vocab",
		"Firefighter : しょうぼうし"},

	{"脚", "leg", "foot",
		" -", "あし",
		"foot; base; leg; foundation",
		"Variant of Chinese: 腳  -  Foot足 or Leg脚, あし"},

	{"先", "future", "ahead",
		"せん", "さき",
		"先週 (last week), 先生 (teacher, master)",
		"before "},

	{"少女", "girl-small", "baby-girl",
		" -", " -",
		" – 少 means \"a little\": but can also mean \"few\"",
		"Little Girl, 私は女の赤ちゃんです:I am a baby girl "},

	{"踵", "heel", "visit",
		"しょう (shō), しゅ (shu)", "かかと",
		"踵 of 足, heel of foot",
		"to follow; to visit; to call on  -  かかと Heel: to follow; to visit; to call on"},

	{"私", "I", "myself",
		" -", " -",
		"私、わたし(I)  私、あたし(I-female)  私、わたくし(I-formal)",
		"僕(I-male ぼく)  俺(I-male-casual おれ)"},

	{"私、わたし", "I", "myself",
		" -", " -",
		"私、わたし(I)  私、あたし(I-female)  私、わたくし(I-formal)",
		"僕(I-male ぼく)  俺(I-male-casual おれ)"},

	{"私、あたし", "I-female", "female-I",
		" -", " -",
		"私は女です (I am a female [softer sounding]) ",
		"Softer ver of \"I\" used by females"},

	{"私、わたくし", "I-formal", "formal-I",
		" -", " -",
		"vocabulary: ",
		"Most formal version of Myself or I"},

	{"僕", "I-male", "male-I",
		" -", "ぼく(boku)",
		"僕は男です (I am a male)",
		"I, myself (mainly used by males)"},

	{"俺", "I-male-casual", "I-male-informal",
		" -", "おれ",
		"俺は男です (I am a male[between friends])",
		"I, myself (mainly used by males [informal/casual])"},

	{"警察官", "police-officer", "cop",
		" -", "けいさつかん",
		" – vocab",
		"Police Officer : けいさつかん"},

	{"歌手", "singer", "vocalist",
		" -", "かしゅ",
		" – vocab",
		" "},

	{"兵士", "soldier", "warrior",
		" -", "へいし",
		" – vocab",
		"soldier : へいし"},

	{"先生", "teacher", "instructor",
		" -", "せんせい",
		" – vocab",
		"Teacher (sensei) : せんせい"},

	{"時", "time", "hour",
		"じ", "とき",
		"時計 (clock, watch)",
		" "},

	{"若者", "youth", "young-person",
		"しゃく, にゃく", "わかい",
		"わかもの、若者 ",
		" "},

	{"唇", "lips", "-",
		"ony: -", "くちびる",
		"-",
		"-"},

	{"脛", "shin", "-",
		"ony: -", "すね、",
		"-",
		"-"},

	{"膝", "knee", "-",
		"ony: -", "ひざ",
		"-",
		"-"},

	{"腿", "thigh", "thighs",
		"ony: -", "もも、",
		"-",
		"-"},

	{"頭", "head", "-",
		"ony: -", "あたま",
		"-",
		"-"},

	{"顔", "face", "-",
		"ony: -", "かお",
		"-",
		"-"},

	{"口", "mouth", "-",
		"ony: -", "くち",
		"-",
		"-"},

	{"歯", "tooth", "teeth, dentures, dentition",
		"ony: -", "は",
		"歯医者 : しかい、、はいしゃ, dentist",
		"-"},

	{"鼻", "nose", "snozz, beak, proboscis",
		"ony: -", "はな、",
		"-",
		"-"},

	{"目", "eye", "-",
		"ony: -", "め",
		"-",
		"-"},

	{"山", "mountain", "Mt",
		"san", "yama",
		"山 (mountain), mound",
		"hill"},

	{"人", "person", "human",
		"じん, にん", "ひと",
		"人々 (people); ひと、人 – person; にんげん、人間 – human; じんるい、人類 – humanity",
		"vocab2"},

	{"川", "river", "stream",
		"sen", "kawa",
		"川 (river)",
		"vocab2"},

	{"木", "tree", "wood",
		"moku　もく, boku　ぼく", "ki　き, ko　こ",
		"木曜日 (Thursday)",
		" "},

	{"安", "cheap", "peace",
		"anあん", "yasuやす(iい)",
		"安い (cheap)",
		" - or; safety"},

	{"耳", "ear", "ears",
		"ji　じ", "mimi　みみ",
		"耳 (ear)",
		" "},

	{"東", "east", "East",
		"どう", "ひがし",
		"東京 (Tokyo)",
		" "},

	{"八", "eight", "Eight",
		"hachi", "yat(tsu), ya",
		"八日 (8th day of the month)",
		" "},

	{"毎", "every", "each",
		"mai　まい", "—",
		"毎日(every day)",
		" "},

	{"目", "eye", "ocular",
		"moku　もく", "me　め",
		"目 (eye)",
		" "},

	{"父", "father", "dad",
		"fu", "chichi",
		"父 (father)",
		" "},

	{"少", "few", "a-little",
		"しょう", "すこ(し), すく(ない)",
		"少ない (few)",
		" "},

	{"花", "flower", "blossom",
		"ka　か", "hana　はな",
		"花火 (fireworks)",
		" "},

	{"半", "half", "middle",
		"han　まん", "naka(ba)　なか(ば)",
		"半分 (half)",
		" "},

	{"手", "hand", "pyka",
		"shu", "te",
		"手紙 (letter)",
		" Russian: pyka"},

	{"生", "life", "to-grow",
		"せい, しょう", "i(kiru), u(mareru), ha(yasu)",
		"生徒 (pupil)",
		"to live, to be born, to grow"},

	{"小", "little", "small",
		"shouしょう", "chiiちい(saiさい), koこ",
		"小さい (little)",
		" "},

	{"男", "man", "male",
		"だん, なん", "おとこ",
		"おとこのこ、男の子 – boy; おとこのひと、男の人 – man; おとこ、男 – male",
		" "},

	{"金", "money", "gold",
		"kin　きん, kon　こん", "kane　かね",
		"金曜日 (Friday)",
		" "},

	{"名", "name", "moniker",
		"mei　めい, myou　みょう", "na　な",
		"名前 (name)",
		" "},

	{"新", "new", "fresh",
		"shinしん", "ataraあたら(shiiしい)",
		"新しい (new), 新聞 (newspaper)",
		" "},

	{"九", "nine", "Nine",
		"kyuu, ku", "kokono(tsu)",
		"九日 (9th day of the month)",
		" "},

	{"午", "noon", "midday",
		"ご", "—",
		"午前 (morning, A.M.)",
		"compare: cow 牛"},

	{"今", "now", "the-present",
		"kon　こん, kin　きん", "ima　いま",
		"今晩 (this evening), 今朝 (this morning)",
		" "},

	{"外", "outside", "External",
		"gai-ge", "soto-hazu(reru)-hoka",
		"外国 (foreign country)",
		", foreign - to deviate"},

	{"南", "south", "South",
		"なん", "みなみ",
		"南 (south)",
		" "},

	{"気", "spirit", "ghost",
		"ki　き, ke　け", "—",
		"元気 (healthy, spirit, fine)",
		" "},

	{"買", "to-buy", "to-purchase",
		"bai ばい", "ka か(uう)",
		"買い物 (shopping)",
		" "},

	{"来", "to-come", "approach",
		"rai　らい", "ku(ru)　く(る)",
		"来月 (next month), 来る (to come)",
		" "},

	{"食", "to-eat", "food",
		"しょく, Meal: しょくじ", "た, To Eat: たべる",
		"食堂 (dining room)",
		"食事 means Meal, ‘食’ being pronounced as Onyomi: しょく, the whole word being しょくじ ‘shokuji’"},

	{"行", "to-go", "carry-out",
		"kou : こう", "い(く), okona : おこな(う)",
		"銀行 (bank)",
		" "},

	{"聞", "to-hear", "to-listen",
		"mon　もん, bun　ぶん", "ki(ku)　き(く)",
		"聞く (to listen, to hear)",
		"to listen, to ask"},

	{"会", "to-meet", "society",
		"kaiかい, eえ", "aあ(uう)",
		"会社 (company)",
		" "},

	{"見", "to-see", "to-show",
		"ken　けん", "mi(ru)　み(る)",
		"見せる (to show)",
		"to be visible, to show"},

	{"立", "to-stand", "stand",
		"ritsu", "ta(tsu)",
		"立つ (to stand)",
		" "},

	{"西", "west", "West",
		"さい, せい", "にし",
		"西 (west)",
		" "},

	{"何", "what", "which",
		"ka　か", "nan　なん, nani　なに",
		"何曜日 (what day of the week)",
		"how many"},

	{"僕は男です", "I am a male", "male-I",
		" -", "ぼく は男です",
		"",
		"I, myself (boku, mainly used by males)"},

	{"弁護士", "lawyer", "attorney",
		" -", "べんごし",
		" – vocab",
		"Lawyer : べんごし"},

	{"看護師", "nurse", "Nurse",
		" -", "かんごし",
		" – vocab",
		" "},

	{"看護婦", "nurse-female", "Female-Nurse",
		" -", "かんごふ",
		" – vocab",
		" "},

	{"政治家", "politician", "statesman",
		" -", "せいじか",
		" – vocab",
		"Politician : せいじか"},

	{"教師", "teacher-academic", "academic-teacher",
		" -", "きょうし",
		" – vocab",
		" "},

	{"時計", "clock", "hour",
		"じ", "とき",
		" 時計 watch",
		" "},

	{"唇", "lips", "-",
		"ony: -", "くちびる",
		"-",
		"-"},

	{"膝", "knee", "-",
		"ony: -", "ひざ",
		"-",
		"-"},

	{"頭", "head", "-",
		"ony: -", "あたま",
		"-",
		"-"},

	{"顔", "face", "-",
		"ony: -", "かお",
		"-",
		"-"},

	{"歯", "tooth", "teeth, dentures, dentition",
		"ony: -", "は",
		"歯医者 : しかい、、はいしゃ, dentist",
		"-"},

	{"髭", "moustache 1", "beard",
		"ony: -", "ひげ",
		"髭、鬚、髯",
		"-"},

	{"鬚", "moustache 2", "beard",
		"ony: -", "ひげ",
		"髭、鬚、髯",
		"-"},

	{"髯", "moustache 3", "beard",
		"ony: -", "ひげ",
		"髭、鬚、髯",
		"-"},

	{"髪", "hair", "-",
		"ony: -", "かみ",
		"-",
		"-"},

	{"耳", "ear", "-",
		"ony: -", "みみ",
		"-",
		"-"},

	{"御腹", "stomach", "-",
		"ony: -", "おなか",
		"-",
		"-"},

	{"腕", "arm", "-",
		"ony: -", "うで",
		"-",
		"-"},

	{"肘", "elbow", "-",
		"ony: -", "ひじ",
		"-",
		"-"},

	{"肩", "shoulder", "-",
		"ony: -", "かた",
		"-",
		"-"},

	{"爪", "nail", "-",
		"ony: -", "つめ",
		"-",
		"-"},

	{"手", "hand", "-",
		"ony: -", "て",
		"-",
		"-"},

	{"手首", "wrist", "-",
		"ony: -", "てくび",
		"-",
		"-"},

	{"掌、手の平", "palm of hand (te-no-hira)", "-",
		"ony: -", "てのひら",
		"-",
		"-"},

	{"指", "finger, toe (yubi)", "-",
		"ony: -", "ゆび",
		"-",
		"-"},

	{"尻", "buttocks (shiri)", "-",
		"ony: -", "しり",
		"-",
		"-"},

	{"お腹 （はら、腹）", "abdomen (o-naka)", "-",
		"ony: -", "おなか",
		"-",
		"-"},

	{"肝臓", "liver (kanzō)", "-",
		"ony: -", "かんぞう",
		"-",
		"-"},

	{"肝", "liver (kimo)", "-",
		"ony: -", "きも",
		"-",
		"-"},

	{"筋肉", "muscle", "-",
		"ony: -", "きんにく",
		"-",
		"-"},

	{"首", "neck", "-",
		"ony: -", "くび",
		"-",
		"-"},

	{"心", "feelings", "heart, feelings",
		"ony: -", "こころ",
		"feelings of the heart",
		"-"},

	{"腰", "waist", "hip",
		"ony: -", "こし",
		"-",
		"-"},

	{"心臓", "heart", "-",
		"ony: -", "しんぞう",
		"しんぞう (shinzō)",
		"-"},

	{"背中", "back", "-",
		"ony: -", "せなか",
		"-",
		"-"},

	{"血", "blood", "-",
		"ony: -", "ち",
		"-",
		"-"},

	{"肉", "meat", "-",
		"ony: -", "にく",
		"-",
		"-"},

	{"肌、膚", "skin", "-",
		"ony: -", "はだ",
		"-",
		"-"},

	{"皮膚", "skin", "-",
		"ony: -", "ひふ",
		"-",
		"-"},

	{"骨", "bone", "-",
		"ony: -", "ほね",
		"-",
		"-"},

	{"胸", "chest", "-",
		"ony: -", "むね",
		"-",
		"-"},

	{"風邪", "cold", "illness",
		"ony: -", "かぜ",
		"-",
		"-"},

	{"下痢", "diarrhea", "-",
		"ony: -", "げり",
		"-",
		"-"},

	{"病気", "illness", "-",
		"ony: -", "びょうき",
		"-",
		"-"},

	{"家族", "family", "-",
		"ony: -", "かぞく",
		"-",
		"-"},

	{"両親", "parents", "-",
		"ony: -", " りょうしん",
		"-",
		"-"},

	{"子供", "children", "child",
		"ony: -", " こども",
		"-",
		"-"},

	{"父", "father (chichi)", "-",
		"ony: -", "ちち",
		"six(\"otou-san\")",
		"-"},

	{"妻", "wife (tsuma)", "-",
		"ony: -", "つま",
		"-",
		"-"},

	{"夫", "husband (otto)", "-",
		"ony: -", "おっと",
		"-",
		"-"},

	{"兄", "older brother (ani)", "-",
		"ony: -", "あに",
		"(onī-san)",
		"-"},

	{"姉", "older sister (ane)", "-",
		"ony: -", "あね",
		"(onē-san)",
		"-"},

	{"弟", "younger brother (otōto)", "-",
		"ony: -", "おとうと",
		"-",
		"-"},

	{"妹", "younger sister (imōto)", "-",
		"ony: -", "いもうと",
		"-",
		"-"},

	{"兄弟", "brothers, siblings (kyōdai)", "-",
		"ony: -", "きょうだい",
		"-",
		"-"},

	{"姉妹", "sisters (shimai)", "-",
		"ony: -", "しまい",
		"-",
		"-"},

	{"祖父", "grandfather (sofu)", "-",
		"ony: -", "そふ",
		"(ojii-san)",
		"-"},

	{"祖母", "grandmother (sobo)", "-",
		"ony: -", "そぼ",
		"(obaa-san)",
		"-"},

	{"孫", "grandchild (mago)", "-",
		"ony: -", "まご",
		"-",
		"-"},

	{"叔父", "uncle (oji) 1", "-",
		"ony: -", "おじ",
		"伯父 (oji-san)",
		"-"},

	{"伯父", "uncle (oji) 2", "-",
		"ony: -", "おじ",
		"叔父 (oji-san)",
		"-"},

	{"叔母", "aunt (oba) 1", "-",
		"ony: -", "おば",
		"伯母 (oba-san)",
		"-"},

	{"伯母", "aunt (oba) 2", "-",
		"ony: -", "おば",
		"叔母 (oba-san)",
		"-"},

	{"従兄弟", "cousin (itoko) 1", "-",
		"ony: -", "いとこ",
		"従兄、従弟、従姉、従妹、従兄弟、従姉妹",
		"-"},

	{"従姉妹", "cousin (itoko) 2", "-",
		"ony: -", "いとこ",
		"従兄、従弟、従姉、従妹、従兄弟、従姉妹",
		"-"},

	{"従兄", "cousin (itoko) 3", "-",
		"ony: -", "いとこ",
		"従兄、従弟、従姉、従妹、従兄弟、従姉妹",
		"-"},

	{"従姉", "cousin (itoko) 4", "-",
		"ony: -", "いとこ",
		"従兄、従弟、従姉、従妹、従兄弟、従姉妹",
		"-"},

	{"従妹", "cousin (itoko) 5", "-",
		"ony: -", "いとこ",
		"従兄、従弟、従姉、従妹、従兄弟、従姉妹",
		"-"},

	{"従弟", "cousin (itoko) 6", "-",
		"ony: -", "いとこ",
		"従兄、従弟、従姉、従妹、従兄弟、従姉妹",
		"-"},

	{"姪", "niece (mei)", "-",
		"ony: -", "めい",
		"-",
		"-"},

	{"甥", "nephew (oi)", "-",
		"ony: -", "おい",
		"-",
		"-"},

	{"生き物", "living creatures (ikimono)", "-",
		"ony: -", "いきもの",
		"-",
		"-"},

	{"化け物", "monster (bakemono)", "-",
		"ony: -", "ばけもの",
		"-",
		"-"},

	{"動物", "animal (dōbutsu)", "-",
		"ony: -", "どうぶつ",
		"-",
		"-"},

	{"犬", "dog (inu)", "-",
		"ony: -", "いぬ",
		"-",
		"-"},

	{"猫", "cat (neko)", "-",
		"ony: -", "ねこ",
		"-",
		"-"},

	{"牛", "cow (ushi)", "-",
		"ony: -", "うし",
		"compare: noon 午",
		"-"},

	{"豚", "pig (buta)", "-",
		"ony: -", "ぶた",
		"-",
		"-"},

	{"馬", "horse (uma)", "-",
		"ony: -", "うま",
		"-",
		"-"},

	{"羊", "sheep (hitsuji)", "-",
		"ony: -", "ひつじ",
		"-",
		"-"},

	{"猿", "monkey (saru)", "-",
		"ony: -", "さる",
		"-",
		"-"},

	{"鼠", "mouse, rat (nezumi)", "-",
		"ony: -", "ねずみ",
		"-",
		"-"},

	{"虎", "tiger (tora)", "-",
		"ony: -", "とら",
		"-",
		"-"},

	{"狼", "wolf (ōkami)", "-",
		"ony: -", "オオカミ",
		"-",
		"-"},

	{"兎", "rabbit (usagi)", "-",
		"ony: -", "うさぎ",
		"-",
		"-"},

	{"竜", "dragon (ryū, tatsu)", "-",
		"ony: -", "りゅう、たつ",
		"-",
		"-"},

	{"鹿", "deer (shika)", "-",
		"ony: -", "しか",
		"-",
		"-"},

	{"蛙", "frog (kaeru)", "-",
		"ony: -", "かえる",
		"-",
		"-"},

	{"蟇", "toad (gama)", "-",
		"ony: -", "がま",
		"-",
		"-"},

	{"獅子", "lion (shishi)", "-",
		"ony: -", "しし",
		"-",
		"-"},

	{"麒麟", "giraffe (kirin)", "-",
		"ony: -", "キリン",
		"-",
		"-"},

	{"象", "elephant (zō)", "-",
		"ony: -", "ぞう",
		"-",
		"-"},

	{"鳥", "bird (tori)", "とり",
		"ony: -", "empty kunyomi field",
		"-",
		"-"},

	{"鶏", "chicken", "-",
		"ony: -", "にわとり、 (niwatori)",
		"-",
		"-"},

	{"子", "child", "Child",
		"し, す", "こ",
		"子供 (child)",
		" "},

	{"国", "country", "State",
		"こく", "くに",
		"外国人 (foreigner)",
		" "},

	{"四", "four", "4",
		"し", "よ(つ), よ, よん",
		"四日 (4th day of the month)",
		" "},

	{"語", "language", "word, to chat",
		"ご", "かた(る)",
		"英語 (English)",
		" "},
}
