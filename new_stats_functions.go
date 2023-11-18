package main

import (
	"fmt"
	"os"
	"strings"
)

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
			fmt.Printf(colorRed)
			fmt.Printf("Meanings: %s, %s ", cardInfoData.FirstMeaningOnRecord, cardInfoData.SecondMeaningOnRecord)
			fmt.Printf(colorCyan)
			fmt.Printf("Freq:")
			fmt.Printf(colorReset)
			fmt.Printf(" %d\n", cardInfoData.CorrectGuessCount)
		}
	}
	fmt.Printf("Number of unique chars: ")
	fmt.Printf(colorPurple)
	fmt.Printf("%d \n\n", numberOfUniqueKanjiCharsHit)
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

var cardCount int
var position int
var secondField charSetStructKanji
var already_used_map2 = make(map[string]int)

func test1() {
	/*
		for position, secondField = range data_file {
			// ^ ^ ^ read data_file :: [for position, secondField = range data_file {]
			// v v v read through the ENTIRE already_used_map2 :: [for length <= len(already_used_map2) { length++ ]
			for length <= len(already_used_map2) {
				if secondField.Kanji is not already in already_used_map2 { // if is_pick_novel2(secondField.Kanji) {
					already_used_map2[secondField.Kanji]++ // then put it in
					fmt.Fprintf(output_file, "%s was unique\n", secondField.Kanji) //
				} else { // it was in the map already, so make a note of it:
					fmt.Fprintf(output_file2, "%s was not unique\n", secondField.Kanji)
				}
			}
		}
	*/
	// length := 0
	output_file, _ := os.OpenFile("all_unique_Cards.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	output_file2, _ := os.OpenFile("all_duplicate_Cards.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	for position, secondField = range data_file { // *
		cardCount++
		// for length <= len(already_used_map2) { // *
		// length++ // *
		if is_pick_novel2(secondField.Kanji) { // * this should ALREADY be reading the ENTIRE map
			already_used_map2[secondField.Kanji]++ // * having read the entire map, put it in
			// break // keep it and print it
			// fmt.Printf("length is:%d\n", length)
			// fmt.Fprintf(output_file, "%d was length\n", length)
			fmt.Fprintf(output_file, "%s was unique\n", secondField.Kanji) // *
			// break *
		} else {
			// fmt.Fprintf(output_file2, "%s was not in the map at position:%d\n", secondField.Kanji, length) // Verified
			fmt.Fprintf(output_file2, "%s was in the map\n", secondField.Kanji)
			// length = 0 // continue the loop ensuring that the entire map is read with this new pick
			// continue // redundant
		}
		// }
	}
	// fmt.Fprintf(output_file, "%s was also unique\n", secondField.Kanji)
	fmt.Printf("Count of cards is:%d\n", cardCount)
	fmt.Printf("position:%d, secondField:%s\n", position, secondField.Kanji)
	already_used_map2 = make(map[string]int)
}
func is_pick_novel2(kanji string) bool {
	for stringFromMap, freqInMap = range already_used_map2 {
		// fmt.Printf("in is_pick_novel str:%s, f:%d \n", stringFromMap, freqInMap)
		if kanji == stringFromMap {
			fmt.Printf("returning false, picked was:%s, map had:%s\n", kanji, stringFromMap)
			// Having either found the pick in the map prior to reading the entire map, or as the map's last element ...
			return false
		}
	} // having read the entire map, and not having found a previously used card in said map ...
	fmt.Printf("returning true, data was:%s, map had:%s\n", kanji, stringFromMap)
	return true
}

func test2() {
	/*
		for position, secondField = range data_file {
			// ^ ^ ^ read data_file :: [for position, secondField = range data_file {]
			// v v v read through the ENTIRE already_used_map2 :: [for length <= len(already_used_map2) { length++ ]
			for length <= len(already_used_map2) {
				if secondField.Kanji is not already in already_used_map2 { // if is_pick_novel2(secondField.Kanji) {
					already_used_map2[secondField.Kanji]++ // then put it in
					fmt.Fprintf(output_file, "%s was unique\n", secondField.Kanji) //
				} else { // it was in the map already, so make a note of it:
					fmt.Fprintf(output_file2, "%s was not unique\n", secondField.Kanji)
				}
			}
		}
	*/
	// length := 0
	output_file, _ := os.OpenFile("all_unique_Cards.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	output_file2, _ := os.OpenFile("all_duplicate_Cards.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	for position, secondField = range data_file100_maybe { // The input dataset
		cardCount++
		// for length <= len(already_used_map2) { // *
		// length++ // *
		if is_pick_novel2(secondField.Kanji) { // * this should ALREADY be reading the ENTIRE map
			already_used_map2[secondField.Kanji]++ // * having read the entire map, put it in
			// break // keep it and print it
			// fmt.Printf("length is:%d\n", length)
			// fmt.Fprintf(output_file, "%d was length\n", length)
			fmt.Fprintf(output_file, "%s was unique\n", secondField.Kanji) // *
			// break *
		} else {
			// fmt.Fprintf(output_file2, "%s was not in the map at position:%d\n", secondField.Kanji, length) // Verified
			fmt.Fprintf(output_file2, "%s was in the map\n", secondField.Kanji)
			// length = 0 // continue the loop ensuring that the entire map is read with this new pick
			// continue // redundant
		}
		// }
	}
	// fmt.Fprintf(output_file, "%s was also unique\n", secondField.Kanji)
	fmt.Printf("Count of cards is:%d\n", cardCount)
	fmt.Printf("position:%d, secondField:%s\n", position, secondField.Kanji)
	already_used_map2 = make(map[string]int)
}
