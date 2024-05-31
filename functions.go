package main

import (
	"fmt"
	"os"
	"time"
)

// DIRECTIVES : --------------------------------------------------------------------------------------------
func resetAllLogs() {
	cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
	cyclicArrayHits = CyclicArrayHits{}
	// Also, flush (clear) the maps
	total_prompts = 0
	kanjiHitMap = make(map[string]CardInfo)
	frequencyMapOf_IsFineOnChars = make(map[string]int)
	frequencyMapOf_need_workOn = make(map[string]int)
	already_used_map = make(map[string]int)
	//
	//goland:noinspection ALL
	fmt.Println("\nArrays and maps flushed:\n")
	fmt.Println("    cyclicArrayOfTheJcharsGottenWrong")
	fmt.Println("    cyclicArrayHits")
	fmt.Println("    frequencyMapOf_IsFineOnChars")
	fmt.Println("    frequencyMapOf_need_workOn")
	fmt.Println("    already_used_map")

	fmt.Println("    kanjiHitMap")
	//goland:noinspection ALL
	fmt.Println("    frequencyMapOf_need_workOn\n")
	if weHaveBeenHereBefore {
		fmt.Println("  All Game values have also been reset; answer this last one from the prior deck")
	} else {
		fmt.Println("  Ready to begin.")
	}
	weHaveBeenHereBefore = true
}

func userHasGivenA_DirectiveIsteadOfGuess(usersSubmission string) bool { // - -
	// Is it a directive?
	if usersSubmission == "?" || // context-sensitive help on the current card
		usersSubmission == "t1" || // test one
		usersSubmission == "t2" || // test two
		usersSubmission == "sdk" || // Switch Deck
		usersSubmission == "fif" || // find in files
		usersSubmission == "lff" || // list chars
		usersSubmission == "rs" || // reset all logs etc
		usersSubmission == "setc" || // set, force, a new card
		usersSubmission == "nt" || // notes
		usersSubmission == "st" || // stats
		usersSubmission == "dir" || // redisplay the menu of directives etc
		usersSubmission == "dirx" || // eXtended Directives list
		usersSubmission == "q" || // quit
		usersSubmission == "frmt" || // format a file
		usersSubmission == "rm" || // Read the Maps
		usersSubmission == "bgs" || // Begin a Game Session
		usersSubmission == "goff" || // Game session Off
		usersSubmission == "gdc" { // set the Game Duration Counter
		// Then:
		return true
	}
	return false
}

