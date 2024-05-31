package main

import (
	"fmt"
	"strings"
)

/*
	i = 0 // i being a convenience global
	for i, lastPull = range cyclicArrayPulls.pulls {
		if lastPull == promptField { // If promptField has been used in recent memory ...
			// Store seen promptField in frequencyMapOfSeenChars prior to getting another that may also not be novel (novel ones get stored later)
			frequencyMapOfSeenChars[promptField]++ // The '++' increments the int value associated with promptField
			fmt.Printf("We've seen the pseudo-random char before lastPull:%s and promptField:%s \n", lastPull, promptField)
			// If a match is found, set a new promptField and return
			promptField, primary_meaning, secondary_meaning = randomize_over_all_decks()
			// Having obtained a replacement promptField for the one we had already seen ...
			i = 0 // Reset the index i so as to check the replacement against the entire contents of the array
			// And ... implicitly continue the loop
		} else {
			// There was no match yet, so continue to end of array
			i++
			// a continue seems to be implied ???????????????
			// continue // Only continue to the end of the array; and then "break"
		}
	} // The loop ends when the promptField has been verified as being a novel one

*/
// indexIntoSlice = 0 // i being a convenience global
// A for loop, per chat suggestion (to avoid the recursion that I had been using, which eventually caused a stack overflow*)
// A standard format :: for i := 0; i < len(cyclicArrayPulls.pulls); i++ {
/*
		This for loop iterates over the elements of the cyclicArrayPulls.pulls slice using the range clause. Here's what happens:

		_, is the index of the current element in the cyclicArrayPulls.pulls slice

		lastPull is the value of the current element at index _,

		The loop iterates through the entire cyclicArrayPulls.pulls slice, and _, is automatically incremented with each iteration

	for _, lastPull := range cyclicArrayPulls.pulls {
		if lastPull == promptField {
			// Store non-novel promptField in frequencyMapOfSeenChars prior to getting another that may also not be novel (novel ones get stored later)
			frequencyMapOfSeenChars[promptField]++ // The '++' increments the int value associated with promptField in the map
			fmt.Printf("We've seen the pseudo-random char before lastPull:%s and promptField:%s \n", lastPull, promptField)
			// If a match is found, set a new promptField (which then, also, needs to be compared to all the members of the slice)
			promptField, primary_meaning, secondary_meaning = randomize_over_all_decks()
			// Having obtained a replacement promptField for the one which was found in the slice ...
			// instead of simply continuing the loop with the new promptField (and last-used index) ...
			// I must be assured that the entire slice is parsed with this new promptField: Do I need to force the range operator to restart ???
		} else {
			// else, do nothing: exit the loop since we have verified that promptField is nowhere to be found in the slice
		}
	}
*/
/*
	// Randomize over all decks and modes (except Guru)  (using global values of the indices to maintain state)
	if current_deck == "randAll" {

		// Randomly chose a deck and a mode
		randDeckAndMode = rand.Intn(8)
		//
		// Randomly access cards // deck and mode 0-3
		if randDeckAndMode == 0 {
			current_deckA = "init"
			randIndex := rand.Intn(len(fileOfCardsInitiate))
			aCard = fileOfCardsInitiate[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
			promptField = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
		} else if randDeckAndMode == 1 {
			current_deckA = "nov"
			randIndex := rand.Intn(len(fileOfCardsNovice))
			aCard = fileOfCardsNovice[randIndex]
			promptField = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
		} else if randDeckAndMode == 2 {
			current_deckA = "grad"
			randIndex := rand.Intn(len(fileOfCardsGraduate))
			aCard = fileOfCardsGraduate[randIndex]
			promptField = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
		} else if randDeckAndMode == 3 {
			current_deckA = "mast"
			randIndex := rand.Intn(len(fileOfCardsMaster))
			aCard = fileOfCardsMaster[randIndex]
			promptField = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
		} else if randDeckAndMode == 4 {
			//
			// Randomly access cards // deck and mode 4-7
			current_deckA = "inits"
			for {
				if indexInitS < len(fileOfCardsInitiate) {
					aCard = fileOfCardsInitiate[indexInitS]
					promptField = aCard.Kanji
					primary_meaning = aCard.Meaning
					secondary_meaning = aCard.Second_Meaning
					indexInitS++
					break // We only pull one card, and then exit
				} else {
					indexInitS = 0 // If end of deck, go through it again
					continue
				}
			}
		} else if randDeckAndMode == 5 {
			current_deckA = "novs"
			for {
				if indexNovS < len(fileOfCardsNovice) {
					aCard = fileOfCardsNovice[indexNovS]
					promptField = aCard.Kanji
					primary_meaning = aCard.Meaning
					secondary_meaning = aCard.Second_Meaning
					indexNovS++
					break
				} else {
					indexNovS = 0
					continue
				}
			}
		} else if randDeckAndMode == 6 {
			current_deckA = "grads"
			for {
				if indexGradS < len(fileOfCardsGraduate) {
					aCard = fileOfCardsGraduate[indexGradS]
					promptField = aCard.Kanji
					primary_meaning = aCard.Meaning
					secondary_meaning = aCard.Second_Meaning
					indexGradS++
					break
				} else {
					indexGradS = 0
					continue
				}
			}
		} else if randDeckAndMode == 7 {
			current_deckA = "masts"
			for {
				if indexMastS < len(fileOfCardsMaster) {
					aCard = fileOfCardsMaster[indexMastS]
					promptField = aCard.Kanji
					primary_meaning = aCard.Meaning
					secondary_meaning = aCard.Second_Meaning
					indexMastS++
					break
				} else {
					indexMastS = 0
					continue
				}
			}
		}
		//

*/
//
/*
	// My original code:
		lastPull = cyclicArrayPulls.pulls[i]
		// fmt.Printf("pulling/pulled:%s, OldPrompt is:%s, promptField is:%s\n", lastPull, oldPrompt, promptField)
		if lastPull == promptField { // seen it, so reset i, get new promptField, i.e. repeat
			fmt.Printf("seen it before as lastPull:%s- and promptField:%s- \n", lastPull, promptField)
			// seen it, so acquire a new promptField, having reset index so as to search array anew
			i = 0
			// The following global var from_recursion is used to alert the caller that some extra tasks must be done if recursion has occurred
			// from_recursion = true // Possibly a clumsy kluge?
			// The single bool argument of pick_aCard_and_assign_fields is not actually being used currently, it was just an initial idea
			// that I was toying with; and since it would require a fair amount of tedious edits to the code base, and I have not entirely
			// abandoned the idea that I may yet be using it, it has stuck around.
			// pick_aCard_and_assign_fields(true) // get yet another new promptField, just once!!!!!
			promptField, primary_meaning, secondary_meaning = randomize_over_all_decks()
			//
			cyclicArrayPulls.InsertKChar(promptField) // Maybe need to do this here too??
			//
			return

*/
// If a match is found, set a new promptField and return
// promptField = getNewPromptField()
//
// cyclicArrayPulls.InsertKChar(promptField) // Maybe need to do this here too??

