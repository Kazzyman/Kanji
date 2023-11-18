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
		"empty Onyomi field", "けんちくか",
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

	{"歯医者 : しかい、、はいしゃ", "dentist", "orthodontist",
		" -", " -",
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

	{"消防士 : しょうぼうし", "firefighter", "fireman",
		" -", " -",
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

	{"私 わたし", "I", "myself",
		" -", " -",
		"vocabulary: ",
		" "},

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
		" -", "おれ、",
		"俺は男です (I am a male[between friends])",
		"I, myself (mainly used by males [informal/casual])"},

	{"警察官 : けいさつかん", "police-officer", "cop",
		" -", " -",
		" – vocab",
		"Police Officer : けいさつかん"},

	{"歌手 : かしゅ", "singer", "vocalist",
		" -", " -",
		" – vocab",
		" "},

	{"兵士 : へいし", "soldier", "warrior",
		" -", " -",
		" – vocab",
		"soldier : へいし"},

	{"先生 : せんせい", "teacher", "instructor",
		" -", " -",
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

	{"くちびる、唇", "lips", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"脛", "shin", "empty second meaning field",
		"empty onyomi field", "すね、",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"ひざ、膝", "knee", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"腿", "thigh", "thighs",
		"empty onyomi field", "もも、",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"あたま、頭", "head", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"かお、顔", "face", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"くち、口", "mouth", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"は、歯", "tooth", "teeth, dentures, dentition",
		"empty onyomi field", "empty kunyomi field",
		"歯医者 : しかい、、はいしゃ, dentist",
		"second empty vocabulary field"},

	{"鼻", "nose", "snozz, beak, proboscis",
		"empty onyomi field", "はな、",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"め、目", "eye", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

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

	{"僕 、ぼく", "I-male", "male-I",
		" -", " -",
		"僕は男です (I am a male)",
		"I, myself (boku, mainly used by males)"},

	{"弁護士 : べんごし", "lawyer", "attorney",
		" -", " -",
		" – vocab",
		"Lawyer : べんごし"},

	{"看護師 : かんごし", "nurse", "Nurse",
		" -", "かんごし",
		" – vocab",
		" "},

	{"看護婦 : かんごふ", "nurse-female", "Female-Nurse",
		" -", " -",
		" – vocab",
		" "},

	{"政治家 : せいじか", "politician", "statesman",
		" -", " -",
		" – vocab",
		"Politician : せいじか"},

	{"教師 : きょうし ", "teacher-academic", "academic-teacher",
		" -", " -",
		" – vocab",
		" "},

	{"時計", "clock", "hour",
		"じ", "とき",
		" 時計 watch",
		" "},

	{"唇 、くちびる", "lips", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"膝 、ひざ", "knee", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"頭 、あたま", "head", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"顔 、かお", "face", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"歯 、は", "tooth", "teeth, dentures, dentition",
		"empty onyomi field", "empty kunyomi field",
		"歯医者 : しかい、、はいしゃ, dentist",
		"second empty vocabulary field"},

	{"髭、鬚、髯 、ひげ", "moustache", "beard",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"髪 、かみ", "hair", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"みみ、耳", "ear", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"御腹 、おなか", "stomach", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"腕 、うで", "arm", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"肘 、ひじ", "elbow", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"肩 、かた", "shoulder", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"つめ、爪", "nail", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"て、手", "hand", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"てくび、手首", "wrist", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"てのひら、掌、手の平", "palm of hand (te-no-hira)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"ゆび、指", "finger, toe (yubi)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"しり、尻", "buttocks (shiri)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"おなか、お腹 （はら、腹）", "abdomen (o-naka)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"かんぞう、肝臓", "liver (kanzō)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"きも、肝", "liver (kimo)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"筋肉 、きんにく", "muscle", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"首 、くび", "neck", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"心 、こころ", "feelings", "heart, feelings",
		"empty onyomi field", "empty kunyomi field",
		"feelings of the heart",
		"second empty vocabulary field"},

	{"腰 、こし", "waist", "hip",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"心臓 、しんぞう", "heart", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"しんぞう (shinzō)",
		"second empty vocabulary field"},

	{"背中 、せなか", "back", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"血 、ち", "blood", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"肉 、にく", "meat", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"肌、膚 、はだ", "skin", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"皮膚 、ひふ", "skin", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"骨 、ほね", "bone", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"胸 、むね", "chest", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"風邪 、かぜ", "cold", "illness",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"下痢 、げり", "diarrhea", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"病気 、びょうき", "illness", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"家族 、かぞく", "family", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"両親 、りょうしん", "parents", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"子供 、こども", "children", "child",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"ちち、父", "father (chichi)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"six(\"otou-san\")",
		"second empty vocabulary field"},

	{"つま、妻", "wife (tsuma)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"おっと、夫", "husband (otto)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"あに、兄", "older brother (ani)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"six (onī-san)",
		"second empty vocabulary field"},

	{"あね、姉", "older sister (ane)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"six (onē-san)",
		"second empty vocabulary field"},

	{"おとうと、弟", "younger brother (otōto)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"いもうと、妹", "younger sister (imōto)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"きょうだい、兄弟", "brothers, siblings (kyōdai)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"しまい、姉妹", "sisters (shimai)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"そふ、祖父", "grandfather (sofu)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"six (ojii-san)",
		"second empty vocabulary field"},

	{"そぼ、祖母", "grandmother (sobo)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"six (obaa-san)",
		"second empty vocabulary field"},

	{"まご、孫", "grandchild (mago)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"おじ、伯父、叔父", "uncle (oji)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"six (oji-san)",
		"second empty vocabulary field"},

	{"おば、伯母、叔母", "aunt (oba)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"six (oba-san)",
		"second empty vocabulary field"},

	{"いとこ、従兄弟、従姉妹、従兄、従弟、従姉、従妹", "cousin (itoko)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"めい、姪", "niece (mei)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"おい、甥", "nephew (oi)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"いきもの、生き物", "living creatures (ikimono)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"ばけもの、化け物", "monster (bakemono)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"どうぶつ、動物", "animal (dōbutsu)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"いぬ、犬", "dog (inu)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"ねこ、猫", "cat (neko)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"うし、牛", "cow (ushi)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"compare: noon 午",
		"second empty vocabulary field"},

	{"ぶた、豚", "pig (buta)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"うま、馬", "horse (uma)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"ひつじ、羊", "sheep (hitsuji)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"さる、猿", "monkey (saru)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"ねずみ、鼠", "mouse, rat (nezumi)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"とら、虎", "tiger (tora)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"オオカミ、狼", "wolf (ōkami)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"うさぎ、兎", "rabbit (usagi)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"竜 、りゅう、たつ", "dragon (ryū, tatsu)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"しか、鹿", "deer (shika)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"かえる、蛙", "frog (kaeru)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"がま、蟇", "toad (gama)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"しし、獅子", "lion (shishi)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"キリン、麒麟", "giraffe (kirin)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"ぞう、象", "elephant (zō)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"とり、鳥", "bird (tori)", "empty second meaning field",
		"empty onyomi field", "empty kunyomi field",
		"first empty vocabulary field",
		"second empty vocabulary field"},

	{"鶏", "chicken", "empty second meaning field",
		"empty onyomi field", "にわとり、 (niwatori)",
		"first empty vocabulary field",
		"second empty vocabulary field"},

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
