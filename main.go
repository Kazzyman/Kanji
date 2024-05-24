package main

import (
	"fmt"
)

// todo remove debugging Printf lines.
func main() {
	fmt.Println()
	countSLOC()
	initialize_stuff()
	display_List_of_Directives()
	switch_the_deck()
	reads := 0
	loopCounter := 0
	// This outermost for loop (the main loop, pun intended) is endless, excepting os.Exit(1) via User-Directive "q".
	for {
		// Cards (and Decks) may be randomized.
		new_prompt, objective, objective_kind, secondary_objective = pick_RandomCard_Assign_fields()
		for reads <= len(already_used_map) { // Read entire map before concluding the pick is novel.
			reads++                 // Used to compare the current length of the map to the number of reads done of it.
			loopCounter++           // Required to determine if an excessive/exhaustive number of secondary picks have been made and tested.
			if loopCounter > 9999 { // 'loopCounter++' being the only way to exit the loop when no novel picks are possible.
				clearMap()
				break // Keep the last pick (could be any randomized card).
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++ // Log the novel pick into the map.
				break                          // Keep the novel card.
			}
			// via the implied else, pick another more-fresh card.
			new_prompt, objective, objective_kind, secondary_objective = pick_RandomCard_Assign_fields()
			reads = 0 // Continue the loop ensuring that the entire map is read, and compared to, the newly-picked card.
		}
		// begin(new_prompt, objective, objective_kind, secondary_objective) // new_prompt is renamed to promptField via the signature.
		begin(new_prompt) // new_prompt is renamed to promptField via the signature.
		// userInput, promptField
		// fmt.Printf("here in main, under the call to begin \n")
	} // This outer-loop of main is always endless and can only be exited via a 'q' Directive from the user.
	// fmt.Printf("execution has landed here 15\n")

}

