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
		"' set the Deck")
	fmt.Println("        Enter '" + colorGreen +
		"gdc" + colorReset +
		"' set the Duration Counter for a Game session ")
	fmt.Println("        Enter '" + colorGreen +
		"bgs" + colorReset +
		"' or 'goff' Begin or end a Game Session ")
	fmt.Println("        Enter '" + colorGreen +
		"?" + colorReset +
		"' context-sensitive help on the current character")
	fmt.Println("        Enter '" + colorGreen +
		"st" + colorReset +
		"' Statistics")
	fmt.Println("        Enter '" + colorGreen +
		"nt" + colorReset +
		"' (notes) an explanation of Onyomi and Kunyomi")
	fmt.Println("        Enter '" + colorGreen +
		"rs" + colorReset +
		"' ReSet (flush or clear) all stats logs etc.")
	fmt.Println("        Enter '" + colorGreen +
		"rm" + colorReset +
		"' Read the current contents of the Maps")
	fmt.Println("        Enter '" + colorGreen +
		"setc" + colorReset +
		"' (set) force the use of a specific card")
	fmt.Println("        Enter '" + colorGreen +
		"frmt" + colorReset +
		"' Format a text file as card elements")
	fmt.Println("        Enter '" + colorGreen +
		"fif" + colorReset +
		"' Find in files")
	fmt.Println("        Enter '" + colorGreen +
		"q" + colorReset +
		"', (quit) terminate the app")
}

// Special prompts for use when soliciting second, or final, guesses
func prompt_interim(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s", promptField)
	fmt.Printf("%s", colorCyan)
	if current_deck == "all" || current_deck == "rand" {
		fmt.Printf(" Meaning? (deck:%s) Help is off, ", current_deckA)
	} else if current_deck == "randAll" {
		fmt.Printf(" Meaning? (mode:%s) Help is off, ", current_deckA)
	} else if current_deck == "fresh" {
		fmt.Printf(" Meaning? (deck:%s) Help is off, ", current_deckA)
	} else {
		fmt.Printf(" Meaning? (deck:%s) Help is off, ", current_deck)
	}
	fmt.Printf(string(colorReset))
	fmt.Printf("you must guess! \n :> ")
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return
}
func prompt_interim2(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s", promptField)
	fmt.Printf("%s", colorCyan)
	if current_deck == "all" || current_deck == "rand" {
		fmt.Printf(" Meaning? (deck:%s) Help is off, you must guess, ", current_deckA)
	} else if current_deck == "randAll" {
		fmt.Printf(" Meaning? (mode:%s) Help is off, you must guess, ", current_deckA)
	} else if current_deck == "fresh" {
		fmt.Printf(" Meaning? (deck:%s) Help is off, you must guess, ", current_deckA)
	} else {
		fmt.Printf(" Meaning? (deck:%s) Help is off, you must guess, ", current_deck)
	}
	fmt.Printf(string(colorReset))
	fmt.Printf("just once more!! \n :> ")
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return
}
func prompt_interim3(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s", promptField)
	fmt.Printf("%s", colorCyan)
	if current_deck == "all" || current_deck == "rand" {
		fmt.Printf(" Meaning? (deck:%s)", current_deckA)
	} else if current_deck == "randAll" {
		fmt.Printf(" Meaning? (mode:%s)", current_deckA)
	} else if current_deck == "fresh" {
		fmt.Printf(" Meaning? (deck:%s)", current_deckA)
	} else {
		fmt.Printf(" Meaning? (deck:%s)", current_deck)
	}
	fmt.Printf(string(colorReset))
	fmt.Printf(" try any substring from the %s", colorRed)
	fmt.Printf("red%s", colorReset)
	fmt.Printf(" text\n :> ")
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return
}

// Initial prompt, to be used when first introducing a new Kanji char
func promptWithDir(prompt string) (usersGuessOrOptionDirective string) { // - -
	fmt.Printf("%s", prompt)
	fmt.Printf("%s", colorCyan)
	if current_deck == "all" || current_deck == "rand" {
		fmt.Printf(" Meaning? (deck:%s), 'dir' or '?' for help with ", current_deckA)
	} else if current_deck == "randAll" {
		fmt.Printf(" Meaning? (mode:%s), 'dir' or '?' for help with ", current_deckA)
	} else if current_deck == "fresh" {
		fmt.Printf(" Meaning? (deck:%s), 'dir' or '?' for help with ", current_deckA)
	} else {
		fmt.Printf(" Meaning? (deck:%s), 'dir' or '?' for help with ", current_deck)
	}
	fmt.Printf(string(colorReset))
	fmt.Printf("%s \n", prompt)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" :> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return
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
	fmt.Printf("Current Deck is: %s \n", current_deck)

	if current_deck == "inits" || current_deck == "novs" || current_deck == "grads" ||
		current_deck == "masts" || current_deck == "gurus" {
		fmt.Println("Order is set to Sequential\n")
	} else if current_deck == "randAll" {
		fmt.Printf("Mode is randAll, and deck is %s\n\n", current_deckA)
	} else {
		fmt.Println("Order is Random\n")
	}
}
