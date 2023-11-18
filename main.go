package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	gameOn = false
	current_deck = "grad" // The default deck, can be changed via the sdk Directive
	display_List_of_Directives()
	// Create and prime vars, kluge?
	new_prompt, objective, objective_kind, secondary_objective := "prime", "prime", "prime", "prime"
	reads := 0
	loopCounter := 0
	// fmt.Printf("reads of map is %d\n", len(already_used_map))
	// main loop (endless)
	for {
		new_prompt, objective, objective_kind, secondary_objective = pick_RandomCard_Assign_fields() // This line is done after each ^^Right!
		for reads <= len(already_used_map) {                                                         // Read the entire map before concluding that the pick is novel
			reads++       // Used to compare the current length of the map to the number of reads done of it
			loopCounter++ // Required to determine if an excessive/exhaustive number of secondary picks have been made and tested
			// 'loopCounter++' being the only way to exit the loop when no novel picks are possible
			if loopCounter > 9999 {
				// fmt.Printf("loopCounter:%d, so doing a clear of the map\n", loopCounter)
				clearMap()
				break // Keep the last pick (could be any card)
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++ // Log the novel pick into the map
				break                          // Keep the novel card
			}
			// else, pick another card
			new_prompt, objective, objective_kind, secondary_objective = pick_RandomCard_Assign_fields()
			reads = 0 // continue the loop ensuring that the entire map is read with this new pick
		}
		begin(new_prompt, objective, objective_kind, secondary_objective)
	}
	// The outer loop of main is always endless and is meant to be exited only via the 'q' Directive
}

var returning_fr_dir bool
var field_to_prompt_from = "kanji" // The default

// The first function that prompts the user for a guess, with a Kanji
func begin(promptField, objective, objective_kind, secondary_objective string) { // May be a Hira, Kata, or Romaji prompt  - -
	returning_fr_dir = false
	if game_loop_counter > game_duration {
		game_off()
	}
	var in string // A var declaration was needed as a ":=" would not work within the conditional because "in" not in signature
	for {
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = promptWithDir(promptField) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}

		if in_list_of_Directives(in) {
			if in == "setc" { // respond_to_UserSuppliedDirective(in) will want to return values if "setc" is switched on
				promptField, objective, objective_kind, secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			} else {
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			}
			continue // ... After "Directive" handling, re-prompt with the same/original promptField
		} else {
			// Passing recursion false [or recall true], means do rightOrOops()
			// ...                                                                  f,f,f Does rightOrOops
			evaluateUsersGuess(in, promptField, objective, objective_kind, false, false, false, secondary_objective)
			break // ... Having finished with all potential guessing, return to main ...
		}
	} // ... Returns to main()'s inner loop; to (usually randomly) select the next card
}

func evaluateUsersGuess(in, promptField, objective, objective_kind string, recursion, recall, skipOops bool, secondary_objective string) {
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
	// and, returns to begin() and, hense, to main() for the next card ???
	// returning_fr_dir = false // Redundant ?????????????
}

