package main

import (
	"fmt"
	"os"
	"time"
)

// DIRECTIVES : --------------------------------------------------------------------------------------------

/*
.
*/

// sdk Directive
func switch_the_deck() {
	for {
		fmt.Println("\nEnter a deck from below for randomized prompting of a specific deck:\n")

		// fmt.Println("    \"all\" ")
		fmt.Printf("    \"all\"      %d total cards (below)\n\n",
			beauty_len+claude_len+current_len+data_len+fresh_len+grad_len+guru_len+init_len+master_len+novice_len)
		// fmt.Println("    beauty")
		fmt.Printf("    \"beauty\"   %d cards in the Beauty terms deck\n", beauty_len)
		// fmt.Println("    claude")
		fmt.Printf("    \"claude\"   %d cards in the Claude deck\n", claude_len)
		// fmt.Println("    current")
		fmt.Printf("    \"current\"  %d cards in the \"Current\" deck \n", current_len)
		// fmt.Println("    data")
		fmt.Printf("    \"data\"     %d cards in the Data deck\n", data_len)
		// fmt.Println("    fresh")
		fmt.Printf("    \"fresh\"    %d cards in the \"Fresh\" deck\n", fresh_len)
		// fmt.Println("    grad")
		fmt.Printf("    \"grad\"     %d cards in the Graduate deck\n", grad_len)
		// fmt.Println("    guru")
		fmt.Printf("    \"guru\"     %d cards in the Guru deck\n", guru_len)
		// fmt.Println("    init")
		fmt.Printf("    \"init\"     %d cards in the Initiate deck\n", init_len)
		// fmt.Println("    mast")
		fmt.Printf("    \"mast\"     %d cards in the Master deck\n", master_len)
		// fmt.Println("    nov \n")
		fmt.Printf("    \"nov\"      %d cards in the Novice deck\n\n", novice_len)
		fmt.Printf("    \"words\"    %d cards in the Words deck\n\n", words_len)

		fmt.Printf(" Here%s:> %s", colorCyan, colorReset)

		_, _ = fmt.Scan(&current_deck) // current_deck is a global var.

		if current_deck == "q" {
			os.Exit(1)
		}
		if current_deck != "all" &&
			current_deck != "beauty" &&
			current_deck != "claude" &&
			current_deck != "current" &&
			current_deck != "data" &&
			current_deck != "fresh" &&
			current_deck != "grad" &&
			current_deck != "guru" &&
			current_deck != "init" &&
			current_deck != "mast" &&
			current_deck != "nov" &&
			current_deck != "words" {
			fmt.Printf("%s\n  \"%s\" is was not a valid deck, you seem to be a novice at this so nov it will be:) \n%s",
				colorRed, current_deck, colorReset)
			current_deck = "nov"
			deck_len = novice_len
			break // Causes the use of both the novice deck, and also kanji prompting vs kun'yomi prompting.
		} else {
			for {
				fmt.Println("\nEnter a field to prompt from:\n")
				fmt.Println("\nkanji or kunyomi\n")
				_, _ = fmt.Scan(&field_to_prompt_from)
				if field_to_prompt_from != "kanji" && field_to_prompt_from != "kunyomi" {
					fmt.Printf("%s\n  \"%s\" is not a valid field, try again: \n%s", colorRed, field_to_prompt_from, colorReset)
					continue
				} else {
					if current_deck == "all" {
						// deck_len = all_len // This is taken care of in pick_RandomCard_Assign_fields()
					}
					if current_deck == "beauty" {
						deck_len = beauty_len
					}
					if current_deck == "words" {
						deck_len = words_len
					}
					if current_deck == "claude" {
						deck_len = claude_len
					}
					if current_deck == "current" {
						deck_len = current_len
					}
					if current_deck == "data" {
						deck_len = data_len
					}
					if current_deck == "fresh" {
						deck_len = fresh_len
					}
					if current_deck == "grad" {
						deck_len = grad_len
					}
					if current_deck == "guru" {
						deck_len = guru_len
					}
					if current_deck == "init" {
						deck_len = init_len
					}
					if current_deck == "mast" {
						deck_len = master_len
					}
					if current_deck == "nov" {
						deck_len = novice_len
					}
					resetAllLogs()
					return
				}
			}
		}
	}
}

