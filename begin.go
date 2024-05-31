package main

/*
// todo remove debugging Printf lines.
// Little-more than a way to get to the func following the next func? Albeit it does implement consecutive directives feature at startup.
func begin() { // Prompting may occur via Hira, Kata, or Romaji chars.  - -
	// returning_fr_dir = false // ::: oddly, the only problem with removing every assignment to false is that setc sporadically fails to update the prompt ???

	// var usersSubmission string // A var is needed because ":=" does not work throughout these conditionals, i.e., "usersSubmission " is not in the signature.
	// var promptField string
	for {
		if field_to_prompt_from == "kanji" {
			// usersSubmission = promptWithDir(promptField) // Note: this prompt takes promptField (rather than the new_kanjiChar variant).
			promptWithDir() // Note: this prompt takes promptField (rather than the new_kanjiChar variant).
			// fmt.Printf("here in beginA\n")
		} else {
			promptWithDir()
		}

		// Handle the potential of a Directive command from user. ::: this allows for consecutive directives, at the start. but only there.
		if userHasGivenA_DirectiveIsteadOfGuess(usersSubmission) { // Checks to see if user's response is a Directive.
			if usersSubmission == "setc" { // Only capture all the values returned by respond_to_UserSuppliedDirective(usersSubmission ) if "setc" was done.
				// fmt.Printf("here in begin1\n") // todo <---------------------
				actual_prompt_string, primary_meaning, secondary_meaning, returning_fr_dir = respond_to_UserSuppliedDirective(usersSubmission)
				// ::: WE NEVER!!!!!!!!!!! EVER!!!!!!!!!! LAND HERE after a setc ????????????? unless it was done at startup.
				// fmt.Printf("execution has landed here 18\n") // todo <---------------------

			} else { // else handle it, but only capture the flag; a flag that is not needed here, but is usually needed in other cases.
				_, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(usersSubmission) // Returns a flag, and re-prompts (above).
			}
			continue // ^ ^ ^ After "Directive" handling, re-prompt (via the top of this for loop) with the same/original promptField value.
		} else { // A non-Directive was given, so assume it to be an attempted guess.
			// Passing recursion=false [or recall=true], to evaluateUsersGuess() causes the subsequent execution of rightOrOops() [via recursion]
			// near the top of evaluateUsersGuess()  ^   v   ^   v   ^   v   ^   v   ^   v   ^   v
			// ...   ^  v  ^  v  ^  v  ^  v  ^  v  ^  v  ^  v  ^  v   ^   v   ^   v   ^    false,   v   false, causes ONE execution of rightOrOops().
			// evaluateUsersGuess()
			// usersSubmission
			break // Exits this func, then ... perhaps much later -- perhaps having exhausted all potential guessing -- return to main() ...
		}
	}
	fmt.Printf("execution has landed here 17\n") // todo <---------------------

} // Returns to main()'s outer loop to (usually randomly) select a fresh card.

*/
