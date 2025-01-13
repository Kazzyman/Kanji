package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var words []string

func prompt_the_user_for_input() { // ::: - -
	if guessLevelCounter == 1 { // ::: ----------- 1 1 1 1 1 1 --------------------
		guessLevelCounter++

		/* 1 19
		1 in Jap2 we did our display-Right lines here, but they are done in the caller here.
		*/
		promptWithDir() // ::: Introduction of a new card.
		gottenHonestly = true

		// Obtain users input.
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		// Split the input into words
		words = strings.Fields(input)

		// fmt.Printf("prompt1 len of guess is %d\n", len(words))

		// Limit to 5 words
		if len(words) > 5 {
			words = words[:5]
		}

		// Join the words back into a single string
		usersSubmission = strings.Join(words, " ")

	} else if guessLevelCounter == 2 { // ::: ---------- 2 2 2 2 2 2 2 2  ---------------------
		guessLevelCounter++

		/* 2 18
		2 in Jap2 we did our display-Right lines here, but they are done by the caller in this Kanji practice app.
		*/
		prompt_interim() // ::: Says: guess again!
		gottenHonestly = true

		// Obtain users input.
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		// Split the input into words
		words = strings.Fields(input)

		// fmt.Printf("prompt2 len of guess is %d\n", len(words))

		// Limit to 5 words
		if len(words) > 5 {
			words = words[:5]
		}

		// Join the words back into a single string
		usersSubmission = strings.Join(words, " ")

	} else if guessLevelCounter == 3 { // ::: --------------- 3 3 3 3 3 3 3 3 3 3 ----------------------------
		guessLevelCounter++

		/* 3 17
		3 in Jap2 we did our display-Right lines here, but they are done in the caller here.
		*/
		prompt_interim2() // ::: Says: Last guess!!
		gottenHonestly = true

		// Obtain users input.
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		// Split the input into words
		words = strings.Fields(input)

		// fmt.Printf("prompt3 len of guess is %d\n", len(words))

		// Limit to 5 words
		if len(words) > 5 {
			words = words[:5]
		}

		// Join the words back into a single string
		usersSubmission = strings.Join(words, " ")

	} else if guessLevelCounter == 4 { // ::: -------------------------- 4 4 4 4 4 4 4 4 4 4 4 4  --------------------

		// ::: That was your last try looser. Here's a clue, just for you: ...
		weHadFailed_And_OnlyGotThisRightBecauseOfTheClue = true
		display_failure_of_final_guess_message_etc(usersSubmission)

		gottenHonestly = false

		log_oops(actual_prompt_string, primary_meaning, usersSubmission)
		guessLevelCounter = 1
		gotLastCardRightSoGetFreshOne = false // Re-prompt with the same card as the one the user totally-bombed on.
		prompt_interim3()                     // Message: " try any substring from the RED text."
	} else if guessLevelCounter > 4 || guessLevelCounter <= -1 {
		fmt.Printf("%s\n - The value of guessLevelCounter is out of range, it is %d - %s\n", colorRed, guessLevelCounter, colorReset)
	}
}

/*
. 4
. 5
*/
func prompt_interim() { //  - -
	if field_to_prompt_from == "yomi" {
		fmt.Printf("O:%s, K:%s", aCard.Onyomi, aCard.Kunyomi)
	} else {
		fmt.Printf("%s%s", colorYellow, actual_prompt_string)
	}
	if current_deck == "all" {
		fmt.Printf("%s Meaning? (deck:%s:%s) ", colorCyan, current_deck, current_deck_B) // todo ?
	} else if current_deck == "quiz_novice" || current_deck == "quiz_comp" || current_deck == "quiz_mast" || current_deck == "quiz_guru" || current_deck == "quiz_japan" {
		fmt.Printf("%s Answer? (deck:%s) ", colorCyan, current_deck)
	} else {
		fmt.Printf("%s Meaning? (deck:%s) ", colorCyan, current_deck)
	}
	fmt.Printf("%sguess again!", colorReset) // ::: guess again!
	if aGameIsRunning {
		fmt.Printf(", %s%s is playing: %s1st:%s%d%s, 2nd:%s%d%s, 3rd:%s%d%s, fails:%s%d, %s%d/%d\n",
			colorCyan, nameOfPlayer, colorRed, colorReset, correctOnFirstAttemptAccumulator,
			colorRed, colorReset, correctOnSecondAttemptAccumulator, colorRed, colorReset, correctOnThirdAttemptAccumulator,
			colorRed, colorReset, failedOnThirdAttemptAccumulator, colorCyan, game_loop_counter, game_duration_set_by_user)
	} else {
		fmt.Println() // This is correct! Trust me.
	}
	if aGameIsRunning {
		fmt.Printf("%s :game> %s", colorYellow, colorReset) // Always reset to white after :> so that users input displays consistent white.
	} else {
		fmt.Printf("%s :> %s", colorYellow, colorReset) // Always reset to white after :> so that users input displays consistent white.
	}
}

/* 5 16

 */