// sdk Directive
func switch_the_deck() {
	for {
		fmt.Println("\nEnter a deck from below for randomized prompting of a specific deck:\n")

		// fmt.Println("    \"all\" ")
		fmt.Printf("    \"all\"      %d total cards (below)\n\n",
			beauty_len+claude_len+current_len+data_len+fresh_len+grad_len+guru_len+init_len+master_len+novice_len)
		// fmt.Println("    beauty")
		fmt.Printf("    \"beauty\"   %d cards in the Beauty terms deck\n", beauty_len)
		// fmt.Println("    claude")
		fmt.Printf("    \"claude\"   %d cards in the Claude deck\n", claude_len)
		// fmt.Println("    current")
		fmt.Printf("    \"current\"  %d cards in the \"Current\" deck \n", current_len)
		// fmt.Println("    data")
		fmt.Printf("    \"data\"     %d cards in the Data deck\n", data_len)
		// fmt.Println("    fresh")
		fmt.Printf("    \"fresh\"    %d cards in the \"Fresh\" deck\n", fresh_len)
		// fmt.Println("    grad")
		fmt.Printf("    \"grad\"     %d cards in the Graduate deck\n", grad_len)
		// fmt.Println("    guru")
		fmt.Printf("    \"guru\"     %d cards in the Guru deck\n", guru_len)
		// fmt.Println("    init")
		fmt.Printf("    \"init\"     %d cards in the Initiate deck\n", init_len)
		// fmt.Println("    mast")
		fmt.Printf("    \"mast\"     %d cards in the Master deck\n", master_len)
		// fmt.Println("    nov \n")
		fmt.Printf("    \"nov\"      %d cards in the Novice deck\n\n", novice_len)
		fmt.Printf("    \"words\"    %d cards in the Words deck\n\n", words_len)

		fmt.Printf(" Here%s:> %s", colorCyan, colorReset)

		_, _ = fmt.Scan(&current_deck) // current_deck is a global var.

		if current_deck == "q" {
			os.Exit(1)
		}
		if current_deck != "all" &&
			current_deck != "beauty" &&
			current_deck != "claude" &&
			current_deck != "current" &&
			current_deck != "data" &&
			current_deck != "fresh" &&
			current_deck != "grad" &&
			current_deck != "guru" &&
			current_deck != "init" &&
			current_deck != "mast" &&
			current_deck != "nov" &&
			current_deck != "words" {
			fmt.Printf("%s\n  \"%s\" is was not a valid deck, you seem to be a novice at this so nov it will be:) \n%s",
				colorRed, current_deck, colorReset)
			current_deck = "nov"
			deck_len = novice_len
			break // Causes the use of both the novice deck, and also kanji prompting vs kun'yomi prompting.
		} else {
			for {
				fmt.Println("\nEnter a field to prompt from:\n")
				fmt.Println("\nkanji or kunyomi\n")
				_, _ = fmt.Scan(&field_to_prompt_from)
				if field_to_prompt_from != "kanji" && field_to_prompt_from != "kunyomi" {
					fmt.Printf("%s\n  \"%s\" is not a valid field, try again: \n%s", colorRed, field_to_prompt_from, colorReset)
					continue
				} else {
					if current_deck == "all" {
						// deck_len = all_len // This is taken care of in pick_RandomCard_Assign_fields()
					}
					if current_deck == "beauty" {
						deck_len = beauty_len
					}
					if current_deck == "words" {
						deck_len = words_len
					}
					if current_deck == "claude" {
						deck_len = claude_len
					}
					if current_deck == "current" {
						deck_len = current_len
					}
					if current_deck == "data" {
						deck_len = data_len
					}
					if current_deck == "fresh" {
						deck_len = fresh_len
					}
					if current_deck == "grad" {
						deck_len = grad_len
					}
					if current_deck == "guru" {
						deck_len = guru_len
					}
					if current_deck == "init" {
						deck_len = init_len
					}
					if current_deck == "mast" {
						deck_len = master_len
					}
					if current_deck == "nov" {
						deck_len = novice_len
					}
					resetAllLogs()
					return
				}
			}
		}
	}
}

