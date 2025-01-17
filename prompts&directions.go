package main

import (
	"fmt"
)

func List_of_Directives() {
	fmt.Println("View source code at https://github.com/Kazzyman/Kanji")
	fmt.Println("    Use Alpha-Numeric (US) input-mode on your system to:")
	fmt.Println("        Enter '" + colorGreen +
		"dir" + colorReset +
		"' redisplay this menu of available Directives")
	fmt.Println("        Enter '" + colorGreen +
		"sdk" + colorReset +
		"' set the Deck, or prompt_field")
	fmt.Println("        Enter '" + colorGreen +
		"?" + colorReset +
		"' context-sensitive help on the current character")
	fmt.Println("        Enter '" + colorGreen +
		"st" + colorReset +
		"' Statistics")
	fmt.Println("        Enter '" + colorGreen +
		"abt" + colorReset +
		"' About this app")
	fmt.Println("        Enter '" + colorGreen +
		"rs" + colorReset +
		"' ReSet (flush or clear) all stats logs etc.")
	fmt.Println("        Enter '" + colorGreen +
		"rm" + colorReset +
		"' Read the current contents of the Maps")
	fmt.Println("        Enter '" + colorGreen +
		"setc" + colorReset +
		"' (set) force the use of a specific card ...")
	fmt.Println("          ... via any English word; found in ANY deck :)")
	fmt.Println("        Enter '" + colorGreen +
		"game" + colorReset +
		"', to begin a session and log stats to a file")
	fmt.Println("        Enter '" + colorGreen +
		"q" + colorReset +
		"', (quit) terminate the app")
}

/*
.
.
*/

func extended_list_of_directives() {
	fmt.Println("        Enter '" + colorGreen +
		"fif" + colorReset +
		"' scan category files, compare \"source\" file. Report")
	fmt.Println("        Enter '" + colorGreen +
		"lff" + colorReset +
		"' create two files. A list of kanji. A list of duplicates")
	fmt.Println("        Enter '" + colorGreen +
		"t1" + colorReset +
		"' test1 : don't recall what I was testing???")
	fmt.Println("        Enter '" + colorGreen +
		"t2" + colorReset +
		"' test2 : don't recall what I was testing???")
	fmt.Println("        Enter '" + colorGreen +
		"t3" + colorReset +
		"' proposed third directive ??")
	fmt.Println("        Enter '" + colorGreen +
		"frmt" + colorReset +
		"' Format a text file as card elements")
}

// 'Directive Menu' ; displays only at inception
func display_List_of_Directives() { // (unique)     - -
	frequencyMapRightOrOops := make(map[string]int)
	for i := 0; i < len(cyclicArrayHits.RightOrOops); i++ {
		str := cyclicArrayHits.RightOrOops[i] // Parse the cyclicArray
		// Apparently; this loads a string into, and increments the frequency of, that particular string, in the map
		// ... one such entry per unique string
		frequencyMapRightOrOops[str]++ // Specifically, the '++' must increment the int value associated with str
	}
	// Then, parse the map which was created, and loaded, above
	// As is the way with maps, the frequency map has only one entry for Right, & one for Oops
	for str, freq := range frequencyMapRightOrOops {
		if str == "Right" { // Finds the one potential entry for Right
			total_prompts = freq // Obviously, total_prompts has been declared as a global, elsewhere
		} else if str == "Oops" { // Finds the one potential entry for Oops
			total_prompts = total_prompts + freq
		}
	}
	fmt.Printf("\n\n\n\n\n")
	List_of_Directives()
	//goland:noinspection ALL
	fmt.Println("\n")
	fmt.Printf("Game counter: %d, Game Duration: %d \n", game_loop_counter, game_duration+2)
	fmt.Printf("Current Prompt Count Total: %d \n", total_prompts)
	if current_deck == "" {
		current_deck = "yet to be determined"
	}
	fmt.Printf("Current Deck is: %s \n", current_deck)
}

// 'Directive Menu' ; displays only in response to "dir" Directive
func re_display_List_of_Directives() { // (unique)     - -
	frequencyMapRightOrOops := make(map[string]int)
	for i := 0; i < len(cyclicArrayHits.RightOrOops); i++ {
		str := cyclicArrayHits.RightOrOops[i] // Parse the cyclicArray
		// Apparently; this loads a string into, and increments the frequency of, that particular string, in the map
		// ... one such entry per unique string
		frequencyMapRightOrOops[str]++ // Specifically, the '++' must increment the int value associated with str
	}
	// Then, parse the map which was created, and loaded, above
	// As is the way with maps, the frequency map has only one entry for Right, & one for Oops
	for str, freq := range frequencyMapRightOrOops {
		if str == "Right" { // Finds the one potential entry for Right
			total_prompts = freq // Obviously, total_prompts has been declared as a global, elsewhere
		} else if str == "Oops" { // Finds the one potential entry for Oops
			total_prompts = total_prompts + freq
		}
	}
	fmt.Printf("\n")
	List_of_Directives()
	//goland:noinspection ALL
	fmt.Printf("Game counter: %d, Game Duration: %d \n", game_loop_counter, game_duration+2)
	fmt.Printf("Current Prompt Count Total: %d \n", total_prompts)
	if current_deck == "all" {
		fmt.Printf("Current Deck is: %s:%s \n", current_deck, current_deck_B)
	} else {
		fmt.Printf("Current Deck is: %s \n", current_deck)
	}
}
