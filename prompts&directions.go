package main

import (
	"fmt"
)

func List_of_Directives() {
	fmt.Println("View source code at https://github.com/Kazzyman/Kanji")
	fmt.Println("    Use Alpha-Numeric (US) input-mode on your system to:")
	fmt.Println("        Enter '" + colorGreen +
		"dir" + colorReset +
		"' redisplay this menu of available Directives")
	fmt.Println("        Enter '" + colorGreen +
		"sdk" + colorReset +
		"' set the Deck")
	fmt.Println("        Enter '" + colorGreen +
		"gdc" + colorReset +
		"' set the Duration Counter for a Game session ")
	fmt.Println("        Enter '" + colorGreen +
		"bgs" + colorReset +
		"' or '" + colorGreen + "goff" + colorReset + "' Begin or end a Game Session ")
	fmt.Println("        Enter '" + colorGreen +
		"?" + colorReset +
		"' context-sensitive help on the current character")
	fmt.Println("        Enter '" + colorGreen +
		"st" + colorReset +
		"' Statistics")
	fmt.Println("        Enter '" + colorGreen +
		"nt" + colorReset +
		"' (notes) an explanation of Onyomi and Kunyomi")
	fmt.Println("        Enter '" + colorGreen +
		"rs" + colorReset +
		"' ReSet (flush or clear) all stats logs etc.")
	fmt.Println("        Enter '" + colorGreen +
		"rm" + colorReset +
		"' Read the current contents of the Maps")
	fmt.Println("        Enter '" + colorGreen +
		"setc" + colorReset +
		"' (set) force the use of a specific card")
	fmt.Println("        Enter '" + colorGreen +
		"dirx" + colorReset +
		"' display extended Directives list")
	fmt.Println("        Enter '" + colorGreen +
		"q" + colorReset +
		"', (quit) terminate the app")
}

