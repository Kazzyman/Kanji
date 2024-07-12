package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func gaming_regulations() {
	if game_loop_counter > game_duration_set_by_user {
		game_loop_counter = 0
		// fmt.Printf("The game_loop_counter was:%d, and the game_duration_set_by_user was:%d", game_loop_counter, game_duration_set_by_user)
		game_off()
		postGame_wrap_up()
	}
	if aGameIsRunning {
		if usersSubmission == "game" {
		} // Do not start a new game if there is already one running.
		if usersSubmission == "q" {
			os.Exit(1)
		} else if usersSubmission == "off" || usersSubmission == "goff" {
			supress_one_oops_message = true
			postGame_wrap_up()
		} else if usersSubmission == "dirg" {
			display_limited_gaming_dir_list()
		}
	} else if usersSubmission == "game" {
		guessLevelCounter = 1
		fmt.Println("Welcome to the game. Dir options: off/goff, setc, q, dirg")
		fmt.Println("What is your name?")
		_, _ = fmt.Scan(&nameOfPlayer)
		fmt.Println("Enter a number for how many prompts there will be in the game")
		_, _ = fmt.Scan(&game_duration_set_by_user)
		display_limited_gaming_dir_list()
		now_using_game_duration_set_by_user = true
		supress_one_oops_message = true
		initialize_game()
	}
}

var supress_one_oops_message bool

/*
.
*/

func is_that_card_really_fresh() {
	reads := 0
	loopCounter := 0
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
}

/*
.
*/

func Recursive_Dir_Handler() { // ::: - -
	detectDirective(usersSubmission)
	if its_a_directive {
		its_a_directive = false

		respond_to_UserSupplied_Directive(usersSubmission)

		promptWithDirAtInception() // Works here because it has its own Scanner.

		detectDirective(usersSubmission)
		if its_a_directive {
			its_a_directive = false
			Recursive_Dir_Handler() // ::: Recursion
		} else {
			return // Need this to get out of the recursive func.
		}
	}
}

/*
.
*/

func detectDirective(in string) { // ::: - -
	if in == "sdk" ||
		in == "setc" ||
		in == "?" || // <-- If it IS a directive
		in == "??" ||
		in == "rs" ||
		in == "st" ||
		in == "dir" ||
		in == "q" ||
		in == "rm" ||
		in == "abt" ||
		in == "help" {
		// Then:
		its_a_directive = true
	}
}

/*
.
*/

func respond_to_UserSupplied_Directive(usersSubmission string) { // ::: - -
	switch usersSubmission {
	// Directives follow:
	// ::: alphabetically (mostly)
	case "?": // context-sensitive help on the current card
		fmt.Printf("\"%s\" is the primaryMeaning of %s\n\"%s\" is the secondaryMeaning of %s\n%s\n%s\n%s\n%s\n\n",
			aCard.Meaning, aCard.Kanji, aCard.Second_Meaning, aCard.Kanji, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab, aCard.Vocab2)
	case "abt":
		// countSLOC()
		// CalcSLOC_new()
		CalcData()
	case "dir": // reDisplay the DIRECTORY OF DIRECTIVES (and instructions):
		re_display_List_of_Directives()
	case "sdk":
		switch_the_deck()
	case "q":
		os.Exit(1)
	case "rm":
		read_map_of_fineOn()
		read_map_of_needWorkOn()
		// read_pulls_not_used_array()
	case "rs":
		resetAllLogs()
	case "setc":
		setc_kanji()
	case "st":
		newHits()
	default:
		// fmt.Println("Directive not found") // Does not ever happen because only existent cases are passed to the switch.
	}
}

/*
.
*/

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
	quiz_comp_len = len(fileOf_Quiz_comp)     //
	quiz_novice_len = len(fileOf_Quiz_novice) //
	quiz_master_len = len(fileOf_Quiz_master) //
	quiz_guru_len = len(fileOf_Quiz_guru)     //

	gameOn = false
	// current_deck = "claude" // The default deck, can be changed via the sdk Directive
	// deck_len = claude_len // see also: functions.go/switch_the_deck() for the current fat-fingered default.
}
