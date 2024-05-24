package main

import (
	"fmt"
	"math/rand"
	"time"
)

// todo remove debugging Printf lines.
// Little-more than a way to get to the func following the next func? Albeit it does implement consecutive directives feature at startup.
func begin(promptField string) { // Prompting may occur via Hira, Kata, or Romaji chars.  - -
	// returning_fr_dir = false // ::: oddly, the only problem with removing every assignment to false is that setc sporadically fails to update the prompt ???
	if game_loop_counter > game_duration {
		game_off()
	}
	var userInput string // A var is needed because ":=" does not work throughout these conditionals, i.e., "userInput " is not in the signature.
	// var promptField string
	for {
		if field_to_prompt_from == "kanji" {
			userInput = promptWithDir(promptField) // Note: this prompt takes promptField (rather than the new_prompt variant).
			// fmt.Printf("here in beginA\n")
		} else {
			userInput = promptWithDir(aCard.Kunyomi)
		}

		// Handle the potential of a Directive command from user. ::: this allows for consecutive directives, at the start. but only there.
		if userHasGivenA_DirectiveIsteadOfGuess(userInput) { // Checks to see if user's response is a Directive.
			if userInput == "setc" { // Only capture all the values returned by respond_to_UserSuppliedDirective(userInput ) if "setc" was done.
				// fmt.Printf("here in begin1\n") // todo <---------------------
				promptField, objective, objective_kind, secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(userInput)
				// ::: WE NEVER!!!!!!!!!!! EVER!!!!!!!!!! LAND HERE after a setc ????????????? unless it was done at startup.
				// fmt.Printf("execution has landed here 18\n") // todo <---------------------

			} else { // else handle it, but only capture the flag; a flag that is not needed here, but is usually needed in other cases.
				_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(userInput) // Returns a flag, and re-prompts (above).
			}
			continue // ^ ^ ^ After "Directive" handling, re-prompt (via the top of this for loop) with the same/original promptField value.
		} else { // A non-Directive was given, so assume it to be an attempted guess.
			// Passing recursion=false [or recall=true], to evaluateUsersGuess() causes the subsequent execution of rightOrOops() [via recursion]
			// near the top of evaluateUsersGuess()  ^   v   ^   v   ^   v   ^   v   ^   v   ^   v
			// ...   ^  v  ^  v  ^  v  ^  v  ^  v  ^  v  ^  v  ^  v   ^   v   ^   v   ^    false,   v   false, causes ONE execution of rightOrOops().
			evaluateUsersGuess(userInput, promptField, false, false)
			break // Exits this func, then ... perhaps much later -- perhaps having exhausted all potential guessing -- return to main() ...
		}
	}
	fmt.Printf("execution has landed here 17\n") // todo <---------------------

} // Returns to main()'s outer loop to (usually randomly) select a fresh card.
/*
.
.
*/
// Little-more than a very messy way to get to the next func!!!
// Or, maybe it is useful for dir handling ???
func evaluateUsersGuess(userInput, promptField string, recursion, recall bool) {
	if game_loop_counter > game_duration {
		game_off()
	}
	// fmt.Printf("here in evaluateUsersGuess 77 \n") // todo <---------------------

	/*
	   .
	   These next constructs may look strange! Because; they may seem to allow for rightOrOops() to be done "twice" -- Oh contraire!
	   Note that rightOrOops() sets-in-motion a chain of events that eventually returns to where it was called EXCEPT after lastTry().
	   These complex constructs are needed to allow for rightOrOops() to be entirely-omitted if evaluateUsersGuess() has been
	   recursively called with recursion=true and recall=false. Thus allowing for consecutive execution of various Directives ::: is this true?.
	   .
	   Essentially; whenever the first of these two conditions is true, and the second false, NO rightOrOops(...) is executed!! Thereby
	   allowing an endlessly-recursive loop (merely) in order to enable the processing of CONSECUTIVE Directive commands from the user.
	   It is (and was), admittedly, a lot of fuss for such a small and seldom-used feature; but the heart wants what the heart wants.
	*/
	if recursion {
		// If recursion is true, do nothing here, and then break-out of the recursion via the call to rightOrOops(...).
	} else {
		// fmt.Printf("here in evaluateUsersGuess 1 \n") // todo <---------------------

		if gameOn {
			game_loop_counter++
		} else {
			game_loop_counter = 0
		} // ::: This (next line) be the only point in this function from which execution can leave this function.
		// rightOrOops(userInput, promptField
		rightOrOops(userInput, promptField) // May call: tryAgain() ... which may call: lastTry().
		// fmt.Printf("Called from evaluateUsersGuess 1 new_prompt is %s and promptField is %s \n", new_prompt, promptField) // todo <---------------------
		// ::: re setc: promptField "wrong" now after bottom of rightOrOops. this would have been the value last passed to this func.

		/*
		   Eventually, we return here (effectively directly to here) from all subsequent functions. Excepting from the end of lastTry(),
		   which returns directly (effectively) to main(). But, in all other cases, we return to here ... with a fresh card, already in hand,
		   ready to re-prompt, and to then handle the potential of a user's Directive (both of which happen just-below here). Then, a
		   recursion brings us back here (to the top of this func) only to be sent out, once again, through-the-mill, via rightOrOops() etc.
		   ...
		   Clearly, when falling from here, the recall bool can ONLY LOGICALLY BE false :: sequential calls to rightOrOops() cannot be made!
		*/
	}
	if recall {
		// fmt.Printf("here in evaluateUsersGuess 2 \n") // todo <---------------------

		if gameOn {
			game_loop_counter++
		} else {
			game_loop_counter = 0
		}
		// ::: Ditto the prior emphasized comment. (this be the only exit point).
		rightOrOops(userInput, promptField) // May call: tryAgain() ... which may call: lastTry().
		// fmt.Printf("Called from evaluateUsersGuess 2, new_prompt is %s and promptField is %s \n", new_prompt, promptField) // todo <---------------------
		/*
		   Eventually, we return here (effectively directly to here) from all subsequent functions. Excepting from the end of lastTry(),
		   which returns directly (effectively) to main(). But, in all other cases, we return to here ... with a fresh card, already in hand,
		   ready to re-prompt, and to then handle the potential of a user's Directive (both of which happen just-below here). Then, a
		   recursion brings us back here (to the top of this func) only to be sent out, once again, through-the-mill, via rightOrOops() etc.
		*/
	} else { // Now, LOGICALLY, (since the foregoing logic ABSOLUTELY precludes it) if the recall bool was false ...
		// ... then the recursion flag must also have been false, allowing execution to fall towards the recursive calls below.
	}
	/*
	   Next, we handle the potential of a Directive command having been issued by the user. If evaluateUsersGuess() is called after
	   handling a "Directive" (upon subsequent recursion) rightOrOops() MUST be entirely omitted. Thusly allowing for the
	   undeniable bliss of being able to CONSECUTIVELY-enter Directives at each INITIAL prompt. But it should go without saying that
	   to allow any Directives (especially that '?' one) at secondary or tertiary prompts would be just Evil! Plain and simple.
	*/
	if returning_fr_dir {
		if field_to_prompt_from == "kanji" {
			// fmt.Printf("Now in evaluateUsersGuess 3 new_prompt is %s and promptField is %s \n", new_prompt, promptField) // todo <---------------------
			// userInput = promptWithDir(promptField)                                                                       // promptField (rather than the new_prompt variant).
			userInput = promptWithDir(new_prompt) // if new_prompt was avail this may be the thing?
		} else {
			userInput = promptWithDir(aCard.Kunyomi)
		}
		// returning_fr_dir = false
	} else { // ::: we get here sporadically from dir handlers todo: FIXIT!
		if field_to_prompt_from == "kanji" {
			// fmt.Printf("here in evaluateUsersGuess 4 ? sporatic ? \n") // todo <---------------------

			userInput = prompt_interim3(promptField) // promptField (rather than the new_prompt variant).
		} else {
			userInput = prompt_interim3(aCard.Kunyomi)
		}
	}
	// Handle the potential of a Directive command from the user.
	if userHasGivenA_DirectiveIsteadOfGuess(userInput) {
		if userInput == "setc" { // See prior comments in regards to this section.
			fmt.Printf("here in begin2\n")

			promptField, objective, objective_kind, secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(userInput)
		} else {
			_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(userInput)
		}
		/*
			   Recursively ...
			   after responding to a Directive, re-prompting is immediately required: obvious one would think.



			via recursion back to the top, while using the
			   same/original promptField etc. With recursion true & recall false, rightOrOops() will NOT be done with said recursion,
			   Upon recursion, we always fall-through to evaluate, one way or the other, after potentially handling a user given Directive.
			   .
			   true, false (immediately below) causes both of the rightOrOops() calls to be bypassed upon recursion.
		*/
		// fmt.Printf("here in evaluateUsersGuess 5 \n")

		evaluateUsersGuess(userInput, promptField, true, false)
	} else {
		/*
		   true, true (below) causes execution of ONE rightOrOops() upon recursion.
		*/
		// fmt.Printf("here in evaluateUsersGuess 6 \n")

		evaluateUsersGuess(userInput, promptField, true, true)
	}
	// fmt.Printf("here in evaluateUsersGuess 66 \n")

	// ::: Your are not leaving this function that easily!
}

