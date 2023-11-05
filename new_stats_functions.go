package main

import (
	"fmt"
	"strings"
)

func test1() {

}

func test2() {

}

// Create a kanjiHitMap map. A map of keys (kanji chars) and/to their associated history data, i.e., the struct: CardInfo
var kanjiHitMap = make(map[string]CardInfo)

// The above is a map of the things below, which is keyed by a kanji char string

// The type of card which will be included in the above map
type CardInfo struct {
	UsersGuess            []string
	FirstMeaningOnRecord  string
	SecondMeaningOnRecord string
	CorrectGuessCount     int
}

func recordGuess(kanji, users_Guess, Meaning_on_record, second_meaning string) {
	// Check if the kanji exists in the map, not in the CardInfo struct, but in the map: kanjiHitMap
	cardInfo, exists := kanjiHitMap[kanji] // Use kanji as key to kanjiHitMap map
	// fmt.Printf("cardInfo is:%v, and exists is:%v \n", cardInfo, exists)
	// cardInfo is a struct object of type CardInfo. [kanji] is FROM the signature of this func, and will be a kanji string which is the key to the map
	if !exists {
		// If a card by the name[kanji] doesn't already exist, create a new blank card of type CardInfo
		cardInfo = CardInfo{} // A new blank card which we will append to (or add to) below
	}

	// Append the user's guess [1] and the meaning that we have on record [2] to either the new blank card or the existing card
	// ... oh, and also increment the count [3]
	cardInfo.UsersGuess = append(cardInfo.UsersGuess, users_Guess) // [1]
	if cardInfo.FirstMeaningOnRecord == "" {
		cardInfo.FirstMeaningOnRecord = Meaning_on_record // [2]
		if cardInfo.SecondMeaningOnRecord == "" {
			cardInfo.SecondMeaningOnRecord = second_meaning
		}
	}
	cardInfo.CorrectGuessCount++ // [3]

	// Update the map with the new or updated cardInfo
	kanjiHitMap[kanji] = cardInfo
}

func newHits() {
	// Create maps to store the frequency of each relevant string for that map
	frequencyMapRightOrOops := make(map[string]int)
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

	// Load the wrongs frequency map // As documented in the foregoing 'for' loop
	for i := 0; i < len(cyclicArrayOfTheJcharsGottenWrong.jchar); i++ {
		str := cyclicArrayOfTheJcharsGottenWrong.jchar[i]
		frequencyMapWrongs[str]++
	}

	// -- PRINT -- 'Right' and its frequency
	for str, freq := range frequencyMapRightOrOops { // The map has only one entry for Right, & one for Oops
		if str == "Right" { // Finds the one potential entry for Right
			fmt.Printf(colorGreen)
			fmt.Printf("\n%s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorGreen)
			fmt.Printf(" %d\n", freq) // Frequency of Right, per the map
			fmt.Printf(colorReset)
			total_prompts = freq
		} else if str == "" {
			// else, it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		}
	}
	//
	//
	// Print the unique Chars and their frequencies (continuing the printout above)
	numberOfUniqueKanjiCharsHit := 0
	for kanjiString, cardInfoData := range kanjiHitMap {
		if cardInfoData.CorrectGuessCount == 0 {
			// it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		} else {
			numberOfUniqueKanjiCharsHit++
			fmt.Printf(" %s ", kanjiString)
			fmt.Printf(colorGreen)
			fmt.Printf("Your guesses:%s, ", cardInfoData.UsersGuess)
			fmt.Printf(colorYellow)
			fmt.Printf("Meanings: %s, %s ", cardInfoData.FirstMeaningOnRecord, cardInfoData.SecondMeaningOnRecord)
			fmt.Printf(colorCyan)
			fmt.Printf("Freq:")
			fmt.Printf(colorReset)
			fmt.Printf(" %d\n", cardInfoData.CorrectGuessCount)
		}
	}
	fmt.Printf("Number of unique chars: ")
	fmt.Printf(colorPurple)
	fmt.Printf("%d \n", numberOfUniqueKanjiCharsHit)
	fmt.Printf(colorReset)

	fmt.Printf("Total prompts:")
	fmt.Printf(colorRed)
	fmt.Printf(" %d\n", total_prompts)
	fmt.Printf(colorReset)

	// -- PRINT -- 'Oops' and its frequency
	for str, freq := range frequencyMapRightOrOops { // The map has only one entry for Right, & one for Oops
		if str == "Oops" { // Finds the one potential entry for Oops
			fmt.Printf(colorRed)
			fmt.Printf("%s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorRed)
			fmt.Printf(" %d\n", freq) // Frequency of Oops, per the map
			fmt.Printf(colorReset)
			total_prompts = total_prompts + freq
		} else if str == "" {
			// else, it is an 'empty' position in the map due to empty uninitialized position in the cyclic array
		}
	}
	//
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
