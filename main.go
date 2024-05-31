package main

import (
	"fmt"
	"os"
)

// todo remove debugging Printf lines.
func main() {
	fmt.Println()
	countSLOC()
	initialize_stuff()
	display_List_of_Directives()
	switch_the_deck()
	gotLastCardRightSoGetFreshOne = true // This will get us an initial card.
	inception = true
	begin_kanji_practice()
}

func begin_kanji_practice() {
	// This outermost for loop (the main loop, pun intended) is endless, excepting os.Exit(1) via User-Directive "q".
	reads := 0
	loopCounter := 0
	for {
		// Cards (and Decks) may be randomized.
		if gotLastCardRightSoGetFreshOne {
			pick_RandomCard_Assign_fields() // This has within it lots of code to assure novel and fresh cards.
			// Is that pick really fresh?
			for reads <= len(already_used_map) { // Read entire map before concluding the pick is novel.
				reads++                 // Used to compare the current length of the map to the number of reads done of it.
				loopCounter++           // Required to determine if an excessive/exhaustive number of secondary picks have been made and tested.
				if loopCounter > 9999 { // 'loopCounter++' being the only way to exit the loop when no novel picks are possible.
					clearMap()
					break // Keep the last pick (could be any randomized card).
				}
				if is_pick_novel(actual_prompt_string) {
					already_used_map[actual_prompt_string]++ // Log the novel pick into the map.
					break                                    // Keep the novel card.
				}
				// via the implied else, pick another more-fresh card.
				pick_RandomCard_Assign_fields()
				reads = 0 // Continue the loop ensuring that the entire map is read, and compared to, the newly-picked card.
			}
			guessLevelCounter = 1
		}

		if inception {
			promptWithDirAtInception() // only do this on inception
			inception = false
		} else {
			// Prompt according to guessLevelCounter, type of displayed prompt, and type of requested response.
			// Each time we do this we increment the guessLevelCounter.
			prompt_the_user_for_input()
		}

		if game_loop_counter > game_duration {
			game_off()
		}
		if theGameIsRunning == true {
			// Skip looking to start a game if one is already running.
		} else if usersSubmission == "game" {
			guessLevelCounter = 1
			fmt.Println("Welcome to the game. Dir options: off/goff, stc, stcr, q, dirg")
			fmt.Println("What is your name?")
			_, _ = fmt.Scan(&nameOfPlayer)
			fmt.Println("Enter a number for how many prompts there will be in the game")
			_, _ = fmt.Scan(&game_duration_set_by_user)
			display_limited_gaming_dir_list()

			now_using_game_duration_set_by_user = true
			the_game_begins()
		}
		if theGameIsRunning != true {
			// skip looking to dirg since it only makes sense when a game is running.
		} else if usersSubmission == "dirg" {
			display_limited_gaming_dir_list()
		}

		if usersSubmission == "setc" { // Force a new card, formerly in dir handler, now only needed here.
			setc_kanji()
		}

		// During gaming, disallow checking for Directives other than q, and goff.
		if theGameIsRunning {
			if usersSubmission == "q" {
				os.Exit(1)
			}
			if usersSubmission == "off" || usersSubmission == "goff" {
				// the_game_ends()
			}
		} else {
			// If user's input is a Directive, handle it.
			frontEnd_Possible_Recursive_DirHandler() // ::: this contains the only other prompt
		}

		// this was only done in func log_right_andUpdateGame(prompt_it_was, in string) { todo in Jap2
		if game_loop_counter > game_duration { // todo, done twice
			game_off()
		}

		Process_users_input_as_a_guess()
	}
}