func prompt_interim2() { //  - -
	if field_to_prompt_from == "yomi" {
		fmt.Printf("O:%s, K:%s", aCard.Onyomi, aCard.Kunyomi)
	} else {
		fmt.Printf("%s%s", colorRed, actual_prompt_string)
	}
	if current_deck == "all" {
		fmt.Printf("%s Meaning? (deck:%s:%s) %sLast guess!!, ", colorCyan, current_deck, current_deck_B, colorRed)
	} else if current_deck == "quiz_novice" || current_deck == "quiz_comp" || current_deck == "quiz_mast" || current_deck == "quiz_guru" || current_deck == "quiz_japan" {
		fmt.Printf("%s Answer? (deck:%s) %sLast guess!!", colorCyan, current_deck, colorRed)
	} else {
		fmt.Printf("%s Meaning? (deck:%s) %sLast guess!!", colorCyan, current_deck, colorRed)
	}
	if aGameIsRunning {
		fmt.Printf(", %s%s is playing: %s1st:%s%d%s, 2nd:%s%d%s, 3rd:%s%d%s, fails:%s%d, %s%d/%d\n",
			colorCyan, nameOfPlayer, colorRed, colorReset, correctOnFirstAttemptAccumulator,
			colorRed, colorReset, correctOnSecondAttemptAccumulator, colorRed, colorReset, correctOnThirdAttemptAccumulator,
			colorRed, colorReset, failedOnThirdAttemptAccumulator, colorCyan, game_loop_counter, game_duration_set_by_user)
	} else {
		fmt.Println()
	}
	if aGameIsRunning {
		fmt.Printf("%s :game> %s", colorRed, colorReset) // Always reset to white after :> so that users input displays consistent white.
	} else {
		fmt.Printf("%s :> %s", colorRed, colorReset) // Always reset to white after :> so that users input displays consistent white.
	}
}

// To be used after the final Oops message.
func prompt_interim3() { //  - -
	fmt.Printf("%s try any substring from the ", colorReset) // ::: ------------- try any substring from the red text ------
	fmt.Printf("%sred", colorRed)
	fmt.Printf("%s text\n%s", colorReset, colorCyan)
}

