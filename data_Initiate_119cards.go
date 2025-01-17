package main

var fileOfCardsInitiate = []charSetStructKanji{

	{"建築家", "architect", "designer",
		" -", " -",
		" – vocab",
		"Architect : けんちくか  "},

	{"赤ちゃん", "baby", "infant",
		" -", " -",
		" – 私は赤女です  あかちゃん、(aka-chan),  (あか-ちゃん)",
		" "},

	{"前", "before", "in front",
		"ぜん", "まえ",
		"名前 (name)",
		" "},

	{"少年", "boy small", "boy",
		" -", " -",
		" – 少 \"a little\": but can also mean \"few\"",
		"Small in years, boy?男, しょうねん"},

	{"歯医者", "dentist", "orthodontist",
		" -", " -",
		" – vocab",
		"Dentist : (shikai, ha-isha),  : しかい、、はいしゃ"},

	{"医者", "doctor", "physician",
		" -", " -",
		" – vocab ",
		"Doctor: いしゃ"},

	{"エンジニア", "engineer", "Engineer",
		" -", " -",
		" – vocab",
		" "},

	{"消防士", "firefighter", "fireman",
		" -", " -",
		" – vocab",
		"Firefighter : しょうぼうし  "},

	{"少女", "girl small", "baby girl",
		" -", " -",
		" – 少 means \"a little\": but can also mean \"few\"",
		"Little Girl, 私は女の赤ちゃんです:I am a baby girl "},

	{"踵", "heel", "visit",
		"しょう (shō), しゅ (shu)", "かかと",
		"踵 of 足, heel of foot",
		"to follow; to visit; to call on  -  かかと Heel: to follow; to visit; to call on"},

	{"私", "I", "myself",
		" -", " -",
		"vocabulary: ",
		"  (わたし)"},

	{"私", "I female", "female I",
		" -", " -",
		"私は女です (I am a female [softer sounding]) ",
		"あたし、Softer ver of \"I\" used by females"},

	{"私", "I formal", "formal I",
		" -", " -",
		"わたくし、: ",
		"Most formal version of Myself or I"},

	{"僕", "I male", "male I",
		" -", " -",
		"僕は男です (I am a male)",
		"I, myself (boku, mainly used by males) ぼく、"},

	{"俺", "I male casual", "I male informal",
		" -", " -",
		"俺は男です (I am a male[between friends])",
		"I, myself (mainly used by males [informal/casual]) おれ、"},

	{"弁護士", "lawyer", "attorney",
		" -", " -",
		" – vocab",
		"Lawyer : べんごし  "},

	{"看護師", "nurse", "Nurse",
		" -", " -",
		" – vocab",
		"  : かんごし"},

	{"看護婦", "nurse female", "Female Nurse",
		" -", " -",
		" – vocab",
		"  : かんごふ"},

	{"警察官", "police officer", "cop",
		" -", " -",
		" – vocab",
		"Police Officer : けいさつかん  "},

	{"政治家", "politician", "statesman",
		" -", " -",
		" – vocab",
		"Politician : せいじか"},

	{"歌手", "singer", "vocalist",
		" -", " -",
		" – vocab",
		"  : かしゅ"},

	{"兵士", "soldier", "warrior",
		" -", " -",
		" – vocab",
		"soldier : へいし "},

	{"先生", "teacher", "instructor",
		" -", " -",
		" – vocab",
		"Teacher (sensei) : せんせい"},

	{"教師", "teacher academic", "academic teacher",
		" -", " -",
		" –  : きょうし ",
		" "},

	{"若者", "youth", "young person",
		"ony: ジャク(jaku), ニャク(nyaku)",
		"わかい, わかもの(wakamono)",
		"若者 (youth), 青年 (young person)",
		"Compare 老 meaning old or aged"},

	{"唇", "lips",
		" ",
		"ony: -",
		"kun: くちびる",
		" ",
		" "},

	{"脛", "shin",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" すね、"},

	{"膝", "knee",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" ひざ、"},

	{"顔", "face",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" かお、"},

	{"歯", "tooth",
		"teeth, dentures, dentition",
		"ony: -",
		"kun: -",
		"歯医者 : しかい、、はいしゃ, dentist",
		" は、"},

	{"鼻", "nose",
		"snozz, beak, proboscis",
		"ony: -",
		"kun: -",
		" ",
		" はな、"},

	{"髭、鬚、髯", "moustache",
		"beard",
		"ony: -",
		"kun: -",
		" ",
		" ひげ、 compare hair 髪"},

	{"髪", "hair",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" かみ、  compare moustache, beard 髭、鬚、髯"},

	{"耳", "ear",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" みみ、"},

	{"御腹", "stomach",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" おなか、"},

	{"腕", "arm",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" うで、"},

	{"肘", "elbow",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" ひじ、"},

	{"肩", "shoulder",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" かた、"},

	{"爪", "nail",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" つめ、"},

	{"手", "hand",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" て、"},

	{"手首", "wrist",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" てくび、"},

	{"掌、手の平", "palm",
		"palm of hand",
		"ony: -",
		"kun: -",
		" ",
		" てのひら、 (te-no-hira)"},

	{"指", "finger",
		"toe",
		"ony: -",
		"kun: -",
		" ",
		" ゆび、 (yubi)"},

	{"尻", "buttocks",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" (shiri) しり、"},

	{"腹", "abdomen",
		" ",
		"ony: -",
		"kun: -",
		"お腹 (o-naka)   （はら、腹）",
		" おなか、"},

	{"肝臓", "liver",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" かんぞう、 (kanzō)"},

	{"肝", "liver",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" きも、 (kimo)"},

	{"筋肉", "muscle",
		" ",
		"ony: -",
		"kun: -",
		" ",
		"きんにく、 (kin'niku)"},

	{"首", "neck",
		" ",
		"ony: -",
		"kun: くび (kubi)",
		" ",
		" "},

	{"心", "heart",
		"heart [as in feelings] (kokoro)",
		"ony: -",
		"kun: -",
		" ",
		"こころ、  (kokoro)"},

	{"腰", "waist",
		"hip",
		"ony: -",
		"kun: -",
		" ",
		" こし、 (koshi)"},

	{"心臓", "heart",
		" ",
		"ony: -",
		"kun: -",
		" ",
		"しんぞう、  (shinzō)"},

	{"背中", "back",
		" ",
		"ony: -",
		"kun: -",
		" ",
		"せなか、  (senaka)"},

	{"血", "blood",
		" ",
		"ony: -",
		"kun: -",
		" ",
		"ち、  (chi)"},

	{"肉", "meat",
		" ",
		"ony: -",
		"kun: -",
		" ",
		"にく、  (niku)"},

	{"肌、膚", "skin",
		" ",
		"ony: -",
		"kun: -",
		" ",
		"はだ、  (hada)"},

	{"皮膚", "skin",
		" ",
		"ony: -",
		"kun: -",
		" ",
		"ひふ、  (hifu)"},

	{"ほね、骨", "bone",
		"hone",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"むね、胸", "chest",
		"mune",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"かぜ、風邪", "cold",
		"illness, kaze",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"げり、下痢", "diarrhea",
		"geri",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"びょうき、病気", "illness (byōki)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"かぞく、家族", "family (kazoku)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"りょうしん、両親", "parents (ryoushin)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"こども、子供", "children, child (kodomo)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"ちち、父", "father (chichi)",
		" ",
		"ony: -",
		"kun: -",
		"six(\"otou-san\")",
		" "},

	{"はは、母", "mother (haha)",
		" ",
		"ony: -",
		"kun: -",
		"six(\"okaa-san\")",
		" "},

	{"つま、妻", "wife (tsuma)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"おっと、夫", "husband (otto)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"あに、兄", "older brother (ani)",
		" ",
		"ony: -",
		"kun: -",
		"six (onī-san)",
		" "},

	{"あね、姉", "older sister (ane)",
		" ",
		"ony: -",
		"kun: -",
		"six (onē-san)",
		" "},

	{"おとうと、弟", "younger brother (otōto)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"いもうと、妹", "younger sister (imōto)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"きょうだい、兄弟", "brothers, siblings (kyōdai)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"しまい、姉妹", "sisters (shimai)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"そふ、祖父", "grandfather (sofu)",
		" ",
		"ony: -",
		"kun: -",
		"six (ojii-san)",
		" "},

	{"そぼ、祖母", "grandmother (sobo)",
		" ",
		"ony: -",
		"kun: -",
		"six (obaa-san)",
		" "},

	{"まご、孫", "grandchild (mago)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"おじ、伯父、叔父", "uncle (oji)",
		" ",
		"ony: -",
		"kun: -",
		"six (oji-san)",
		" "},

	{"おば、伯母、叔母", "aunt (oba)",
		" ",
		"ony: -",
		"kun: -",
		"six (oba-san)",
		" "},

	{"いとこ、従兄弟、従姉妹、従兄、従弟、従姉、従妹", "cousin (itoko)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"めい、姪", "niece (mei)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"おい、甥", "nephew (oi)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"いきもの、生き物", "living creatures (ikimono)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"ばけもの、化け物", "monster (bakemono)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"どうぶつ、動物", "animal (dōbutsu)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"チーター", "cheetah (chītā)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"いぬ、犬", "dog (inu)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"ねこ、猫", "cat (neko)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"うし、牛", "cow (ushi)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"ぶた、豚", "pig (buta)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"うま、馬", "horse (uma)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"ひつじ、羊", "sheep (hitsuji)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"さる、猿", "monkey (saru)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"ねずみ、鼠", "mouse, rat (nezumi)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"とら、虎", "tiger (tora)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"オオカミ、狼", "wolf (ōkami)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"うさぎ、兎", "rabbit (usagi)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"りゅう、たつ、竜", "dragon (ryū, tatsu)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"しか、鹿", "deer (shika)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"かえる、蛙", "frog (kaeru)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"がま、蟇", "toad (gama)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"しし、獅子", "lion (shishi)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"キリン、麒麟", "giraffe (kirin)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"ぞう、象", "elephant (zō)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"とり、鳥", "bird (tori)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"にわとり、鶏", "chicken (niwatori)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},
	// to here

	{"じかん、時間", "time (jikan)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"とき、じ、時", "~hours (toki, ji)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"ふん、分", "minute (fun)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"ひ、にち、日", "day (hi, nichi)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"つき、がつ、月", "month (tsuki, gatsu)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},

	{"とし、ねん、年", "year (toshi, nen)",
		" ",
		"ony: -",
		"kun: -",
		" ",
		" "},
}
