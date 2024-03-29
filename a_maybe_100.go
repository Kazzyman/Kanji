package main

var data_file100_maybe = []charSetStructKanji{
	{"一", "one", "first, single",
		"ony: イチ(ichi)", "kun: ひと(hito)",
		"一日(ichinichi; one day), 一人(hitori; one person), 一回(ikkai; one time)",
		"初めて(hajimete; the first time), 一番(ichiban; number one), 一つ(hitotsu; one thing)"},

	{"二", "two", "second, double",
		"ony: ニ(ni)", "kun: ふた(uta)",
		"二日(futsuka; two days), 二人(futari; two people), 二回(nikai; two times)",
		"二重(nijuu; double), 二等(nitou; second class), 二月(nigatsu; February)"},

	{"三", "three", "triple, third",
		"ony: サン(san)", "kun: み(mi)",
		"三日(mikka; three days), 三人(sannin; three people), 三回(sankai; three times)",
		"三角(sankaku; triangle), 三味線(shamisen; 3-stringed instrument), 三度(sandan; third degree)"},

	{"四", "four", "quadruple",
		"ony: シ(shi)", "kun: よ(yo)",
		"四日(yokka; four days), 四人(yonin; four people), 四回(shikai; four times)",
		"四角(shikaku; square), 四月(shigatsu; April), 四季(shiki; four seasons)"},

	{"五", "five", "quintuple",
		"ony: ゴ(go)", "kun: いつ(itsu)",
		"五日(itsuka; five days), 五人(gonin; five people), 五回(gokai; five times)",
		"五角(gokaku; pentagon), 五月(gogatsu; May), 五十(isso; fifty)"},

	{"六", "six", "sextuple",
		"ony: ロク(roku)", "kun: む(mu)",
		"六日(muika; six days), 六人(rokunin; six people), 六回(rokkai; six times)",
		"六角(rokkaku; hexagon), 六月(rokugatsu; June), 六十(rokujuu; sixty)"},

	{"七", "seven", "seventh",
		"ony: シチ(shichi)", "kun: なな(nana)",
		"七日(nanoka; seven days), 七人(shichinin; seven people), 七回(nanakai; seven times)",
		"七月(shichigatsu; July), 七十(nanajuu; seventy), 七色(nanairo; seven colors)"},

	{"八", "eight", "octuple",
		"ony: ハチ(hachi)", "kun: や(ya)",
		"八日(youka; eight days), 八人(hachinin; eight people), 八回(hakkai; eight times)",
		"八月(hachigatsu; August), 八十(hachijuu; eighty), 八百(happyaku; eight hundred)"},

	{"九", "nine", "ninth",
		"ony: キュウ(kyuu)", "kun: ここの(kokono)",
		"九日(kokonoka; nine days), 九人(kyuunin; nine people), 九回(kyuukai; nine times)",
		"九月(kugatsu; September), 九十(kyuujuu; ninety), 九州(kyuushuu; Kyushu)"},

	{"百", "hundred", "many, percent",
		"ony: ヒャク(hyaku)", "kun: -",
		"百日(hyakunichi; hundred days), 百人(hyakunin; hundred people), 百回(hyakkai; hundred times)",
		"百パーセント(hyaku paasento; one hundred percent), 二百(nihyaku; two hundred), 三百(sanbyaku; three hundred)"},

	{"万", "ten thousand", "myriad, all",
		"ony: マン(man)", "kun: -",
		"一万(ichiman; ten thousand), 二万(niman; twenty thousand), 三万(sanman; thirty thousand)",
		"十万(juuman; hundred thousand), 百万(hyakuman; million), 千万(senman; ten million)"},

	{"王", "king", "ruler, crown",
		"ony: オウ(ou)", "kun: -",
		"王様(ousama; king), 女王(joou; queen), 王国(oukoku; kingdom)",
		"王子(ouji; prince), 王冠(oukan; crown), 王道(oudou; kingship)"},

	{"玉", "jewel; ball", "precious, spherical",
		"ony: ギョク(gyoku)", "kun: たま(tama)",
		"宝玉(hougyoku; jewel), 玉石(gyokuseki; gem), 金玉(kin gyoku; gold nugget)",
		"玉子(tamago; egg), 野球玉(yakyuudama; baseball), 玉ねぎ(tamanegi; onion)"},

	{"花", "flower", "bloom, beauty",
		"ony: カ(ka)", "kun: はな(hana)",
		"花瓶(kabin; vase), 花束(hanataba; bouquet), 花見(hanami; flower viewing)",
		"桜の花(sakura no hana; cherry blossoms), 花火(hanabi; fireworks), 花屋(hanaya; florist)"},

	{"草", "grass", "plant, green",
		"ony: ソウ(sou)", "kun: くさ(kusa)",
		"草原(sougen; grassland), 野草(nogusa; wild grass), 芝草(shibafu; lawn grass)",
		"草花(kusabana; grass and flowers), 雑草(zassou; weed), 芝生(shibafu; lawn)"},

	{"木", "tree", "wood, forest",
		"ony: モク(moku)", "kun: き(ki)",
		"樹木(jumoku; tree), 木材(mokuzai; timber), 竹(take; bamboo)",
		"森林(shinrin; forest), 家具(kagu; furniture), 椅子(isu; chair)"},

	{"林", "grove", "many trees",
		"ony: リン(rin)", "kun: はやし(hayashi)",
		"森林(shinrin; forest), 松林(matsubayashi; pine grove), 竹林(chikurin; bamboo grove)",
		"山林(sanrin; mountain forest), 林間(hayashi ma; among trees), 人工林(jinkou rinka; artificial forest)"},

	{"山", "mountain", "nature, climbing",
		"ony: サン(san)", "kun: やま(yama)",
		"富士山(fujisan; Mt. Fuji), 山頂(sancho; mountain top), 火山(kazan; volcano)",
		"山岳(sangaku; alpinism), 山男(yamabito; mountain man), 山菜(sansai; mountain vegetables)"},

	{"川", "river", "water, stream",
		"ony: セン(sen)", "kun: かわ(kawa)",
		"河川(kasen; river), 川岸(kawagishi; riverside), 川底(kawazoko; riverbed)",
		"川魚(kawauo; river fish), 川崎(Kawasaki), 川沿い(kawayoi; along the river)"},

	{"海", "sea", "ocean, beach",
		"ony: カイ(kai)", "kun: うみ(umi)",
		"海岸(kaigan; seashore), 海水(kaisui; seawater), 海外(kaigai; overseas)",
		"海産物(kaisanbutsu; marine products), 海洋(kaiyou; ocean), 海辺(umibe; seaside)"},

	{"金", "gold", "money, metal",
		"ony: キン(kin)", "kun: かね(kane)",
		"金地金(kinjinkin; gold bullion), 金メダル(kinmedaru; gold medal), 金貨(kinga; gold coin)",
		"お金(okane; money), 金庫(kinko; safe), 銀行(ginkou; bank)"},

	{"銀", "silver", "metal, shine",
		"ony: ギン(gin)", "kun: しろがね(shirogane)",
		"銀地金(ginjinkin; silver bullion), 銀メダル(ginmedaru; silver medal), 銀貨(ginga; silver coin)",
		"銀行(ginkou; bank), 銀色(giniro; silvery color), 銀杏(ichou; ginkgo)"},

	{"星", "star", "night sky, shine",
		"ony: セイ(sei)", "kun: ほし(hoshi)",
		"明星(myousei; Venus), 流星(ryuusei; shooting star), 北極星(hokkyokusei; Polaris)",
		"夜空の星(yozora no hoshi; stars in the night sky), 星座(seiza; constellation), 星宿(seishuku; star cluster)"},

	{"天", "heaven; sky", "above, celestial",
		"ony: テン(ten)", "kun: あま(ama)",
		"天国(tengoku; heaven), 天気(tenki; weather), 天使(tenshi; angel)",
		"空(sora; sky), 天皇(tennou; emperor), 天井(tenjou; ceiling)"},

	{"雨", "rain", "wet, umbrella",
		"ony: ウ(u)", "kun: あめ(ame)",
		"豪雨(gouu; heavy rain), 雨季(uuki; rainy season), 台風の雨(taifuu no ame; typhoon rain)",
		"傘(kasa; umbrella), 雨具(amagu; rainwear), 雨天(uenten; rainy weather)"},

	{"雪", "snow", "winter, cold",
		"ony: セツ(setsu)", "kun: ゆき(yuki)",
		"大雪(oosetsu; heavy snow), 降雪(kousetsu; snowfall), 積雪(sekisetsu; snow cover)",
		"雪国(yukiguni; snowy country), 雪景色(yukigeshiki; snowscape), 雪遊び(yukiasobi; playing in snow)"},

	{"音", "sound", "noise, tone",
		"ony: オン(on)", "kun: ね(ne)",
		"音楽(ongaku; music), 音量(onryou; volume), 発音(hatsuon; pronunciation)",
		"耳鳴り(miminari; ringing in ears), 音程(nechou; musical pitch), 音色(neiro; tone color)"},

	{"学", "learning", "study, school",
		"ony: ガク(gaku)", "kun: まな(mana)",
		"学校(gakkou; school), 学生(gakusei; student), 学習(gakushuu; study)",
		"勉強(benkyou; study), 学者(gakusha; scholar), 学部(gakubu; faculty)"},

	{"力", "strength", "power, ability",
		"ony: リキ(riki)", "kun: ちから(chikara)",
		"力量(rikiryou; ability), 体力(tairyoku; physical strength), 投力(touriki; throwing power)",
		"勢力(seiryoku; influence), 能力(nouryoku; ability), 力不足(chikarafusoku; lack of strength)"},

	{"公", "public", "government, official",
		"ony: コウ(kou)", "kun: おおやけ(ooyake)",
		"公園(kouen; park), 公務員(koumuin; civil servant), 公用語(kouyougo; official language)",
		"公平(kouhei; fair), 公式(koushiki; official), 公安(kouan; public peace)"},

	{"司", "manage", "direct, rule",
		"ony: シ(shi)", "kun: つかさど(tsukasado)",
		"管理(kanri; management), 指揮(shiki; command), 司令(shirei; command)",
		"司会(shikai; master of ceremonies), 司法(shihou; justice), 司直(shinao; being on duty)"},

	{"夏", "summer", "hot, season",
		"ony: カ(ka)", "kun: なつ(natsu)",
		"夏休み(natsuyasumi; summer vacation), 夏場(natsuba; summer season), 夏物(natsumono; summer clothing)",
		"涼しい(suzushii; cool), 暑い(atsui; hot), 夏祭り(natsumatsuri; summer festival)"},

	{"秋", "autumn", "season, harvest",
		"ony: シュウ(shuu)", "kun: あき(aki)",
		"秋分(shubun; autumn equinox), 秋口(akiguchi; early autumn), 秋晴れ(akibare; clear autumn weather)",
		"紅葉(kouyou; autumn leaves), 秋祭り(akisai; autumn festival), 秋風(akikaze; autumn breeze)"},

	{"冬", "winter", "cold, season",
		"ony: トウ(tou)", "kun: ふゆ(fuyu)",
		"冬休み(fuyuyasumi; winter vacation), 冬将軍(fuyu shogun; cold winter weather), 冬支度(fuyu shitaku; winter preparations)",
		"寒い(samui; cold), 雪(yuki; snow), 冬景色(fuyu gesshiki; winter scenery)"},

	{"春", "spring", "season, new",
		"ony: シュン(shun)", "kun: はる(haru)",
		"春分(shunbun; spring equinox), 春先(harusaki; early spring), 春の訪れ(haru no otozure; arrival of spring)",
		"春風(harukaze; spring breeze), 春眠(harumin; spring slumber), 春祭り(harumatsuri; spring festival)"},

	{"国", "country", "nation, state",
		"ony: コク(koku)", "kun: くに(kuni)",
		"国民(kokumin; national), 他国(takoku; other country), 国土(kokudo; national land)",
		"外国(gaikoku; foreign country), 祖国(sokoku; motherland), 国家(kokka; nation)"},

	{"外", "outside", "exterior, foreign",
		"ony: ガイ(gai)", "kun: そと(soto)",
		"外国(gaikoku; foreign country), 外出(gaishutsu; going out), 外観(gaikan; exterior)",
		"屋外(okugai; outdoors), 海外(kaigai; overseas), 表(omote; surface)"},

	{"新", "new", "refresh, renew",
		"ony: シン(shin)", "kun: あたら(atara)",
		"新品(shinpin; new article), 新築(shinchiku; new construction), 新設(shinsetsu; new establishment)",
		"新人(shinjin; new employee), 更新(koushin; renewal), 新鮮(shinsen; freshness)"},

	{"古", "old", "former, antiquity",
		"ony: コ(ko)", "kun: ふる(furui)",
		"古典(koten; classic), 古代(kodai; ancient times), 古文(kobun; classical literature)",
		"古い(furui; old), 古ぼける(fubokeru; grow old), 古着(furugi; old clothes)"},

	{"少", "few", "little, scarce",
		"ony: ショウ(shou)", "kun: すく(suku)",
		"少数(shousuu; few), 少量(shouryou; small quantity), 少し(sukoshi; little)",
		"年少(nenshou; young age), 少年(shounen; boy), 少女(shoujo; girl)"},

	{"多", "many", "much, numerous",
		"ony: タ(ta)", "kun: おお(oo)",
		"多数(tasuu; many), 多岐(taki; manifold), 多量(tairyo; large quantity)",
		"豊富(houfu; abundance), 沢山(takusan; a lot), たくさん(takusan; a lot)"},

	{"上", "above", "upper, superior",
		"ony: ジョウ(jou)", "kun: うえ(ue)",
		"上体(joutai; upper part of body), 上品(jouhin; high quality), 上手(ujou; skill)",
		"上がる(agaru; rise), 上る(noboru; ascend), 上野(Ueno)"},

	{"下", "below", "lower, descend",
		"ony: カ(ka)", "kun: した(shita)",
		"下半身(kahanmi; lower half of body), 下品(gehin; low quality), 下手(heta; unskillful)",
		"下がる(sagaru; drop), 下る(kudaru; descend), 下町(shitamachi; downtown)"},

	{"左", "left", "direction",
		"ony: サ(sa)", "kun: ひだり(hidari)",
		"左手(hidarite; left hand), 左翼(sayoku; left wing), 左折(sasetsu; turn left)",
		"左側(hidarigawa; left side), 坐る(suwaru; sit)"},

	{"右", "right", "direction",
		"ony: ウ(u)", "kun: みぎ(migi)",
		"右手(migite; right hand), 右翼(uyoku; right wing), 右折(usettsu; turn right)",
		"右側(migigawa; right side), 坐る(suwaru; sit)"},

	{"南", "south", "direction",
		"ony: ナン(nan)", "kun: みなみ(minami)",
		"南極(nankyoku; South Pole), 南西(nansei; southwest), 南東(touhoku; southeast)",
		"南口(minamiguchi; south exit), 南国(nangoku; southern country), 北南(hokunan; north-south)"},

	{"西", "west", "direction",
		"ony: セイ; サイ(sei; sai)", "kun: にし(nishi)",
		"西日(nishibi; setting sun), 西洋(seiyo; the West), 西南(seinan; southwest)",
		"西口(nishiguchi; west exit)"},

	{"東", "east", "direction",
		"ony: トウ(tou)", "kun: ひがし(higashi)",
		"東日(toujitsu; rising sun), 東洋(touyou; the East), 東北(touhoku; northeast)",
		"東口(higashiguchi; east exit)"},

	{"語", "language", "speech, word",
		"ony: ゴ(go)", "kun: かた(kata)",
		"言語(gengo; language), 外国語(gaikokugo; foreign language), 方言(hougen; dialect)",
		"言葉(kotoba; word), 丁寧語(teineigo; polite language), 漢語(kango; Chinese-derived words)"},

	{"主", "main", "primary, lord",
		"ony: シュ(shu)", "kun: あるじ(arugi); ぬし(nushi)",
		"主人(shujin; master), 主要(shuyo; main), 主役(shuyaku; leading part)",
		"主(nushi; owner), 主(aruji; lord), 主婦(shufu; housewife)"},

	{"体", "body", "substance, reality",
		"ony: タイ(tai)", "kun: からだ(karada)",
		"身体(shintai; body), 人体(jintai; human body), 容体(youtai; condition of body)",
		"体重(taijuu; body weight), 実体(jittai; reality), 実体験(jittiken; real experience)"},

	{"色", "color", "hue, tone",
		"ony: ショク(shoku)", "kun: いろ(iro)",
		"原色(genshoku; primary color), 色彩(shikisai; coloration), 色調(shikichou; tone of color)",
		"赤色(sekishoku; red color), 色合い(iroai; tint), 肌色(hadairo; skin color)"},

	{"楽", "comfort", "ease, enjoy",
		"ony: ラク(raku); ガク(gaku)", "kun: たの(tanoshii)",
		"安楽(anraku; comfort), 楽園(rakuen; paradise), 楽屋(gakuya; dressing room)",
		"楽しい(tanoshii; enjoyable), 楽勝(rakushou; easy win), 手軽(tegaru; easy)"},

	{"運", "fortune", "transport",
		"ony: ウン(un)", "kun: はこ(hako)",
		"運命(unmei; fate), 運転(unten; driving), 運搬(unpan; transport)",
		"運がいい(un ga ii; be lucky), 運ぶ(hakobu; transport), 運動(undou; exercise)"},

	{"空", "sky", "emptiness, vacuum",
		"ony: クウ(kuu); ソラ(sora)", "kun: あく(aku); あき(aki); から(kara)",
		"空港(kuukou; airport), 宇宙空間(uchuu kukan; outer space), 空腹(kuufuku; hunger)",
		"空っぽ(karappo; empty), 空く(aku; become empty), 空(sora; sky)"},

	{"院", "institution", "temple, hospital",
		"ony: イン(in)", "kun: -",
		"大学院(daigakuin; graduate school), 研究所(kenkyuujo; research institute), 病院(byouin; hospital)",
		"寺院(jiin; temple), 政府機関(seifu kikan; government institution), 裁判所(saibansho; courthouse)"},

	{"市", "city", "town, market",
		"ony: シ(shi)", "kun: いち(ichi)",
		"都市(toshi; city), 市役所(shiyakusho; city hall), 市場(ichiba; marketplace)",
		"町(machi; town), 開市(kaishi; opening of market), 閉市(heishi; closing of market)"},

	{"区", "ward", "district, zone",
		"ony: ク(ku)", "kun: -",
		"区役所(kuyakusho; ward office), 住宅地区(jutakuchiiki; residential district), 校区(kouku; school district)",
		"区切り(kugiri; division), 区分(kubun; classification), 特別区(tokubetsuku; special ward)"},

	{"政", "government", "politics, policy",
		"ony: セイ(sei)", "kun: まつ(matsu)",
		"政府(seifu; government), 政策(seisaku; policy), 政治(seiji; politics)",
		"行政(gyousei; administration), 治政(chisei; governing), 政党(seitou; political party)"},

	{"正", "correct", "justice, 10th month",
		"ony: セイ(sei)", "kun: ただ(tada); まさ(masa)",
		"正解(seikai; right answer), 正確(seikaku; accurate), 正月(shougatsu; New Year's)",
		"正しい(tadashii; correct), 正直(shoujiki; honest), 正午(shougo; noon)"},

	{"立", "stand", "establish, consist",
		"ony: リツ(ritsu); リュウ(ryuu)", "kun: た(ta)",
		"立候補(rikouho; to stand for election), 立証(risshou; proof), 独立(dokuritsu; independence)",
		"立つ(tatsu; stand up), 建てる(tateru; build up), 存在(sonzai; existence)"},

	{"方", "direction", "person, way",
		"ony: ホウ(hou)", "kun: かた(kata)",
		"方法(houhou; method), 四方(shihou; four directions), 方向(houkou; direction)",
		"一方(ippo; one side), やり方(yarikata; way of doing), 人(hito; person)"},

	{"社", "company", "society, shrine",
		"ony: シャ(sha)", "kun: やしろ(yashiro)",
		"会社(kaisha; company), 出版社(shuppan sha; publisher), 神社(jinja; shrine)",
		"公社(kousha; public corporation), 社員(shain; employee), 社交(shakou; social intercourse)"},

	{"真", "true", "reality, perfect",
		"ony: シン(shin)", "kun: ま(ma)",
		"真実(shinjitsu; truth), 真珠(shinju; pearl), 真空(shinkuu; vacuum)",
		"真面目(majime; serious), 真っ直ぐ(massugu; straight), 真(makoto; truth)"},

	{"切", "cut", "slit, sever",
		"ony: セツ(setsu)", "kun: き(ki)",
		"切手(kitte; postage stamp), 切符(kippu; ticket), 切断(setsudan; disruption)",
		"切る(kiru; cut), 切開(kikai; cut open), 切り取る(kiri toru; cut off)"},

	{"然", "natural", "so, unconsciously",
		"ony: ゼン(zen); ジ(ji)", "kun: しか(shika)",
		"当然(touzen; natural), 必然(hitsuzen; inevitable), 無意識(muishiki; unconsciously)",
		"然るべき(arubeki; proper), 然し(shikashi; however), 然も (shikamo; nevertheless)"},

	{"引", "pull", "subtract, refer",
		"ony: イン(in)", "kun: ひ(hi)",
		"引っ張る(hipparu; pull), 引く(hiku; draw back), 引退(intai; retirement)",
		"引用(inyou; citation), 引き出す(hikidasu; deduce), 引き分け(hikiwake; tie game)"},

	{"止", "stop", "halt, resign",
		"ony: シ(shi)", "kun: と(to)",
		"停止(teishi; stop), 中止(chuushi; suspension), 禁止(kinshi; prohibition)",
		"止まる(tomaru; stop), 止める(yameru; stop), 辞める(yameru; quit)"},

	{"見", "see", "outlook, view",
		"ony: ケン(ken)", "kun: み(mi)",
		"見解(kenkai; view), 山頂(sanchou; mountain top), 考え方(kangaekata; way of thinking)",
		"見る(miru; see), 見える(mieru; be visible), 一見(ikken; glance)"},

	{"思", "think", "thought",
		"ony: シ(shi)", "kun: おも(omo)",
		"思想(shisou; thought), 思考(shikou; thinking), 思い出(omoida; memory)",
		"思う(omou; think), 思い込む(omoidkomu; be convinced), 思い切って(omoikkiri; boldly)"},

	{"以", "by means of", "because, compared to",
		"ony: イ(i)", "kun: もっ(motte)",
		"以下(ika; below), 以内(inai; within), 以外(igai; except)",
		"以上(ijou; more than), 以来(irai; since), 以って(motte; by means of)"},

	{"取", "take", "fetch, choose",
		"ony: シュ(shu)", "kun: と(to)",
		"取る(toru; take), 受取(uketoru; receive), 取り出す(toridasu; take out)",
		"取引(torihiki; transaction), 取材(shuzai; coverage), 携帯(keitai; portable)"},

	{"親", "parent", "intimate, close",
		"ony: シン(shin)", "kun: おや(oya); したし(shitashi)",
		"親子(oyako; parent and child), 親友(shin'yu; close friend), 親しい(shitashii; intimate)",
		"養父(youfu; foster father), 養母(youbo; foster mother), 養親(youshin; foster parent)"},

	{"自", "oneself", "automatic, spontaneous",
		"ony: ジ(ji); シ(shi)", "kun: みずか(mizuka)",
		"自分(jibun; oneself), 自動(jidou; automatic), 自然(shizen; nature)",
		"自宅(jutaku; one's home), 自発(jihatsu; spontaneous), 自力(jirki; own power)"},

	{"界", "world", "limit, boundary",
		"ony: カイ(kai); ケ(ke)", "kun: -",
		"世界(sekai; world), 学界(gakkai; academic circle), 業界(gyoukai; industry)",
		"境界(kyoukai; border), 分界(bunkai; division), 定義(teigi; definition)"},

	{"令", "order", "rule, decree",
		"ony: リョウ(ryou); レイ(rei)", "kun: みことのり(mikotonori)",
		"命令(meirei; order), 規則(kisoku; rule), 法令(houritsu; law)",
		"令状(reijou; warrant), 命令(meirei; command), 勅令(chokurei; imperial edict)"},

	{"変", "change", "strange, unusual",
		"ony: ヘン(hen)", "kun: か(ka)",
		"変化(henka; change), 変更(henkou; modification), 変異(hen'i; variation)",
		"変わる(kawaru; change), 変な(hen na; strange), 異変(ibiten; unusual event)"},

	{"完", "perfect", "completion, end",
		"ony: カン(kan)", "kun: -",
		"完璧(kanpeki; flawless), 完成(kansei; completion), 完全 (kanzen; perfect)",
		"完了(kanryou; completion), 完遂(kannsui; accomplishment), 完結(kanketsu; conclusion)"},

	{"数", "number", "figure, several",
		"ony: スウ(suu); シュ(shu); ス(su)", "kun: かず(kazu)",
		"数字(suuji; number), 数学(suugaku; mathematics), 数少ない(sukunai; few)",
		"勘定(kanjou; bill), 多数(tasuu; majority), 幾つ(ikutsu; how many)"},

	{"合", "join", "fit, suit",
		"ony: ゴウ(gou); カッ(katsu); ガッ(gatu)", "kun: あ(a); あう(au)",
		"合同(godou; joint), 合成(gousei; synthesis), 結合(ketsugou; combination)",
		"合う(au; fit), 合意(goui; agreement), 手合わせ(teawase; applause)"},

	{"動", "move", "motion, activity",
		"ony: ドウ(dou)", "kun: うご(ugo)",
		"移動(idou; movement), 動作(dousa; movement), 反動(handou; reaction)",
		"動く(ugoku; move), 動機(douki; motive), 感動(kandou; being moved)"},

	{"度", "degree", "time, occasion",
		"ony: ド(do); ト(to)", "kun: たび(tabi)",
		"温度(ondo; temperature), 三度(sand; three times), 十度(juutodo; ten degrees)",
		"再度(saido; again), 度々(tabitabi; often), 幾度(ikutabi; how many times)"},

	{"刀", "sword", "knife, blade",
		"ony: トウ(tou)", "kun: かたな(katana)",
		"刀剣(touken; sword), 小刀(shoutou; knife), 短刀(tantou; short sword)",
		"刃(ha; blade), 刀身(toumi; sword blade), 刀工(toukou; swordsmith)"},

	{"友", "friend", "companion",
		"ony: ユウ(yuu)", "kun: とも(tomo)",
		"友人(yuujin; friend), 友好(yuukou; friendship), 親友(shin'yuu; close friend)",
		"仲間(nakama; companion), 取引先(torihikisaki; business connection), 共同(kyoudou; joint)"},

	{"品", "goods", "dignity, quality",
		"ony: ヒン(hin); ホン(hon)", "kun: しな(shina)",
		"商品(shouhin; goods), 品物(shinamono; article), 食品(shokuhin; food products)",
		"品格(hinkaku; dignity), 品質(hinshitsu; quality), 工芸品(kougeihin; craft goods)"},

	{"直", "straight", "direct",
		"ony: チョク(choku); ジキ(jiki)", "kun: なお(nao); ただ(tada)",
		"直線(chokusen; straight line), 正直(shoujiki; honest), 直後(chokugo; immediately after)",
		"直ぐ(sugu; immediately), 直る(naoru; be corrected), まっすぐ(massugu; straight)"},

	{"買", "buy", "purchase",
		"ony: バイ(bai)", "kun: か(ka)",
		"購入(kounyuu; purchase), 買い物(kaimono; shopping), 買取(kaitori; buying)",
		"買う(kau; buy), 売買(baibai; buying and selling), 代金(daikin; payment)"},

	{"売", "sell", "selling",
		"ony: バイ(bai)", "kun: う(u)",
		"販売(hanbai; sale), 売店(uriten; shop), 売り切れ(uri kire; sold out)",
		"売る(uru; sell), 売買(baibai; buying and selling), 売上(uriage; sales)"},

	{"返", "return", "revert, repay",
		"ony: ヘン(hen)", "kun: かえ(kae)",
		"返事(henji; reply), 帰り道(kaerimichi; return route), 返品(henpin; returned goods)",
		"返す(kaesu; return), 返る(kaeru; go back), 振り返る(furikaeru; look back)"},

	{"始", "begin", "start, opening",
		"ony: シ(shi)", "kun: はじ(haji)",
		"開始(kaishi; start), 始発(shuppatsu; first train), 始業(shigyou; start of work)",
		"始める(hajimeru; begin), 始まる(hajimaru; begin), 初日(shonichi; first day)"},

	{"信", "faith", "belief, letter",
		"ony: シン(shin)", "kun: まこと(makoto)",
		"信頼(shinrai; trust), 信仰(shinkou; religious belief), 信号(shingou; traffic signal)",
		"信じる(shinjiru; believe in), 誠実(seijitsu; sincere), 宛名(ate name; addressee)"},

	{"君", "you", "ruler",
		"ony: クン(kun)", "kun: きみ(kimi)",
		"君主(kunshu; monarch), 君臨(kunrin; reign), 君民(kunmin; the people)",
		"お前(omae; you), 貴方(anata; you), あなた(anata; you)"},

	{"治", "govern", "cure, remedy",
		"ony: チ(chi); ジ(ji)", "kun: おさ(osa)",
		"政治(seiji; politics), 治療(chiryou; medical treatment), 整備 (seibi; preparation)",
		"治す(osameru; put in order), 治る(naoru; get well), 治安(chian; public peace)"},

	{"具", "tool", "equipment, ingredient",
		"ony: グ(gu)", "kun: そな(sona)",
		"用具(yougu; tool), 調理器具(chouri kigu; cooking utensil), 具材(guzai; ingredient)",
		"装備(soubai; equipment), 装置(souchi; apparatus), 道具(dougu; tool)"},

	{"界", "world", "limit, boundary",
		"ony: カイ(kai); ケ(ke)", "kun: -",
		"世界(sekai; world), 学界(gakkai; academic circle), 業界(gyoukai; industry)",
		"境界(kyoukai; border), 分界(bunkai; division), 定義(teigi; definition)"},

	{"党", "party", "faction, clique",
		"ony: トウ(tou)", "kun: -",
		"政党(seitou; political party), 与党(yotou; ruling party), 野党(yatou; opposition party)",
		"党首(toushu; party leader), 党員(touin; party member), 民主党(minshutou; Democratic Party)"},

	{"集", "gather", "collection, crowd",
		"ony: シュウ(shuu)", "kun: あつ(atsu)",
		"集合(shugou; assembly), 収集(shushuu; collection), 人集り(hitoatsumari; crowd of people)",
		"集める(atsumeru; gather), 集中(shuchuu; concentration), 図書館(toshokan; library)"},

	{"派", "faction", "dispatch, school",
		"ony: ハ(ha); パイ(hai)", "kun: -",
		"政治派閥(seiji habatsu; political faction), 派遣(haken; dispatch), 芸術派(geijutsu ha; art school)",
		"派生(hassei; derivation), 支部(shibu; branch), 分派(bunpa; splinter group)"},

	{"区", "ward", "district, zone",
		"ony: ク(ku)", "kun: -",
		"区役所(kuyakusho; ward office), 住宅地区(jutakuchiiki; residential district), 校区(kouku; school district)",
		"区切り(kugiri; division), 区分(kubun; classification), 特別区(tokubetsuku; special ward)"},

	{"住", "dwell", "reside, live",
		"ony: ジュウ(jyu)", "kun: す(su)",
		"住居(jukyo; dwelling), 居住区(kyojyuku; residential area), 住所(juusho; address)",
		"住む(sumu; live), 居住(kyojyu; residence), 寄宿舎(kishusha; dormitory)"},

	{"字", "character", "letter, symbol",
		"ony: ジ(ji)", "kun: あざ(aza)",
		"文字(moji; letter), 字母(ji; alphabet), 字画(jikaku; stroke of a character)",
		"単字(tanji; single character), 語字(goji; word character), 熟字訓(jukujikun; Japanese reading of Chinese characters)"},

	{"帽", "hat", "cap",
		"ony: ボウ(bou)", "kun: かぶ(kabu)",
		"帽子(boushi; hat), ハット(hatto; hat), 野球帽(yakyuubou; baseball cap)",
		"被る(kiru; put on head), かぶる(kaburu; wear on head), 冠(kanmuri; crown)"},

	{"紙", "paper", "sheet",
		"ony: シ(shi)", "kun: かみ(kami)",
		"新聞紙(shinbunshi; newspaper), 和紙(washi; Japanese paper), 包装紙(housoushi; wrapping paper)",
		"紙片(kamihada; scrap of paper), 紙くず(kamikuzu; waste paper), 文書(bunsho; document)"},

	{"知", "know", "wisdom, intellect",
		"ony: チ(chi)", "kun: し(shi)",
		"知識(chishiki; knowledge), 知能(chinou; intellect), 知恵(chie; wisdom)",
		"知る(shiru; know), 認知(ninchi; cognition), 情報(jouhou; information)"},

	{"兄", "elder brother", "senior",
		"ony: ケイ(kei); キョウ(kyou)", "kun: あに(ani)",
		"兄弟(kyoudai; brothers), 兄貴(aniki; elder brother), 義兄(gikei; brother-in-law)",
		"お兄さん(oniisan; elder brother), 上の兄弟(ue no kyoudai; older sibling), 后(ane; elder sister)"},

	{"弟", "younger brother", "junior",
		"ony: テイ(tei); ダイ(dai); デ(de)", "kun: おとうと(otouto)",
		"弟子(deshi; pupil), 小弟(shoutei; younger brother), 妹(imouto; younger sister)",
		"お弟子(otouto; apprentice), 下の弟(shita no otouto; younger brother)"},

	{"妹", "younger sister", "female junior",
		"ony: マイ(mai)", "kun: いもうと(imouto)",
		"姉妹(shimai; sisters), 妹思い(imouto omoi; sister complex), 義妹(gimai; sister-in-law)",
		"お妹さん(imoutosan; younger sister), 妹分(imoutobun; sisterly figure), 妹(ane; elder sister)"},

	{"姉", "elder sister", "female senior",
		"ony: シ(shi)", "kun: あね(ane)",
		"姉妹(shimai; sisters), 長姉(choushi; eldest sister), 義姉(giane; sister-in-law)",
		"お姉さん(neesan; elder sister), 姉上(anesue; elder sister), 姉ちゃん(neechan; big sis)"},

	{"父", "father", "parent",
		"ony: フ(fu); ブ(bu)", "kun: ちち(chichi)",
		"父親(fubo; father), 父上(chichiue; father), 義父(gifu; foster father)",
		"お父さん(otosan; father), 父ちゃん(tochan; dad), 爸(papa; dad)"},

	{"母", "mother", "parent",
		"ony: ボ(bo)", "kun: はは(haha)",
		"母親(hahaoya; mother), 母上(hahaue; mother), 義母(gibou; mother-in-law)",
		"お母さん(okaasan; mother), 母ちゃん(kachan; mom), ママ(mama; mom)"},

	{"伯", "uncle", "senior",
		"ony: ハク(haku)", "kun: -",
		"叔父(ojiki; uncle), 伯父(ojiki; uncle), 大伯父(dai ojiki; granduncle)",
		"おじ(oji; uncle), 伯母(obasan; aunt), 伯父さん(oji-san; uncle)"},

	{"叔", "aunt", "junior",
		"ony: シュク(shuku)", "kun: -",
		"叔母(oba; aunt), 小叔母(obasan; aunt), 大叔母(dai obasan; grandaunt)",
		"おば(oba; aunt), 叔父(ojiki; uncle), 叔母さん(oba-san; aunt)"},

	{"息", "breath", "rest, break",
		"ony: ソク(soku)", "kun: いき(iki)",
		"呼吸(kokyuu; breathing), 一息(hitoiki; breath), 休息(kyuuikei; rest)",
		"息子(musuko; son), 息災(sokusai; taking a rest), 休む(yasumu; take a break)"},

	{"方", "person", "direction",
		"ony: ホウ(hou)", "kun: かた(kata)",
		"一方(ippo; one side), やり方(yarikata; way of doing), 人(hito; person)",
		"方法(houhou; method), 四方(shihou; four directions), 方向(houkou; direction)"},

	{"指", "finger", "point",
		"ony: シ(shi)", "kun: ゆび(yubi)",
		"人差指(hitosashiyubi; index finger), 指先(yubisaki; fingertip), 指図(shizu; directions)",
		"指す(sashu; point at), 指差す(yubi sasu; point at), 指導(shidou; guidance)"},

	{"肉", "meat", "flesh",
		"ony: ニク(niku)", "kun: -",
		"肉体(nikutai; flesh), 牛肉(gyuniku; beef), 鶏肉(toriniku; chicken)",
		"筋肉(kinniku; muscle), 身体(karada; body), 体(tei; body)"},

	{"学", "learning", "study",
		"ony: ガク(gaku)", "kun: まな(mana)",
		"学校(gakkou; school), 学生(gakusei; student), 学習(gakushuu; study)",
		"勉強(benkyou; study), 学者(gakusha; scholar), 学部(gakubu; faculty)"},

	{"文", "writing", "sentence",
		"ony: ブン(bun); モン(mon)", "kun: ふみ(fumi)",
		"文書(bunsho; document), 文化(bunka; culture), 文学(bungaku; literature)",
		"文章(bunshou; sentence), 文字(moji; character), 文房具(bumbutsu; stationery)"},

	{"方", "direction", "person, way",
		"ony: ホウ(hou)", "kun: かた(kata)",
		"一方(ippo; one side), やり方(yarikata; way of doing), 人(hito; person)",
		"方法(houhou; method), 四方(shihou; four directions), 方向(houkou; direction)"},

	{"直", "straight", "direct",
		"ony: チョク(choku); ジキ(jiki)", "kun: なお(nao); ただ(tada)",
		"直線(chokusen; straight line), 正直(shoujiki; honest), 直後(chokugo; immediately after)",
		"直ぐ(sugu; immediately), 直る(naoru; be corrected), まっすぐ(massugu; straight)"},

	{"区", "ward", "district",
		"ony: ク(ku)", "kun: -",
		"区役所(kuyakusho; ward office), 住宅地区(jutakuchiiki; residential district), 校区(kouku; school district)",
		"区切り(kugiri; division), 区分(kubun; classification), 特別区(tokubetsuku; special ward)"},

	{"止", "stop", "halt",
		"ony: シ(shi)", "kun: と(to)",
		"停止(teishi; stop), 中止(chuushi; suspension), 禁止(kinshi; prohibition)",
		"止まる(tomaru; stop), 止める(yameru; stop), 辞める(yameru; quit)"},

	{"見", "see", "view",
		"ony: ケン(ken)", "kun: み(mi)",
		"見解(kenkai; view), 山頂(sanchou; mountain top), 考え方(kangaekata; way of thinking)",
		"見る(miru; see), 見える(mieru; be visible), 一見(ikken; glance)"},

	{"思", "think", "thought",
		"ony: シ(shi)", "kun: おも(omo)",
		"思想(shisou; thought), 思考(shikou; thinking), 思い出(omoida; memory)",
		"思う(omou; think), 思い込む(omoidkomu; be convinced), 思い切って(omoikkiri; boldly)"},

	{"以", "by means of", "because",
		"ony: イ(i)", "kun: もっ(motte)",
		"以下(ika; below), 以内(inai; within), 以外(igai; except)",
		"以上(ijou; more than), 以来(irai; since), 以って(motte; by means of)"},

	{"取", "take", "choose",
		"ony: シュ(shu)", "kun: と(to)",
		"取る(toru; take), 受取(uketoru; receive), 取り出す(toridasu; take out)",
		"取引(torihiki; transaction), 取材(shuzai; coverage), 携帯(keitai; portable)"},

	{"引", "pull", "subtract",
		"ony: イン(in)", "kun: ひ(hi)",
		"引っ張る(hipparu; pull), 引く(hiku; draw back), 引退(intai; retirement)",
		"引用(inyou; citation), 引き出す(hikidasu; deduce), 引き分け(hikiwake; tie game)"},

	{"変", "change", "strange",
		"ony: ヘン(hen)", "kun: か(ka)",
		"変化(henka; change), 変更(henkou; modification), 変異(hen'i; variation)",
		"変わる(kawaru; change), 変な(hen na; strange), 異変(ibiten; unusual event)"},

	{"完", "perfect", "completion",
		"ony: カン(kan)", "kun: -",
		"完璧(kanpeki; flawless), 完成(kansei; completion), 完全 (kanzen; perfect)",
		"完了(kanryou; completion), 完遂(kannsui; accomplishment), 完結(kanketsu; conclusion)"},

	{"数", "number", "figure",
		"ony: スウ(suu); シュ(shu); ス(su)", "kun: かず(kazu)",
		"数字(suuji; number), 数学(suugaku; mathematics), 数少ない(sukunai; few)",
		"勘定(kanjou; bill), 多数(tasuu; majority), 幾つ(ikutsu; how many)"},

	{"合", "join", "fit",
		"ony: ゴウ(gou); カッ(katsu); ガッ(gatu)", "kun: あ(a); あう(au)",
		"合同(godou; joint), 合成(gousei; synthesis), 結合(ketsugou; combination)",
		"合う(au; fit), 合意(goui; agreement), 手合わせ(teawase; applause)"},

	{"動", "move", "motion",
		"ony: ドウ(dou)", "kun: うご(ugo)",
		"移動(idou; movement), 動作(dousa; movement), 反動(handou; reaction)",
		"動く(ugoku; move), 動機(douki; motive), 感動(kandou; being moved)"},

	{"度", "degree", "time",
		"ony: ド(do); ト(to)", "kun: たび(tabi)",
		"温度(ondo; temperature), 三度(sand; three times), 十度(juutodo; ten degrees)",
		"再度(saido; again), 度々(tabitabi; often), 幾度(ikutabi; how many times)"},

	{"刀", "sword", "knife",
		"ony: トウ(tou)", "kun: かたな(katana)",
		"刀剣(touken; sword), 小刀(shoutou; knife), 短刀(tantou; short sword)",
		"刃(ha; blade), 刀身(toumi; sword blade), 刀工(toukou; swordsmith)"},

	{"友", "friend", "companion",
		"ony: ユウ(yuu)", "kun: とも(tomo)",
		"友人(yuujin; friend), 友好(yuukou; friendship), 親友(shin'yuu; close friend)",
		"仲間(nakama; companion), 取引先(torihikisaki; business connection), 共同(kyoudou; joint)"},

	{"品", "goods", "quality",
		"ony: ヒン(hin); ホン(hon)", "kun: しな(shina)",
		"商品(shouhin; goods), 品物(shinamono; article), 食品(shokuhin; food products)",
		"品格(hinkaku; dignity), 品質(hinshitsu; quality), 工芸品(kougeihin; craft goods)"},

	{"返", "return", "repay",
		"ony: ヘン(hen)", "kun: かえ(kae)",
		"返事(henji; reply), 帰り道(kaerimichi; return route), 返品(henpin; returned goods)",
		"返す(kaesu; return), 返る(kaeru; go back), 振り返る(furikaeru; look back)"},

	{"始", "begin", "start",
		"ony: シ(shi)", "kun: はじ(haji)",
		"開始(kaishi; start), 始発(shuppatsu; first train), 始業(shigyou; start of work)",
		"始める(hajimeru; begin), 始まる(hajimaru; begin), 初日(shonichi; first day)"},

	{"信", "faith", "belief",
		"ony: シン(shin)", "kun: まこと(makoto)",
		"信頼(shinrai; trust), 信仰(shinkou; religious belief), 信号(shingou; traffic signal)",
		"信じる(shinjiru; believe in), 誠実(seijitsu; sincere), 宛名(ate name; addressee)"},

	{"君", "you", "ruler",
		"ony: クン(kun)", "kun: きみ(kimi)",
		"君主(kunshu; monarch), 君臨(kunrin; reign), 君民(kunmin; the people)",
		"お前(omae; you), 貴方(anata; you), あなた(anata; you)"},

	{"治", "govern", "cure",
		"ony: チ(chi); ジ(ji)", "kun: おさ(osa)",
		"政治(seiji; politics), 治療(chiryou; medical treatment), 整備 (seibi; preparation)",
		"治す(osameru; put in order), 治る(naoru; get well), 治安(chian; public peace)"},

	{"具", "tool", "equipment",
		"ony: グ(gu)", "kun: そな(sona)",
		"用具(yougu; tool), 調理器具(chouri kigu; cooking utensil), 具材(guzai; ingredient)",
		"装備(soubai; equipment), 装置(souchi; apparatus), 道具(dougu; tool)"},

	{"党", "party", "faction",
		"ony: トウ(tou)", "kun: -",
		"政党(seitou; political party), 与党(yotou; ruling party), 野党(yatou; opposition party)",
		"党首(toushu; party leader), 党員(touin; party member), 民主党(minshutou; Democratic Party)"},

	{"集", "gather", "collection",
		"ony: シュウ(shuu)", "kun: あつ(atsu)",
		"集合(shugou; assembly), 収集(shushuu; collection), 人集り(hitoatsumari; crowd of people)",
		"集める(atsumeru; gather), 集中(shuchuu; concentration), 図書館(toshokan; library)"},

	{"派", "faction", "dispatch",
		"ony: ハ(ha); パイ(hai)", "kun: -",
		"政治派閥(seiji habatsu; political faction), 派遣(haken; dispatch), 芸術派(geijutsu ha; art school)",
		"派生(hassei; derivation), 支部(shibu; branch), 分派(bunpa; splinter group)"},

	{"将", "general", "leader",
		"ony: ショウ(shou)", "kun: まさ(masa)",
		"将軍(shougun; general), 将来(shourai; future), 棋士(kishi; go/shogi player)",
		"指揮官(shikikan; commander), 社長(shachou; company president), 将(masu; piece in shogi)"},

	{"消", "erase", "use up",
		"ony: ショウ(shou)", "kun: け(ke); き(ki)",
		"消去(shoukyo; erasing), 消費(shouhi; consumption), 消化(shouka; digestion)",
		"消す(kesu; erase), 消える(kieru; disappear), 消耗(shoumou; wear and tear)"},

	{"写", "copy", "photograph",
		"ony: シャ(sha)", "kun: うつ(utsu)",
		"写真(shashin; photo), 複写(fukusha; duplication), 映写(eisha; projection)",
		"写す(utsusu; copy), 写る(utsuru; be photographed), 模写(mo-sha; imitation)"},

	{"作", "make", "prepare",
		"ony: サ(sa); ソウ(sou)", "kun: つく(tsuku)",
		"作成(sakusei; production), 著作(chosaku; literary work), 作者(sakusha; author)",
		"作る(tsukuru; make), 仕事(shigoto; work), 製作(seisaku; production)"},

	{"切", "cut", "sever",
		"ony: セツ(setsu)", "kun: き(ki)",
		"切手(kitte; postage stamp), 切符(kippu; ticket), 切断(setsudan; disruption)",
		"切る(kiru; cut), 切開(kikai; cut open), 切り取る(kiri toru; cut off)"},

	{"然", "natural", "unconscious",
		"ony: ゼン(zen); ジ(ji)", "kun: しか(shika)",
		"当然(touzen; natural), 必然(hitsuzen; inevitable), 無意識(muishiki; unconsciously)",
		"然るべき(arubeki; proper), 然し(shikashi; however), 然も (shikamo; nevertheless)"},

	{"引", "pull", "subtract",
		"ony: イン(in)", "kun: ひ(hi)",
		"引っ張る(hipparu; pull), 引く(hiku; draw back), 引退(intai; retirement)",
		"引用(inyou; citation), 引き出す(hikidasu; deduce), 引き分け(hikiwake; tie game)"},

	{"止", "stop", "halt",
		"ony: シ(shi)", "kun: と(to)",
		"停止(teishi; stop), 中止(chuushi; suspension), 禁止(kinshi; prohibition)",
		"止まる(tomaru; stop), 止める(yameru; stop), 辞める(yameru; quit)"},

	{"語", "language", "speech",
		"ony: ゴ(go)", "kun: かた(kata)",
		"言語(gengo; language), 外国語(gaikokugo; foreign language), 方言(hougen; dialect)",
		"言葉(kotoba; word), 丁寧語(teineigo; polite language), 漢語(kango; Chinese-derived words)"},

	{"主", "main", "primary",
		"ony: シュ(shu)", "kun: あるじ(arugi); ぬし(nushi)",
		"主人(shujin; master), 主要(shuyo; main), 主役(shuyaku; leading part)",
		"主(nushi; owner), 主(aruji; lord), 主婦(shufu; housewife)"},

	{"体", "body", "substance",
		"ony: タイ(tai)", "kun: からだ(karada)",
		"身体(shintai; body), 人体(jintai; human body), 容体(youtai; condition of body)",
		"体重(taijuu; body weight), 実体(jittai; reality), 実体験(jittiken; real experience)"},

	{"色", "color", "hue",
		"ony: ショク(shoku)", "kun: いろ(iro)",
		"原色(genshoku; primary color), 色彩(shikisai; coloration), 色調(shikichou; tone of color)",
		"赤色(sekishoku; red color), 色合い(iroai; tint), 肌色(hadairo; skin color)"},

	{"楽", "comfort", "ease",
		"ony: ラク(raku); ガク(gaku)", "kun: たの(tanoshii)",
		"安楽(anraku; comfort), 楽園(rakuen; paradise), 楽屋(gakuya; dressing room)",
		"楽しい(tanoshii; enjoyable), 楽勝(rakushou; easy win), 手軽(tegaru; easy)"},

	{"運", "fortune", "transport",
		"ony: ウン(un)", "kun: はこ(hako)",
		"運命(unmei; fate), 運転(unten; driving), 運搬(unpan; transport)",
		"運がいい(un ga ii; be lucky), 運ぶ(hakobu; transport), 運動(undou; exercise)"},

	{"空", "sky", "emptiness",
		"ony: クウ(kuu); ソラ(sora)", "kun: あく(aku); あき(aki); から(kara)",
		"空港(kuukou; airport), 宇宙空間(uchuu kukan; outer space), 空腹(kuufuku; hunger)",
		"空っぽ(karappo; empty), 空く(aku; become empty), 空(sora; sky)"},

	{"院", "institution", "temple",
		"ony: イン(in)", "kun: -",
		"大学院(daigakuin; graduate school), 研究所(kenkyuujo; research institute), 病院(byouin; hospital)",
		"寺院(jiin; temple), 政府機関(seifu kikan; government institution), 裁判所(saibansho; courthouse)"},

	{"市", "city", "town",
		"ony: シ(shi)", "kun: いち(ichi)",
		"都市(toshi; city), 市役所(shiyakusho; city hall), 市場(ichiba; marketplace)",
		"町(machi; town), 開市(kaishi; opening of market), 閉市(heishi; closing of market)"},

	{"区", "ward", "district",
		"ony: ク(ku)", "kun: -",
		"区役所(kuyakusho; ward office), 住宅地区(jutakuchiiki; residential district), 校区(kouku; school district)",
		"区切り(kugiri; division), 区分(kubun; classification), 特別区(tokubetsuku; special ward)"},

	{"政", "government", "politics",
		"ony: セイ(sei)", "kun: まつ(matsu)",
		"政府(seifu; government), 政策(seisaku; policy), 政治(seiji; politics)",
		"行政(gyousei; administration), 治政(chisei; governing), 政党(seitou; political party)"},

	{"正", "correct", "justice",
		"ony: セイ(sei)", "kun: ただ(tada); まさ(masa)",
		"正解(seikai; right answer), 正確(seikaku; accurate), 正月(shougatsu; New Year's)",
		"正しい(tadashii; correct), 正直(shoujiki; honest), 正午(shougo; noon)"},

	{"立", "stand", "establish",
		"ony: リツ(ritsu); リュウ(ryuu)", "kun: た(ta)",
		"立候補(rikouho; to stand for election), 立証(risshou; proof), 独立(dokuritsu; independence)",
		"立つ(tatsu; stand up), 建てる(tateru; build up), 存在(sonzai; existence)"},

	{"方", "direction", "person",
		"ony: ホウ(hou)", "kun: かた(kata)",
		"方法(houhou; method), 四方(shihou; four directions), 方向(houkou; direction)",
		"一方(ippo; one side), やり方(yarikata; way of doing), 人(hito; person)"},

	{"社", "company", "society",
		"ony: シャ(sha)", "kun: やしろ(yashiro)",
		"会社(kaisha; company), 出版社(shuppan sha; publisher), 神社(jinja; shrine)",
		"公社(kousha; public corporation), 社員(shain; employee), 社交(shakou; social intercourse)"},

	{"真", "true", "real",
		"ony: シン(shin)", "kun: ま(ma)",
		"真実(shinjitsu; truth), 真珠(shinju; pearl), 真空(shinkuu; vacuum)",
		"真面目(majime; serious), 真っ直ぐ(massugu; straight), 真(makoto; truth)"},

	{"切", "cut", "slit",
		"ony: セツ(setsu)", "kun: き(ki)",
		"切手(kitte; postage stamp), 切符(kippu; ticket), 切断(setsudan; disruption)",
		"切る(kiru; cut), 切開(kikai; cut open), 切り取る(kiri toru; cut off)"},

	{"然", "natural", "unconscious",
		"ony: ゼン(zen); ジ(ji)", "kun: しか(shika)",
		"当然(touzen; natural), 必然(hitsuzen; inevitable), 無意識(muishiki; unconsciously)",
		"然るべき(arubeki; proper), 然し(shikashi; however), 然も (shikamo; nevertheless)"},

	{"引", "pull", "subtract",
		"ony: イン(in)", "kun: ひ(hi)",
		"引っ張る(hipparu; pull), 引く(hiku; draw back), 引退(intai; retirement)",
		"引用(inyou; citation), 引き出す(hikidasu; deduce), 引き分け(hikiwake; tie game)"},

	{"止", "stop", "halt",
		"ony: シ(shi)", "kun: と(to)",
		"停止(teishi; stop), 中止(chuushi; suspension), 禁止(kinshi; prohibition)",
		"止まる(tomaru; stop), 止める(yameru; stop), 辞める(yameru; quit)"},

	{"見", "see", "view",
		"ony: ケン(ken)", "kun: み(mi)",
		"見解(kenkai; view), 山頂(sanchou; mountain top), 考え方(kangaekata; way of thinking)",
		"見る(miru; see), 見える(mieru; be visible), 一見(ikken; glance)"},

	{"思", "think", "thought",
		"ony: シ(shi)", "kun: おも(omo)",
		"思想(shisou; thought), 思考(shikou; thinking), 思い出(omoida; memory)",
		"思う(omou; think), 思い込む(omoidkomu; be convinced), 思い切って(omoikkiri; boldly)"},

	{"以", "by means of", "because",
		"ony: イ(i)", "kun: もっ(motte)",
		"以下(ika; below), 以内(inai; within), 以外(igai; except)",
		"以上(ijou; more than), 以来(irai; since), 以って(motte; by means of)"},

	{"取", "take", "choose",
		"ony: シュ(shu)", "kun: と(to)",
		"取る(toru; take), 受取(uketoru; receive), 取り出す(toridasu; take out)",
		"取引(torihiki; transaction), 取材(shuzai; coverage), 携帯(keitai; portable)"},

	{"親", "parent", "intimate",
		"ony: シン(shin)", "kun: おや(oya); したし(shitashi)",
		"親子(oyako; parent and child), 親友(shin'yu; close friend), 親しい(shitashii; intimate)",
		"養父(youfu; foster father), 養母(youbo; foster mother), 養親(youshin; foster parent)"},

	{"自", "oneself", "automatic",
		"ony: ジ(ji); シ(shi)", "kun: みずか(mizuka)",
		"自分(jibun; oneself), 自動(jidou; automatic), 自然(shizen; nature)",
		"自宅(jutaku; one's home), 自発(jihatsu; spontaneous), 自力(jirki; own power)"},

	{"界", "world", "limit",
		"ony: カイ(kai); ケ(ke)", "kun: -",
		"世界(sekai; world), 学界(gakkai; academic circle), 業界(gyoukai; industry)",
		"境界(kyoukai; border), 分界(bunkai; division), 定義(teigi; definition)"},

	{"令", "order", "rule",
		"ony: リョウ(ryou); レイ(rei)", "kun: みことのり(mikotonori)",
		"命令(meirei; order), 規則(kisoku; rule), 法令(houritsu; law)",
		"令状(reijou; warrant), 命令(meirei; command), 勅令(chokurei; imperial edict)"},

	{"変", "change", "strange",
		"ony: ヘン(hen)", "kun: か(ka)",
		"変化(henka; change), 変更(henkou; modification), 変異(hen'i; variation)",
		"変わる(kawaru; change), 変な(hen na; strange), 異変(ibiten; unusual event)"},

	{"完", "perfect", "completion",
		"ony: カン(kan)", "kun: -",
		"完璧(kanpeki; flawless), 完成(kansei; completion), 完全 (kanzen; perfect)",
		"完了(kanryou; completion), 完遂(kannsui; accomplishment), 完結(kanketsu; conclusion)"},

	{"数", "number", "figure",
		"ony: スウ(suu); シュ(shu); ス(su)", "kun: かず(kazu)",
		"数字(suuji; number), 数学(suugaku; mathematics), 数少ない(sukunai; few)",
		"勘定(kanjou; bill), 多数(tasuu; majority), 幾つ(ikutsu; how many)"},

	{"合", "join", "fit",
		"ony: ゴウ(gou); カッ(katsu); ガッ(gatu)", "kun: あ(a); あう(au)",
		"合同(godou; joint), 合成(gousei; synthesis), 結合(ketsugou; combination)",
		"合う(au; fit), 合意(goui; agreement), 手合わせ(teawase; applause)"},

	{"動", "move", "motion",
		"ony: ドウ(dou)", "kun: うご(ugo)",
		"移動(idou; movement), 動作(dousa; movement), 反動(handou; reaction)",
		"動く(ugoku; move), 動機(douki; motive), 感動(kandou; being moved)"},

	{"度", "degree", "time",
		"ony: ド(do); ト(to)", "kun: たび(tabi)",
		"温度(ondo; temperature), 三度(sand; three times), 十度(juutodo; ten degrees)",
		"再度(saido; again), 度々(tabitabi; often), 幾度(ikutabi; how many times)"},

	{"刀", "sword", "knife",
		"ony: トウ(tou)", "kun: かたな(katana)",
		"刀剣(touken; sword), 小刀(shoutou; knife), 短刀(tantou; short sword)",
		"刃(ha; blade), 刀身(toumi; sword blade), 刀工(toukou; swordsmith)"},

	{"友", "friend", "companion",
		"ony: ユウ(yuu)", "kun: とも(tomo)",
		"友人(yuujin; friend), 友好(yuukou; friendship), 親友(shin'yuu; close friend)",
		"仲間(nakama; companion), 取引先(torihikisaki; business connection), 共同(kyoudou; joint)"},

	{"品", "goods", "quality",
		"ony: ヒン(hin); ホン(hon)", "kun: しな(shina)",
		"商品(shouhin; goods), 品物(shinamono; article), 食品(shokuhin; food products)",
		"品格(hinkaku; dignity), 品質(hinshitsu; quality), 工芸品(kougeihin; craft goods)"},

	{"返", "return", "repay",
		"ony: ヘン(hen)", "kun: かえ(kae)",
		"返事(henji; reply), 帰り道(kaerimichi; return route), 返品(henpin; returned goods)",
		"返す(kaesu; return), 返る(kaeru; go back), 振り返る(furikaeru; look back)"},

	// 273 ? 1033/5=207

	// Here are 39 additional Japanese kanji flash cards that I think would be useful for beginning Japanese learners,
	// based on reviewing the cards I previously provided:

	{"大", "big", "large",
		"ony: ダイ(dai); タイ(tai)", "kun: おお(oo)",
		"大きい(ookii; big), 大人(otona; adult), 大学(daigaku; university)",
		"最大(saidai; biggest), 大量(tairyo; large quantity), 巨大(kyodai; huge)"},

	{"小", "small", "little",
		"ony: ショウ(shou); コ(ko)", "kun: ちい(chiisai)",
		"小さい(chiisai; small), 小学校(shougakkou; elementary school), 少女(shoujo; girl)",
		"最小(saishou; smallest), 少量(shouryou; small amount), 細かい(komakai; tiny)"},

	{"短", "short", "brief",
		"ony: タン(tan)", "kun: みじか(mijika)",
		"短い(mijikai; short), 短縮(tanshuku; reduction), 短所(tansho; shortcoming)",
		"短期(tanki; short time), 短距離(tankyori; short distance), 簡短(kantan; brief)"},

	{"高", "high", "tall",
		"ony: コウ(kou); タカ(taka)", "kun: たか(taka)",
		"高い(takai; high), 最高(saikou; highest), 高原(kougen; plateau)",
		"高校(koukou; high school), 高層(kousou; high-rise), 高速(kousoku; high speed)"},

	{"低", "low", "short",
		"ony: テイ(tei)", "kun: ひく(hiku)",
		"低い(hikui; low), 最低(saitei; lowest), 低地(teichi; lowland)",
		"低価(teika; low price), 低レベル(teireberu; low level), 低下(teika; decline)"},

	{"多", "many", "much",
		"ony: タ(ta)", "kun: おお(oo)",
		"多い(ooi; many), 多数(tasuu; majority), 多量(tairyo; large quantity)",
		"多く(ooku; many), 沢山(takusan; a lot), 豊富(houfu; abundance)"},

	{"少", "few", "little",
		"ony: ショウ(shou)", "kun: すく(suku)",
		"少ない(sukunai; few), 少数(shousuu; minority), 少量(shouryou; small amount)",
		"わずか(wazuka; little), 少し(sukoshi; a little), 乏しい(toboshii; scarce)"},

	{"新", "new", "refresh",
		"ony: シン(shin)", "kun: あたら(atara)",
		"新品(shinpin; new article), 新築(shinchiku; new construction), 新設(shinsetsu; new establishment)",
		"新人(shinjin; new employee), 更新(koushin; renewal), 新鮮(shinsen; freshness)"},

	{"旧", "old", "former",
		"ony: キュウ(kyuu); コ(ko)", "kun: ふる(furui)",
		"旧友(kyuuyuu; old friend), 旧姓(kyuusei; maiden name), 旧弊(kyuuhai; old practice)",
		"昔(mukashi; in the past), 古い(furui; old), 用いられた(mochiirareru; used)"},

	{"白", "white", "unmarked",
		"ony: ハク(haku); ビャク(byaku)", "kun: しろ(shiro)",
		"白色(hakushoku; white color), 純白(junpaku; pure white), 白人(hakujin; Caucasian)",
		"空白(kuuhaku; blank space), 白状(hakujou; confession), 白日(byakujitsu; daylight)"},

	{"黒", "black", "dark",
		"ony: コク(koku)", "kun: くろ(kuro)",
		"黒色(koku shoku; black color), 黒板(kokuban; blackboard), 真っ黒(makkuro; pitch black)",
		"暗黒(ankoku; darkness), 黒い(kuroi; black), 黒雲(kurokumo; black cloud)"},

	{"赤", "red", "flush",
		"ony: セキ(seki); シャク(shaku)", "kun: あか(aka)",
		"赤色(sekshoku; red color), 真っ赤(makka; bright red), 赤道(sekidou; equator)",
		"赤信号(akashingou; red traffic light), 赤み(akamimi; reddish), 赤面(sekimen; flushing)"},

	{"青", "blue", "green",
		"ony: セイ(sei); ショウ(shou)", "kun: あお(ao)",
		"青色(aoshoku; blue color), 深青(shinkou; deep blue), 青年(seinen; youth)",
		"青信号(aoshin gou; green traffic light), 青々(aoao; verdant), 青臭い(aokusai; fishy)"},

	{"男", "man", "male",
		"ony: ダン(dan); ナン(nan)", "kun: おとこ(otoko)",
		"男性(dansei; male), 男子(danshi; boy), 男っ気(otokogokoro; manliness)",
		"男の人(otoko no hito; man), 男らしい(otokorashii; manly), 男子学生(danshigakusei; male student)"},

	{"女", "woman", "female",
		"ony: ジョ(nyo); ニョ(nyo)", "kun: おんな(onna)",
		"女性(josei; female), 女子(joshi; girl), 女房(nyoubou; wife)",
		"女の人(onna no hito; woman), 女らしい(onnarashii; feminine), 女子学生(joshigakusei; female student)"},

	{"車", "car", "vehicle",
		"ony: シャ(sha)", "kun: くるま(kuruma)",
		"自動車(jidousha; automobile), 電車(densha; train), 車体(shatai; car body)",
		"新車(shinchya; new car), 車輪(sharin; wheel), 三輪車(sanrinsha; tricycle)"},

	{"本", "book", "origin",
		"ony: ホン(hon)", "kun: もと(moto)",
		"本屋(honya; bookstore), 本物(honmono; real thing), 日本(nihon; Japan)",
		"元(moto; origin), 本(hon; book), 本当(hontou; real)"},

	{"火", "fire", "Tuesday",
		"ony: カ(ka)", "kun: ひ(hi)",
		"火事(kaji; fire), 火曜日(kayoubi; Tuesday), 炎(honoo; flame)",
		"火山(kazan; volcano), 火星(kasei; Mars), 花火(hanabi; fireworks)"},

	{"水", "water", "Thursday",
		"ony: スイ(sui)", "kun: みず(mizu)",
		"水泳(suiei; swimming), 水曜日(suiyoubi; Wednesday), 水分(suibun; moisture)",
		"飲料水(inryousui; drinking water), 水産物(suisanbutsu; seafood), 水兵(suihei; sailor)"},

	{"木", "tree", "wood",
		"ony: モク(moku)", "kun: き(ki)",
		"木曜日(mokuyoubi; Thursday), 木材(mokuzai; lumber), 森(mori; forest)",
		"樹木(ju moku; tree), 家具(kagu; furniture), 椅子(isu; chair)"},

	{"金", "gold", "money",
		"ony: キン(kin)", "kun: かね(kane)",
		"金曜日(kin'youbi; Friday), 金メダル(kinmedaru; gold medal), お金(okane; money)",
		"金庫(kinko; safe), 金魚(kingyo; goldfish), 銀行(ginkou; bank)"},

	{"土", "soil", "Saturday",
		"ony: ド(do); ト(to)", "kun: つち(tsuchi)",
		"土曜日(doyoubi; Saturday), 土地(tochi; land), 土産(miyage; souvenir)",
		"大地(daichi; earth), 土台(dodai; foundation), 黒土(kurotsuchi; black soil)"},

	{"耳", "ear", "hearing",
		"ony: ジ(ji)", "kun: みみ(mimi)",
		"耳鳴り(miminari; ringing in ears), 耳栓(mimisen; earplugs), 耳飾り(mimi kazari; earring)",
		"聴覚(choukaku; hearing), 聞こえる(kikoeru; audible), 聞く(kiku; listen)"},

	// Here are some additional common kanji that are often seen in Japanese train stations:

	{"入", "enter", "insert",
		"ony: ニュウ(nyuu)", "kun: い(i)る",
		"入口(irikuchi; entrance), 入場(nyujou; admission), 入力(nyuryoku; input)",
		"入学(nyugaku; enrollment), 入れる(ireru; insert), 入る(hairu; enter)"},

	{"出", "exit", "leave",
		"ony: シュツ(shitsu); シュ(shu)", "kun: だ(da)す",
		"出口(deguchi; exit), 出発(shuppatsu; departure), 出席(shusseki; attendance)",
		"出す(dasu; take out), 出る(deru; leave), 出かける(dekakeru; go out)"},

	{"乗", "ride", "board",
		"ony: ジョウ(jou)", "kun: の(no)る",
		"乗客(joukyaku; passenger), 乗り物(norimono; vehicle), 乗車(jousha; boarding)",
		"乗る(noru; ride), 乗せる(noseru; give a ride), 乗り換え(norikae; transfer)"},

	{"降", "descend", "alight",
		"ony: コウ(kou)", "kun: お(o)りる",
		"降車(kousha; getting off), 降水(kousui; precipitation), 降参(kousan; visit)",
		"降る(oriru; alight), 降ろす(orosu; unload), 降り注ぐ(furishizuku; pour down)"},

	{"行", "go", "travel",
		"ony: コウ(kou); ギョウ(gyou)", "kun: い(i)く; ゆ(yu)く",
		"行く(iku; go), 行列(gyouretsu; line), 航行(koukou; navigation)",
		"旅行(ryokou; travel), 帰行(kigyou; return trip), 出発(shuppatsu; departure)"},

	{"来", "come", "next",
		"ony: ライ(rai)", "kun: く(ku)る",
		"来客(raikyaku; visitor), 来月(raigetsu; next month), 来年(rainen; next year)",
		"来る(kuru; come), 来て(kite; please come), 参上(sanjou; coming humbly)"},

	{"当", "right", "hit",
		"ony: トウ(tou)", "kun: あ(a)たる",
		"当てる(ateru; allot), 当日(toujitsu; that day), 当直(toujik; duty)",
		"当たり(atari; hit), 当たる(ataru; be applicable), 駅当(ekiatari; for the station)"},

	{"次", "next", "order",
		"ony: ジ(ji)", "kun: つぎ(tsugi)",
		"次第(shidai; order), 次回(jikai; next time), 次男(jinan; second son)",
		"次に(tsugi ni; next), 次から次へ(tsugi kara tsugi e; one after another), 続く(tsuzuku; follow)"},

	{"改", "revision", "improvement",
		"ony: カイ(kai)", "kun: あら(ara)たる",
		"改善(kaizen; improvement), 改札(kaisatsu; ticket gate), 改装(kaisou; renovation)",
		"改正(kaisei; amendment), 改める(arata meru; revise), 新しい(atarashii; new)"},

	{"番", "number", "one's turn",
		"ony: バン(ban)", "kun: -",
		"番号(bangou; number), 順番(junban; turn), 番組(bangumi; TV/radio program)",
		"番(tsugi; turn), 番人(bannin; watchman), 係番(kakariban; attendant)"},

	{"線", "line", "railway line",
		"ony: セン(sen)", "kun: -",
		"路線(rosen; route), 鉄道線(tetsudou sen; railway line), 直線(chokusen; straight line)",
		"境界線(kyoukaisen; border line), ライン(rain; line), 線路(senro; railroad track)"},

	{"号", "number", "counter for vehicles",
		"ony: ゴウ(gou)", "kun: - ",
		"列車号(ressha gou; train number), 便名(binmei; service name), 航空便(koukuu bin; airline service)",
		"番号(bangou; number), 100号車(hyaku gou sha; car number 100), 3号線(san gou sen; line number 3)"},

	// the kanji for entrance/exit, ride/alight, go/come, next, station names, lines, and train/service numbers are very common in train stations.

	// Here are some common kanji related to taxis:

	{"走", "run", "drive",
		"ony: ソウ(sou)", "kun: はし(hashi)る",
		"走行(soukou; driving), 走査(sousa; survey), 全力走行(zenryoku soukou; driving at full speed)",
		"走る(hashiru; run), 運転する(unten suru; drive), 疾走(shissou; dashing)"},

	{"運", "transport", "luck",
		"ony: ウン(un)", "kun: はこ(hako)ぶ",
		"運転(unten; driving), 運行(unkou; operation), 運命(unmei; fate)",
		"運ぶ(hakobu; transport), 運搬(unpan; freight), 運が良い(un ga ii; be lucky)"},

	{"料", "fee", "ingredients",
		"ony: リョウ(ryou)", "kun: -",
		"料金(ryoukin; fee), 入場料(nyuuryou ryou; admission fee), 材料(zairyou; ingredient)",
		"請求料金(seikyuu ryoukin; billed fee), 配達料金(haitatsu ryoukin; delivery fee)"},

	{"税", "tax", "duty",
		"ony: ゼイ(zei)", "kun: -",
		"消費税(shouhi zei; consumption tax), 申告税額(shinkoku zeigaku; declared tax amount), 関税(kanzei; customs duty)",
		"付加価値税(fukka kachi zei; value added tax), 税抜き(zeibaraki; tax excluded), 非課税(hikazei; tax free)"},

	{"計", "measure", "meter",
		"ony: ケイ(kei)", "kun: はか(haka)る",
		"計算(keisan; calculation), 計画(keikaku; plan), 体温計(taionkei; thermometer)",
		"計る(hakaru; measure), 距離計(kyori kei; odometer), 料金計算機(ryoukin keisanki; fare meter)"},

	{"配", "distribute", "dispatch",
		"ony: ハイ(hai)", "kun: くば(kuba)る",
		"配達(haitatsu; delivery), 配布(haifu; distribution), 配置(haichi; arrangement)",
		"配る(kubaru; distribute), 配慮(hairyo; consideration), 配当(haitou; dividend)"},

	// kanji for drive, transport, fee, tax, meter, and distribute/dispatch relate to taxis.

	// Here are some common kanji related to restaurants:

	{"食", "eat", "food",
		"ony: ショク(shoku)", "kun: く(ku)う",

		"食事(shokuji; meal), 食品(shokuhin; food), 食堂(shokudou; dining hall)",
		"食べる(taberu; eat), 朝食(choushoku; breakfast), 献立(kondate; menu)"},

	{"料", "fee", "ingredients",
		"ony: リョウ(ryou)", "kun: -",

		"料金(ryoukin; fee), 食料(shokuryou; provisions), 原料(genryou; ingredient)",
		"材料(zairyou; ingredient), 調理料理(chouri ryouri; cooked food), 配膳料(haizen ryou; table charge)"},

	{"店", "store", "shop",
		"ony: テン(ten)", "kun: みせ(mise)",

		"店舗(tenpo; shop), 書店(shoten; bookstore), 八百屋(yaoya; greengrocer's)",
		"売店(uriba; shop), 喫茶店(kissaten; cafe), 飲食店(inshokuten; restaurant)"},

	{"酒", "alcohol", "sake",

		"ony: シュ(shu)", "kun: さけ(sake)",
		"酒場(sakaba; bar), 日本酒(nihonshu; sake), ビール(biiru; beer)",
		"酒類(shurui; alcoholic beverages), 酒飲み(sakenomi; drinking alcohol), 酔う(you; get drunk)"},

	{"注", "pour", "order",
		"ony: チュウ(chuu)", "kun: そそ(soso)ぐ",

		"注文(chuumon; order), 注意(chuui; attention), 注入(chunyu; injection)",
		"注ぐ(sosogu; pour), 注ぎ口(sogikuchi; spout), 追注(tsuichuu; additional order)"},

	{"席", "seat", "occasion",

		"ony: セキ(seki)", "kun: -",
		"席順(seki jun; seating order), 席取り(seki dori; seat reservation), 出席(shusseki; attendance)",
		"空席(kuseki; vacancy), 指定席(shitei seki; reserved seat), カウンター席(kaunta seki; counter seat)"},

	{"報", "report", "information",
		"ony: ホウ(hou)", "kun: -",
		"報道(houdou; news report), 報告(houkoku; report), 情報(jouhou; information)",
		"速報(sokuhou; bulletin), ニュース(nyuusu; news), 天気予報(tenki yohou; weather forecast)"},

	{"気", "feeling", "spirit",
		"ony: キ(ki)", "kun: -",
		"気分(kibun; feeling), 気持ち(kimochi; feeling), 気象(kishou; weather)",
		"大気(taiki; atmosphere), 気圧(kiatsu; atmospheric pressure), 気候(kikou; climate)"},

	{"地", "earth", "place",
		"ony: チ(chi); ジ(ji)", "kun: -",
		"地震(jishin; earthquake), 地域(chiiki; region), 土地(tochi; land)",

		"現地(genchi; site), 出身地(shusshin chi; hometown), 地名(chimei; place name)"},

	{"国", "country", "homeland",
		"ony: コク(koku)", "kun: くに(kuni)",

		"国家(kokka; country), 他国(takoku; foreign country), 祖国(sokoku; homeland)",
		"外国(gaikoku; foreign country), 国境(kokkyou; national border), 国際(kokusai; international)"},

	{"政", "government", "politics",
		"ony: セイ(sei)", "kun: まつ(matsu)",
		"政治(seiji; politics), 政府(seifu; government), 治政(chisei; governing)",

		"行政(gyousei; administration), 政党(seitou; political party), 政策(seisaku; policy)"},

	{"経", "sutra", "longitude",
		"ony: ケイ(kei); キョウ(kyou)", "kun: -",

		"経済(keizai; economy), 経験(keiken; experience), 人生経験(jinsei keiken; life experience)",
		"経理(keiri; accounting), 経過(keika; progress), 経営(keiei; management)"},

	// Here are some common kanji related to Japanese cities and places:

	{"東", "east", "eastern",
		"ony: トウ(tō)", "kun: ひがし(higashi)",
		"東京(Tōkyō; Tokyo), 東北(Tōhoku; northeast region), 東海道(Tōkaidō; Tokaido region)",
		"東口(Higashiguchi; east exit), 東側(higashigawa; eastern side), 東向き(higashimuki; facing east)"},

	{"西", "west", "western",
		"ony: セイ(sei); サイ(sai)", "kun: にし(nishi)",
		"西日本(Nishi-Nihon; western Japan), 西洋(Seiyō; the West), 西南(Seinan; southwest)",
		"西口(Nishiguchi; west exit), 西側(nishigawa; western side), 西向き(nishimuki; facing west)"},

	{"南", "south", "southern",
		"ony: ナン(nan)", "kun: みなみ(minami)",
		"南国(Nangoku; southern country), 南極(Nankyoku; South Pole), 南北(Nanboku; north-south)",
		"南口(Minamiguchi; south exit), 南側(minamigawa; southern side), 南向き(minamimuki; facing south)"},

	{"京", "capital", "metropolis",
		"ony: キョウ(kyō); ケイ(kei)", "kun: -",
		"東京(Tōkyō; Tokyo), 京都(Kyōto; Kyoto), 大阪(Ōsaka; Osaka)",
		"首都(Shuto; capital city), 特別区(Tokubetsuku; special ward), 首長(Shuchō; mayor)"},

	{"市", "city", "town",
		"ony: シ(shi)", "kun: いち(ichi)",
		"都市(Toshi; city), 市場(Ichiba; marketplace), 開市(Kaishi; opening of market)",
		"町(Machi; town), 自治体(Jichitai; municipality), 市長(Shichō; mayor)"},

	{"区", "ward", "district",
		"ony: ク(ku)", "kun: -",
		"区役所(Kuyakusho; ward office), 特別区(Tokubetsuku; special ward), 校区(Kōku; school district)",
		"分区(Bunku; division), 地区(Chiku; local area), 区長(Kuchō; ward mayor)"},

	{"町", "town", "street",
		"ony: チョウ(chō)", "kun: まち(machi)",
		"街路(Kairō; street), 町中(Machinaka; downtown), 町方言(Machigotokukai; local dialect)",
		"町人(Chōnin; townspeople), 町工場(Machikōjō; local factory), 町長(Chōchō; town mayor)"},

	// Here are the Shinjuku area kanji again with the proper formatting:

	{"新", "new", "refresh, renew",
		"ony: シン(shin)",
		"kun: あら(ara)",
		"新宿(Shinjuku), 新大久保(Shin-Ōkubo), 新南口(Shin-Minamiguchi)",
		"新築(shinchiku; new construction), 新設(shinsetsu; new establishment), 新参者(shinzanza; new entrant)",
	},

	{"宿", "inn", "lodge, lodgement",
		"ony: シュク(shuku)",
		"kun: やど(yado)",
		"新宿(Shinjuku), 宿泊(Shukuhaku), 旅館(Ryokan)",
		"宿場(shukuba; lodging place), 宿坊(shukubō; temple lodging), 宿直(shukujitsu; night duty)",
	},

	{"垢", "deepest", "dirt, scum",
		"ony: コ(ko)",
		"kun: あか(aka)",
		"歌舞伎町(Kabukichō), 最低の場所(Saitei no basho), 赤線地区(Akasen chiku)",
		"汚名(osen; disgrace), 汚点(oten; stain), 最底辺(saiteihen; rock bottom)"},

	{"谷", "valley", "ravine, gorge",
		"ony: コク(kok); タニ(tani)",
		"kun: や(ya)",
		"渋谷(Shibuya), 谷底(Tanisoko), 渓谷(Keikoku)",
		"渓谷(keikoku; ravine), 山間(sankan; mountain recesses), 盆地(bonchi; basin)"},

	{"ヶ", "counter for items", "-",
		"ony: ケ(ke); カ(ka)",
		"kun: -",
		"一ヶ月(Ikka getsu), 三ヶ国語(Sanka kokugo), 一ヶ所(Ikasho)",
		"",
	},

	{"原", "meadow", "field, wilderness",
		"ony: ゲン(gen)",
		"kun: はら(hara)",
		"代々木原(Yoyogi-hara), 平原(Heigen), 草原(Sōgen)",
		"原野(hara; field), 原(hara; meadow), 野原(nohara; field)"},

	{"西", "west", "western",
		"ony: セイ(sei); サイ(sai)",
		"kun: にし(nishi)",
		"西新宿(Nishi-Shinjuku), 西方(Seihō), 西日(Nishibi)",
		"西洋(seiyō; the West), 西暦(seireki; Western calendar), 西口(nishiguchi; west exit)"},

	{"口", "entrance", "opening, hole",
		"ony: コウ(kō); ク(ku)",
		"kun: ぐち(guchi)",
		"新南口(Shin-Minamiguchi), 東口(Higashiguchi), 駅口(Ekiguchi)",
		"入口(irikuchi; entrance), 出口(deguchi; exit), 穴口(anaguchi; hole opening)"},

	{"猫", "cat", "pet",
		"ony: ビョウ(byou)",
		"kun: ねこ(neko)",
		"猫(neko), 猫ちゃん(neko-chan), 子猫(koneko)",
		"ねこねこ(nekoneko), 飼い猫(kaineko), ネコアレルギー(nekoarerugii)"},

	{"犬", "dog", "pet",
		"ony: ケン(ken); ク(ku)",
		"kun: いぬ(inu)",
		"犬(inu), 子犬(koinu), わんわん(wanwan)",
		"飼い犬(kaikenu), 盲導犬(moudouken), イヌアレルギー(inuarergii)"},

	{"料理", "cooking", "cuisine",
		"ony: リョウリ(ryouri)",
		"kun: -",
		"料理(ryouri), 和食(washoku), 中華料理(chuuka ryouri)",
		"調理(chouri), 献立(kondate), レシピ(resipi)"},

	{"考える", "think", "ponder",
		"ony: コウエル(koueru)",
		"kun: かんがえる(kangaeru)",
		"考える(kangaeru), 思考(shikou), 熟考(jukkou)",
		"思い巡らす(omoidmegurasu), 推論(suiron), 思索(shisaku)"},

	{"笑う", "laugh", "smile",
		"ony: ショウ(shou)",
		"kun: わら(wara)う",
		"笑う(warau), 笑顔(egao), 苦笑(nigawarai)",
		"吹き出す(fukisadasu), 愉快(yukai), 楽しげ(tanoshige)"},

	{"走る", "run", "dash",
		"ony: ソウ(sou)",
		"kun: はし(hashi)る",
		"走る(hashiru), 全力走行(zenryoku soukou), 走り回る(hashirimawaru)",
		"駆ける(kakeru), マラソン(marason), ランニング(ran'ningu)"},

	{"読む", "read", "peruse",
		"ony: ドク(doku); トウ(tou)",
		"kun: よ(yo)む",
		"読む(yomu), 読書(dokusho), 披露(hiryou)",
		"文学(bungaku), 読者(dokusha), 朗読(roudoku)"},

	{"会う", "meet", "interview",
		"ony: カイ(kai); ガイ(gai); エ(e)",
		"kun: あ(a)う",
		"会う(au), 出会う(deau), 面会(menkai)",
		"面接(mensetsu), デート(deeto), 盲目(moumoku)"},

	{"買う", "buy", "purchase",
		"ony: バイ(bai)",
		"kun: か(ka)う",
		"買う(kau), 購入(kounyuu), 買い物(kaimono)",
		"消費者(shouhisha), セール(seeru), 購買(koubaai)"},

	{"住む", "live", "reside",
		"ony: ジュウ(juu)",
		"kun: す(su)む",
		"住む(sumu), 居住(kyojuu), 定住(teijuu)",
		"生活(seikatsu), 一人暮らし(hitorigurashi), 住民(juumin)"},

	{"風邪", "cold", "flu",
		"ony: フウジャ(fuuja)",
		"kun: かぜ(kaze)",
		"風邪(kaze), 感冒(kanbou), 発熱(hatsunetsu)",
		"病気(byouki), 予防接種(yobousesshu), 症状(shoujou)"},

	{"家", "house", "home",
		"ony: カ(ka); ケ(ke)",
		"kun: いえ(ie); や(ya)",
		"家(ie), 住宅(juutaku), 我が家(wagaya)",
		"屋敷(yashiki), 建物(tatemono), 家具(kagu)"},

	{"学校", "school", "campus",
		"ony: ガッコウ(gakkou)",
		"kun: -",
		"学校(gakkou), 中学校(chuugakkou), 大学(daigaku)",
		"教室(kyoushitsu), 制服(seifuku), 入学試験(nyuugaku shiken)"},

	{"国", "country", "state",
		"ony: コク(koku); クニ(kuni)",
		"kun: くに(kuni)",
		"国(kuni), 他国(takoku), アメリカ合衆国(amerika gasshukoku)",
		"国内(kokunai), 外国(gaikoku), 国旗(kokki)"},

	{"赤", "red", "scarlet",
		"ony: セキ(seki); シャク(shaku)",
		"kun: あか(aka)",
		"赤(aka), 真っ赤(makka), 赤信号(akashingou)",
		"紅葉(kouyou), 赤ちゃん(akachan), 赤十字(sekijuuji)"},

	{"青", "blue", "green",
		"ony: セイ(sei); ショウ(shou)",
		"kun: あお(ao)",
		"青(ao), 深青(shinkou), 青信号(aoshin)",
		"鮮やか(azayaka), 青年(seinen), 青春(seishun)"},

	{"言う", "say", "speak",
		"ony: ゲン(gen); ゴン(gon)",
		"kun: い(i)う",
		"言う(iu), 発言(hatsugen), 言いたい(iitai)",
		"語る(kataru), 主張(shuuchou), 表現(hyougen)"},

	{"山", "mountain", "nature",
		"ony: サン(san)",
		"kun: やま(yama)",
		"山(yama), 富士山(fujisan), 山頂(sanchou)",
		"登山(tozan), 高原(kougen), 山盛り(yama mori)"},

	{"川", "river", "stream",
		"ony: セン(sen)",
		"kun: かわ(kawa)",
		"川(kawa), 河川(kasen), 川岸(kawagishi)",
		"渓流(keiryuu), 用水路(yousuiro), 河口(kakou)"},

	{"田", "rice field", "paddy",
		"ony: デン(den)",
		"kun: た(ta)",
		"田(ta), 田んぼ(tanbo), 農場(noujou)",
		"稲作(inaku), 田園(den'en), 自給自足(jikyuu jisoku)"},

	{"雨", "rain", "precipitation",
		"ony: ウ(u)",
		"kun: あめ(ame); あま(ama)",
		"雨(ame), 大雨(ooame), 梅雨(tsuyu)",
		"降水(kousui), 雨乞い(amagoi), 打たれ強し(utare tsuyoshi)"},

	{"花", "flower", "blossom",
		"ony: カ(ka)",
		"kun: はな(hana)",
		"花(hana), 花瓶(kabin), 花見(hanami)",
		"草花(kusabana), 花言葉(hanakotoba), 花壇(hanadan)"},

	{"人", "person", "human",
		"ony: ジン(jin); ニン(nin)",
		"kun: ひと(hito)",
		"人(hito), 人口(jinkou), 人数(ninzuu)",
		"個人(kojin), 人生(jinsei), 人形(ningyou)"},

	// Here are 10 example kanji flashcards related to Japanese culture:

	{"歌", "song", "sing",
		"ony: カ(ka)",
		"kun: うた(uta)",
		"歌う(utau), 歌声(utagoe), 歌手(kashu)",
		"合唱(gasshou), 歌詞(kashi), 歌謡(kayou)"},

	{"舞", "dance", "perform",
		"ony: ブ(bu)",
		"kun: ま(ma)う",
		"舞う(mau), 舞踊(buyō), 舞台(butai)",
		"演舞(enbu), 踊り子(odoriko), 盆踊り(bon odori)"},

	{"茶", "tea", "ceremony",
		"ony: チャ(cha)",
		"kun: -",
		"お茶(ocha), 茶道(sadō), 茶屋(chaya)",
		"茶室(chashitsu), 抹茶(matcha), 茶請け(chadōgu)"},

	{"華", "flower", "splendor",
		"ony: カ(ka)",
		"kun: はな(hana)",
		"華道(kadō), 花華(hanaba), 華麗(karei)",
		"華美(kabi), 華やか(hanayaka), 花びら(hanabira)"},

	{"酒", "alcohol", "sake",
		"ony: シュ(shu)",
		"kun: さけ(sake)",
		"日本酒(nihonshu), 酒場(sakaba), 酒宴(sakuen)",
		"酔う(you), 酒量(shuryou), 酒席(shaseki)"},

	{"刀", "sword", "katana",
		"ony: トウ(tō)",
		"kun: かたな(katana)",
		"刀剣(tōken), 小刀(shōtō), 日本刀(nihontō)",
		"刀工(tōkō), 刀身(tōmi), 刀鍛冶(katana kaji)"},

	{"形", "form", "shape",
		"ony: ケイ(kei); ギョウ(gyō)",
		"kun: かた(kata); かたち(katachi)",
		"形状(keijō), 形式(keishiki), 人形(ningyō)",
		"模様(moyō), 型(kata), 儀式(gishiki)"},

	{"花", "flower", "blossom",
		"ony: カ(ka)",
		"kun: はな(hana)",
		"花道(hanamichi), 花見(hanami), 花壇(hanadan)",
		"花火(hanabi), 花びら(hanabira), 草花(kusabana)"},

	{"神", "god", "deity",
		"ony: シン(jin); ジン(jin)",
		"kun: かみ(kami)",
		"神社(jinja), 神道(shintō), 神様(kamisama)",
		"八百万の神(yaoyorozu no kami), 祈願(kigan), 祈祷(kitō)"},

	// Here are some additional foundational or high-frequency kanji we have not covered yet:

	{"大", "big", "large",
		"ony: ダイ(dai); タイ(tai)",
		"kun: おお(oo)",
		"大きい(ookii), 大人(otona), 大学(daigaku)",
		"最大(saidai), 大量(tairyo), 巨大(kyodai)"},

	{"小", "small", "little",
		"ony: ショウ(shou); コ(ko)",
		"kun: ちい(chiisai)",
		"小さい(chiisai), 小学校(shougakkou), 少女(shoujo)",
		"最小(saishou), 少量(shouryou), 細かい(komakai)"},

	{"短", "short", "brief",
		"ony: タン(tan)",
		"kun: みじか(mijika)",
		"短い(mijikai), 短縮(tanshuku), 短所(tansho)",
		"短期(tanki), 短距離(tankyori), 簡短(kantan)"},

	{"高", "high", "tall",
		"ony: コウ(kou); タカ(taka)",
		"kun: たか(taka)",
		"高い(takai), 最高(saikou), 高原(kougen)",
		"高校(koukou), 高層(kousou), 高速(kousoku)"},

	{"低", "low", "short",
		"ony: テイ(tei)",
		"kun: ひく(hiku)",
		"低い(hikui), 最低(saitei), 低地(teichi)",
		"低価(teika), 低レベル(teireberu), 低下(teika)"},

	{"多", "many", "much",
		"ony: タ(ta)",
		"kun: おお(oo)",
		"多い(ooi), 多数(tasuu), 多量(tairyo)",
		"豊富(houfu), 沢山(takusan), たくさん(takusan)"},

	{"少", "few", "little",
		"ony: ショウ(shou)",
		"kun: すく(suku)",
		"少ない(sukunai), 少数(shousuu), 少量(shouryou)",
		"わずか(wazuka), 少し(sukoshi), 乏しい(toboshii)"},

	{"新", "new", "refresh",
		"ony: シン(shin)",
		"kun: あら(ara)たら",
		"新品(shinpin), 新築(shinchiku), 新設(shinsetsu)",
		"新人(shinjin), 更新(koushin), 新鮮(shinsen)"},

	{"旧", "old", "former",
		"ony: キュウ(kyuu); コ(ko)",
		"kun: ふる(furui)",
		"旧友(kyuuyuu), 旧姓(kyuusei), 旧弊(kyuuhai)",
		"昔(mukashi), 古い(furui), 用いられた(mochiirareru)"},

	/*
			From the initial Shinjuku area kanji:
			新, 宿, 垢, 谷, ヶ, 原, 西, 口

			From the kanji relating to transportation:

			入, 出, 乗, 降, 行, 来, 当, 次, 改, 番, 線, 号

			From the kanji relating to restaurants:
			食, 料, 店, 酒, 注, 席

			From the kanji relating to news/TV:
			報, 気, 地, 国, 政, 経

			From the kanji relating to geography:
			東, 西, 南, 北, 京, 市, 区, 町

			From the kanji relating to culture:
			歌, 舞, 茶, 華, 酒, 刀, 書, 形, 花, 神

			From the foundational kanji:

			大, 小, 長, 短, 高, 低, 多, 少, 新, 旧

		From the people/relationships kanji:
		猫, 犬, 料理, 考える, 笑う, 走る, 書く, 読む, 会う, 買う, 住む, 風邪, 家, 学校, 国, 赤, 青, 言う, 書く, 山, 川, 田, 雨, 花, 人

		From the culture kanji:

		歌, 舞, 茶, 華, 酒, 刀, 書, 形, 花, 神

		From the foundational kanji:
		大, 小, 長, 短, 高, 低, 多, 少, 新, 旧

		From the Shinjuku area kanji:
		新, 宿, 垢, 谷, ヶ, 原, 西, 口

		From the transportation kanji:
		入, 出, 乗, 降, 行, 来, 当, 次, 改, 番, 線, 号

		From the restaurant kanji:
		食, 料, 店, 酒, 注, 席

		From the news/TV kanji:
		報, 気, 地, 国, 政, 経

		From the geography kanji:

		東, 西, 南, 北, 京, 市, 区, 町
	*/

	{"老", "old", "aged",
		"ony: ロウ(rou)",

		"kun: お(o)いる",
		"老人(roujin), 高齢者(koureisha), 老齢(rouireki)",
		"老若男女(rounyounan), 老体(otai), 老衰(rousui)"},

	{"耄", "aged", "long-lived",
		"ony: ヒ(hi); ホウ(hou)",
		"kun: としよ(toshiyo)",
		"耄碌(hiruko), 高齢(kourei), 長寿(chouju)",

		"老人(roujin), 寿命(jumyou), 老衰(rousui)"},

	// Here are some example kanji related to retirement and immigration:

	{"退", "retreat", "retire",
		"ony: タイ(tai)",
		"kun: の(no)く; し(shi)る",
		"退職(taishoku), 退院(taiiin), 引退(intai)",
		"やめる(yameru), 去る(saru), 下がる(sagaru)"},

	{"休", "rest", "retire",
		"ony: キュウ(kyū)",
		"kun: やす(yasu)む",
		"休暇(kyūka), 休日(kyūjitsu), 休業(kyūgyō)",
		"休む(yasumu), 打ち切る(uchikiru), 止める(yameru)"},

	{"社", "company", "firm",
		"ony: シャ(sha)",
		"kun: -",
		"会社(kaisha), 企業(kigyō), 株式会社(kabushikigaisha)",
		"商社(shōsha), 事業所(jigyōsho), 会見(kaiken)"},

	{"移", "move", "transfer",
		"ony: イ(i)",
		"kun: うつ(utsu)る",
		"移動(idō), 移民(imin), 移管(ikan)",
		"引越(hikkoshi), 交替(kōtai), 移す(utusus)"},

	{"民", "people", "populace",
		"ony: ミン(min)",
		"kun: たみ(tami)",
		"国民(kokumin), 民族(minzoku), 市民(shimin)",
		"人民(jinmin), 民衆(minshū), 民生(minsei)"},

	{"外", "outside", "foreign",
		"ony: ガイ(gai)",
		"kun: そと(soto); ほ(ho)か",
		"外国(gaikoku), 外出(gaishutsu), 外観(gaikan)",
		"外来(gailai), 外交(gaikō), 外部(gaihu)"},

	{"国", "country", "nation",
		"ony: コク(koku); ク(ku)ニ",
		"kun: くに(kuni)",
		"国家(kokka), 他国(takoku), 国土(kokudo)",
		"国境(kokkyō), 国民(kokumin), 国際(kokusai)"},

	// Here are some example kanji related to technology and engineering:

	{"機", "machine", "mechanism",
		"ony: キ(ki)",
		"kun: はた(hata)",
		"機械(kikai), 機能(kinou), 器機(kiki)",
		"装置(souchi), 御器(mikomori), 器具(kigu)"},

	{"車", "car", "vehicle",
		"ony: シャ(sha)",
		"kun: くるま(kuruma)",
		"自動車(jidousha), 電車(densya), 車両(sharyou)",
		"新車(shinsha), 車輪(sharin), 三輪車(sanrinsha)"},

	{"船", "ship", "boat",
		"ony: セン(sen)",
		"kun: ふね(fune)",
		"船舶(senbaku), 漁船(gyosen), 客船(kyakusen)",
		"乗船券(joufasenken), 船出(funade), 海船(kaisen)"},

	{"電", "electricity", "electronics",
		"ony: デン(den)",
		"kun: -",
		"電気(denki), 電子(denshi), 電力(denryoku)",
		"電源(dengen), 電波(denpa), 電卓(dentaku)"},

	{"計", "meter", "measure",
		"ony: ケイ(kei)",
		"kun: はか(haka)る",
		"計画(keikaku), 計算(keisan), 測定(sokutei)",
		"計る(hakaru), 体温計(taionkei), 計時(keiji)"},

	{"工", "construction", "manufacturing",
		"ony: コウ(kou); ク(ku)",
		"kun: -",
		"工学(kougaku), 工場(koujou), 工事(kouji)",
		"建設(kensetsu), 製造(seizou), 工業(kougyou)"},

	{"学", "study", "science",
		"ony: ガク(gaku)",
		"kun: まな(mana)ぶ",
		"学問(gakumon), 学者(gakusha), 学位(gakui)",
		"学ぶ(manabu), 修学(shugaku), 学習(gakushuu)"},

	// Here are some kanji related to automotive mechanics:

	{"車", "car", "vehicle",
		"ony: シャ(sha)",
		"kun: くるま(kuruma)",
		"自動車(jidousha), 車両(sharyou), 車体(shatai)",
		"車輪(sharin), 車軸(shajiku), ハンドル(handoru)"},

	{"修", "repair", "maintain",
		"ony: シュウ(shu)",
		"kun: おさ(osa)める",
		"修理(shuri), 修繕(shusen), 整備(seibi)",
		"直す(naosu), 手入れ(teire), 修学(shugaku)"},

	{"改", "reform", "modify",
		"ony: カイ(kai)",
		"kun: あら(ara)たる",
		"改造(kaizo), 改良(kairyo), 改善(kaizen)",
		"改める(arata meru), 変更(henko), 修正(shusei)"},

	{"整", "adjust", "arrange",
		"ony: セイ(sei)",
		"kun: とと(toto)のえる",
		"整備(seibi), 整理(seiri), 調整(chosei)",
		"整える(totonou), そろえる(sorou), 仕上げる(shiageru)"},

	{"配", "distribute", "arrange",
		"ony: ハイ(hai)",
		"kun: くば(kuba)る",
		"配列(hairetsu), 配線(haisen), 分配(bunpai)",
		"配る(kubaru), 配置(haichi), 割り当て(wariate)"},

	{"器", "apparatus", "implement",
		"ony: キ(ki)",
		"kun: うつわ(utsuwa)",
		"器具(kigu), 器械(kikai), 器物(kibutsu)",
		"道具(dogu), 用具(yougu), 器(utsuwa)"},

	{"具", "ingredient", "equipment",
		"ony: グ(gu)",
		"kun: そな(sona)える",
		"部品(buhin), 付属品(fuzokuhin), 資材(shizai)",
		"装備(soubi), 用具(yougu), 道具(dougu)"},

	// Here are some example kanji related to human physiology:

	{"体", "body", "substance",
		"ony: タイ(tai)",
		"kun: からだ(karada)",
		"身体(shintai), 人体(jitai), 容体(youtai)",
		"体重(taijuu), 体液(taieki), 体系(taikei)"},

	{"血", "blood", "lineage",
		"ony: ケツ(ketsu)",
		"kun: ち(chi)",
		"血液(ketsueki), 出血(shukketsu), 血族(ketsuzoku)",
		"血縁(kechien), 血筋(chisuji), 血友病(keppu byou)"},

	{"骨", "bone", "skeleton",
		"ony: コツ(kotsu)",
		"kun: ほね(hone)",
		"骨折(kossetsu), 骨髄(kotsuzui), 骨格(kokkaku)",
		"骨盤(kotsuban), 骨片(kotsuben), 骨を埋める(hone wo umeru)"},

	{"肉", "meat", "flesh",
		"ony: ニク(niku)",
		"kun: -",
		"肉体(nikutai), 牛肉(gyuuniku), 鶏肉(toriniku)",
		"筋肉(kinniku), 身体(karada), 体(tai)"},

	{"皮", "skin", "hide",
		"ony: ヒ(hi)",
		"kun: かわ(kawa)",
		"皮膚(hifu), 毛皮(kegawa), 皮革(kawa)",
		"皮下(hika), 手の甲(tenohira), 皮肉(hiniku)"},

	{"脳", "brain", "cerebrum",
		"ony: ノウ(nou)",
		"kun: -",
		"脳梗塞(noukouzou), 脳波(nouha), 脳内(nounai)",
		"頭脳(zunou), 知能(chinou), 思考(shikou)"},

	{"心", "heart", "spirit",
		"ony: シン(shin)",
		"kun: こころ(kokoro)",
		"心臓(shinzou), 心情(shinjou), 心境(shinkyou)",
		"情(jou), 気持ち(kimochi), 気分(kibun)"},

	// Here are some kanji related to residential construction:

	{"建", "build", "establish",
		"ony: ケン(ken); コン(kon)",
		"kun: た(ta)てる",
		"建設(kensetsu), 建築(kenchiku), 建物(tatemono)",
		"建てる(tateru), 設立(setsuritsu), 創建(souken)"},

	{"宅", "house", "residence",
		"ony: タク(taku)",
		"kun: -",
		"住宅(jūtaku), 宅地(takuchi), 一戸建(ittokuken)",
		"住居(jūkyo), 居宅(kyotaku), 家(ie)"},

	{"設", "establish", "institute",
		"ony: セツ(setsu); ショ(sho)",
		"kun: もう(mou)ける",
		"設計(sekkei), 設置(sechi), 設備(setsubi)",
		"設ける(mookeru), 創設(sousetsu), 開設(kaisestu)"},

	{"区", "ward", "district",
		"ony: ク(ku)",
		"kun: -",
		"区画(kukaku), 住宅地区(jūtakuchiiki), 区長(kuchou)",
		"区域(kuiki), 分区(bunku), 地区(chiku)"},

	{"造", "make", "construct",
		"ony: ゾウ(zou); ズ(zu)",
		"kun: つく(tsuku)る",
		"建造(kenzou), 造船(zousen), 造園(zouen)",
		"造る(tsukuru), 製造(seizou), 創造(souzou)"},

	{"所", "place", "establishment",
		"ony: ショ(sho); ジョ(jo)",
		"kun: ところ(tokoro)",
		"場所(basho), 所在(shozai), 事務所(jimusho)",
		"処(sho), ところ(tokoro), 辺(hen)"},

	// Here are some kanji related to clothing items:

	{"服", "clothing", "garment",
		"ony: フク(fuku); フカ(fuka)",
		"kun: -",
		"衣服(ifuku), 服装(fukusou), 軍服(gunfuku)",
		"洋服(youfuku), 手芸品(shugeihin), 織物(orimono)"},

	{"帽", "hat", "cap",
		"ony: ボウ(bou)",
		"kun: かぶ(kabu)",
		"帽子(boushi), ハット(hatto), 野球帽(yakyuubou)",
		"被る(kiru), かぶる(kaburu), 冠(kanmuri)"},

	{"靴", "shoes", "footwear",
		"ony: カ(ka)",
		"kun: くつ(kutsu)",
		"靴下(kutsushita), 運動靴(undouka), 長靴(naga-gutsu)",
		"履く(haku), 靴音(kutsuoto), 靴紐(kutsuhimo)"},

	{"袋", "bag", "sack",
		"ony: タイ(tai)",
		"kun: ふくろ(fukuro)",
		"手袋(tebukuro), 旅行かばん(ryokoukaban), ビニール袋(bini-ru fukuro)",
		"入れ物(iremono), かばん(kaban), 袋(fukuro)"},

	{"帯", "belt", "sash",
		"ony: タイ(tai)",
		"kun: お(o)び",
		"腰帯(koshiobi), 帯地(obiji), 黒帯(kuroobi)",
		"締める(shimeru), 巻く(maku), 帯刀(taito)"},

	{"衣", "clothing", "garment",
		"ony: イ(i)",
		"kun: ころも(koromo); きぬ(kinu)",
		"衣料(iryou), 衣服(ifuku), 衣類(irui)",
		"着物(kimono), 一枚(hitomae), 服(fuku)"},

	// Here are some example kanji related to religion and government:

	{"教", "teach", "religion",
		"ony: キョウ(kyou); オシ(oshi)エ(e)る",
		"kun: -",
		"教育(kyouiku), 教会(kyoukai), 教師(kyoushi)",
		"教える(oshieru), 教訓(kyoukun), 宗教(shukyou)"},

	{"仏", "Buddha", "Buddhism",
		"ony: ブツ(butsu); ブ(bu)",
		"kun: ほとけ(hotoke)",
		"仏教(bukkyou), 仏陀(butsuzo), 仏舎利(bussari)",
		"仏(hotoke), 寺院(jiin), 僧侶(souryo)"},

	{"神", "god", "deity",
		"ony: シン(jin); ジン(jin)",
		"kun: かみ(kami)",
		"神社(jinja), 神道(shinto), 神様(kamisama)",
		"八百万の神(yaoyorozu no kami), 祈願(kigan), 祈祷(kito)"},

	{"政", "politics", "government",
		"ony: セイ(sei)",
		"kun: まつ(matsu)る",
		"政治(seiji), 政府(seifu), 政策(seisaku)",
		"行政(gyosei), 治政(chisei), 政党(seito)"},

	{"官", "government", "bureau",
		"ony: カン(kan)",
		"kun: -",
		"官庁(kantyo), 官公庁(kankocho), 官吏(kanri)",
		"役人(yakunin), 公務員(komuin), 官僚(kanryo)"},

	{"法", "law", "rule",
		"ony: ホウ(hou); ハッ(hatsu)",
		"kun: のり(nori)",
		"法律(horitsu), 法令(horei), 法案(hopan)",
		"規則(kisoku), 方法(hoho), 条例(jorei)"},

	// Here are some example kanji for names of famous scientists:

	{"菅", "Kanehiro", "Nobel laureate",
		"ony: カン(Kan)",
		"kun: すが(Suga)",
		"菅野(Kanno), 菅原(Kanbara), 菅谷(Kan'ya)",
		"菅直人(Kan Naoto; former PM), 菅原道真(Sugawara no Michizane), 菅沼宏志(Kan'numa Hiroshi; mathematician)",
	},

	{"湯", "Yukawa", "Nobel laureate",
		"ony: ユ(Yu)",
		"kun: -",
		"湯川(Yukawa), 湯浅(Yuasa), 湯本(Yumoto)",
		"湯川秀樹(Yukawa Hideki; physicist), 湯沢(Yuzawa), 湯木(Yunoki), 湯村(Yumura)"},

	{"朝", "Asahi", "Olympic medalist",
		"ony: チョウ(Chō)",
		"kun: あさ(Asa)",
		"朝日(Asahi), 朝倉(Asakura), 朝比奈(Asahina)",
		"朝日新聞(Asahi Shimbun; newspaper), 朝食(chōshoku; breakfast), 朝青龍明徳(Asashōryū Akinori; sumo wrestler)",
	},

	{"益", "Masatoshi", "Nobel laureate",
		"ony: エキ(Eki)",
		"kun: ます(Masu)",
		"益子(Masuko), 益智(Masuchi), 益田(Masuda)",
		"益々(masu masu; more and more), 有益(yūeki; beneficial), 小柴昌俊(Koshiba Masatoshi; physicist)",
	},

	{"中", "Nakaya", "physicist",
		"ony: チュウ(Chū)",
		"kun: なか(Naka)",
		"中里(Nakazato), 中山(Nakayama), 中村(Nakamura)",
		"中道(chūdō; middle way), 中心(chūshin; center), 中野目慎太郎 (Nakanome Shintarō; physicist)",
	},
}
