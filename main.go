package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	countSLOC()
	initialize_stuff()
	display_List_of_Directives()
	switch_the_deck()
	inception = true                     // Because the first time is always so special.
	gotLastCardRightSoGetFreshOne = true // Gets us the initial card.
	begin_kanji_practice()
}

func begin_kanji_practice() {
	for {
		// Obtain a card from one of the decks.
		if gotLastCardRightSoGetFreshOne {
			pick_RandomCard_Assign_fields() // Cards (and Decks) may both be randomized.
			is_that_card_really_fresh()
			if emptyCardCounter > 0 {
				fmt.Printf("%s--- %d times an empty card was skipped in %s%s\n", colorRed, emptyCardCounter, current_deck, colorReset) // Verified
				emptyCardCounter = 0
			}
			guessLevelCounter = 1
		}

		// Prompt the user.
		if inception {
			promptWithDirAtInception() // Differs in that it terminates with fmt.Scan(&usersSubmission).
			inception = false
		} else {
			prompt_the_user_for_input() // ... according to guessLevelCounter; Scans, etc.
		}

		// Enforce gaming regulations etc.
		gaming_regulations()

		// Directive commands from the user can be tricky to handle elegantly.
		// During gaming, entirely-disallow checking for standard Directive commands from the user.
		if !aGameIsRunning {
			Recursive_Dir_Handler() // Trap and handle specific strings on the list of Directives.
			// The forgoing func contains a method to re-prompt so as to allow for consecutive commands.
		}
		// fmt.Printf("%sThe game_loop_counter was:%d, and the game_duration_set_by_user was:%d%s\n", colorRed, game_loop_counter, game_duration_set_by_user, colorReset)

		// If the users-Submission was not trapped and handled by Recursive_Dir_Handler()
		Process_users_input_as_a_guess()

	}
}

/*
.
*/

// Process_users_input_as_a_guess trys several methods to match a card to the user's guess, i.e., it is friendly.
func Process_users_input_as_a_guess() { // - -

	// The instance of a simple and complete match of the primary meaning field is straight-forward.
	if strings.EqualFold(usersSubmission, aCard.Meaning) { // Think: if usersSubmission == aCard.Meaning // but ignoring case
		// (These Display "Right" messages -- Unlike those in Jap2 -- are kept here for this Kanji version).
		// Display an appropriate response:
		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf(" %s \"%s%s%s\" ",
			actual_prompt_string, colorRed, usersSubmission, colorGreen)
		fmt.Printf("or \"%s%s%s\" %s \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, actual_prompt_string, colorReset)

		gotLastCardRightSoGetFreshOne = true
		log_right(usersSubmission, actual_prompt_string)
	} else // an else if ::: without this "else if" the following else block would execute following a match above.

	// And, the instance of a complete match of the second meaning field is also straight-forward.
	if strings.EqualFold(usersSubmission, aCard.Second_Meaning) {
		// Display an appropriate response:
		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf(" %s \"%s%s%s\" or \"%s%s%s\" %s \n\n%s",
			actual_prompt_string, colorReset, aCard.Meaning, colorGreen, colorRed, aCard.Second_Meaning, colorGreen, actual_prompt_string, colorReset)

		gotLastCardRightSoGetFreshOne = true
		log_right(usersSubmission, actual_prompt_string)
	} else { // ::: this naked else executes only if BOTH of the prior if statements FAIL to match.

		// if the user has guessed "to xxxx" or "the xxxx" we want to omit the first field before continuing.
		if len(words) > 1 { // words is a global slice which is populated in prompt_the_user_for_input() from usersSubmission
			// fmt.Printf("length was greater than 1\nand so we now check for \"to\" or \"the\"\n")
			if words[0] == "the" || words[0] == "to" {
				usersSubmission = words[1] // Reset usersSubmission equal to a more-significant field prior to sub-match searches.
				// fmt.Printf("after trimming we have %s as the new userSubmission\n", usersSubmission)
			}
		}
		// Having failed to fully-match the users-Submission to one of the two meaning fields ...
		// things will now have to get fancy. First, we'll strategically check for a sub-match WITHIN the primary-meaning field ...
		if check_for_match_within_primary_field(usersSubmission) {
			// Display an appropriate response if a reasonable sub-match was found within the primary meaning field:
			fmt.Printf("%s", colorReset)
			fmt.Printf("    ^^%spartly-correct, within the primary meaning: \"%s%s%s\" \n", colorGreen, colorReset, aCard.Meaning, colorGreen)
			fmt.Printf("   %s \n   %s \n%s%s\n", aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
			fmt.Printf("  \"%s%s%s\" ",
				colorReset, aCard.Meaning, colorGreen)
			fmt.Printf("or \"%s%s%s\" \n\n%s",
				colorReset, aCard.Second_Meaning, colorGreen, colorReset)

			gotLastCardRightSoGetFreshOne = true
			log_right(usersSubmission, actual_prompt_string)

			// ... failing that, we'll strategically check for a sub-match WITHIN the secondary-meaning field.
		} else if check_for_match_within_secondary_field(usersSubmission) {
			// Display an appropriate response if a reasonable sub-match was found with the secondary meaning field:
			fmt.Printf("%s", colorReset)
			fmt.Printf("    ^^%spartly-correct, within the secondary meaning: \"%s%s%s\" \n", colorGreen, colorReset, aCard.Second_Meaning, colorGreen)
			fmt.Printf("   %s \n   %s \n%s%s\n", aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
			fmt.Printf("  \"%s%s%s\" ",
				colorReset, aCard.Meaning, colorGreen)
			fmt.Printf("or \"%s%s%s\" \n\n%s",
				colorReset, aCard.Second_Meaning, colorGreen, colorReset)

			gotLastCardRightSoGetFreshOne = true
			log_right(usersSubmission, actual_prompt_string)

			// Failing the forgoing strategy, intelligently peruse the remaining fields of the card via a custom parsing algorithm.
		} else if check_for_match_in_other_fields(usersSubmission) {
			// Display an appropriate response if a "match" was found:
			// fmt.Printf("%s", colorGreen)

			fmt.Printf("\n%s%s\n%s\n%s\n%s\n%s", colorGreen,
				aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)
			fmt.Printf("\n%sfound:%s \"%s%s%s\" in an auxiliary field\n\n%s",
				colorPurple, colorGreen, colorReset, usersSubmission, colorGreen, colorReset)

			gotLastCardRightSoGetFreshOne = false
			// and then return to caller as if no match had been found, i.e.,
			// the next statement to be executed is: prompt_the_user_for_input() // ... according to guessLevelCounter; Scans, etc.

		} else {
			// Having failed every-which-way to secure a reasonable "match" of any sort, we concede with a simple Oops message ... or not
			if gottenHonestly { // there MAY be an Oops. // todo ??
				if supress_one_oops_message {
					// Suppress it. // ::: if a one-time request to suppress has been posted todo there will be NO Oops message
				} else {
					fmt.Printf("%s  　^^Oops! \n", colorRed) // todo, providing that it was gotten honestly, and there are no objections
					supress_one_oops_message = false
				}
			}
			// In any case:
			log_oops(actual_prompt_string, primary_meaning, usersSubmission)
			gotLastCardRightSoGetFreshOne = false
		}
	}
}
