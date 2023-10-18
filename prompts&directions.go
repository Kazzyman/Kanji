package main

import (
	"fmt"
)

// Special prompts for use when soliciting second, or final, guesses
func promptForRomaji(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s", promptField)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Romaji input-mode expected,\n Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

func promptForRomajiWithDir(prompt string) (usersGuessOrOptionDirective string) { // - -
	fmt.Printf("%s", prompt)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Romaji input-mode expected, or '?' for help with: %s \n", prompt)
	fmt.Printf(" or, type '??' for help with something else ... \n")
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
	fmt.Println("View source code at https://github.com/Kazzyman/Kanji")
	fmt.Println("    Use Alpha-Numeric (US) input-mode on your system to:")
	fmt.Println("        Enter 'dir' to redisplay this menu of available directives")
	fmt.Println("        Enter 'gamed to set the Duration counter for a game session ")
	fmt.Println("        Enter 'gameon' to begin a game session ")
	fmt.Println("        Enter '?' for context-sensitive help on the current character")
	fmt.Println("        Enter 'stats' or 'st' for statistics re what you have done")
	fmt.Println("        Enter 'notes' for an explanation of Onyomi and Kunyomi")
	fmt.Println("        Enter 'reset' to reset (flush or clear) all stats logs etc.")
	fmt.Println("        Enter 'rm' to Read the current contents of the Maps")
	fmt.Println("        Enter 'set' to force the use of a specific card")
	fmt.Println("        Enter 'format' to format a text file as card elements")
	fmt.Println("        Enter 'exit' or 'quit', 'ex' or 'q', to terminate this app")
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
	fmt.Printf("\n")
	fmt.Println("View source code at https://github.com/Kazzyman/Kanji")
	fmt.Println("    Use Alpha-Numeric (US) input-mode on your system to:")
	fmt.Println("        Enter 'dir' to redisplay this menu of available directives")
	fmt.Println("        Enter 'gamed to set the Duration counter for a game session ")
	fmt.Println("        Enter 'gameon' to begin a game session ")
	fmt.Println("        Enter '?' for context-sensitive help on the current character")
	fmt.Println("        Enter 'stats' or 'st' for statistics re what you have done")
	fmt.Println("        Enter 'notes' for an explanation of Onyomi and Kunyomi")
	fmt.Println("        Enter 'reset' to reset (flush or clear) all stats logs etc.")
	fmt.Println("        Enter 'rm' to Read the current contents of the Maps")
	fmt.Println("        Enter 'set' to force the use of a specific card")
	fmt.Println("        Enter 'format' to format a text file as card elements")
	fmt.Println("        Enter 'exit' or 'quit', 'ex' or 'q', to terminate this app")
	//goland:noinspection ALL
	fmt.Printf("Game counter: %d, Game Duration: %d \n", game_loop_counter, game_duration+2)
	fmt.Printf("Current Prompt Count Total: %d \n", total_prompts)
}