/*
-
-
- ::: Only helper functions follow ...
-
*/

// todo: instead of all this fuss, what happens if we just allow respond_to_UserSuppliedDirective() to handle setc dirs unaided??
// ::: currently, setc is not updating the prompt. working now.
// ::: this is clobbering the prompt field????

// new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = directiveHandler
func directiveHandler(userInput, inew_prompt string) (new_prompt string, returning_fr_dir bool) {
	directiveHandlerHasBeenRun = true
	var setcWasRun bool
	setcWasRun = false
	if userHasGivenA_DirectiveIsteadOfGuess(userInput) {
		if userInput == "setc" {
			setcWasRun = true
			// fmt.Printf("here in setc3\n")

			new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(userInput)
		} else {
			setcWasRun = false
			_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(userInput)
		}
		returning_fr_dir = true // todo this is moot since we explicitly return true.
		// Either way, evaluateUsersGuess(...)
		// evaluateUsersGuess(userInput, inew_prompt, inew_objective, inew_objective_kind, new_secondary_objective, true, false)
	} else { // ::: else this be a non-dir very bona-fide guess!!
		returning_fr_dir = true // todo this is moot since we explicitly return true.
		// if setc was NOT run ::: this be a non-dir very bona-fide guess!!
		// ::: BUT WHAT IF WE JUST TAKES OUR CHANCES ?? // todo it was bad!!
		// evaluateUsersGuess(userInput, inew_prompt, inew_objective, inew_objective_kind, new_secondary_objective, true, true)
		// evaluateUsersGuess(userInput, inew_prompt, inew_objective, inew_objective_kind, new_secondary_objective, true, true)
		evaluateUsersGuess(userInput, inew_prompt, true, true)

	}

	// Determine which values to actually return; originals, i.e., internals/locals, i.e., inew_prompt etc. Or, the newly-minted??
	if setcWasRun {
		// fmt.Printf("here in directiveHandler setcWasRun\n")
		// ::: how do we go from here to (evaluateUsersGuess 3) ? because, we have to return from here to
		// return the forced new card PLUS the returning_fr_dir flag.
		return new_prompt, true // returning_fr_dir
	} else { // this runs but the returned values are lost after they are collected because they are local vars in the caller: rightOrOops
		// fmt.Printf("here in directiveHandler setc Was Not Run\n")
		/*
			here in directiveHandler setc Was Not Run
			execution has landed here 1, promptField is 年, new_prompt is 週 ::: and i did want week
			execution has landed here 14, promptField is 年, valueFromreSet_aCard_andThereBy_reSet_thePromptString is
			was i called from here 2
			here in evaluateUsersGuess 3
			年 : oldVar directiveHandlerHasBeenRun+ Meaning?
		*/
		// return originals, thereby keeping -- via not changing them.  And, of course, the returning_fr_dir flag.
		// new_prompt, new_objective, new_objective_kind, new_secondary_objective = inew_prompt, inew_objective, inew_objective_kind, inew_secondary_objective
		// storedValForPromp = inew_prompt + " rick says so"
		// return inew_prompt, inew_objective, inew_objective_kind, inew_secondary_objective, true // returning_fr_dir // todo unsure of what to do here.
		return inew_prompt, true // returning_fr_dir // todo unsure of what to do here.
		// return "rick", "rick", "rick", "rick", true // returning_fr_dir
		// return new_prompt, new_objective, new_objective_kind, new_secondary_objective, true // ::: am i returning the original, or ?
		// return new_prompt, new_objective, new_objective_kind, new_secondary_objective, true // returning_fr_dir
	}
}

