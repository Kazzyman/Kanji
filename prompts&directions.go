package main

import (
	"fmt"
)

func contentOf_List_of_Directives() {
	fmt.Println("View source code at https://github.com/Kazzyman/Kanji")
	fmt.Println("    Use Alpha-Numeric (US) input-mode on your system to:")
	fmt.Println("        Enter 'dir' redisplay this menu of available Directives")
	fmt.Println("        Enter 'gdc set the Duration Counter for a Game session ")
	fmt.Println("        Enter 'bgs' Begin a Game Session ")
	fmt.Println("        Enter '?' context-sensitive help on the current character")
	fmt.Println("        Enter 'st' Statistics")
	fmt.Println("        Enter 'nt' (notes) an explanation of Onyomi and Kunyomi")
	fmt.Println("        Enter 'rs' ReSet (flush or clear) all stats logs etc.")
	fmt.Println("        Enter 'rm' Read the current contents of the Maps")
	fmt.Println("        Enter 'st' (set) force the use of a specific card")
	fmt.Println("        Enter 'frmt' Format a text file as card elements")
	fmt.Println("        Enter 'q', (quit) terminate the app")
}

// Special prompts for use when soliciting second, or final, guesses
func promptForRomaji(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s", promptField)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Type meaning,\n Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

func promptForRomajiWithDir(prompt string) (usersGuessOrOptionDirective string) { // - -
	fmt.Printf("%s", prompt)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Type meaning, 'dir' or '?' for help with: %s \n", prompt)
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
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
	contentOf_List_of_Directives()
	//goland:noinspection ALL
	fmt.Println("\n")
	fmt.Printf("Game counter: %d, Game Duration: %d \n", game_loop_counter, game_duration+2)
	fmt.Printf("Current Prompt Count Total: %d \n", total_prompts)
}

// 'Directive Menu' ; displays only in response to "Dir" Directive
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
	/*
		dk, show the name of the current deck
		chdk, display a menu of the available decks : and then switch to a selected deck
	*/
	fmt.Printf("\n")
	contentOf_List_of_Directives()
	//goland:noinspection ALL
	fmt.Printf("Game counter: %d, Game Duration: %d \n", game_loop_counter, game_duration+2)
	fmt.Printf("Current Prompt Count Total: %d \n", total_prompts)
}
