package main

import (
	"fmt"
	"os"
)

func prompt_the_user_for_input() { // ::: - -
	if guessLevelCounter == 1 { // ::: --------- 1 1 1 1 1 1 ------------------
		guessLevelCounter++

		/*
			in Jap2 we did our display-Right lines here, but they are done in the caller here.
		*/
		promptWithDir()
		gottenHonestly = true

	} else if guessLevelCounter == 2 { // ::: ------------------------------------------ 2 2 2 2 2 2 2 2 2 2 2 -------------------
		guessLevelCounter++

		/*
			in Jap2 we did our display-Right lines here, but they are done in the caller here.
		*/
		prompt_interim() // 3a, 3b
		gottenHonestly = true

		// Obtain users input.
		_, _ = fmt.Scan(&usersSubmission)

	} else if guessLevelCounter == 3 { // ::: -------------------3 3 3 3 3 3 3 3 3 3 --------------------------------------
		guessLevelCounter++

		/*
			in Jap2 we did our display-Right lines here, but they are done in the caller here.
		*/
		prompt_interim2() // 4a, 4b
		// gottenHonestly = false // todo ??

		// Obtain users input.
		_, _ = fmt.Scan(&usersSubmission)

	} else if guessLevelCounter > 3 { // ::: --------------------------  > 3  > 3  > 3  > 3  > 3  > 3  ---- i.e. 4  --------------------

		display_failure_of_final_guess_message_etc(usersSubmission)
		fmt.Printf("%s  　^^Oops! ", colorRed)

		log_oops(actual_prompt_string, primary_meaning, usersSubmission)
		weHadFailed_And_OnlyGotThisRightBecauseOfTheClue = true
		guessLevelCounter = 1 // ::: was 1, why 1 ??, prob because it is set to 1 twice in main
		begin_kanji_practice()

	} else if guessLevelCounter >= 4 || guessLevelCounter <= -1 {
		fmt.Printf("The value of guessLevelCounter is out of range, it is %d \n", guessLevelCounter)
	} else {
	}
}

/*
.
.
*/
func prompt_interim() { //  - -
	fmt.Printf("%s%s", actual_prompt_string, colorCyan)
	if current_deck == "all" {
		fmt.Printf(" Meaning? (deck:%s:%s) Help is off, %s", current_deck, current_deck_B, colorReset)
	} else {
		fmt.Printf(" Meaning? (deck:%s) Help is off, %s", current_deck, colorReset)
	}
	fmt.Printf("you must guess! \n%s :> %s", colorCyan, colorReset) // ::: ------- "you must guess!" ---------
}
func prompt_interim2() { //  - -
	fmt.Printf("%s%s", actual_prompt_string, colorCyan)
	if current_deck == "all" {
		fmt.Printf(" Meaning? (deck:%s:%s) Help is off, you must guess, %s", current_deck, current_deck_B, colorReset)
	} else {
		fmt.Printf(" Meaning? (deck:%s) Help is off, you must guess, %s", current_deck, colorReset)
	}
	fmt.Printf("just once more!! \n%s :> %s", colorCyan, colorReset) // ::: ------ "just once more!!" --------
}
func prompt_interim3() { //  - -
	fmt.Printf("%s%s", actual_prompt_string, colorCyan)
	if current_deck == "all" {
		fmt.Printf(" Meaning? (deck:%s:%s)%s", current_deck, current_deck_B, colorReset)
	} else {
		fmt.Printf(" Meaning? (deck:%s)%s", current_deck, colorReset)
	}
	fmt.Printf(" Meaning? (deck:%s)%s", current_deck, colorReset)
	fmt.Printf(" try any substring from the %s", colorRed) // ::: ------------- try any substring from the red text ------
	fmt.Printf("red%s", colorReset)
	fmt.Printf(" text\n %s:> %s", colorCyan, colorReset)
	// _, _ = fmt.Scan(&usersSubmission)
	// return
}

/*
.
.

*/

// Initial prompt, to be used when first introducing a new Kanji char at inception.
func promptWithDirAtInception() { // - -

	// Create a map to store the frequency of the string for the map
	frequencyMapRightOrOops := make(map[string]int)

	//
	// Parse the cyclic array to check_for_match_in_secondary_field the strings and put them into the map:
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

	// ::: Determine if a setc has been run and we therefore need to prompt from the global value of the newly forced/set card.
	if setcHasBeenrunGlobal {
		fmt.Printf("%s%s", valueFromreSet_aCard_andThereBy_reSet_thePromptString, colorCyan)
		setcHasBeenrunGlobal = false
	} else if directiveHandlerHasBeenRun {
		// ::: latest,last: fmt.Printf("%s%s : oldVar directiveHandlerHasBeenRun alt: new_kanjiChar is %s", prompt, colorCyan, new_kanjiChar)
		fmt.Printf("%s%s", actual_prompt_string, colorCyan)
		directiveHandlerHasBeenRun = false
	} else {
		fmt.Printf("%s%s", actual_prompt_string, colorCyan)
	}

	if current_deck == "all" {
		fmt.Printf(" Meaning?pwd1 (deck:%s:%s, cards in deck:%s%d%s,%d,%d), \n'dir' or '?' for help with %s",
			current_deck, current_deck_B, colorReset, deck_len, colorCyan, numberOfUniqueKanjiCharsHit, total_prompts, colorReset)
	} else {
		fmt.Printf(" Meaning? (deck:%s,len:%s%d%s; #unique:%s%d%s, #ofPrompts:%s%d%s), \n'dir' or '?' for help with %s",
			current_deck, colorReset, deck_len, colorCyan, colorReset, numberOfUniqueKanjiCharsHit, colorCyan, colorReset, total_prompts, colorCyan, colorReset)
	}
	fmt.Printf("%s \n%s", actual_prompt_string, colorCyan)
	fmt.Printf(" here:> %s", colorReset)
	_, _ = fmt.Scan(&usersSubmission)
	// return
}

