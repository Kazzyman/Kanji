package main

import (
	"fmt"
	"strings"
)

func hits() { // - -
	// Create maps to store the frequency of each relevant string for that map
	frequencyMapRightOrOops := make(map[string]int)
	frequencyMapChar := make(map[string]int) // These, apparently, create a map to associate a unique string with an int
	frequencyMapWrongs := make(map[string]int)

	//
	// Parse the relevant cyclic array to check_for_match_in_other_fields the strings and put them into the relevant map:
	//
	// Load the RightOrOops frequency map
	for i := 0; i < len(cyclicArrayHits.RightOrOops); i++ {
		str := cyclicArrayHits.RightOrOops[i]
		// Apparently; this loads a string into, and increments the frequency of, that particular string, in the map
		frequencyMapRightOrOops[str]++ // Specifically, the '++' must increment the int value associated with str
	}
	// Load the char frequency map // As documented in the foregoing 'for' loop
	for i := 0; i < len(cyclicArrayHits.jchar); i++ {
		str := cyclicArrayHits.jchar[i]
		frequencyMapChar[str]++
	}
	// Load the wrongs frequency map // As documented in the foregoing 'for' loop
	for i := 0; i < len(cyclicArrayOfTheJcharsGottenWrong.jchar); i++ {
		str := cyclicArrayOfTheJcharsGottenWrong.jchar[i]
		frequencyMapWrongs[str]++
	}

	// -- PRINT -- the 'Right' and 'Oops' and their frequencies (Right or Oops) (top of printout)
	for str, freq := range frequencyMapRightOrOops { // The map has only one entry for Right, & one for Oops
		if str == "Right" { // Finds the one potential entry for Right
			fmt.Printf(colorGreen)
			fmt.Printf(" %s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorGreen)
			fmt.Printf(" %d\n", freq) // Frequency of Right, per the map
			fmt.Printf(colorReset)
			total_prompts = freq
		} else if str == "Oops" { // Finds the one potential entry for Oops
			fmt.Printf(colorRed)
			fmt.Printf(" %s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorRed)
			fmt.Printf(" %d\n", freq) // Frequency of Oops, per the map
			fmt.Printf(colorReset)
			total_prompts = total_prompts + freq
		} else if str == "" {
			// else, it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		}
	}
	// Print the unique Chars and their frequencies (continuing the printout above)
	numberOfUniqueCharsHit := 0
	for str, freq := range frequencyMapChar {
		if str == "" {
			// else, it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		} else {
			numberOfUniqueCharsHit++
			fmt.Printf(" %s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorReset)
			fmt.Printf(" %d\n", freq)
		}
	}
	fmt.Printf("Number of unique chars: ")
	fmt.Printf(colorPurple)
	fmt.Printf("%d \n\n", numberOfUniqueCharsHit)
	fmt.Printf(colorReset)

	fmt.Printf("Total prompts:")
	fmt.Printf(colorRed)
	fmt.Printf(" %d\n", total_prompts)
	fmt.Printf(colorReset)

	// Print the ones gotten wrong  (continuing the printout above)
	for str, freq := range frequencyMapWrongs {
		if str == "" {
			// it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		} else {
			// Print "Gotten Wrong:" + 'str' as multicolored (below)
			fmt.Printf("Gotten Wrong:")            // (in color White)
			fieldsOfStr := strings.Split(str, ":") // Print 'str' as multicolored (below)
			//                              // Gotten Wrong: (in color White)
			fmt.Printf(colorRed)
			fmt.Printf("%s", fieldsOfStr[0]) // KataOrWhatEver (in color Red)
			fmt.Printf(colorReset)
			fmt.Printf(" %s: ", fieldsOfStr[1]) // it was (in color White)
			fmt.Printf(colorRed)
			fmt.Printf("%s", fieldsOfStr[2]) // RomajiOrWhatEver (in color Red)
			fmt.Printf(colorReset)
			fmt.Printf(" %s: ", fieldsOfStr[3]) // but you had guessed:_ (in color White) _ is a space char
			fmt.Printf(colorRed)
			fmt.Printf("%s ", fieldsOfStr[4]) // the bad guess_ (in color Red) _ is a space char
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:") // Frequency: (in color Cyan)
			fmt.Printf(colorReset)
			fmt.Printf(" %d \n", freq) // 'a number' (in color White)
		}
	}
	fmt.Println("")
}

// from line parsing file:
// Older versions and segments of code:
/*
pattern := `(?i).*` + strings.Join(strings.Split(our_guess_in_lower_case, ""), `.*`) + `.*`
    reg := regexp.MustCompile(pattern)

    for _, str := range strings_from_card {
        if reg.MatchString(str) {
            found_one = true
            break
        }
    }
*/

/*
func customMatch(our_guess, str string) bool {
	ourGuessLower := strings.ToLower(our_guess)
	strLower := strings.ToLower(str)

	// Debugging info:
	fmt.Printf("strLower:%s, ourGuessLower:%s, %v\n", strLower, ourGuessLower, strings.Contains(strLower, ourGuessLower))
	fmt.Printf("ourGuessLower:%s, strLower:%s, %v\n\n", ourGuessLower, strLower, strings.Contains(ourGuessLower, strLower))

	return strings.Contains(strLower, ourGuessLower) ||
		strings.Contains(ourGuessLower, strLower)
	// Add more code here to parse strLower for any (non-alpha delimited) substrings in it that match any substring in ourGuessLower ...
	// ... such that if strLower is: "one, two, three, sam, sick, water", it will match if ourGuessLower is: "xxxxwaterxxxx" where x is any char
}

func check_for_match_in_other_fields(in string) (found_one bool) {
	our_guess_in_lower_case := strings.ToLower(in)
	strings_from_card := []string{aCard.Meaning, aCard.Second_Meaning, aCard.Kunyomi, aCard.Onyomi, aCard.Vocab, aCard.Vocab2}

	for _, str := range strings_from_card {
		if customMatch(our_guess_in_lower_case, str) {
			found_one = true
			break
		}
	}

		// Debugging: Print whether a match was found
		fmt.Printf("Search String: %s\n", our_guess_in_lower_case)
		fmt.Printf("Fields: %v\n", strings_from_card)
		fmt.Printf("Matches: %v\n", found_one)

		return found_one
}

*/