// ::: setc now updates the prompt incorrectly. But everything else works well.

func promptCase1(new_prompt, field_to_prompt_from string) (userInput string) {
	// This prompt, deployed by new_objective_kind, takes new_prompt
	if field_to_prompt_from == "kanji" {
		// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
		userInput = promptWithDir(new_prompt) // Get user's input, from a randomly selected prompt
	}
	if field_to_prompt_from == "kunyomi" {
		userInput = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
	}
	return userInput
}

func freshCard(inew_prompt, inew_objective, inew_objective_kind, inew_secondary_objective string, already_used_map map[string]int) (new_prompt, new_objective, new_objective_kind, new_secondary_objective string) {
	var actuallyReturnNewValues bool
	actuallyReturnNewValues = false
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
		if is_pick_novel(inew_prompt) {
			already_used_map[inew_prompt]++
			break // keep it
		}
		// else [because of the double breaks above], pick another
		new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
		length = 0 // continue the loop ensuring that the entire map is read with this new pick
		actuallyReturnNewValues = true
	}
	if actuallyReturnNewValues {
		return new_prompt, new_objective, new_objective_kind, new_secondary_objective // return new pick.
	} else {
		return inew_prompt, inew_objective, inew_objective_kind, inew_secondary_objective // return old original values: keep first draw.
	}
}

func initialize_stuff() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with the current time in nanoseconds.

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
	// deck_len = claude_len // see also: functions.go/switch_the_deck() for the current fat-fingered default.
}
