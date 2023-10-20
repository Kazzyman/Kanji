package main

/*

	{"ぼく、僕", "I-male", "I, myself (boku, mainly used by males)", " -", " -",
		"僕は男です (I am a male)", " "},

	{"おれ、俺", "I-male-casual", "I, myself (mainly used by males [informal/casual])", " -", " -",
		"俺は男です (I am a male[between friends])", " "},

	{"あたし、私", "I-female", "Softer ver of \"I\" used by females", " -", " -",
		"私は女です (I am a female [softer sounding]) ", " "},

	{"赤ちゃん (aka-chan)", "baby", "or: infant", " -", " -", " – 私は赤女です  あかちゃん、(akachan)", " "},

	{"少女", "girl-small", "Little Girl, 私は女の赤ちゃんです:I am a baby girl ", " -", " -",
		" – 少 means \"a little\": but can also mean \"few\"", " "},

	{"少年", "boy-small", "Small in years, boy?男, しょうねん", " -", " -",
		" – 少 \"a little\": but can also mean \"few\"", " "},

	{"医者", "doctor", "Doctor: いしゃ", " -", " -", " – vocab ", " "},

	{"看護師 : かんごし", "nurse", "Nurse: かんごし", " -", " -", " – vocab", " "},

	{"看護婦 : かんごふ", "nurse-female", "Female Nurse", " -", " -", " – vocab", " "},

	{"歯医者 : しかい、、はいしゃ", "dentist", "Dentist : (shikai, ha-isha)", " -", " -", " – vocab", " "},

	{"政治家 : せいじか", "politician", "Politician : せいじか", " -", " -", " – vocab", " "},

	{"弁護士 : べんごし", "lawyer", "Lawyer : べんごし", " -", " -", " – vocab", " "},

	{"消防士 : しょうぼうし", "firefighter", "Firefighter : しょうぼうし", " -", " -", " – vocab", " "},

	{"警察官 : けいさつかん", "police-officer", "Police Officer : けいさつかん", " -", " -", " – vocab", " "},

	{"兵士 : へいし", "soldier", "soldier : へいし", " -", " -", " – vocab", " "},

	{"建築家 : けんちくか ", "architect", "Architect : けんちくか", " -", " -", " – vocab", " "},

	{"先生 : せんせい", "teacher", "Teacher : せんせい", " -", " -", " – vocab", " "},

	{"教師 : きょうし ", "teacher-academic", "Teacher (academic)", " -", " -", " – vocab", " "},

	{"歌手 : かしゅ", "singer", "Singer", " -", " -", " – vocab", " "},

	{"エンジニア", "engineer", "Engineer", " -", " -", " – vocab", " "},


*/
var fileOfCardsInitiate = []charSetStructKanji{
	{"若者", "youth", "or, Young person", "しゃく, にゃく", "わかい",
		"わかもの、若者 ",
		" "},

	{"時", "time", "or, hour", "じ", "とき",
		"時計 (clock, watch)",
		" "},

	{"私 (わたし)", "I", "Myself, I", " -", " -",
		"vocabulary: ",
		" "},

	{"わたくし、私", "I-formal", "Most formal version of Myself or I", " -", " -",
		"vocabulary: ",
		" "},

	{"ぼく、僕", "I-male", "I, myself (boku, mainly used by males)", " -", " -",
		"僕は男です (I am a male)",
		" "},

	{"おれ、俺", "I-male-casual", "I, myself (mainly used by males [informal/casual])", " -", " -",
		"俺は男です (I am a male[between friends])",
		" "},

	{"あたし、私", "I-female", "Softer ver of \"I\" used by females", " -", " -",
		"私は女です (I am a female [softer sounding]) ",
		" "},

	{"赤ちゃん (aka-chan)", "baby", "or: infant", " -", " -",
		" – 私は赤女です  あかちゃん、(akachan)",
		" "},

	{"少女", "girl-small", "Little Girl, 私は女の赤ちゃんです:I am a baby girl ", " -", " -",
		" – 少 means \"a little\": but can also mean \"few\"",
		" "},

	{"少年", "boy-small", "Small in years, boy?男, しょうねん", " -", " -",
		" – 少 \"a little\": but can also mean \"few\"",
		" "},

	{"医者", "doctor", "Doctor: いしゃ", " -", " -",
		" – vocab ",
		" "},

	{"看護師 : かんごし", "nurse", "Nurse: かんごし", " -", " -",
		" – vocab",
		" "},

	{"看護婦 : かんごふ", "nurse-female", "Female Nurse", " -", " -",
		" – vocab",
		" "},

	{"歯医者 : しかい、、はいしゃ", "dentist", "Dentist : (shikai, ha-isha)", " -", " -",
		" – vocab",
		" "},

	{"政治家 : せいじか", "politician", "Politician : せいじか", " -", " -",
		" – vocab",
		" "},

	{"弁護士 : べんごし", "lawyer", "Lawyer : べんごし", " -", " -",
		" – vocab",
		" "},

	{"消防士 : しょうぼうし", "firefighter", "Firefighter : しょうぼうし", " -", " -",
		" – vocab",
		" "},

	{"警察官 : けいさつかん", "police-officer", "Police Officer : けいさつかん", " -", " -",
		" – vocab",
		" "},

	{"兵士 : へいし", "soldier", "soldier : へいし", " -", " -",
		" – vocab",
		" "},

	{"建築家 : けんちくか ", "architect", "Architect : けんちくか", " -", " -",
		" – vocab",
		" "},

	{"先生 : せんせい", "teacher", "Teacher (sensei) : せんせい", " -", " -",
		" – vocab",
		" "},

	{"教師 : きょうし ", "teacher-academic", "Teacher (academic)", " -", " -",
		" – vocab",
		" "},

	{"歌手 : かしゅ", "singer", "Singer", " -", " -",
		" – vocab",
		" "},

	{"エンジニア", "engineer", "Engineer", " -", " -",
		" – vocab",
		" "},

	{"先", "before", "or; ahead, future", "せん", "さき",
		"先週 (last week), 先生 (teacher, master)",
		" "},

	{"前", "before", "or, in front of", "ぜん", "まえ",
		"名前 (name)",
		" "},
}