/*
.
.
*/
// todo remove debugging Printf lines.
func rightOrOops(userInput, promptField string) { // - -
	// ...
	// Note absence of Directive Handler 'DH' preceding future guess evaluations. In this particular case, DH was already done above ...
	// ... just prior to the recursive calls at the end of evaluateUsersGuess(...).
	//
	if userInput == objective { // So, the guess was correct
		log_right(promptField, userInput)
		// promptField ~ aCard.Kanji
		// in ~ users guess
		// objective ~ aCard.Meaning, secondary_objective is second meaning
		// fmt.Printf("this is from A rightOrOops()\n")
		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" ",
			colorReset, userInput, colorGreen)
		fmt.Printf("or \"%s%s%s\" \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset)

		/*
			var new_prompt string
			var objective string
			var objective_kind string
			var secondary_objective string
			var new_objective string
			var new_objective_kind, new_secondary_objective string
			var field_to_prompt_from = "kanji" // The default of kanji vs kun'yomi (and use of a novice deck) will be effected in a fat-fingers event.
			var returning_fr_dir bool
		*/
		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)

		userInput = promptCase1(new_prompt, field_to_prompt_from)

		new_prompt, returning_fr_dir = directiveHandler(userInput, new_prompt)

		// fmt.Printf("execution has landed here 1, promptField is %s, new_prompt is %s \n", promptField, new_prompt)

	} else if userInput == secondary_objective {
		// promptField ~ aCard.Kanji
		// in ~ users guess
		// objective ~ aCard.Meaning, secondary_objective is second meaning
		log_right(promptField, userInput)

		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" or \"%s%s%s\"\n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset, aCard.Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		userInput = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, returning_fr_dir = directiveHandler(userInput, new_prompt)
		// fmt.Printf("execution has landed here 2\n")

		// Check for a match within the primary field
	} else if check_for_match_within_primary_field(userInput) { // If any of the fields of the card contain a match via our custom parsing algorithm.
		log_right(promptField, userInput)

		fmt.Printf("%s", colorReset)
		fmt.Printf("    ^^%spartly-correct, within the primary meaning: \"%s%s%s\" \n", colorGreen, colorReset, aCard.Meaning, colorGreen)
		fmt.Printf("   %s \n   %s \n%s%s\n", aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" ",
			colorReset, aCard.Meaning, colorGreen)
		fmt.Printf("or \"%s%s%s\" \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		userInput = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, returning_fr_dir = directiveHandler(userInput, new_prompt)
		// fmt.Printf("execution has landed here 3, promptField is %s, new_prompt is %s \n", promptField, new_prompt)

	} else if check_for_match_in_secondary_field(userInput) { // If any of the fields of the card contain a match via our custom parsing algorithm.
		log_right(promptField, userInput)

		fmt.Printf("%s", colorGreen)
		fmt.Printf("     %s^^%spossibly-correct, in another field\n", colorReset, colorGreen)
		fmt.Printf("\"%s%s%s\" was its primary meaning\n\"%s%s%s\" was its secondary meaning\n%s\n%s\n%s\n%s\n\n%s",
			colorReset, aCard.Meaning, colorGreen, colorReset, aCard.Second_Meaning, colorGreen,
			aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)
		fmt.Printf("%sfound:%s \"%s%s%s\" in secondary meaning field\n\n%s",
			colorPurple, colorGreen, colorReset, userInput, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		userInput = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, returning_fr_dir = directiveHandler(userInput, new_prompt)
		// fmt.Printf("execution has landed here 4\n")

	} else { // it is not a Right, and is therefor an Oops, no new aCard has been fetched etc.

		log_oops(promptField, objective, userInput)
		fmt.Printf("%s  　^^Oops! ", colorRed)
		tryAgain(userInput, promptField) // Try again (with same card) passing the old original values.
	}
	// fmt.Printf("execution has landed here 14, promptField is %s, new_prompt is %s \n", promptField, new_prompt)
	// ::: and (in the case of setc) the next logged line is evaluateUsersGuess 3
} // end of rightOrOops(...)   returning via tail of evaluateUsersGuess(...) --> the tail of begin(...)  and finally to main().
// ::: Not True!!  ^ ^ ^ ^ ^ instead, rightOrOops tends to return EXACTLY to where it was called, obviously!
/*
.
.
*/
// todo remove debugging Printf lines.
func tryAgain(userInput, promptField string) { // - -

	fmt.Printf("Try again \n%s", colorReset)
	// fmt.Printf("%s", string(colorReset))
	// var userInput string // var declaration needed as a ":=" would not work within the conditional because "userInput" not in signature
	// **** Now that we are trying again, after a failed guess, prompts do not solicit Directives:(currently inoperative)
	// ... so, these prompts, deployed by objective_kind, take promptField (rather than the new_prompt variant)
	// userInput = prompt_interim(promptField) // Get user's guess
	if field_to_prompt_from == "kanji" {
		// fmt.Printf("careful here, new_prompt is %s \n", new_prompt)
		// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
		if returning_fr_dir || directiveHandlerHasBeenRun {
			userInput = prompt_interim(new_prompt) // Get user's input, from a randomly selected prompt
		} else {
			userInput = prompt_interim(promptField) // Get user's input, from a randomly selected prompt
			// fmt.Printf("execution has landed here 5\n")
		}
	}
	if field_to_prompt_from == "kunyomi" {
		userInput = prompt_interim(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
	}
	// ...
	// Note the lack of a Directive handler to precede checking for a match.
	//
	if userInput == objective {
		log_right(promptField, userInput)

		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" ",
			colorReset, userInput, colorGreen)
		fmt.Printf("or \"%s%s%s\" \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		userInput = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, returning_fr_dir = directiveHandler(userInput, new_prompt)
		// fmt.Printf("execution has landed here 6\n")

	} else if userInput == secondary_objective {
		log_right(promptField, userInput)

		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" or \"%s%s%s\"\n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset, aCard.Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		userInput = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, returning_fr_dir = directiveHandler(userInput, new_prompt)
		// fmt.Printf("execution has landed here 7\n")

	} else if check_for_match_within_primary_field(userInput) { // If any of the fields of the card contain a match via our custom parsing algorithm.
		log_right(promptField, userInput)

		fmt.Printf("%s", colorReset)
		fmt.Printf("    ^^%spartly-correct, within the primary meaning: \"%s%s%s\" \n", colorGreen, colorReset, aCard.Meaning, colorGreen)
		fmt.Printf("   %s \n   %s \n%s%s\n", aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" ",
			colorReset, aCard.Meaning, colorGreen)
		fmt.Printf("or \"%s%s%s\" \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		userInput = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, returning_fr_dir = directiveHandler(userInput, new_prompt)
		// fmt.Printf("execution has landed here 8\n")

	} else if check_for_match_in_secondary_field(userInput) { // If any of the fields of the card contain a match via our custom parsing algorithm.
		log_right(promptField, userInput)

		fmt.Printf("%s", colorGreen)
		fmt.Printf("     %s^^%spossibly-correct, in another field\n", colorReset, colorGreen)
		fmt.Printf("\"%s%s%s\" was its primary meaning\n\"%s%s%s\" was its secondary meaning\n%s\n%s\n%s\n%s\n\n%s",
			colorReset, aCard.Meaning, colorGreen, colorReset, aCard.Second_Meaning, colorGreen,
			aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)
		fmt.Printf("%sfound:%s \"%s%s%s\" in secondary meaning field\n\n%s",
			colorPurple, colorGreen, colorReset, userInput, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		userInput = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, returning_fr_dir = directiveHandler(userInput, new_prompt)
		// fmt.Printf("execution has landed here 9\n")

	} else {
		log_oops(promptField, objective, userInput)
		fmt.Printf("%s  　^^Oops Again! ", colorRed)
		lastTry(userInput, promptField) // Try yet again, again with the same card.
	}
	// fmt.Printf("execution has landed here 13\n")

} // end of tryAgain(...)  returning via tails of rightOrOops(...) -> evaluateUsersGuess(...) -> begin(...)  and finally to main().
// ::: True, because the only place tryAgain() was called WAS from rightOrOops().
/*
.
.
*/
// todo remove debugging Printf lines.
func lastTry(userInput, promptField string) { // - -
	fmt.Printf("Last Try! \n%s", colorReset)
	// var userInput string // var declaration needed as a ":=" would not work within the conditional ~ "userInput" not in signature
	// **** Now that we are trying again, after a failed guess, prompts do not solicit Directives:(currently inoperative)
	// ... so, these prompts, deployed by objective_kind, take promptField (rather than the new_prompt variant)
	// userInput = prompt_interim2(promptField) // Get user's guess
	if field_to_prompt_from == "kanji" {
		// fmt.Printf("careful here, new_prompt is %s \n", new_prompt)
		// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
		if returning_fr_dir || directiveHandlerHasBeenRun {
			userInput = prompt_interim(new_prompt) // Get user's input, from a randomly selected prompt
		} else {
			userInput = prompt_interim(promptField) // Get user's input, from a randomly selected prompt
			// fmt.Printf("execution has landed here 5\n")
		}
	}
	if field_to_prompt_from == "kunyomi" {
		userInput = prompt_interim2(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
	}
	// ...
	// Note the lack of a Directive handler to precede checking for a match.
	//
	if userInput == objective {
		log_right(promptField, userInput)

		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" ",
			colorReset, userInput, colorGreen)
		fmt.Printf("or \"%s%s%s\" \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		userInput = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, returning_fr_dir = directiveHandler(userInput, new_prompt)

	} else if userInput == secondary_objective {
		log_right(promptField, userInput)

		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" or \"%s%s%s\"\n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset, aCard.Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		userInput = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, returning_fr_dir = directiveHandler(userInput, new_prompt)
		// fmt.Printf("execution has landed here 11\n")

	} else if check_for_match_within_primary_field(userInput) { // If any of the fields of the card contain a match via our custom parsing algorithm.
		log_right(promptField, userInput)

		fmt.Printf("%s", colorReset)
		fmt.Printf("    ^^%spartly-correct, within the primary meaning: \"%s%s%s\" \n", colorGreen, colorReset, aCard.Meaning, colorGreen)
		fmt.Printf("   %s \n   %s \n%s%s\n", aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" ",
			colorReset, aCard.Meaning, colorGreen)
		fmt.Printf("or \"%s%s%s\" \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		userInput = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, returning_fr_dir = directiveHandler(userInput, new_prompt)

	} else if check_for_match_in_secondary_field(userInput) { // If any of the fields of the card contain a match via our custom parsing algorithm.
		log_right(promptField, userInput)

		fmt.Printf("%s", colorGreen)
		fmt.Printf("     %s^^%spossibly-correct, in another field\n", colorReset, colorGreen)
		fmt.Printf("\"%s%s%s\" was its primary meaning\n\"%s%s%s\" was its secondary meaning\n%s\n%s\n%s\n%s\n\n%s",
			colorReset, aCard.Meaning, colorGreen, colorReset, aCard.Second_Meaning, colorGreen,
			aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)
		fmt.Printf("%sfound:%s \"%s%s%s\" in secondary meaning field\n\n%s",
			colorPurple, colorGreen, colorReset, userInput, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		userInput = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, returning_fr_dir = directiveHandler(userInput, new_prompt)

	} else {
		log_oops(aCard.Kanji, aCard.Meaning, userInput)
		fmt.Printf("%s\"%s\" is the primaryMeaning of %s\n\"%s\" is the secondaryMeaning of %s\n%s\n%s\n%s\n%s\n\n%s", colorRed,
			aCard.Meaning, aCard.Kanji, aCard.Second_Meaning, aCard.Kanji,
			aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)
		// Note that we do not try again with the same card and instead ... we bounce our way directly back to main() for a fresh one.
	}
	// fmt.Printf("execution has landed here 12\n")

} // end lastTry(...)  returning via tails of tryAgain(...) -> rightOrOops(...) -> evaluateUsersGuess(...) -> begin(...)  to main().
// ::: True, because the only place lastTry() was called WAS from tryAgain().