// var submission_already_processed_above bool // is global
func Process_users_input_as_a_guess() { // - -

	if usersSubmission == aCard.Meaning { // So, the guess was correct

		// Display right message (keep these here for Kanji
		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" ",
			colorReset, usersSubmission, colorGreen)
		fmt.Printf("or \"%s%s%s\" \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset)

		submission_already_processed_above = true
		submission_already_processed_above = true // ::: because we are doing it now.

		gotLastCardRightSoGetFreshOne = true
		// todo, write to a file, as illustrated below (how Jap2 does it)
		log_right(usersSubmission, actual_prompt_string)

	} else if submission_already_processed_above == true {
		// ... it was already done above, and there is nothing more to do, else
	} else if usersSubmission == aCard.Second_Meaning {

		// Display right message (keep these here for Kanji
		fmt.Printf("%s    %s \n    %s \n%s\n%s \n", // Indented Ony & Kun, out-dented vocabs (all in green)
			colorGreen, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" or \"%s%s%s\"\n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset, aCard.Meaning, colorGreen, colorReset)

		submission_already_processed_above = true // ::: because we are doing it now.

		gotLastCardRightSoGetFreshOne = true
		// todo, write to a file, as illustrated below (how Jap2 does it)
		log_right(usersSubmission, actual_prompt_string)

		// Check for a match within the primary field
	} else if submission_already_processed_above == true {
		// ... it was already done above, and there is nothing more to do, else
	} else if check_for_match_within_primary_field(usersSubmission) { // If any of the fields of the card contain a match via our custom parsing algorithm.

		// Display right message (keep these here for Kanji
		fmt.Printf("%s", colorReset)
		fmt.Printf("    ^^%spartly-correct, within the primary meaning: \"%s%s%s\" \n", colorGreen, colorReset, aCard.Meaning, colorGreen)
		fmt.Printf("   %s \n   %s \n%s%s\n", aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
		fmt.Printf("  \"%s%s%s\" ",
			colorReset, aCard.Meaning, colorGreen)
		fmt.Printf("or \"%s%s%s\" \n\n%s",
			colorReset, aCard.Second_Meaning, colorGreen, colorReset)

		submission_already_processed_above = true // ::: because we are doing it now.

		gotLastCardRightSoGetFreshOne = true
		// todo, write to a file, as illustrated below (how Jap2 does it)
		log_right(usersSubmission, actual_prompt_string)

	} else if submission_already_processed_above == true {
		// ... it was already done above, and there is nothing more to do, else
	} else if check_for_match_in_secondary_field(usersSubmission) { // If any of the fields of the card contain a match via our custom parsing algorithm.

		// Display right message (keep these here for Kanji
		fmt.Printf("%s", colorGreen)
		fmt.Printf("     %s^^%spossibly-correct, in another field\n", colorReset, colorGreen)
		fmt.Printf("\"%s%s%s\" was its primary meaning\n\"%s%s%s\" was its secondary meaning\n%s\n%s\n%s\n%s\n\n%s",
			colorReset, aCard.Meaning, colorGreen, colorReset, aCard.Second_Meaning, colorGreen,
			aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2, colorReset)
		fmt.Printf("%sfound:%s \"%s%s%s\" in secondary meaning field\n\n%s",
			colorPurple, colorGreen, colorReset, usersSubmission, colorGreen, colorReset)

		submission_already_processed_above = true // ::: because we are doing it now.

		gotLastCardRightSoGetFreshOne = true
		// todo, write to a file, as illustrated below (how Jap2 does it)
		log_right(usersSubmission, actual_prompt_string)

	} else { // it is not a Right, and is therefor an Oops, no new aCard has been fetched etc.
		fmt.Printf("%s  　^^Oops! \n", colorRed)

		log_oops(actual_prompt_string, primary_meaning, usersSubmission)
		gotLastCardRightSoGetFreshOne = false
	}
	// todo this may be an issue:
	submission_already_processed_above = false // So, reset it for the next round.

	/* this is how Jap2 does file logging:
	if usersSubmission == actual_objective {
			if guessLevelCounter == 3 {
				fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				check_error(err)
				_, err1 := fmt.Fprintf(fileHandle,
					"\nUser may have mistyped==%s:%s:%s", aCard.Romaji, aCard.Hira, aCard.Kata) // mistyped is a word?
				check_error(err1)
			} else if guessLevelCounter == 4 {
				fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				check_error(err)
				_, err1 := fmt.Fprintf(fileHandle,
					"\nUser had a some difficulty with==%s:%s:%s", aCard.Romaji, aCard.Hira, aCard.Kata)
				check_error(err1)
			}
			gotLastCardRightSoGetFreshOne = true // done in Kanji
			displayRight_logRight(usersSubmission, actual_prompt_char, actual_objective_type) // done differently in Kanji
		} else {
			gotLastCardRightSoGetFreshOne = false
			// logging etc. of Oops is being done in log_oops() by way of prompt_the_user_for_input()
		}
	*/

}