// ::: the only reason I added return arguments in this signature was in consideration of the setc dir. Was that wise/needed???
// todo, well, I tried running this in place of the new respond_to_UserSupplied_Directive() and things were messy. Because the returns in this
// stc handler are only local vars instantiated via the signature of this func.
func respond_to_UserSuppliedDirective(ingOnUsersGuess string) (prompt, primary_meaning, secondary_meaning string, returning_fr_dir bool) { // - -
	var count int
	switch ingOnUsersGuess { // Was "in/usersSubmission" but ingOnUsersGuess is more fun!
	case "t1":
		test1()
	case "t2":
		test2()
	case "fif":
		find_in_files()
	case "lff":
		list_from_files()
	case "sdk":
		// Switch Decks
		switch_the_deck()
	case "gdc": // set the Game Duration Counter
		fmt.Println("Enter a number for how many prompts there will be in the game")
		_, _ = fmt.Scan(&count)
		game_duration = count - 2
	case "bgs": // Begin a Game Session
		// game_loop_counter ++
		game_on()
	case "goff": // Game session Off
		game_off()
	case "rs": // reset all logs etc
		resetAllLogs()
		/*
			// Flush (clear) the old stats and hits arrays
			cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
			cyclicArrayHits = CyclicArrayHits{}
			// Also, flush (clear) the maps
			total_prompts = 0
			//
			//goland:noinspection ALL
			fmt.Println("\nArrays and maps flushed:\n")
			fmt.Println("    cyclicArrayOfTheJcharsGottenWrong")
			fmt.Println("    cyclicArrayHits")
			fmt.Println("    frequencyMapOf_IsFineOnChars")
			//goland:noinspection ALL
			fmt.Println("    frequencyMapOf_need_workOn\n")
			fmt.Println("  And, all Game values have also been reset")

		*/
	case "q": // quit
		os.Exit(1)
	case "?": // context-sensitive help on the current card
		fmt.Printf("\"%s\" is the primaryMeaning of %s\n\"%s\" is the secondaryMeaning of %s\n%s\n%s\n%s\n%s\n\n",
			aCard.Meaning, aCard.Kanji, aCard.Second_Meaning, aCard.Kanji, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
	case "st": // stats
		newHits()
	case "frmt": // format a file
		formatter()
	case "dir": // reDisplay the DIRECTORY OF DIRECTIVES (and instructions):
		countSLOC()
		re_display_List_of_Directives()
	case "dirx":
		extended_list_of_directives()
	case "rm": // Read the Maps
		read_map_of_fineOn()
		read_map_of_needWorkOn()
		// display_already_used_map()
		read_already_used_map()

	case "nt":
		fmt.Println("\nokurigana (hiragana written next to the kanji)\n" +
			"Onyomi (音読み, おにょみ): the reading based on the character's original Chinese pronunciation. \n" +
			"    Onyomi readings are typically used when kanji characters are combined to form compound words, especially in words with \n" +
			"a more formal or academic context.\n" +
			"    The Onyomi reading of a kanji character in Japanese was derived from the character's original Chinese pronunciation. \n" +
			"However, over time, the pronunciation of characters has evolved differently in the two languages; so, there are often \n" +
			"significant differences in how they are now pronounced in the two languages. There might be some similarities due to \n" +
			"their shared origin, but the sounds are usually no longer identical; largely because, the pronunciation in Japanese \n" +
			"has been influenced by the Japanese language's phonetic structure and historical linguistic developments.\n" +
			"    In addition, a kanji character may have a certain pronunciation in Mandarin Chinese, but its Onyomi reading in Japanese \n" +
			"may have undergone changes in consonants, vowels, or tones. And, the meanings associated with the characters \n" +
			"might also have diverged between the two languages.\n" +
			"    So, while there may be historical and etymological connections, the modern pronunciation of Onyomi in Japanese \n" +
			"does not necessarily closely resemble how a Mandarin Chinese speaker would pronounce the corresponding characters.\n\n" +
			"Kunyomi (訓読み, くにょみ): an alternative reading based on the character's native Japanese pronunciation. \n" +
			"    When Chinese characters were introduced, the Japanese adapted them to fit their language by assigning Kunyomi readings \n" +
			"based on existing Japanese words that had similar meanings. Kunyomi readings are commonly used when the kanji character \n" +
			"appears in isolation or is combined with Hiragana characters to form native Japanese words -- allowing for the creation of \n" +
			"new terms and expressions with extra nuance. \n" +
			"    Distinguishing Homophones: As more characters were borrowed from Chinese, it became essential to distinguish between \n" +
			"homophones and provide context-specific meanings. Kunyomi readings helped in achieving this, as they often carried specific \n" +
			"nuances and meanings in the Japanese language.\n" +
			"    Influence of Japanese Phonetics: The Japanese language has a different phonetic structure from Chinese. The Kunyomi \n" +
			"readings reflect this phonetic structure, including differences in syllables, vowel sounds, and pitch accent, making the \n" +
			"pronunciation more natural for native Japanese speakers.\n" +
			"    Development of Native Vocabulary: Over centuries, Japan developed its own vocabulary and linguistic expressions. \n" +
			"Kunyomi readings evolved to accommodate these developments and ensure that Chinese characters were effectively integrated \n" +
			"into the Japanese language.\n" +
			"    In summary, the creation and use of Kunyomi readings in Japanese kanji were motivated by a desire to adapt Chinese \n" +
			"characters to the linguistic, cultural, and semantic needs of the Japanese language. This process allowed for the \n" +
			"coexistence of both Onyomi and Kunyomi readings, contributing to the rich and nuanced nature of the Japanese writing system.")
		fmt.Println()
	default:
		// fmt.Println("Directive not found") // Does not work because only existent cases are passed to the switch
	}
	returning_fr_dir = true
	return prompt, primary_meaning, secondary_meaning, returning_fr_dir
}

func setc_kanji() {
	setcHasBeenrunGlobal = true
	actual_prompt_string, primary_meaning, secondary_meaning = reSet_aCard_andThereBy_reSet_thePromptString()
	// setcHasBeenrunGlobal = false
	// this needs to store our new prompt field in valueFromreSet_aCard_andThereBy_reSet_thePromptString
	// ::: and no it does that. Like this:
	// valueFromreSet_aCard_andThereBy_reSet_thePromptString = aCard.Kanji
	if foundElement == nil {
		fmt.Printf(" Setting to \"west\" :: \n%s", colorRed)
		// Show the user exactly what is about to be done:
		fmt.Println(runeOfCode)
		fmt.Printf(colorReset)
		silentlyLocateCard("west") // Set the Convenience-global: foundElement
		if foundElement != nil {
			aCard = *foundElement // Set the global var-object 'aCard'
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
		} else {
			fmt.Printf("\"west\" Not found : respond_to_UserSuppliedDirective()\n\n")
		}
	}
}

// Handles the Directive 'setc'
func reSet_aCard_andThereBy_reSet_thePromptString() (prompt, primary_meaning, secondary_meaning string) { //  - -
	var theMeaningOfCardToSilentlyLocate string

	fmt.Printf("\nEnter an English Meaning \"using US chars\" to")
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" reSet the prompt via the corresponding Kanji card :> ")
	fmt.Printf("%s", colorReset) //
	_, _ = fmt.Scan(&theMeaningOfCardToSilentlyLocate)

	// Tentatively, prepare to Scan for user's input and attempt locating a matching 'aCard'

	// Confidently, go-looking for user's input: locate matching 'aCard'
	silentlyLocateCard(theMeaningOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
	if foundElement != nil {
		aCard = *foundElement // Set the global var-object 'aCard'
		prompt = aCard.Kanji
		if setcHasBeenrunGlobal {
			valueFromreSet_aCard_andThereBy_reSet_thePromptString = aCard.Kanji
		}
		primary_meaning = aCard.Meaning
		secondary_meaning = aCard.Second_Meaning
	} else {
		fmt.Println("Error, foundElement is nil in reSet_aCard_andThereBy_reSet_thePromptString()")
	}
	// primary_meaning_kind = "Romaji"
	fmt.Println("")

	return prompt, primary_meaning, secondary_meaning
}

func game_on() (game string) { // - -
	game = "on"
	gameOn = true
	fmt.Println("The game is on")

	// startBeforeCall = time.Now()
	currentTime := time.Now()
	TimeOfStartFromTop = time.Now()

	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	_, err1 := fmt.Fprintf(fileHandle,
		"\n The game began at: %s \n",
		currentTime.Format("15:04:05 on Monday 01-02-2006"))
	check_error(err1)
	return game
}
func game_off() (game string) { // - -
	game = "off"
	gameOn = false
	game_duration = 998

	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	currentTime := time.Now()

	// calculate elapsed time
	t_s2 := time.Now()
	elapsed := t_s2.Sub(TimeOfStartFromTop)

	// cast time durations to a String type for Fprintf "formatted print"
	TotalRun := elapsed.String()

	fmt.Printf("\nRun time was: %s, gameOn is: %t \n\n", TotalRun, gameOn)

	// End timer and report elapsed time
	_, err1 := fmt.Fprintf(fileHandle,
		"\n The game ended at: %s  Total prompts was: %d \n",
		currentTime.Format("15:04:05 on Monday 01-02-2006"), total_prompts)
	check_error(err1)
	_, err2 := fmt.Fprintf(fileHandle,
		"\n Elapsed time of game was: %s \n",
		TotalRun)
	check_error(err2)
	return game
}

// Creates a func named check_error which takes one parameter "e" of type error
func check_error(e error) { //    - -
	if e != nil {
		panic(e) // use panic() to display error code
	}
}