/*
;

*/
// Initial prompt, to be used when first introducing a new Kanji char at inception.
func promptWithDir() { // - -

	// Create a map to store the frequency of the string for the map
	frequencyMapRightOrOops := make(map[string]int)

	//
	// Parse the cyclic array to check_for_match_in_secondary_field the strings and put them into the map:
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

	// ::: Determine if a setc has been run and we therefore need to prompt from the global value of the newly forced/set card.
	if setcHasBeenrunGlobal {
		fmt.Printf("%s%s", valueFromreSet_aCard_andThereBy_reSet_thePromptString, colorCyan)
		setcHasBeenrunGlobal = false
	} else if directiveHandlerHasBeenRun {
		// ::: latest,last: fmt.Printf("%s%s : oldVar directiveHandlerHasBeenRun alt: new_kanjiChar is %s", prompt, colorCyan, new_kanjiChar)
		fmt.Printf("%s%s", actual_prompt_string, colorCyan)
		directiveHandlerHasBeenRun = false
	} else {
		fmt.Printf("%s%s", actual_prompt_string, colorCyan)
	}

	if current_deck == "all" {
		fmt.Printf(" Meaning?pwd1 (deck:%s:%s, cards in deck:%s%d%s,%d,%d), \n'dir' or '?' for help with %s",
			current_deck, current_deck_B, colorReset, deck_len, colorCyan, numberOfUniqueKanjiCharsHit, total_prompts, colorReset)
	} else {
		fmt.Printf(" Meaning? (deck:%s,len:%s%d%s; #unique:%s%d%s, #ofPrompts:%s%d%s), \n'dir' or '?' for help with %s",
			current_deck, colorReset, deck_len, colorCyan, colorReset, numberOfUniqueKanjiCharsHit, colorCyan, colorReset, total_prompts, colorCyan, colorReset)
	}
	fmt.Printf("%s \n%s", actual_prompt_string, colorCyan)
	fmt.Printf(" here:> %s", colorReset)
	_, _ = fmt.Scan(&usersSubmission)
	// return
}

/*
.
.
.
*/
func display_failure_of_final_guess_message_etc(userInput string) { // ::: - -
	log_oops_andUpdateGame(aCard.Kanji, aCard.Kunyomi, userInput)
	fmt.Printf("%s", colorRed)
	fmt.Printf("     ^^Oops! That was your last try looser. Here's a clue, just for you: ...\n %s", colorReset)
	fmt.Printf("\n%s\n%s\n\n", aCard.Vocab, aCard.Vocab2)

	fileHandle, err := os.OpenFile("KanjiLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)
	_, err1 := fmt.Fprintf(fileHandle,
		"\nUser had a REAL ISSUE with==%s:%s:%s", aCard.Meaning, aCard.Second_Meaning, aCard.Kanji)
	check_error(err1)
}
func log_oops_andUpdateGame(prompt_it_was, field_it_was, guess string) { // - -
	if theGameIsRunning {
		failedOnThirdAttemptAccumulator++
	}
	logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(prompt_it_was)
	logHits_in_cyclicArrayHits("Oops", prompt_it_was)
	logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(prompt_it_was +
		":it was:" + field_it_was + ":but you had guessed:" + guess)
}

/*
-
-
-
*/
func display_limited_gaming_dir_list() {
	fmt.Println("        Enter '" + colorGreen +
		"dirg" + colorReset +
		"' DirGame, Display this list")
	fmt.Println("        Enter '" + colorGreen +
		"off" + colorReset +
		"' End this game early")
	fmt.Println("        Enter '" + colorGreen +
		"stc" + colorReset +
		"' (Set-Card) force the use of a specific card (Hira input)")
	fmt.Println("        Enter '" + colorGreen +
		"stcr" + colorReset +
		"' (Set-Card) force the use of a specific card (Roma input)")
	fmt.Println("        Enter '" + colorGreen +
		"q" + colorReset +
		"', (quit) terminate the app")
}