/* 6 15
. 6
. 7

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
		if field_to_prompt_from == "yomi" {
			fmt.Printf("O:%s, K:%s%s", aCard.Onyomi, aCard.Kunyomi, colorCyan)
		} else {
			fmt.Printf("%s%s", actual_prompt_string, colorCyan)
		}
		directiveHandlerHasBeenRun = false
	} else {
		if field_to_prompt_from == "yomi" {
			fmt.Printf("O:%s, K:%s%s", aCard.Onyomi, aCard.Kunyomi, colorCyan)
		} else {
			fmt.Printf("%s%s%s", colorGreen, actual_prompt_string, colorCyan)
		}
	}

	if current_deck == "all" {
		fmt.Printf(" Meaning? (deck:%s:%s, cards in deck:%s%d%s,%d,%d), \n'dir' or '?' for help with %s",
			current_deck, current_deck_B, colorReset, deck_len, colorCyan, numberOfUniqueKanjiCharsHit, total_prompts, colorReset)
	} else if current_deck == "quiz_novice" || current_deck == "quiz_comp" || current_deck == "quiz_mast" || current_deck == "quiz_guru" || current_deck == "quiz_japan" {
		fmt.Printf(" Answer? (deck:%s,len:%s%d%s; #unique:%s%d%s, #ofPrompts:%s%d%s), \n'dir' or '?' for help with %s",
			current_deck, colorReset, deck_len, colorCyan, colorReset, numberOfUniqueKanjiCharsHit, colorCyan, colorReset, total_prompts, colorCyan, colorReset)
	} else {
		fmt.Printf(" Meaning? (deck:%s,len:%s%d%s; #unique:%s%d%s, #ofPrompts:%s%d%s), \n'dir' or '?' for help with %s",
			current_deck, colorReset, deck_len, colorCyan, colorReset, numberOfUniqueKanjiCharsHit, colorCyan, colorReset, total_prompts, colorCyan, colorReset)
	}
	if field_to_prompt_from == "yomi" {
		fmt.Printf("O:%s, K:%s%s", aCard.Onyomi, aCard.Kunyomi, colorCyan)
	} else {
		fmt.Printf("%s%s", actual_prompt_string, colorCyan)
	}
	// fmt.Printf("%s%s", actual_prompt_string, colorCyan)
	if aGameIsRunning {
		fmt.Printf(", %s%s is playing: %s1st:%s%d%s, 2nd:%s%d%s, 3rd:%s%d%s, fails:%s%d, %s%d/%d%s\n",
			colorCyan, nameOfPlayer, colorRed, colorReset, correctOnFirstAttemptAccumulator,
			colorRed, colorReset, correctOnSecondAttemptAccumulator, colorRed, colorReset, correctOnThirdAttemptAccumulator,
			colorRed, colorReset, failedOnThirdAttemptAccumulator, colorCyan, game_loop_counter, game_duration_set_by_user, colorReset)
	} else {
		fmt.Println()
	}
	fmt.Printf("%s :> %s", colorGreen, colorReset)

	// Obtain users input.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Split the input into words
	words = strings.Fields(input)

	// Limit to 5 words
	if len(words) > 5 {
		words = words[:5]
	}

	// Join the words back into a single string
	usersSubmission = strings.Join(words, " ")

}

/* 7 14
; 8
. 9
*/
// Initial prompt, to be used when first introducing a new Kanji char here.
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
		if field_to_prompt_from == "yomi" {
			fmt.Printf("O:%s, K:%s%s", aCard.Onyomi, aCard.Kunyomi, colorCyan)
		} else {
			fmt.Printf("%s%s", actual_prompt_string, colorCyan)
		}
		directiveHandlerHasBeenRun = false
	} else {
		if field_to_prompt_from == "yomi" {
			fmt.Printf("O:%s, K:%s%s", aCard.Onyomi, aCard.Kunyomi, colorCyan)
		} else {
			fmt.Printf("%s%s%s", colorGreen, actual_prompt_string, colorCyan) // prior to "Meaning?"
		}
	}

	if current_deck == "all" {
		fmt.Printf(" Meaning? (deck:%s:%s, cards in deck:%s%d%s,%d,%d), \n'dir' or '?' for help with %s",
			current_deck, current_deck_B, colorReset, deck_len, colorCyan, numberOfUniqueKanjiCharsHit, total_prompts, colorReset)
	} else if current_deck == "quiz_novice" || current_deck == "quiz_comp" || current_deck == "quiz_mast" || current_deck == "quiz_guru" || current_deck == "quiz_japan" {
		fmt.Printf(" Answer? (deck:%s,len:%s%d%s; #unique:%s%d%s, #ofPrompts:%s%d%s), \n'dir' or '?' for help with %s",
			current_deck, colorReset, deck_len, colorCyan, colorReset, numberOfUniqueKanjiCharsHit, colorCyan, colorReset, total_prompts, colorCyan, colorReset)
	} else {
		fmt.Printf(" Meaning? (deck:%s,len:%s%d%s; #unique:%s%d%s, #ofPrompts:%s%d%s), \n'dir' or '?' for help with %s",
			current_deck, colorReset, deck_len, colorCyan, colorReset, numberOfUniqueKanjiCharsHit, colorCyan, colorReset, total_prompts, colorCyan, colorReset)
	}
	if field_to_prompt_from == "yomi" {
		fmt.Printf("O:%s, K:%s%s", aCard.Onyomi, aCard.Kunyomi, colorCyan)
	} else {
		fmt.Printf("%s%s%s", colorGreen, actual_prompt_string, colorCyan) // on the "for help with" line
	}
	if aGameIsRunning {
		fmt.Printf(", %s%s is playing: %s1st:%s%d%s, 2nd:%s%d%s, 3rd:%s%d%s, fails:%s%d, %s%d/%d%s\n",
			colorCyan, nameOfPlayer, colorRed, colorReset, correctOnFirstAttemptAccumulator,
			colorRed, colorReset, correctOnSecondAttemptAccumulator, colorRed, colorReset, correctOnThirdAttemptAccumulator,
			colorRed, colorReset, failedOnThirdAttemptAccumulator, colorCyan, game_loop_counter, game_duration_set_by_user, colorReset)
	} else {
		fmt.Println()
	}
	if aGameIsRunning {
		fmt.Printf("%s :game> %s", colorGreen, colorReset)
	} else {
		fmt.Printf("%s :> %s", colorGreen, colorReset)
	}
}

/*
. 10
. 11
*/
func display_failure_of_final_guess_message_etc(userInput string) { // ::: - -
	log_oops_andUpdateGame(aCard.Kanji, aCard.Kunyomi, userInput)
	fmt.Printf("%s", colorReset)
	fmt.Printf("     That was your last try looser. Here's a clue, just for you: ...\n %s", colorRed)
	fmt.Printf("\n%s, %s\n%s\n%s\n\n%s", aCard.Meaning, aCard.Second_Meaning, aCard.Vocab, aCard.Vocab2, colorReset)

	fileHandle, err := os.OpenFile("Kanji-newLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)
	_, err1 := fmt.Fprintf(fileHandle,
		"\n%s had a REAL ISSUE with==%s:%s:%s", nameOfPlayer, aCard.Meaning, aCard.Second_Meaning, aCard.Kanji)
	check_error(err1)
}

func log_oops_andUpdateGame(prompt_it_was, field_it_was, guess string) { // - -
	if aGameIsRunning {
		failedOnThirdAttemptAccumulator++ // ::: update game.
	}
	logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(prompt_it_was)
	logHits_in_cyclicArrayHits("Oops", prompt_it_was)
	logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(prompt_it_was +
		":it was:" + field_it_was + ":but you had guessed:" + guess)
}

/* 9 13
- 12
*/

func display_limited_gaming_dir_list() {
	fmt.Println("        Enter '" + colorGreen +
		"dirg" + colorReset +
		"' DirGame, Display this list")
	fmt.Println("        Enter '" + colorGreen +
		"off" + colorReset +
		"' End this game early")
	fmt.Println("        Enter '" + colorGreen +
		"q" + colorReset +
		"', (quit) terminate the app")
}
