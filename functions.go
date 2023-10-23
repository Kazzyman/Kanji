package main

import (
	"fmt"
	"os"
	"time"
)

// DIRECTIVES : --------------------------------------------------------------------------------------------
//

func testForDirective(in string) (result bool) { // - -
	// if it IS a directive
	if in == "?" || // context-sensitive help on the current card
		in == "sdk" || // Switch Deck
		in == "fif" || // find in files
		in == "rs" || // reset all logs etc
		in == "setc" || // set, force, a new card
		in == "nt" || // notes
		in == "st" || // stats
		in == "dir" || // redisplay the menu of directives etc
		in == "q" || // quit
		in == "frmt" || // format a file
		in == "rm" || // Read the Maps
		in == "bgs" || // Begin a Game Session
		in == "goff" || // Game session Off
		in == "gdc" { // set the Game Duration Counter
		// Then:
		result = true
	}
	return result
}

// sdk Directive
func switch_the_deck() {
	for {
		fmt.Println("\nEnter a deck from below for randomized prompting:\n")

		fmt.Println("    init")
		fmt.Println("    nov")
		fmt.Println("    grad")
		fmt.Println("    mast")
		fmt.Println("    guru \n")

		fmt.Println("Or, a deck from below for sequential prompting:\n")

		fmt.Println("    inits")
		fmt.Println("    novs")
		fmt.Println("    grads")
		fmt.Println("    masts")
		fmt.Printf("    gurus \n\n Here:> ")

		_, _ = fmt.Scan(&current_deck)

		if current_deck != "init" && current_deck != "nov" && current_deck != "grad" && current_deck != "mast" &&
			current_deck != "guru" && current_deck != "inits" && current_deck != "novs" && current_deck != "grads" &&
			current_deck != "masts" && current_deck != "gurus" {
			fmt.Printf("%s", string(colorRed))
			fmt.Printf("\n  \"%s\" is not a valid deck, try again: \n", current_deck)
			fmt.Printf("%s", string(colorReset))
			continue
		} else {
			break
		}
	}
	// Tried to get too fancy here. But it fucked things up. So, just live with the bug you were trying to fix.
	// ... which was just that after a sdk Dir you would get one last prompt from the prior deck, big F'n deal!
	/*
		new_prompt, _, _ := pick_RandomCard_Assign_fields()
		promptWithDir(new_prompt)

	*/
	/*
		new_prompt, _, _ := pick_RandomCard_Assign_fields()
		in = promptWithDir(new_prompt)
		return in

	*/
}

var current_deck string

func respond_to_UserSuppliedDirective(in string) (prompt, objective, kind string) { // - -
	var count int
	switch in {
	case "fif":
		find_in_files()
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
	case "q": // quit
		os.Exit(1)
	case "?": // context-sensitive help on the current card
		fmt.Printf("\n%s\n%s\n%s\n%s\n%s\n%s\n\n", aCard.Kanji, aCard.Meaning, aCard.Long_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab)
	case "st": // stats
		hits()
	case "frmt": // format a file
		formatter()
	case "dir": // reDisplay the DIRECTORY OF DIRECTIVES (and instructions):
		re_display_List_of_Directives()
	case "rm": // Read the Maps
		read_map_of_fineOn()
		read_map_of_needWorkOn()
	case "setc": // set, force, a new card
		prompt, objective, kind = reSet_aCard_andThereBy_reSet_thePromptString()
		if foundElement == nil {
			fmt.Println(" Setting to \"west\" :: ")
			fmt.Printf(string(colorRed))
			// Show the user exactly what is about to be done:
			runeOfCode := `
    silentlyLocateCard("west") // Set the Convenience-global: foundElement
    if foundElement != nil {
        aCard = *foundElement // Set the global var-object 'aCard'
        prompt = aCard.Kanji
        objective = aCard.Meaning
    }
`
			fmt.Println(runeOfCode)
			fmt.Printf(string(colorReset))
			silentlyLocateCard("west") // Set the Convenience-global: foundElement
			if foundElement != nil {
				aCard = *foundElement // Set the global var-object 'aCard'
				prompt = aCard.Kanji
				objective = aCard.Meaning
			} else {
				fmt.Printf("\"west\" Not found : respond_to_UserSuppliedDirective()\n\n")
			}
		}

	case "nt":
		fmt.Println("\nOnyomi (音読み, おにょみ): the reading based on the character's original Chinese pronunciation. \n" +
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
			"coexistence of both Onyomi and Kunyomi readings, contributing to the rich and nuanced nature of the Japanese writing system.\n")
	default:
		// fmt.Println("Directive not found") // Does not work because only existent cases are passed to the switch
	}
	return prompt, objective, kind
}

// Handles the Directive 'setc'
func reSet_aCard_andThereBy_reSet_thePromptString() (prompt, objective, objective_kind string) { //  - -
	var theMeaningOfCardToSilentlyLocate string

	fmt.Printf("\nEnter a Meaning to")
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
		objective = aCard.Meaning
	} else {
		fmt.Println("Error, foundElement is nil in reSet_aCard_andThereBy_reSet_thePromptString()")
	}
	objective_kind = "Romaji"
	fmt.Println("")

	return prompt, objective, objective_kind
}

func game_on() (game string) { // - -
	game = "on"
	gameOn = true
	fmt.Println("The game is on")

	startBeforeCall = time.Now()
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