/*
.
*/

// Handles the Directive 'setc' (part 1 of 2)
func setc_kanji() {
	setcHasBeenrunGlobal = true
	actual_prompt_string, primary_meaning, secondary_meaning = reSet_aCard_andThereBy_reSet_thePromptString()
	// setcHasBeenrunGlobal = false
	// this needs to store our new prompt field in valueFromreSet_aCard_andThereBy_reSet_thePromptString
	// ::: and no it does that. Like this:
	// valueFromreSet_aCard_andThereBy_reSet_thePromptString = aCard.Kanji
	if foundElement == nil {
		fmt.Printf(" Setting to \"west\" :: \n%s", colorRed)
		// Show the user exactly what is about to be done:
		fmt.Println(runeOfCode)
		fmt.Printf(colorReset)
		silentlyLocateCard("west") // Set the Convenience-global: foundElement
		if foundElement != nil {
			aCard = *foundElement // Set the global var-object 'aCard'
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
		} else {
			fmt.Printf("\"west\" Not found : respond_to_UserSuppliedDirective()\n\n")
		}
	}
}

// Handles the Directive 'setc' (part 2 of 2)
func reSet_aCard_andThereBy_reSet_thePromptString() (prompt, primary_meaning, secondary_meaning string) { //  - -
	var theMeaningOfCardToSilentlyLocate string

	fmt.Printf("\nEnter an English Meaning \"using US chars\" to")
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" reSet the prompt via the corresponding Kanji card :> ")
	fmt.Printf("%s", colorReset) //
	_, _ = fmt.Scan(&theMeaningOfCardToSilentlyLocate)

	// Tentatively, prepare to Scan for user's input and attempt locating a matching 'aCard'

	// Confidently, go-looking for user's input: locate matching 'aCard'
	silentlyLocateCard(theMeaningOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
	if foundElement != nil {
		aCard = *foundElement // Set the global var-object 'aCard'
		prompt = aCard.Kanji
		if setcHasBeenrunGlobal {
			valueFromreSet_aCard_andThereBy_reSet_thePromptString = aCard.Kanji
		}
		primary_meaning = aCard.Meaning
		secondary_meaning = aCard.Second_Meaning
	} else {
		fmt.Println("Error, foundElement is nil in reSet_aCard_andThereBy_reSet_thePromptString()")
	}
	// primary_meaning_kind = "Romaji"
	fmt.Println("")

	return prompt, primary_meaning, secondary_meaning
}

/*
.
*/

func game_off() (game string) { // - -
	game = "off"
	gameOn = false
	game_duration = 998

	fileHandle, err := os.OpenFile("Kanji-newLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	currentTime := time.Now()

	// calculate elapsed time
	t_s2 := time.Now()
	elapsed := t_s2.Sub(TimeOfStartFromTop)

	// cast time durations to a String type for Fprintf "formatted print"
	TotalRun := elapsed.String()

	fmt.Printf("\nRun time was: %s, gameOn is: %t \n\n", TotalRun, gameOn)

	// End timer and report elapsed time
	_, err1 := fmt.Fprintf(fileHandle,
		"\n The game ended at: %s  Total prompts was: %d \n",
		currentTime.Format("15:04:05 on Monday 01-02-2006"), total_prompts)
	check_error(err1)
	_, err2 := fmt.Fprintf(fileHandle,
		"\n Elapsed time of game was: %s \n",
		TotalRun)
	check_error(err2)
	return game
}

// Creates a func named check_error which takes one parameter "e" of type error
func check_error(e error) { //    - -
	if e != nil {
		panic(e) // use panic() to display error code
	}
}