// Special prompts for use when soliciting second, or final, guesses
func prompt_interim(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s%s", promptField, colorCyan)
	if current_deck == "all" {
		fmt.Printf(" Meaning? (deck:%s:%s) Help is off, %s", current_deck, current_deck_B, colorReset)
	} else {
		fmt.Printf(" Meaning? (deck:%s) Help is off, %s", current_deck, colorReset)
	}
	fmt.Printf("you must guess! \n%s :> %s", colorCyan, colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return
}
func prompt_interim2(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s%s", promptField, colorCyan)
	if current_deck == "all" {
		fmt.Printf(" Meaning? (deck:%s:%s) Help is off, you must guess, %s", current_deck, current_deck_B, colorReset)
	} else {
		fmt.Printf(" Meaning? (deck:%s) Help is off, you must guess, %s", current_deck, colorReset)
	}
	fmt.Printf("just once more!! \n%s :> %s", colorCyan, colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return
}
func prompt_interim3(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s%s", promptField, colorCyan)
	if current_deck == "all" {
		fmt.Printf(" Meaning? (deck:%s:%s)%s", current_deck, current_deck_B, colorReset)
	} else {
		fmt.Printf(" Meaning? (deck:%s)%s", current_deck, colorReset)
	}
	fmt.Printf(" Meaning? (deck:%s)%s", current_deck, colorReset)
	fmt.Printf(" try any substring from the %s", colorRed)
	fmt.Printf("red%s", colorReset)
	fmt.Printf(" text\n %s:> %s", colorCyan, colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return
}

/*
	init_len = len(fileOfCardsInitiate)
	guru_len = len(fileOfCardsGuru)
	master_len = len(fileOfCardsMaster)
	novice_len = len(fileOfCardsNovice)
	grad_len = len(fileOfCardsGraduate)
	fresh_len = len(fileOf_fresh)
	current_len = len(fileOf_Current)
	data_len = len(data_file)
	claude_len = len(claude)
*/
// Initial prompt, to be used when first introducing a new Kanji char
func promptWithDir(prompt string) (usersGuessOrOptionDirective string) { // - -
	// Create a map to store the frequency of the string for the map
	frequencyMapRightOrOops := make(map[string]int)

	//
	// Parse the cyclic array to check_for_match_in_other_fields the strings and put them into the map:
	//
	// Load the RightOrOops frequency map
	for i := 0; i < len(cyclicArrayHits.RightOrOops); i++ {
		str := cyclicArrayHits.RightOrOops[i]
		// Apparently; this loads a string into, and increments the frequency of, that particular string, in the map
		frequencyMapRightOrOops[str]++ // Specifically, the '++' must increment the int value associated with str
	}

	// -- Update the total_prompts var (its frequency)
	for str, freq := range frequencyMapRightOrOops { // The map has only one entry for Right, & one for Oops
		if str == "Right" { // Finds the one potential entry for Right
			total_prompts = freq
		} else if str == "Oops" { // Finds the one potential entry for Oops
			total_prompts = total_prompts + freq
		} else if str == "" {
			// else, it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		}
	}
	// Determine the unique Chars and their frequencies
	numberOfUniqueKanjiCharsHit := 0
	for _, cardInfoData := range kanjiHitMap {
		if cardInfoData.CorrectGuessCount == 0 {
			// it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		} else {
			numberOfUniqueKanjiCharsHit++
		}
	}
	fmt.Printf("%s%s", prompt, colorCyan)
	if current_deck == "all" {
		fmt.Printf(" Meaning? (deck:%s:%s, cards in deck:%s%d%s,%d,%d), 'dir' or '?' for help with %s",
			current_deck, current_deck_B, colorReset, deck_len, colorCyan, numberOfUniqueKanjiCharsHit, total_prompts, colorReset)
	} else {
		fmt.Printf(" Meaning? (deck:%s, cards in deck:%s%d%s,%d,%d), 'dir' or '?' for help with %s",
			current_deck, colorReset, deck_len, colorCyan, numberOfUniqueKanjiCharsHit, total_prompts, colorReset)
	}
	fmt.Printf("%s \n%s", prompt, colorCyan)
	fmt.Printf(" :> %s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return
}

func extended_list_of_directives() {
	fmt.Println("        Enter '" + colorGreen +
		"fif" + colorReset +
		"' scan category files, compare \"source\" file. Report")
	fmt.Println("        Enter '" + colorGreen +
		"lff" + colorReset +
		"' create two files. A list of kanji. A list of duplicates")
	fmt.Println("        Enter '" + colorGreen +
		"t1" + colorReset +
		"' test1 : don't recall what I was testing???")
	fmt.Println("        Enter '" + colorGreen +
		"t2" + colorReset +
		"' test2 : don't recall what I was testing???")
	fmt.Println("        Enter '" + colorGreen +
		"t3" + colorReset +
		"' proposed third directive ??")
	fmt.Println("        Enter '" + colorGreen +
		"frmt" + colorReset +
		"' Format a text file as card elements")
}

// 'Directive Menu' ; displays only at inception
func display_List_of_Directives() { // (unique)     - -
	frequencyMapRightOrOops := make(map[string]int)
	for i := 0; i < len(cyclicArrayHits.RightOrOops); i++ {
		str := cyclicArrayHits.RightOrOops[i] // Parse the cyclicArray
		// Apparently; this loads a string into, and increments the frequency of, that particular string, in the map
		// ... one such entry per unique string
		frequencyMapRightOrOops[str]++ // Specifically, the '++' must increment the int value associated with str
	}
	// Then, parse the map which was created, and loaded, above
	// As is the way with maps, the frequency map has only one entry for Right, & one for Oops
	for str, freq := range frequencyMapRightOrOops {
		if str == "Right" { // Finds the one potential entry for Right
			total_prompts = freq // Obviously, total_prompts has been declared as a global, elsewhere
		} else if str == "Oops" { // Finds the one potential entry for Oops
			total_prompts = total_prompts + freq
		}
	}
	fmt.Printf("\n\n\n\n\n")
	List_of_Directives()
	//goland:noinspection ALL
	fmt.Println("\n")
	fmt.Printf("Game counter: %d, Game Duration: %d \n", game_loop_counter, game_duration+2)
	fmt.Printf("Current Prompt Count Total: %d \n", total_prompts)
	fmt.Printf("Current Deck is: %s \n", current_deck)
}

// 'Directive Menu' ; displays only in response to "dir" Directive
func re_display_List_of_Directives() { // (unique)     - -
	frequencyMapRightOrOops := make(map[string]int)
	for i := 0; i < len(cyclicArrayHits.RightOrOops); i++ {
		str := cyclicArrayHits.RightOrOops[i] // Parse the cyclicArray
		// Apparently; this loads a string into, and increments the frequency of, that particular string, in the map
		// ... one such entry per unique string
		frequencyMapRightOrOops[str]++ // Specifically, the '++' must increment the int value associated with str
	}
	// Then, parse the map which was created, and loaded, above
	// As is the way with maps, the frequency map has only one entry for Right, & one for Oops
	for str, freq := range frequencyMapRightOrOops {
		if str == "Right" { // Finds the one potential entry for Right
			total_prompts = freq // Obviously, total_prompts has been declared as a global, elsewhere
		} else if str == "Oops" { // Finds the one potential entry for Oops
			total_prompts = total_prompts + freq
		}
	}
	fmt.Printf("\n")
	List_of_Directives()
	//goland:noinspection ALL
	fmt.Printf("Game counter: %d, Game Duration: %d \n", game_loop_counter, game_duration+2)
	fmt.Printf("Current Prompt Count Total: %d \n", total_prompts)
	if current_deck == "all" {
		fmt.Printf("Current Deck is: %s:%s \n", current_deck, current_deck_B)
	} else {
		fmt.Printf("Current Deck is: %s \n", current_deck)
	}
}
