package main

import (
	"fmt"
	"math/rand"
	"time"
)

func initialize_stuff() {
	future_len = len(fileOf_Future)
	data_file100_maybe_len = len(data_file100_maybe)

	beauty_len = len(data_beauty)
	words_len = len(fileOf_Words)
	init_len = len(fileOfCardsInitiate)
	guru_len = len(fileOfCardsGuru)
	master_len = len(fileOfCardsMaster)
	novice_len = len(fileOfCardsNovice)
	grad_len = len(fileOfCardsGraduate)
	fresh_len = len(fileOf_fresh)
	current_len = len(fileOf_Current)
	data_len = len(data_file)
	claude_len = len(claude)

	gameOn = false
	// current_deck = "claude" // The default deck, can be changed via the sdk Directive
	// deck_len = claude_len
}

func main() {
	fmt.Println()
	countSLOC()
	initialize_stuff()
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with the current time in nanoseconds.
	display_List_of_Directives()
	switch_the_deck()
	reads := 0
	loopCounter := 0
	// This main for loop is endless, excepting os.Exit(1) via Directive "q".
	for {
		new_prompt, objective, objective_kind, secondary_objective := pick_RandomCard_Assign_fields()
		for reads <= len(already_used_map) { // This line ^ is done after each match // Read entire map before concluding the pick is novel.
			reads++                 // Used to compare the current length of the map to the number of reads done of it
			loopCounter++           // Required to determine if an excessive/exhaustive number of secondary picks have been made and tested
			if loopCounter > 9999 { // 'loopCounter++' being the only way to exit the loop when no novel picks are possible.
				clearMap()
				break // Keep the last pick (could be any card).
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++ // Log the novel pick into the map.
				break                          // Keep the novel card.
			}
			// else, pick another more-fresh card.
			new_prompt, objective, objective_kind, secondary_objective = pick_RandomCard_Assign_fields()
			reads = 0 // continue the loop ensuring that the entire map is read, and compared to, this newly-picked card.
		}
		begin(new_prompt, objective, objective_kind, secondary_objective)
	} // The outer loop of main is always endless and is meant to be exited only via the 'q' Directive.
}

var field_to_prompt_from = "kanji" // This default (and the default of a novice deck) will be effected in event of fat fingers.

var returning_fr_dir bool

func begin(promptField, objective, objective_kind, secondary_objective string) { // May be a Hira, Kata, or Romaji prompt  - -
	returning_fr_dir = false
	if game_loop_counter > game_duration {
		game_off()
	}
	var usersGuess string // A var declaration was needed as a ":=" would not work within the conditional because "usersGuess" not in signature
	for {
		if field_to_prompt_from == "kanji" {
			// These prompts take promptField (rather than the new_prompt variant).
			usersGuess = promptWithDir(promptField)
		} else if field_to_prompt_from == "kunyomi" {
			usersGuess = promptWithDir(aCard.Kunyomi)
		} else {
			usersGuess = promptWithDir(promptField) // A convenient default.
		}

		if in_list_of_Directives(usersGuess) {
			if usersGuess == "setc" { // respond_to_UserSuppliedDirective(usersGuess) will want to return values if "setc" is switched on
				promptField, objective, objective_kind, secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(usersGuess)
			} else {
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(usersGuess)
			}
			continue // ... After "Directive" handling, re-prompt with the same/original promptField
		} else {
			// Passing recursion false [or recall true], means do rightOrOops()
			// ...                                                                  f,f,f Does rightOrOops
			evaluateUsersGuess(usersGuess, promptField, objective, objective_kind, false, false, false, secondary_objective)
			break // ... Having finished with all potential guessing, return to main ...
		}
	} // ... Returns to main()'s inner loop; to (usually randomly) select the next card
}