/*
	for i := 0; i < len(cyclicArrayPulls.pulls); i++ {
		str := cyclicArrayPulls.pulls[i]
		// Apparently; this loads a string into, and increments the frequency of, that particular string, in the map
		frequencyMapOfSeenChars[str]++ // Specifically, the '++' must increment the int value associated with str
	}
*/
// return // ??????????????????????????????? maybe

/*
	} else {
		// there was no match yet, so continue to end of array
		i++
		// fmt.Println("continue")
		continue // only continue to end of array, then "break"
	}

*/
/*

	//
	/* chat suggested:
	func pick_aCard_and_assign_fields() {
	    for i, lastPull := range cyclicArrayPulls.pulls {
	        if lastPull == promptField {
	            // If a match is found, set a new promptField and return
	            promptField = getNewPromptField()
	            return
	        }
	    }
	}

*/
//
//
//
/*
	// fmt.Printf("acquired a new promptField1:%s\n", promptField)
	// now that we have a newly-acquired promptField, check if we have seen it recently
	if recursion {
		// do nothing, i.e., skip the else block because only one recursion is desired
		fmt.Println("recursion detected")
		from_recursion = true
	} else { // do this block only if here not by recursion
		// fmt.Printf("recursion should happen first, and should be false:%v\n", recursion)
		i = 0
		for i < len(cyclicArrayPulls.pulls) {
			lastPull = cyclicArrayPulls.pulls[i]
			// fmt.Printf("pulling/pulled:%s, OldPrompt is:%s, promptField is:%s\n", lastPull, oldPrompt, promptField)
			if lastPull == promptField { // seen it, so reset i, get new promptField, i.e. repeat
				fmt.Printf("seen it before as lastPull:%s- and promptField:%s- \n", lastPull, promptField)
				// seen it, so acquire a new promptField, having reset index so as to search array anew
				i = 0
				from_recursion = true
				pick_aCard_and_assign_fields(true) // get yet another new promptField, just once!!!!!
				return

			} else {
				// there was no match yet, so continue to end of array
				i++
				// fmt.Println("continue")
				continue // only continue to end of array, then "break"
			}
		}
	}

*/

func hits() { // - - in comments there are
	// Create maps to store the frequency of each relevant string for that map
	frequencyMapRightOrOops := make(map[string]int)
	frequencyMapChar := make(map[string]int) // These, apparently, create a map to associate a unique string with an int
	frequencyMapWrongs := make(map[string]int)

	//
	// Parse the relevant cyclic array to check_for_match_in_secondary_field the strings and put them into the relevant map:
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

func check_for_match_in_secondary_field(in string) (found_one bool) {
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
