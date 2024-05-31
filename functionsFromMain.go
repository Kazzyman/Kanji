package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func frontEnd_Possible_Recursive_DirHandler() { // ::: - -
	detectDirective(usersSubmission)
	if its_a_directive {
		its_a_directive = false

		respond_to_UserSupplied_Directive(usersSubmission)

		prompt_the_user_for_input()       // re-prompt using same card as before.
		_, _ = fmt.Scan(&usersSubmission) // todo ?

		detectDirective(usersSubmission)
		if its_a_directive {
			its_a_directive = false
			frontEnd_Possible_Recursive_DirHandler() // ::: Recursion
		} else {
			return // Need this to get out of the recursive func.
		}
	}
}
func detectDirective(in string) { // ::: - -
	if in == "stc" ||
		in == "stcr" ||
		in == "?" || // <-- If it IS a directive
		in == "??" ||
		in == "rs" ||
		in == "st" ||
		in == "dir" ||
		in == "nts" ||
		in == "q" ||
		in == "rm" ||
		in == "abt" ||
		in == "exko" ||
		in == "exkf" ||
		in == "konly" ||
		in == "honly" ||
		in == "ronly" ||
		in == "donly" ||
		in == "hko" ||
		in == "help" {
		// Then:
		its_a_directive = true
	}
}
func respond_to_UserSupplied_Directive(usersSubmission string) { // ::: - -
	switch usersSubmission {
	// Directives follow:
	// ::: alphabetically (mostly)
	case "?":
		fmt.Printf("\n%s\n%s\n%s\n\n", aCard.Vocab, aCard.Vocab2, aCard.Kanji)

	case "dir": // reDisplay the DIRECTORY OF DIRECTIVES (and instructions):
		re_display_List_of_Directives()

	case "q":
		os.Exit(1)
	case "rm":
		read_map_of_fineOn()
		read_map_of_needWorkOn()
		// read_pulls_not_used_array()
	case "rs":
		// reset_all_data()
	case "st":
		// st_stats_function()
	case "stc":
		// reSet_via_a_hira_aCard_andThereBy_reSet_thePromptString()
	case "stcr":
		// reSet_aCard_via_a_romaji_andThereBy_reSet_thePromptString()
	default:
		// fmt.Println("Directive not found") // Does not ever happen because only existent cases are passed to the switch.
	}
}

/*
func promptCase1(new_kanjiChar, field_to_prompt_from string) (usersSubmission string) {
	// This prompt, deployed by new_takes new_kanjiChar
	if field_to_prompt_from == "kanji" {
		// This prompt, deployed by takes promptField (rather than the new_kanjiChar variant)
		usersSubmission = promptWithDir(aCard.Kanji) // Get user's input, from a randomly selected prompt
	}
	if field_to_prompt_from == "kunyomi" {
		usersSubmission = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
	}
	return usersSubmission
}

*/

func freshCard() {
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
		if is_pick_novel(actual_prompt_string) {
			already_used_map[actual_prompt_string]++
			break // keep it
		}
		// else [because of the double breaks above], pick another
		pick_RandomCard_Assign_fields()
		length = 0 // continue the loop ensuring that the entire map is read with this new pick
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
