package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	gameOn = false
	current_deck = "mast"
	display_List_of_Directives()
	for {
		if gameOn {
			// game_loop_counter++
		} else {
			game_loop_counter = 0
		}
		if game_loop_counter > game_duration {
			game_off()
			fmt.Println("Game has ended")
			// Do game stats etc.
		}
		new_prompt, objective, objective_kind, secondary_objective := pick_aCard_and_assign_fields() // This line is done after each ^^Right!
		begin(new_prompt, objective, objective_kind, secondary_objective)
	} // Endlessly loop: check game status, pick another card, and begin again
}

// The first function that prompts the user
func begin(promptField, objective, objective_kind, secondary_objective string) { // May be a Hira, Kata, or Romaji prompt  - -
	if game_loop_counter > game_duration {
		game_off()
	}
	var in string // A var declaration was needed as a ":=" would not work within the conditional because "in" not in signature
	for {
		// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
		in = promptWithDir(promptField) // Get user's guess

		if in_list_of_Directives(in) {
			if in == "setc" { // respond_to_UserSuppliedDirective(in) will want to return values if "setc" is switched on
				promptField, objective, objective_kind, secondary_objective = respond_to_UserSuppliedDirective(in)
			} else {
				respond_to_UserSuppliedDirective(in)
			}
			continue // ... After "Directive" handling, re-prompt with the same/original promptField
		} else {
			// Passing recursion false [or recall true], means do rightOrOops()
			// ...                                                                  f,f,f Does rightOrOops
			evaluateUsersGuess(in, promptField, objective, objective_kind, false, false, false, secondary_objective)
			// from_last_failed_attempt = true
			break // ... Having finished with all potential guessing, return to main ...
		}
	} // ... Returns to main()'s inner loop; to obtain a new card from which to prompt
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
		// If a recursion has occurred, then do nothing
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

	if from_last_failed_attempt {
		from_last_failed_attempt = false
		in = prompt_interim3(promptField)
	} else {
		// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
		in = promptWithDir(promptField) // Get user's guess
	}

	if in_list_of_Directives(in) {
		if in == "setc" { // See prior comments
			promptField, objective, objective_kind, secondary_objective = respond_to_UserSuppliedDirective(in)
		} else {
			respond_to_UserSuppliedDirective(in)
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
	// Returns to here from all subsequent functions ...
	// and, then returns to begin() and, hense, to main() for the next card ???
}

func rightOrOops(in, promptField, objective string, skipOops bool, secondary_objective string) { // - -

	if in == objective {
		log_right(promptField, in)
		fmt.Printf("%s", colorReset)
		fmt.Printf("%s", in)
		fmt.Printf("%s", colorGreen)
		fmt.Printf(" <--Right!  exactly right! \n")
		fmt.Printf("%s\n%s\n%s\n%s\n%s\n\n", aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("%s", colorReset)
		// Since this was "^^Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_aCard_and_assign_fields()
		// This prompt, deployed by new_objective_kind, takes new_prompt
		in = promptWithDir(new_prompt) // Get user's guess

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective = respond_to_UserSuppliedDirective(in)
			} else {
				respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else if in == secondary_objective {
		log_right(promptField, in)
		fmt.Printf("%s", colorReset)
		fmt.Printf("%s", in)
		fmt.Printf("%s", colorGreen)
		fmt.Printf(" <--Also Right! But as second meaning\n")
		fmt.Printf("%s\n%s\n%s\n%s\n%s\n\n", aCard.Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("%s", colorReset)
		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_aCard_and_assign_fields()
		// This prompt, deployed by new_objective_kind, take new_prompt
		in = promptWithDir(new_prompt) // Get user's guess

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective = respond_to_UserSuppliedDirective(in)
			} else {
				respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else if check_for_match_in_other_fields(in) { // If any of the fields of the card contain a match via our custom parsing algorithm
		log_right(promptField, in)
		fmt.Printf("%s", colorReset)
		fmt.Printf("%s", in)
		fmt.Printf("%s", colorGreen)
		fmt.Printf(" <--somewhat Right! in other fields, or as a substring in objective\n")
		fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n\n", aCard.Meaning, aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("%s", colorReset)
		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_aCard_and_assign_fields()
		// This prompt, deployed by new_objective_kind, takes new_prompt
		in = promptWithDir(new_prompt) // Get user's guess

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective = respond_to_UserSuppliedDirective(in)
			} else {
				respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else { // It is not a Right, and is therefor an Oops, no new aCard has been fetched etc.
		if skipOops {
			// Then do nothing
		} else {
			log_oops(promptField, objective, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^Oops! ")
		}
		tryAgain(promptField, objective, secondary_objective) // passing the old original values
		// Returns here from both tryAgain() and lastTry()
	}
}

func tryAgain(promptField, objective, secondary_objective string) { // - -

	fmt.Printf("Try again \n")
	fmt.Printf("%s", string(colorReset))
	var in string // var declaration needed as a ":=" would not work within the conditional because "in" not in signature
	// **** Now that we are trying again, after a failed guess, prompts do not solicit Directives, which are currently inoperative
	// ... so, these prompts, deployed by objective_kind, take promptField (rather than the new_prompt variant)
	in = prompt_interim(promptField) // Get user's guess

	// ...
	// Note the lack of a Directive handling section which normally follows prompting, Directives are currently inoperative
	//

	if in == objective {
		log_right(promptField, in)
		fmt.Printf("%s", colorReset)
		fmt.Printf("%s", in)
		fmt.Printf("%s", colorGreen)
		fmt.Printf(" <--Right!  exactly right! \n")
		fmt.Printf("%s\n%s\n%s\n%s\n\n", aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_aCard_and_assign_fields()
		fmt.Printf("%s", colorReset)
		// This prompt, deployed by new_objective_kind, takes new_prompt
		in = promptWithDir(new_prompt) // Get user's guess

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective = respond_to_UserSuppliedDirective(in)
			} else {
				respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else if in == secondary_objective {
		log_right(promptField, in)
		fmt.Printf("%s", colorReset)
		fmt.Printf("%s", in)
		fmt.Printf("%s", colorGreen)
		fmt.Printf(" <--Also Right! But as second meaning\n")
		fmt.Printf("%s\n%s\n%s\n%s\n\n", aCard.Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab)
		fmt.Printf("%s", colorReset)
		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_aCard_and_assign_fields()
		// This prompt, deployed by new_objective_kind, take new_prompt
		in = promptWithDir(new_prompt) // Get user's guess

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective = respond_to_UserSuppliedDirective(in)
			} else {
				respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else if check_for_match_in_other_fields(in) { // If any of the fields of the card contain a match via our custom parsing algorithm
		log_right(promptField, in)
		fmt.Printf("%s", colorReset)
		fmt.Printf("%s", in)
		fmt.Printf("%s", colorGreen)
		fmt.Printf(" <--somewhat Right! in other fields, or as a substring in objective\n")
		fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n\n", aCard.Meaning, aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("%s", colorReset)
		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_aCard_and_assign_fields()
		// This prompt, deployed by new_objective_kind, take new_prompt
		in = promptWithDir(new_prompt) // Get user's guess

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective = respond_to_UserSuppliedDirective(in)
			} else {
				respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else {
		log_oops(promptField, objective, in)
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops Again! ")
		lastTry(promptField, objective, secondary_objective)
		// Returns to caller:
	}
}

func lastTry(promptField, objective, secondary_objective string) { // - -
	fmt.Printf("Last Try! \n")
	fmt.Printf("%s", string(colorReset))
	var in string // var declaration needed as a ":=" would not work within the conditional ~ "in" not in signature
	// **** Now that we are trying again, after a failed guess, prompts do not solicit Directives, which are currently inoperative
	// ... so, these prompts, deployed by objective_kind, take promptField (rather than the new_prompt variant)
	in = prompt_interim2(promptField) // Get user's guess

	// ...
	// Note the lack of a Directive handling section which normally follows prompting, Directives are currently inoperative
	//

	if in == objective {
		log_right(promptField, in)
		fmt.Printf("%s", colorReset)
		fmt.Printf("%s", in)
		fmt.Printf("%s", colorGreen)
		fmt.Printf(" <--Right!  exactly right! \n")
		// fmt.Printf("%s", colorReset)
		fmt.Printf("%s\n%s\n%s\n%s\n\n", aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab)
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_aCard_and_assign_fields()
		fmt.Printf("%s", colorReset)
		// This prompt, deployed by new_objective_kind, takes new_prompt
		in = promptWithDir(new_prompt) // Get user's guess

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective = respond_to_UserSuppliedDirective(in)
			} else {
				respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else if in == secondary_objective {
		log_right(promptField, in)
		fmt.Printf("%s", colorReset)
		fmt.Printf("%s", in)
		fmt.Printf("%s", colorGreen)
		fmt.Printf(" <--Also Right! But as second meaning\n")
		fmt.Printf("%s\n%s\n%s\n%s\n\n", aCard.Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab)
		fmt.Printf("%s", colorReset)
		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_aCard_and_assign_fields()
		// This prompt, deployed by new_objective_kind, takes new_prompt
		in = promptWithDir(new_prompt) // Get user's guess

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective = respond_to_UserSuppliedDirective(in)
			} else {
				respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else if check_for_match_in_other_fields(in) { // If any of the fields of the card contain a match via our custom parsing algorithm
		log_right(promptField, in)
		fmt.Printf("%s", colorReset)
		fmt.Printf("%s", in)
		fmt.Printf("%s", colorGreen)
		fmt.Printf(" <--somewhat Right! in other fields, or as a substring in objective\n")
		fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n\n", aCard.Meaning, aCard.Second_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("%s", colorReset)
		// Since this was "^^somewhat Right!", next we obtain new values in-preparation of "returning" to caller
		new_prompt, new_objective, new_objective_kind, new_secondary_objective := pick_aCard_and_assign_fields()
		// This prompt, deployed by new_objective_kind, takes new_prompt
		in = promptWithDir(new_prompt) // Get user's guess

		// Refer to the previous comments re the following mirrored section:
		if in_list_of_Directives(in) {
			if in == "setc" {
				new_prompt, new_objective, new_objective_kind, new_secondary_objective = respond_to_UserSuppliedDirective(in)
			} else {
				respond_to_UserSuppliedDirective(in)
			}
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true, new_secondary_objective)
		} else {
			evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false, new_secondary_objective)
		}
	} else {
		log_oops(aCard.Kanji, aCard.Meaning, in)
		fmt.Printf("%s", colorRed)
		fmt.Printf("\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n\n", aCard.Kanji, aCard.Meaning, aCard.Second_Meaning,
			aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("%s", string(colorReset))
		from_last_failed_attempt = true
	}
	// Returns to caller:
}
