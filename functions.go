package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"
)

// DIRECTIVES : --------------------------------------------------------------------------------------------
//

// Handles the Directive 'set'
func reSet_aCard_andThereBy_reSet_thePromptString() (prompt, objective, objective_kind string) { //  - -
	var theHiraganaOfCardToSilentlyLocate string
	var isAlphanumeric bool

	fmt.Printf("\nEnter a Hiragana to")
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" reSet the prompt :> ")
	fmt.Printf("%s", colorReset) //
	_, _ = fmt.Scan(&theHiraganaOfCardToSilentlyLocate)

	// Determine if the user has entered a valid Hiragana char (instead of, accidentally, an alpha char or string)
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true {
	case findAlphasIn.MatchString(theHiraganaOfCardToSilentlyLocate):
		isAlphanumeric = true
	default:
		isAlphanumeric = false
	}
	// Tentatively, prepare to Scan for user's input and attempt locating a matching 'aCard'
	if isAlphanumeric == true {
		fmt.Println("Are you in alphanumeric input mode?")
		fmt.Printf("... if so, change it to Hiragana (or I mignt die)\n")
		fmt.Printf("%s", colorRed) //
		fmt.Printf(" cautiously ")
		fmt.Printf("%s", colorCyan)
		fmt.Printf("re-enter your selection, in Hiragana mode :> ")
		fmt.Printf("%s", colorReset)
		_, _ = fmt.Scan(&theHiraganaOfCardToSilentlyLocate)
		// May yet send an Alpha string to the next func, which will itself deal with it elegantly
		silentlyLocateCard(theHiraganaOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                                 // Set the global var-object 'aCard'
		// new_prompt, new_objective, new_objective_kind
		prompt = aCard.Kanji
		objective = aCard.Onyomi
		objective_kind = "Romaji"
		fmt.Println("")
	} else {
		// Confidently, go-looking for user's input: locate matching 'aCard'
		silentlyLocateCard(theHiraganaOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                                 // Set the global var-object 'aCard'
		prompt = aCard.Kanji
		objective = aCard.Onyomi
		objective_kind = "Romaji"
		fmt.Println("")
	}
	return prompt, objective, objective_kind
}

// end of DIRECTIVES -----------------------------------------------------------------------------------

// Creates a func named check_error which takes one parameter "e" of type error
func check_error(e error) { //    - -
	if e != nil {
		panic(e) // use panic() to display error code
	}
}

func testForDirective(in string) (result bool) { // - -
	if in == "set" ||
		in == "?" || // <-- If it IS a directive
		in == "reset" ||
		in == "stat" ||
		in == "st" ||
		in == "dir" ||
		in == "quit" ||
		in == "q" ||
		in == "exit" ||
		in == "ex" ||
		in == "stats" ||
		in == "rm" ||
		in == "gameon" ||
		in == "gameoff" ||
		in == "gamed" {
		// Then:
		result = true
	}
	return result
}

func game_on() (game string) { // - -
	game = "on"
	gameOn = true
	fmt.Println("The game is on")

	startBeforeCall = time.Now()
	currentTime := time.Now()
	TimeOfStartFromTop = time.Now()

	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	_, err1 := fmt.Fprintf(fileHandle,
		"\n The game began at: %s \n",
		currentTime.Format("15:04:05 on Monday 01-02-2006"))
	check_error(err1)
	return game
}
func game_off() (game string) { // - -
	game = "off"
	gameOn = false
	game_duration = 998

	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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

func respond_to_UserSuppliedDirective(in, objective_kind string) (prompt, objective, kind string) { // - -
	var count int
	switch in {
	case "gamed":
		fmt.Println("Enter a number for how many prompts there will be in the game")
		_, _ = fmt.Scan(&count)
		game_duration = count - 2
	case "gameon":
		// game_loop_counter ++
		game_on()
	case "gameoff":
		game_off()
	case "reset":
		// Flush (clear) the old stats and hits arrays
		cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
		cyclicArrayHits = CyclicArrayHits{}
		// Also, flush (clear) the maps
		total_prompts = 0
		//
		//goland:noinspection ALL
		fmt.Println("\nArrays and maps flushed:\n")
		fmt.Println("    cyclicArrayOfTheJcharsGottenWrong")
		fmt.Println("    cyclicArrayHits")
		fmt.Println("    frequencyMapOf_IsFineOnChars")
		//goland:noinspection ALL
		fmt.Println("    frequencyMapOf_need_workOn\n")
		fmt.Println("  And, all Game values have also been reset")
	case "quit":
		os.Exit(1)
	case "q":
		os.Exit(1)
	case "exit":
		os.Exit(1)
	case "ex":
		os.Exit(1)
	case "?":
		fmt.Printf("\n%s\n%s\n%s\n\n", aCard.Kanji, aCard.Meaning)
	case "set":
		prompt, objective, kind = reSet_aCard_andThereBy_reSet_thePromptString()
	case "stat":
		hits()
	case "st":
		hits()
	case "stats":
		hits()
	case "dir": // reDisplay the DIRECTORY OF DIRECTIVES (and instructions):
		re_display_List_of_Directives()
	case "rm":
		read_map_of_fineOn()
		read_map_of_needWorkOn()
	default:
		// fmt.Println("Directive not found") // Does not work because only existent cases are passed to the switch
	}
	return prompt, objective, kind
}

func pick_RandomCard_Assign_fields() (promptField, objective, objective_kind string) { // - -
	/*
		Kanji   string
		Meaning string
		Onyomi  string aCard.Onyomi done
		Kunyomi string
		Vocab   string

		Kata   string
		Hira   string
		Romaji string
		HiraHint string
		KataHint string
	*/
	randIndex := rand.Intn(len(fileOfCardsK))
	aCard = fileOfCardsK[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
	promptField = aCard.Kanji
	objective = aCard.Meaning
	objective_kind = "Romaji"

	return promptField, objective, objective_kind
}