// Called by evaluateUsersGuess()
func rightOrOops(in, promptField, objective string, skipOops bool, secondary_objective string) { // - -

	if in == objective {
		log_right(promptField, in)
		// promptField ~ aCard.Kanji
		// in ~ users guess
		// objective ~ aCard.Meaning, secondary_objective is second meaning

		fmt.Printf("%s%s%s", colorReset, in, colorGreen)
		fmt.Printf(" <--Right!\n")
		fmt.Printf("%s\n%s\n%s\n%s\n%s\n\n%s", aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)

		// Since this was "^^Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		length := 0
		loopCounter := 0
		// fmt.Printf("length of map is %d\n", len(already_used_map))
		for length <= len(already_used_map) { // See comments in main()
			length++
			loopCounter++
			if loopCounter > 9999 {
				// fmt.Printf("loopCounter:%d, so doing a clear of the map\n", loopCounter)
				clearMap()
				break
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++
				break // keep it
			}
			// else, pick another
			new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
			length = 0 // continue the loop ensuring that the entire map is read with this new pick
		}
		// This prompt, deployed by new_objective_kind, takes new_prompt
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = promptWithDir(new_prompt) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			} else {
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else if in == secondary_objective {
		log_right(promptField, in)

		fmt.Printf("%s%s%s", colorReset, in, colorGreen)
		fmt.Printf(" <--Also Right! But as second meaning\n")
		fmt.Printf("%s\n%s\n%s\n%s\n%s\n\n%s", aCard.Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)

		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		length := 0
		loopCounter := 0
		// fmt.Printf("length of map is %d\n", len(already_used_map))
		for length <= len(already_used_map) {
			length++
			loopCounter++
			if loopCounter > 9999 {
				// fmt.Printf("loopCounter:%d, so doing a clear of the map\n", loopCounter)
				clearMap()
				break
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++
				break // keep it
			}
			// else, pick another
			new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
			length = 0 // continue the loop ensuring that the entire map is read with this new pick
		}
		// This prompt, deployed by new_objective_kind, take new_prompt
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = promptWithDir(new_prompt) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			} else {
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
		// check_for_match_within_primary_field
	} else if check_for_match_within_primary_field(in) { // If any of the fields of the card contain a match via our custom parsing algorithm
		log_right(promptField, in)

		fmt.Printf("%s%s%s", colorReset, in, colorGreen)
		fmt.Printf(" <--somewhat Right! within primary field\n")
		fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n\n%s",
			aCard.Meaning, aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)

		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		length := 0
		loopCounter := 0
		// fmt.Printf("length of map is %d\n", len(already_used_map))
		for length <= len(already_used_map) {
			length++
			loopCounter++
			if loopCounter > 9999 {
				// fmt.Printf("loopCounter:%d, so doing a clear of the map\n", loopCounter)
				clearMap()
				break
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++
				break // keep it
			}
			// else, pick another
			new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
			length = 0 // continue the loop ensuring that the entire map is read with this new pick
		}

		// This prompt, deployed by new_objective_kind, takes new_prompt
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = promptWithDir(new_prompt) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			} else {
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else if check_for_match_in_other_fields(in) { // If any of the fields of the card contain a match via our custom parsing algorithm
		log_right(promptField, in)

		fmt.Printf("%s%s%s", colorReset, in, colorGreen)
		fmt.Printf(" <--somewhat Right! in other fields\n")
		fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n\n%s",
			aCard.Meaning, aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)

		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		length := 0
		loopCounter := 0
		// fmt.Printf("length of map is %d\n", len(already_used_map))
		for length <= len(already_used_map) {
			length++
			loopCounter++
			if loopCounter > 9999 {
				// fmt.Printf("loopCounter:%d, so doing a clear of the map\n", loopCounter)
				clearMap()
				break
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++
				break // keep it
			}
			// else, pick another
			new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
			length = 0 // continue the loop ensuring that the entire map is read with this new pick
		}

		// This prompt, deployed by new_objective_kind, takes new_prompt
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = promptWithDir(new_prompt) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			} else {
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else { // it is not a Right, and is therefor an Oops, no new aCard has been fetched etc.
		if skipOops {
			// Then do nothing
		} else {
			log_oops(promptField, objective, in)
			// fmt.Printf("%s", colorRed)
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
	// Note the lack of a Directive handling section which normally follows prompting, Directives are currently inoperative
	//

	if in == objective {
		log_right(promptField, in)
		fmt.Printf("%s%s%s", colorReset, in, colorGreen)
		fmt.Printf(" <--Right!\n")
		fmt.Printf("%s\n%s\n%s\n%s\n\n%s", aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, colorReset)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		length := 0
		loopCounter := 0
		// fmt.Printf("length of map is %d\n", len(already_used_map))
		for length <= len(already_used_map) {
			length++
			loopCounter++
			if loopCounter > 9999 {
				// fmt.Printf("loopCounter:%d, so doing a clear of the map\n", loopCounter)
				clearMap()
				break
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++
				break // keep it
			}
			// else, pick another
			new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
			length = 0 // continue the loop ensuring that the entire map is read with this new pick
		}

		// This prompt, deployed by new_objective_kind, takes new_prompt
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = promptWithDir(new_prompt) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			} else {
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
		// insert here?
	} else if in == secondary_objective {
		log_right(promptField, in)

		fmt.Printf("%s%s%s", colorReset, in, colorGreen)
		fmt.Printf(" <--Also Right! But as second meaning\n")
		fmt.Printf("%s\n%s\n%s\n%s\n\n%s", aCard.Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, colorReset)

		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		length := 0
		loopCounter := 0
		// fmt.Printf("length of map is %d\n", len(already_used_map))
		for length <= len(already_used_map) {
			length++
			loopCounter++
			if loopCounter > 9999 {
				// fmt.Printf("loopCounter:%d, so doing a clear of the map\n", loopCounter)
				clearMap()
				break
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++
				break // keep it
			}
			// else, pick another
			new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
			length = 0 // continue the loop ensuring that the entire map is read with this new pick
		}

		// This prompt, deployed by new_objective_kind, take new_prompt
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = promptWithDir(new_prompt) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			} else {
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else if check_for_match_within_primary_field(in) { // If any of the fields of the card contain a match via our custom parsing algorithm
		log_right(promptField, in)

		fmt.Printf("%s%s%s", colorReset, in, colorGreen)
		fmt.Printf(" <--somewhat Right! within primary field\n")
		fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n\n%s",
			aCard.Meaning, aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)

		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		length := 0
		loopCounter := 0
		// fmt.Printf("length of map is %d\n", len(already_used_map))
		for length <= len(already_used_map) {
			length++
			loopCounter++
			if loopCounter > 9999 {
				// fmt.Printf("loopCounter:%d, so doing a clear of the map\n", loopCounter)
				clearMap()
				break
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++
				break // keep it
			}
			// else, pick another
			new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
			length = 0 // continue the loop ensuring that the entire map is read with this new pick
		}

		// This prompt, deployed by new_objective_kind, takes new_prompt
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = promptWithDir(new_prompt) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			} else {
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else if check_for_match_in_other_fields(in) { // If any of the fields of the card contain a match via our custom parsing algorithm
		log_right(promptField, in)

		fmt.Printf("%s%s%s", colorReset, in, colorGreen)
		fmt.Printf(" <--somewhat Right! in other fields\n")
		fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n\n%s",
			aCard.Meaning, aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)

		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		length := 0
		loopCounter := 0
		// fmt.Printf("length of map is %d\n", len(already_used_map))
		for length <= len(already_used_map) {
			length++
			loopCounter++
			if loopCounter > 9999 {
				// fmt.Printf("loopCounter:%d, so doing a clear of the map\n", loopCounter)
				clearMap()
				break
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++
				break // keep it
			}
			// else, pick another
			new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
			length = 0 // continue the loop ensuring that the entire map is read with this new pick
		}

		// This prompt, deployed by new_objective_kind, take new_prompt
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = promptWithDir(new_prompt) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			} else {
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
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
	// Note the lack of a Directive handling section which normally follows prompting, Directives are currently inoperative
	//

	if in == objective {
		log_right(promptField, in)
		fmt.Printf("%s%s%s", colorReset, in, colorGreen)
		fmt.Printf(" <--Right!\n")
		fmt.Printf("%s\n%s\n%s\n%s\n\n%s", aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, colorReset)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		length := 0
		loopCounter := 0
		// fmt.Printf("length of map is %d\n", len(already_used_map))
		for length <= len(already_used_map) {
			length++
			loopCounter++
			if loopCounter > 9999 {
				// fmt.Printf("loopCounter:%d, so doing a clear of the map\n", loopCounter)
				clearMap()
				break
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++
				break // keep it
			}
			// else, pick another
			new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
			length = 0 // continue the loop ensuring that the entire map is read with this new pick
		}

		// This prompt, deployed by new_objective_kind, takes new_prompt
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = promptWithDir(new_prompt) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			} else {
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else if in == secondary_objective {
		log_right(promptField, in)
		fmt.Printf("%s%s%s", colorReset, in, colorGreen)
		fmt.Printf(" <--Also Right! But as second meaning\n")
		fmt.Printf("%s\n%s\n%s\n%s\n\n%s", aCard.Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, colorReset)
		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		length := 0
		loopCounter := 0
		// fmt.Printf("length of map is %d\n", len(already_used_map))
		for length <= len(already_used_map) {
			length++
			loopCounter++
			if loopCounter > 9999 {
				// fmt.Printf("loopCounter:%d, so doing a clear of the map\n", loopCounter)
				clearMap()
				break
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++
				break // keep it
			}
			// else, pick another
			new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
			length = 0 // continue the loop ensuring that the entire map is read with this new pick
		}

		// This prompt, deployed by new_objective_kind, takes new_prompt
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = promptWithDir(new_prompt) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			} else {
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else if check_for_match_within_primary_field(in) { // If any of the fields of the card contain a match via our custom parsing algorithm
		log_right(promptField, in)

		fmt.Printf("%s%s%s", colorReset, in, colorGreen)
		fmt.Printf(" <--somewhat Right! within primary field\n")
		fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n\n%s",
			aCard.Meaning, aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)

		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		length := 0
		loopCounter := 0
		// fmt.Printf("length of map is %d\n", len(already_used_map))
		for length <= len(already_used_map) {
			length++
			loopCounter++
			if loopCounter > 9999 {
				// fmt.Printf("loopCounter:%d, so doing a clear of the map\n", loopCounter)
				clearMap()
				break
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++
				break // keep it
			}
			// else, pick another
			new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
			length = 0 // continue the loop ensuring that the entire map is read with this new pick
		}

		// This prompt, deployed by new_objective_kind, takes new_prompt
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = promptWithDir(new_prompt) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			} else {
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else if check_for_match_in_other_fields(in) { // If any of the fields of the card contain a match via our custom parsing algorithm
		log_right(promptField, in)
		fmt.Printf("%s%s%s", colorReset, in, colorGreen)
		fmt.Printf(" <--somewhat Right! in other fields\n")
		fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n\n%s",
			aCard.Meaning, aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)
		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_RandomCard_Assign_fields()
		length := 0
		loopCounter := 0
		// fmt.Printf("length of map is %d\n", len(already_used_map))
		for length <= len(already_used_map) {
			length++
			loopCounter++
			if loopCounter > 9999 {
				// fmt.Printf("loopCounter:%d, so doing a clear of the map\n", loopCounter)
				clearMap()
				break
			}
			if is_pick_novel(new_prompt) {
				already_used_map[new_prompt]++
				break // keep it
			}
			// else, pick another
			new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
			length = 0 // continue the loop ensuring that the entire map is read with this new pick
		}

		// This prompt, deployed by new_objective_kind, takes new_prompt
		if field_to_prompt_from == "kanji" {
			// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
			in = promptWithDir(new_prompt) // Get user's input, from a randomly selected prompt
		}
		if field_to_prompt_from == "kunyomi" {
			in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
		}

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			} else {
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else {
		log_oops(aCard.Kanji, aCard.Meaning, in)
		// fmt.Printf("      　^^That was your last try, Oops! ")
		fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n\n%s", colorRed, aCard.Kanji, aCard.Meaning, aCard.Second_Meaning,
			aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)
	}
	// Returns to caller:
	// Keep value of returning_fr_dir ?????????????????
}