func evaluateUsersGuess(in, promptField, objective, objective_kind string,
	recursion, recall, skipOops bool, secondary_objective string) {
	if game_loop_counter > game_duration {
		game_off()
	}

	/*
		This next construct is strange! Because; it seems to allow for rightOrOops() to be done "twice" -- but it does not!
		Since rightOrOops() sets in motion a chain of events that never returns directly here where it was called ...
		duplication here never occurs. The construct is needed to allow for rightOrOops() to be omitted if
		evaluateUsersGuess() has been called with recursion true and recall false
	*/
	if recursion {
		// If recursion is true, then do nothing
	} else {
		if gameOn {
			game_loop_counter++
		} else {
			game_loop_counter = 0
		}
		rightOrOops(in, promptField, objective, skipOops, secondary_objective) // This func may call: tryAgain() ... which may call: lastTry()
	}
	if recall {
		if gameOn {
			game_loop_counter++
		} else {
			game_loop_counter = 0
		}
		rightOrOops(in, promptField, objective, skipOops, secondary_objective) // This func may call: tryAgain() ... which may call: lastTry()
	} else {
		// If recall is false, then do nothing
	}
	// ^ ^ ^ If evaluateUsersGuess() has been called after handling a "Directive" then rightOrOops() is omitted entirely
	// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
	// in = promptWithDir(promptField) // Get user's input, from a randomly selected prompt

	if returning_fr_dir {
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = promptWithDir(promptField) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}
		returning_fr_dir = false
	} else {
		// in = prompt_interim3(promptField)
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = prompt_interim3(promptField) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = prompt_interim3(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}
	}
	if in_list_of_Directives(in) {
		if in == "setc" { // See prior comments
			promptField, objective, objective_kind, secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
		} else {
			_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
		}
		/*
			Recursively ...
			after responding to a Directive, prompt via recursion, from the same/original promptField
			With recursion true & recall false, rightOrOops() will NOT be done with this recursion,
			and, with skipOops true [and passed to rightOrOops()] rightOrOops() will not say ^^Oops! after this recursion
			...                                                                        t, f, t Skips rightOrOops
		*/
		// Do a recursive call after handling a Directive
		evaluateUsersGuess(in, promptField, objective, objective_kind, true, false, true, secondary_objective)
	} else {
		/*
			Do a normal, i.e., unconditional recursion with skipOops set to false
			via recall==true & skipOops==false [recursion false or recall true, means do rightOrOops]
		*/
		// Do a recursive call after handling a Directive
		evaluateUsersGuess(in, promptField, objective, objective_kind, true, true, false, secondary_objective)
		// t,t,f Does rightOrOops
	}
	// Returns to here from all subsequent functions ???
	// and, returns to begin() and, hence, to main() to fetch the next fresh card ???
}

