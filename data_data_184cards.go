package main

var data_file = []charSetStructKanji{
	{"後", "after", "later, behind",
		"ご, こう", "kun: あと",
		"午後 (afternoon, P.M.)",
		" "},

	{"分", "minute", "part, to understand, to divide",
		"ぶん, ぶ, ふん", "kun: わ(かる)",
		"三十分 (thirty minutes), 自分 (oneself)",
		" "},

	{"学 or 校 or 学校", "school", "learning",
		"がく, こう", "kun: まな(ぶ)",
		"大学 (university)",
		" "},

	{"間", "time", "span of time, Time frame",
		"かん, けん", "kun: あいだ",
		"時間 (time, hours)",
		" "},

	{"週", "week", "Week",
		"しゅう", "kun: —",
		"毎週 (every week)",
		" "},

	{"年", "year", "Year",
		"ねん", "kun: とし",
		"今年 (this year), 去年 (last year)",
		" "},

	{"左", "left", "Left",
		"sa さ", "kun: hidari ひだり",
		"左 (left)",
		" "},

	{"母", "mother", "Mother",
		"bo ぼ", "kun: haha はは",
		"母 (mother)",
		" "},

	{"古", "old", "Old",
		"ko こ", "kun: furu ふる(iい)",
		"古い (old)",
		" "},

	{"建築家", "architect", "designer",
		"ony: -", "kun: けんちくか",
		"empty vocabulary field [60 char max]",
		"empty vocabulary field [60 char max]"},

	{"赤ちゃん", "baby", "infant",
		"ony: -", "kun: あか-ちゃん",
		"私は赤女です (Watashi wa aka onna desu) \"I am a red/intense woman\" ;  あかちゃん、(aka-chan)",
		"赤 (read: \"aka\" or \"seki\") means \"red\" and is used metaphorically for something intense or vivid"},

	{"前", "before", "in front",
		"ぜん", "kun: まえ",
		"名前 (name)",
		"empty vocabulary field [60 char max]"},

	{"少年", "boy small", "small boy",
		" -", "kun: -",
		" – 少 \"a little\": but can also mean \"few\"",
		"Small in years, boy?男, しょうねん"},

	{"歯医者", "dentist", "orthodontist",
		" -", "kun: しかい、はいしゃ",
		" – vocab",
		"Dentist : (shikai, ha-isha)"},

	{"医者", "doctor", "physician",
		" -", "kun: -",
		" – vocab ",
		"Doctor: いしゃ"},

	{"エンジニア", "engineer", "Engineer",
		" -", "kun: -",
		" – vocab",
		" "},

	{"消防士", "firefighter", "fireman",
		" -", "kun: しょうぼうし",
		" – vocab",
		"Firefighter : しょうぼうし"},

	{"少女", "girl small", "baby girl",
		" -", "kun: -",
		" – 少 means \"a little\": but can also mean \"few\"",
		"Little Girl, 私は女の赤ちゃんです:I am a baby girl "},

	{"踵", "heel", "visit",
		"しょう (shō), しゅ (shu)", "kun: かかと",
		"踵 of 足, heel of foot",
		"to follow; to visit; to call on  -  かかと Heel: to follow; to visit; to call on"},

	{"私", "I-1", "myself",
		" -", "kun: -",
		"私、わたし(I)  私、あたし(I-female)  私、わたくし(I-formal)",
		"僕(I-male ぼく)  俺(I-male-casual おれ)"},

	{"私、わたし", "I-2", "myself",
		" -", "kun: -",
		"私、わたし(I)  私、あたし(I-female)  私、わたくし(I-formal)",
		"僕(I-male ぼく)  俺(I-male-casual おれ)"},

	{"私、あたし", "I female", "female I",
		" -", "kun: -",
		"私は女です (I am a female [softer sounding]) ",
		"Softer ver of \"I\" used by females"},

	{"私、わたくし", "I formal", "formal I",
		" -", "kun: -",
		"vocabulary: ",
		"Most formal version of Myself or I"},

	{"僕", "I male", "male I",
		" -", "kun: ぼく(boku)",
		"僕は男です (I am a male)",
		"I, myself (mainly used by males)"},

	{"俺", "I male casual", "I male informal",
		" -", "kun: おれ",
		"俺は男です (I am a male[between friends])",
		"I, myself (mainly used by males [informal/casual])" +
			"\ncompare dragon:竜 and note that they are only superficially similar"},

	{"警察官", "police officer", "cop",
		" -", "kun: けいさつかん",
		" – vocab",
		"Police Officer : けいさつかん"},

	{"歌手", "singer", "vocalist",
		" -", "kun: かしゅ",
		" – vocab",
		" "},

	{"兵士", "soldier", "warrior",
		" -", "kun: へいし",
		" – vocab",
		"soldier : へいし"},

	{"先生", "teacher", "instructor",
		" -", "kun: せんせい",
		" – vocab",
		"Teacher (sensei) : せんせい"},

	{"脛", "shin", "vulnerable leg bone",
		"ony: -", "kun: すね、",
		" ",
		" "},

	{"膝", "knee", "kneecap",
		"ony: -", "kun: ひざ",
		" ",
		" "},

	{"腿", "thigh", "thighs",
		"ony: -", "kun: もも、",
		" ",
		" "},

	{"顔", "face", "mug",
		"ony: -", "kun: かお",
		" ",
		" "},

	{"歯", "tooth", "teeth, dentures, dentition",
		"ony: -", "kun: は",
		"歯医者 : しかい、、はいしゃ, dentist",
		" "},

	{"鼻", "nose", "snozz, beak, proboscis",
		"ony: -", "kun: はな、",
		" ",
		" "},

	{"安", "cheap", "peace",
		"anあん", "kun: yasuやす(iい)",
		"安い (cheap),  安 - 2 is peace",
		" - or; safety\n gal wearing a Cheap hat, not a Safety Peace of attire"},

	{"耳", "ear", "ears",
		"ji　じ", "kun: mimi　みみ",
		"耳 (ear)",
		" "},

	{"東", "east", "East",
		"どう", "kun: ひがし",
		"東京 (Tokyo)",
		" "},

	{"八", "eight", "Eight",
		"hachi", "kun: yat(tsu), ya",
		"八日 (8th day of the month)",
		" "},

	{"毎", "every", "each",
		"mai　まい", "kun: -",
		"毎日(every day)",
		" "},

	{"父", "father", "dad",
		"fu", "kun: chichi",
		"父 (father)",
		" "},

	{"少", "few", "a little",
		"しょう", "kun: すこ(し), すく(ない)",
		"少ない (few)",
		" "},

	{"花", "flower", "blossom",
		"ka　か", "kun: hana　はな",
		"花火 (fireworks)",
		" "},

	{"半", "half", "middle",
		"han　まん", "kun: naka(ba)　なか(ば)",
		"半分 (half)",
		" "},

	{"手", "hand", "pyka",
		"shu", "kun: te",
		"手紙 (letter)",
		" Russian: pyka"},

	{"生", "life", "to grow",
		"せい, しょう", "kun: i(kiru), u(mareru), ha(yasu)",
		"生徒 (pupil)",
		"to live, to be born, to grow"},

	{"小", "little", "small",
		"shouしょう", "kun: chiiちい(saiさい), koこ",
		"小さい (little)",
		" "},

	{"男", "man", "male",
		"だん, なん", "kun: おとこ",
		"おとこのこ、男の子 – boy; おとこのひと、男の人 – man; おとこ、男 – male",
		" "},

	{"金", "money", "gold",
		"kin　きん, kon　こん", "kun: kane　かね",
		"金曜日 (Friday)",
		" "},

	{"名", "name", "moniker",
		"mei　めい, myou　みょう", "kun: na　な",
		"名前 (name)",
		" "},

	{"新", "new", "fresh",
		"shinしん", "kun: ataraあたら(shiiしい)",
		"新しい (new), 新聞 (newspaper)",
		" "},

	{"九", "nine", "9",
		"kyuu, ku", "kun: kokono(tsu)",
		"九日 (9th day of the month)",
		" "},

	{"午", "noon", "midday",
		"ご", "kun: -",
		"午前 (morning, A.M.)",
		"\ncompare: cow 牛"},

	{"今", "now", "the present",
		"kon　こん, kin　きん", "kun: ima　いま",
		"今晩 (this evening), 今朝 (this morning)",
		" "},

	{"外", "outside", "External",
		"gai-ge", "kun: soto-hazu(reru)-hoka",
		"外国 (foreign country)",
		", foreign - to deviate"},

	{"南", "south", "South",
		"なん", "kun: みなみ",
		"南 (south)",
		" "},

	{"気", "spirit", "ghost",
		"ki　き, ke　け", "kun: -",
		"元気 (healthy, spirit, fine)",
		" "},

	{"買", "to buy", "to purchase",
		"bai ばい", "kun: ka か(uう)",
		"買い物 (shopping)",
		" "},

	{"来", "to come", "approach",
		"rai　らい", "kun: ku(ru)　く(る)",
		"来月 (next month), 来る (to come)",
		" "},

	{"食", "to eat", "food",
		"しょく, Meal: しょくじ", "kun: た, To Eat: たべる",
		"食堂 (dining room)",
		"食事 means Meal, ‘食’ being pronounced as Onyomi: しょく, the whole word being しょくじ ‘shokuji’"},

	{"行", "to go", "carry out",
		"kou : こう", "kun: い(く), okona : おこな(う)",
		"銀行 (bank)",
		" "},

	{"聞", "to hear", "to listen",
		"ony: mon　もん, bun　ぶん", "kun: ki(ku)　き(く)",
		"聞く (to listen, to hear)",
		"to listen, to ask"},

	{"会", "to meet", "society",
		"kaiかい, eえ", "kun: aあ(uう)",
		"会社 (company)",
		" "},

	{"見", "to see", "to show",
		"ken　けん", "kun: mi(ru)　み(る)",
		"見せる (to show)",
		"to be visible, to show"},

	{"立", "to stand", "stand",
		"ritsu", "kun: ta(tsu)",
		"立つ (to stand)",
		" "},

	{"西", "west", "West",
		"さい, せい", "kun: にし",
		" ",
		" \ncompare parents:両親 noting only a superficial similarity"},

	{"何", "what", "which",
		"ka　か", "kun: nan　なん, nani　なに",
		"何曜日 (what day of the week)",
		"how many"},

	{"僕は男です", "I am a male", "male I",
		" -", "kun: ぼく は男です",
		"",
		"I, myself (boku, mainly used by males)"},

	{"弁護士", "lawyer", "attorney",
		" -", "kun: べんごし",
		" – vocab",
		"Lawyer : べんごし"},

	{"看護師", "nurse", "Nurse",
		" -", "kun: かんごし",
		" – vocab",
		" "},

	{"看護婦", "nurse female", "Female Nurse",
		" -", "kun: かんごふ",
		" – vocab",
		" "},

	{"政治家", "politician", "statesman",
		" -", "kun: せいじか",
		" – vocab",
		"Politician : せいじか"},

	{"教師", "teacher academic", "academic teacher",
		" -", "kun: きょうし",
		" – vocab",
		" "},

	{"膝", "knee", "kneecap",
		"ony: -", "kun: ひざ",
		" ",
		" "},

	{"顔", "face", "mug",
		"ony: -", "kun: かお",
		" ",
		" "},

	{"歯", "tooth", "teeth, dentures, dentition",
		"ony: -", "kun: は",
		"歯医者 : しかい、、はいしゃ, dentist",
		" "},

	{"髭", "moustache", "beard",
		"ony: -", "kun: ひげ",
		"髭、鬚、髯",
		" "},

	{"鬚", "moustache", "beard",
		"ony: -", "kun: ひげ",
		"髭、鬚、髯",
		" "},

	{"髯", "moustache", "beard",
		"ony: -", "kun: ひげ",
		"髭、鬚、髯",
		" "},

	{"髪", "hair", "wig",
		"ony: -", "kun: かみ",
		" ",
		" "},

	{"耳", "ear", "auditory",
		"ony: -", "kun: みみ",
		" ",
		" "},

	{"御腹", "stomach", "gastric",
		"ony: -", "kun: おなか",
		" ",
		" "},

	{"腕", "arm", "Arm",
		"ony: -", "kun: うで",
		" ",
		" "},

	{"肘", "elbow", "funny bone",
		"ony: -", "kun: ひじ",
		" ",
		" "},

	{"肩", "shoulder", "Shoulder",
		"ony: -", "kun: かた",
		" ",
		" "},

	{"爪", "nail", "finger blade",
		"ony: -", "kun: つめ",
		" ",
		" "},

	{"手", "hand", "Pyka",
		"ony: -", "kun: て",
		" ",
		" "},

	{"手首", "wrist", "hand something, complete this card",
		"ony: -", "kun: てくび",
		" ",
		" "},

	{"掌、手の平", "palm of hand (te-no-hira)", "complete this card",
		"ony: -", "kun: てのひら",
		" ",
		" "},

	{"指", "finger", "toe",
		"ony: -", "kun: ゆび",
		"finger, toe (yubi)",
		" "},

	{"尻", "buttocks", "glut",
		"ony: -", "kun: しり",
		" shiri",
		" "},

	{"お腹 （はら、腹）", "abdomen", "abs",
		"ony: -", "kun: おなか",
		"  (o-naka)",
		" "},

	{"肝臓", "liver", "belly organ, kanzō",
		"ony: -", "kun: かんぞう",
		"(kanzō)",
		" "},

	{"肝", "liver", "kimo",
		"ony: -", "kun: きも",
		" kimo",
		" "},

	{"筋肉", "muscle", "complete this card",
		"ony: -", "kun: きんにく",
		" ",
		" "},

	{"心", "feelings", "heart",
		"ony: -", "kun: こころ",
		"feelings of the heart",
		" "},

	{"腰", "waist", "hip",
		"ony: -", "kun: こし",
		" ",
		" "},

	{"心臓", "heart", "complete this card",
		"ony: -", "kun: しんぞう",
		"しんぞう (shinzō)",
		" "},

	{"背中", "back", "complete this card",
		"ony: -", "kun: せなか",
		" ",
		" "},

	{"血", "blood", "liquid tissue",
		"ony: -", "kun: ち",
		" ",
		" "},

	{"肉", "meat", "complete this card",
		"ony: -", "kun: にく",
		" ",
		" "},

	{"肌、膚", "skin", "largest organ",
		"ony: -", "kun: はだ",
		" ",
		" "},

	{"皮膚", "skin", "complete this card",
		"ony: -", "kun: ひふ",
		" ",
		" "},

	{"骨", "bone", "complete this card",
		"ony: -", "kun: ほね",
		" ",
		" "},

	{"胸", "chest", "complete this card",
		"ony: -", "kun: むね",
		" ",
		" "},

	{"風邪", "cold", "illness",
		"ony: -", "kun: かぜ",
		" ",
		" "},

	{"下痢", "diarrhea", "liquid poo",
		"ony: -", "kun: げり",
		" ",
		" "},

	{"病気", "illness", "complete this card",
		"ony: -", "kun: びょうき",
		" ",
		" "},

	{"家族", "family", "complete this card",
		"ony: -", "kun: かぞく",
		" ",
		" "},

	{"両親", "parents", "complete this card",
		"ony: -", "kun: りょうしん",
		" ",
		"\ncompare west:西 "},

	{"子供", "children", "child",
		"ony: -", "kun: こども",
		" ",
		" "},

	{"父", "father", "chichi",
		"ony: -", "kun: ちち",
		"six(\"otou-san\")",
		"chichi"},

	{"妻", "wife", "tsuma",
		"ony: -", "kun: つま",
		" ",
		" "},

	{"夫", "husband", "otto",
		"ony: -", "kun: おっと",
		" ",
		" "},

	{"兄", "older brother (ani)", "onī-san",
		"ony: -", "kun: あに",
		"(onī-san)",
		" "},

	{"姉", "older sister", "ane",
		"ony: -", "kun: あね",
		"(onē-san)",
		" "},

	{"弟", "younger brother", "otōto",
		"ony: -", "kun: おとうと",
		" ",
		" "},

	{"妹", "younger sister", "imōto",
		"ony: -", "kun: いもうと",
		" ",
		" "},

	{"兄弟", "brothers, siblings", "kyōdai",
		"ony: -", "kun: きょうだい",
		" ",
		" "},

	{"姉妹", "sisters", "shimai",
		"ony: -", "kun: しまい",
		" ",
		" "},

	{"祖父", "grandfather", "sofu",
		"ony: -", "kun: そふ",
		"(ojii-san)",
		" "},

	{"祖母", "grandmother (sobo)", " ",
		"ony: -", "kun: そぼ",
		"(obaa-san)",
		" "},

	{"孫", "grandchild (mago)", " ",
		"ony: -", "kun: まご",
		" ",
		" "},

	{"叔父", "uncle (oji) 1", " ",
		"ony: -", "kun: おじ",
		"伯父 (oji-san)",
		" "},

	{"伯父", "uncle (oji) 2", " ",
		"ony: -", "kun: おじ",
		"叔父 (oji-san)",
		" "},

	{"叔母", "aunt (oba) 1", " ",
		"ony: -", "kun: おば",
		"伯母 (oba-san)",
		" "},

	{"伯母", "aunt (oba) 2", " ",
		"ony: -", "kun: おば",
		"叔母 (oba-san)",
		" "},

	{"従兄弟", "cousin (itoko) 1", " ",
		"ony: -", "kun: いとこ",
		"従兄、従弟、従姉、従妹、従兄弟、従姉妹",
		" "},

	{"従姉妹", "cousin (itoko) 2", " ",
		"ony: -", "kun: いとこ",
		"従兄、従弟、従姉、従妹、従兄弟、従姉妹",
		" "},

	{"従兄", "cousin (itoko) 3", " ",
		"ony: -", "kun: いとこ",
		"従兄、従弟、従姉、従妹、従兄弟、従姉妹",
		" "},

	{"従姉", "cousin (itoko) 4", " ",
		"ony: -", "kun: いとこ",
		"従兄、従弟、従姉、従妹、従兄弟、従姉妹",
		" "},

	{"従妹", "cousin (itoko) 5", " ",
		"ony: -", "kun: いとこ",
		"従兄、従弟、従姉、従妹、従兄弟、従姉妹",
		" "},

	{"従弟", "cousin (itoko) 6", " ",
		"ony: -", "kun: いとこ",
		"従兄、従弟、従姉、従妹、従兄弟、従姉妹",
		" "},

	{"姪", "niece (mei)", " ",
		"ony: -", "kun: めい",
		" ",
		" "},

	{"甥", "nephew (oi)", " ",
		"ony: -", "kun: おい",
		" ",
		" "},

	{"生き物", "living creatures (ikimono)", " ",
		"ony: -", "kun: いきもの",
		" ",
		" "},

	{"化け物", "monster (bakemono)", " ",
		"ony: -", "kun: ばけもの",
		" ",
		" "},

	{"動物", "animal (dōbutsu)", " ",
		"ony: -", "kun: どうぶつ",
		" ",
		" "},

	{"犬", "dog (inu)", " ",
		"ony: -", "kun: いぬ",
		" ",
		" "},

	{"猫", "cat (neko)", " ",
		"ony: -", "kun: ねこ",
		" ",
		" "},

	{"牛", "cow", " ",
		"ony: -", "kun: うし (ushi)",
		"\ncompare: noon 午, thousand 千",
		" "},

	{"豚", "pig (buta)", " ",
		"ony: -", "kun: ぶた",
		" ",
		" "},

	{"馬", "horse (uma)", " ",
		"ony: -", "kun: うま",
		" ",
		" "},

	{"羊", "sheep (hitsuji)", " ",
		"ony: -", "kun: ひつじ",
		" ",
		" "},

	{"猿", "monkey (saru)", " ",
		"ony: -", "kun: さる",
		" ",
		" "},

	{"鼠", "mouse, rat (nezumi)", " ",
		"ony: -", "kun: ねずみ",
		" ",
		" "},

	{"虎", "tiger (tora)", " ",
		"ony: -", "kun: とら",
		" ",
		" "},

	{"狼", "wolf (ōkami)", " ",
		"ony: -", "kun: オオカミ",
		" ",
		"the white Wolf is a long-standing friend, beside a mexican in a sombrero"},

	{"兎", "rabbit (usagi)", " ",
		"ony: -", "kun: うさぎ",
		" ",
		" "},

	{"竜", "dragon", " ",
		"ony: -", "kun: りゅう、たつ (ryū, tatsu)",
		"  ",
		"\ncompare casual:俺 and note that they are only superficially similar"},

	{"鹿", "deer (shika)", " ",
		"ony: -", "kun: しか",
		" ",
		" "},

	{"蛙", "frog (kaeru)", " ",
		"ony: -", "kun: かえる",
		" ",
		" "},

	{"蟇", "toad (gama)", " ",
		"ony: -", "kun: がま",
		" ",
		" "},

	{"獅子", "lion (shishi)", " ",
		"ony: -", "kun: しし",
		" ",
		" "},

	{"麒麟", "giraffe (kirin)", " ",
		"ony: -", "kun: キリン",
		" ",
		" "},

	{"象", "elephant (zō)", " ",
		"ony: -", "kun: ぞう",
		" ",
		" "},

	{"鳥", "bird (tori)", "とり",
		"ony: -", "kun: -",
		"\ncompare fish:魚",
		" "},

	{"鶏", "chicken", "food bird",
		"ony: -", "kun: にわとり、 (niwatori)",
		" ",
		" "},

	{"子", "child", "Child",
		"し, す", "kun: こ",
		"子供 (child)",
		" "},

	{"国", "country", "State",
		"こく", "kun: くに",
		"外国人 (foreigner)",
		" "},

	{"四", "four", "4",
		"し", "kun: よ(つ), よ, よん",
		"四日 (4th day of the month)",
		" "},

	{"語", "language", "word, to chat",
		"ご", "kun: かた(る)",
		"英語 (English)",
		" "},
}
