package main

var data_beauty = []charSetStructKanji{

	{"美しすぎる", "too beautiful", "exquisite",
		"ony: ウツクシスギル",
		"kun: -",
		"絶世の美女(zesse no bijo), 天上天下無双の美(tenjō tenka musō no utsukushisa), 余人をもって代え難い(yojin wo motte kaegataki), " +
			"言葉に尽くし難い(kotoba ni tsukushikataki)",
		"超絶的なる美(chōzetsu-teki naru utsukushisa), 驚異的とも言えるほど(kyōgiteki tomo ieru hodo)"},

	{"衝撃的な", "startling", "shocking",
		"ony: ショウゲキテキナ",
		"kun: -",
		"驚くべき(odorokubeki), 愕然(gakuzen), すさまじい(susamajii)",
		"予期せぬ(yoki se nu), 衝迫(shouhaku), 破廉恥(harenchi)"},

	{"鮮やかな", "vivid", "bright",
		"ony: アザヤカナ",
		"kun: -",
		"明るい(akarui), 鋭い(surudoi), 強い(tsuyoi)",
		"色鮮やか(iro azayaka), 艶やか(tsuyayaka), 輝く(kagayaku)"},

	{"素晴らしい", "wonderful", "splendid",
		"ony: スバラシイ",
		"kun: -",
		"すばらしい(subarashii), 素敵(suteki), 見事(migoto)",
		"めでたい(medetai), 優れている(sugurete iru), 秀逸(shūitsu)"},

	{"ファンタスティックな", "fantastic", "extraordinary",
		"ony: ファンタスティックナ",
		"kun: -",
		"素晴らしい(subarashii), すごい(sugoi), 驚異的(kyōgiteki)",
		"抜群の(bakkugun no), 飛び抜けた(tobinuketa), 逸品(ippin)"},

	{"素敵な", "charming", "lovely",
		"ony: ステキナ",
		"kun: -",
		"可憐(kareru), 美しい(utsukushii), 良い(yoi)",
		"気持ち良い(kimochi yoī), すてき(suteki), 魅力(miryoku)"},

	{"超美人", "extremely beautiful woman", "gorgeous",
		"ony: チョウビジン",
		"kun: -",
		"美人(bijin), きれい(kirei), 素晴らしい(subarashii), 容姿(youshi), 容色(keshiki), 風采(fuu-sai)",
		"容姿(youshi), 容色(keshiki), 風采(fuu-sai), 見目麗しい(mimomezurashii), 媚び(kobiru), 麗美(reibi)"},

	{"絢爛", "gorgeous", "lavish",
		"ony: ケンラン",
		"kun: -",
		"美しい(utsukushii), 壮大(soudai), 豪華(gouka), 絵画的(kaiga-teki), 見事(migoto), 鮮やか(azayaka)",
		"絵画的(kaiga-teki), 見事(migoto), 鮮やか(azayaka), 艶やか(tsuyayaka), 魅力(miryoku), 豪壮(gousou)"},

	{"魅力的な", "captivating", "fascinating",
		"ony: ミリョクテキナ",
		"kun: -",
		"魅力(miryoku), 魅せる(miseru), 誘惑(yūwaku), 魅惑的(miwaku-tekina), 虜にする(toriko ni suru), 魔力(maryoku)",
		"惹きつける(hikitsukeru), 魅了(miryō), 誘う(sasou)"},

	{"信じらほど", "unbelievably", "incredibly, improbably, ridiculously",
		"ony: シンジラレナイホド",
		"kun: -",
		"とても(totemo), すごく(sugoku), 非常に(hijōni)",
		"想像を絶する(sōzō wo tayasu), 余りにも(amari ni mo), 途方もなく(tohō mo naku)"},

	/*
		{"衝撃的な", "startling","shocking",
			"ony: ショウゲキテキナ",
			"kun: -",
			"驚くべき(odorokubeki), 愕然(gakuzen), すさまじい(susamajii)",
			"予期せぬ(yoki se nu), 衝迫(shouhaku), 破廉恥(harenchi)",
			"開口一番(kaikou ichi ban), 震撼(shinkan), 駭然(kyouzen)"
		},

		{"鮮やかな", "vivid", "bright",
			"ony: アザヤカナ",

			"kun: -",
			"明るい(akarui), 鋭い(surudoi), 強い(tsuyoi)",
			"色鮮やか(iro azayaka), 艶やか(tsuyayaka), 輝く(kagayaku)",

			"色合い(iroai), 映える(haeru), 鮮明(senmei)"
		},

		{"素晴らしい", "wonderful", "splendid",
			"ony: スバラシイ",
			"kun: -",
			"すばらしい(subarashii), 素敵(suteki), 見事(migoto)",
			"めでたい(medetai), 優れている(sugurete iru), 秀逸(shūitsu)",
			"申し分ない(mōshiwake nai), 絶品(zeppin), 抜群の(bakugun no)"
		},

		{"ファンタスティックな", "fantastic", "extraordinary",
			"ony: ファンタスティックナ",

			"kun: -",
			"素晴らしい(subarashii), すごい(sugoi), 驚異的(kyōgiteki)",
			"抜群の(bakkugun no), 飛び抜けた(tobinuketa), 逸品(ippin)",
			"大変良い(taihen yoi), 非の打ち所がない(hi no uchidokoro ga nai), 絶妙(zetsumyō)"
		},

		{"素敵な", "charming", "lovely",
			"ony: ステキナ",
			"kun: -",

			"可憐(kareru), 美しい(utsukushii), 良い(yoi)",
			"気持ち良い(kimochi yoī), すてき(suteki), 魅力(miryoku)",
			"愛らしい(airashī), すばらしい(subarashī), 美(bi)"
		},

	*/

	/*
		超美人 (chō bijin) - Extremely beautiful woman

		Bijin originally referred to cultured courtesans. Now indicates exceptional beauty.
		Chō intensifies bijin, implying appearance exceeds the norm. Connotes unrealistic ideal.
		Has connotation of dramatic, showstopping attractiveness. Risk of promoting exaggerated standards.
		絢爛 (kenran) - Gorgeous, flamboyant

		Originally described colorful Buddhist tapestries. Indicates ornate, lavish beauty.
		Evokes blooming flowers and the height of their vibrant splendor.
		Has theatrical nuance, not a term applied discreetly. Implies striking sensuality.
		衝撃的な (shōgeki-tekina) - Startling, shocking

		Indicates powerful, intense impression on the senses.
		In terms of beauty, suggests visceral, breathtaking impact beyond expectations.
		Risk of promoting notion of beauty only in extremes rather than everyday appreciation.
		鮮やかな (azayaka na) - Vivid, brightly-colored

		Indicates colors and impressions that are fresh and vibrant.
		In terms of beauty, implies eye-catching radiance.
		Risk of over-emphasizing external qualities vs inner ones.
		素晴らしい (subarashii) - Wonderful, splendid

		General term of praise that can apply to inner or outer qualities.
		Indicates admiration without necessarily specifying beauty only.
		Arguably overused in modern discourse. Promotes positive perspective.
		ファンタスティックな (fantasutikku na) - Fantastic, wonderful

		Exaggerated term implying extraordinariness beyond reality.
		Evokes sense of the fantastical. Risk of overstatement.
		English loanword, not innate to Japanese but commonly borrowed.
		素敵な (suteki na) - Lovely, charming

		A sweet way to describe appealing attractiveness.
		Can apply to personality as well as looks. More about aura than features.
		Indicates admiration without obsessive connotations.
		魅力的な (miryoku-tekina) - Captivating, fascinating

		Focuses on compelling magnetism and charm.
		Can apply to inner qualities as much as outer.
		Allows appreciation without judgement or objectification.
		信じらrenaiほど (shinjirarenai hodo) - Unbelievably, incredibly

		Indicates exceeding the imagination. Risk of overstatement.
		In terms of beauty, implies inhuman levels beyond reality.
		美しすぎる (utsukushisugiru) - Exquisitely beautiful

		Indicates beauty exceeding normal bounds, beyond expectation.
		Risk of idealizing unrealistic standards, forgetting humanity.
		//
				上品な (jouhin-na) - refined, graceful
			上品な (jouhin-na) is an important concept in Japanese aesthetics conveying grace, refinement and tasteful simplicity. Some insights on this term:

			Roots in courtly culture - Originally described the elegant beauty of the imperial court. Linked to qualities like dignity, decorum.
			A sense of restraint - The appeal is in what is not shown overtly. Favors subtlety and suggests modesty.
			Valuing naturalness - Unaffected elegance is prized over ornamentation. Relates to wabi-sabi appreciation of rustic imperfections.
			Propriety and manners - 上品 relates to proper speech and conduct, composure and polish. Carries connotations of moral uprightness.
			Art and aesthetics - Used for artistic styles emphasizing purity, balance, understatement. Calligraphy, flower arrangement, tea ceremony all value 上品.
			Personal comportment - Someone described as 上品 has taste, dignity and admirable character. It suggests style without arrogance.
			Nuance - Imparts grace and calmness rather than just conventional prettiness. Has cultural resonance difficult to translate fully.
				//
						//
						自然な、控えめな美

						(shizen-na, hikaeme-na bi)

						This translates the phrase "natural, understated beauty" as:

						自然な (shizen-na) - natural

						控えめな (hikaeme-na) - understated, restrained, modest

						美 (bi) - beauty
						//
								絢爛 (kenran - exquisite, gorgeous)
					//
							美人 (bijin - beauty) is still used respectfully
					//
								{「超美人」,  "extremely beautiful", "gorgeous"}

								{「絢爛」, "exquisite", "gorgeous"}

								{「衝撃的な」, "startling", "shocking"}

								{「鮮やかな」, "striking", "vivid"}

								{「素晴らしい」, "fabulous", "wonderful"}

								{「ファンタスティックな」, "fantastic", "excellent"}

								{「素敵な」, "lovely", "charming"}

								{「魅力的な」, "delightful", "charming"}

								{「信じられないほど」, "incredible", "unbelievable"}

								{「美しすぎる」, "impossibly pretty", "exquisitely beautiful"}

	*/

}