// Called by evaluateUsersGuess()
func rightOrOops(in, promptField, objective string, skipOops bool, secondary_objective string) { // - -
	// if in == ず and objective == zu then duplicate the following if block ?
	// fmt.Printf("in is: %s and objective is: %s \n", in, objective)

	// ...
	// Note the lack of a Directive handler to precede checking for a match. In this particular case, such was already done above ...
	// ... just prior to the recursive calls at the end of evaluateUsersGuess().
	//
	if in == objective { // So, the guess was correct
		log_right(promptField, in)
		// promptField ~ aCard.Kanji
		// in ~ users guess
		// objective ~ aCard.Meaning, secondary_objective is second meaning
		// fmt.Printf("this is from A rightOrOops()\n")
		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" ",
			colorReset, in, colorGreen)
		fmt.Printf("or \"%s%s%s\" \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		in = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective, // cont.
			returning_fr_dir = directiveHandler(in, new_prompt, new_objective, new_objective_kind, new_secondary_objective)

	} else if in == secondary_objective {
		// promptField ~ aCard.Kanji
		// in ~ users guess
		// objective ~ aCard.Meaning, secondary_objective is second meaning
		log_right(promptField, in)

		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" or \"%s%s%s\"\n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset, aCard.Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		in = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective, // cont.
			returning_fr_dir = directiveHandler(in, new_prompt, new_objective, new_objective_kind, new_secondary_objective)

		// check_for_match_within_primary_field
	} else if check_for_match_within_primary_field(in) { // If any of the fields of the card contain a match via our custom parsing algorithm.
		log_right(promptField, in)

		fmt.Printf("%s", colorReset)
		fmt.Printf("    ^^%spartly-correct, within the primary meaning: \"%s%s%s\" \n", colorGreen, colorReset, aCard.Meaning, colorGreen)
		fmt.Printf("   %s \n   %s \n%s%s\n", aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" ",
			colorReset, aCard.Meaning, colorGreen)
		fmt.Printf("or \"%s%s%s\" \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		in = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective, // cont.
			returning_fr_dir = directiveHandler(in, new_prompt, new_objective, new_objective_kind, new_secondary_objective)

	} else if check_for_match_in_secondary_field(in) { // If any of the fields of the card contain a match via our custom parsing algorithm.
		log_right(promptField, in)

		fmt.Printf("%s", colorGreen)
		fmt.Printf("     %s^^%spossibly-correct, in another field\n", colorReset, colorGreen)
		fmt.Printf("\"%s%s%s\" was its primary meaning\n\"%s%s%s\" was its secondary meaning\n%s\n%s\n%s\n%s\n\n%s",
			colorReset, aCard.Meaning, colorGreen, colorReset, aCard.Second_Meaning, colorGreen,
			aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)
		fmt.Printf("%sfound:%s \"%s%s%s\" in secondary meaning field\n\n%s",
			colorPurple, colorGreen, colorReset, in, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		in = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective, // cont.
			returning_fr_dir = directiveHandler(in, new_prompt, new_objective, new_objective_kind, new_secondary_objective)

	} else { // it is not a Right, and is therefor an Oops, no new aCard has been fetched etc.
		if skipOops {
			// Then do nothing
		} else {
			fmt.Printf("this is from b rightOrOops()\n") // todo: we have to had looked at the wrong card on subsequent visits to here.
			log_oops(promptField, objective, in)
			fmt.Printf("%s  　^^Oops! ", colorRed)
		}
		tryAgain(promptField, objective, secondary_objective) // passing the old original values
		// Returns here from both tryAgain() and lastTry()
		// Keep value of returning_fr_dir ??????????????????????
	}
}

func tryAgain(promptField, objective, secondary_objective string) { // - -

	fmt.Printf("Try again \n%s", colorReset)
	// fmt.Printf("%s", string(colorReset))
	var in string // var declaration needed as a ":=" would not work within the conditional because "in" not in signature
	// **** Now that we are trying again, after a failed guess, prompts do not solicit Directives:(currently inoperative)
	// ... so, these prompts, deployed by objective_kind, take promptField (rather than the new_prompt variant)
	// in = prompt_interim(promptField) // Get user's guess
	if field_to_prompt_from == "kanji" {
		// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
		in = prompt_interim(promptField) // Get user's input, from a randomly selected prompt
	}
	if field_to_prompt_from == "kunyomi" {
		in = prompt_interim(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
	}
	// ...
	// Note the lack of a Directive handler to precede checking for a match.
	//
	if in == objective {
		log_right(promptField, in)

		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" ",
			colorReset, in, colorGreen)
		fmt.Printf("or \"%s%s%s\" \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		in = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective, // cont.
			returning_fr_dir = directiveHandler(in, new_prompt, new_objective, new_objective_kind, new_secondary_objective)

	} else if in == secondary_objective {
		log_right(promptField, in)

		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" or \"%s%s%s\"\n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset, aCard.Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		in = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective, // cont.
			returning_fr_dir = directiveHandler(in, new_prompt, new_objective, new_objective_kind, new_secondary_objective)

	} else if check_for_match_within_primary_field(in) { // If any of the fields of the card contain a match via our custom parsing algorithm.
		log_right(promptField, in)

		fmt.Printf("%s", colorReset)
		fmt.Printf("    ^^%spartly-correct, within the primary meaning: \"%s%s%s\" \n", colorGreen, colorReset, aCard.Meaning, colorGreen)
		fmt.Printf("   %s \n   %s \n%s%s\n", aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" ",
			colorReset, aCard.Meaning, colorGreen)
		fmt.Printf("or \"%s%s%s\" \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		in = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective, // cont.
			returning_fr_dir = directiveHandler(in, new_prompt, new_objective, new_objective_kind, new_secondary_objective)

	} else if check_for_match_in_secondary_field(in) { // If any of the fields of the card contain a match via our custom parsing algorithm.
		log_right(promptField, in)

		fmt.Printf("%s", colorGreen)
		fmt.Printf("     %s^^%spossibly-correct, in another field\n", colorReset, colorGreen)
		fmt.Printf("\"%s%s%s\" was its primary meaning\n\"%s%s%s\" was its secondary meaning\n%s\n%s\n%s\n%s\n\n%s",
			colorReset, aCard.Meaning, colorGreen, colorReset, aCard.Second_Meaning, colorGreen,
			aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)
		fmt.Printf("%sfound:%s \"%s%s%s\" in secondary meaning field\n\n%s",
			colorPurple, colorGreen, colorReset, in, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		in = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective, // cont.
			returning_fr_dir = directiveHandler(in, new_prompt, new_objective, new_objective_kind, new_secondary_objective)

	} else {
		log_oops(promptField, objective, in)
		fmt.Printf("%s  　^^Oops Again! ", colorRed)
		lastTry(promptField, objective, secondary_objective)
		// Returns to caller:
		// Keep value of returning_fr_dir ??????????????????
	}
}

func lastTry(promptField, objective, secondary_objective string) { // - -
	fmt.Printf("Last Try! \n%s", colorReset)
	var in string // var declaration needed as a ":=" would not work within the conditional ~ "in" not in signature
	// **** Now that we are trying again, after a failed guess, prompts do not solicit Directives:(currently inoperative)
	// ... so, these prompts, deployed by objective_kind, take promptField (rather than the new_prompt variant)
	// in = prompt_interim2(promptField) // Get user's guess
	if field_to_prompt_from == "kanji" {
		// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
		in = prompt_interim2(promptField) // Get user's input, from a randomly selected prompt
	}
	if field_to_prompt_from == "kunyomi" {
		in = prompt_interim2(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
	}
	// ...
	// Note the lack of a Directive handler to precede checking for a match.
	//
	if in == objective {
		log_right(promptField, in)

		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" ",
			colorReset, in, colorGreen)
		fmt.Printf("or \"%s%s%s\" \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		in = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective, // cont.
			returning_fr_dir = directiveHandler(in, new_prompt, new_objective, new_objective_kind, new_secondary_objective)

	} else if in == secondary_objective {
		log_right(promptField, in)

		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" or \"%s%s%s\"\n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset, aCard.Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		in = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective, // cont.
			returning_fr_dir = directiveHandler(in, new_prompt, new_objective, new_objective_kind, new_secondary_objective)

	} else if check_for_match_within_primary_field(in) { // If any of the fields of the card contain a match via our custom parsing algorithm.
		log_right(promptField, in)

		fmt.Printf("%s", colorReset)
		fmt.Printf("    ^^%spartly-correct, within the primary meaning: \"%s%s%s\" \n", colorGreen, colorReset, aCard.Meaning, colorGreen)
		fmt.Printf("   %s \n   %s \n%s%s\n", aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" ",
			colorReset, aCard.Meaning, colorGreen)
		fmt.Printf("or \"%s%s%s\" \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		in = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective, // cont.
			returning_fr_dir = directiveHandler(in, new_prompt, new_objective, new_objective_kind, new_secondary_objective)

	} else if check_for_match_in_secondary_field(in) { // If any of the fields of the card contain a match via our custom parsing algorithm.
		log_right(promptField, in)

		fmt.Printf("%s", colorGreen)
		fmt.Printf("     %s^^%spossibly-correct, in another field\n", colorReset, colorGreen)
		fmt.Printf("\"%s%s%s\" was its primary meaning\n\"%s%s%s\" was its secondary meaning\n%s\n%s\n%s\n%s\n\n%s",
			colorReset, aCard.Meaning, colorGreen, colorReset, aCard.Second_Meaning, colorGreen,
			aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)
		fmt.Printf("%sfound:%s \"%s%s%s\" in secondary meaning field\n\n%s",
			colorPurple, colorGreen, colorReset, in, colorGreen, colorReset)

		// Since that was "correct", next we [conditionally] obtain new values in-preparation of "returning" to caller.
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		// Do we need a different card? freshCard(...) will make it so, if need be.
		new_prompt, new_objective, new_objective_kind, // cont.
			new_secondary_objective = freshCard(new_prompt, new_objective, new_objective_kind, new_secondary_objective, already_used_map)
		in = promptCase1(new_prompt, field_to_prompt_from)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective, // cont.
			returning_fr_dir = directiveHandler(in, new_prompt, new_objective, new_objective_kind, new_secondary_objective)

	} else {
		log_oops(aCard.Kanji, aCard.Meaning, in)
		fmt.Printf("%s\"%s\" is the primaryMeaning of %s\n\"%s\" is the secondaryMeaning of %s\n%s\n%s\n%s\n%s\n\n%s", colorRed,
			aCard.Meaning, aCard.Kanji, aCard.Second_Meaning, aCard.Kanji,
			aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)
	}
	// Returns to caller:
}
