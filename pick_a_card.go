package main

import (
	"fmt"
	"math/rand"
)

// var lastPull string
var frequencyMapOfSeenChars = make(map[string]int)

func randomize_over_all_decks() (promptField, objective, objective_kind, secondary_objective string) {
	// Randomize over all decks and modes (except Guru)  (using global values of the indices to maintain state)
	// Randomly chose a deck and a mode
	randDeckAndMode = rand.Intn(8)
	//
	// Randomly access cards // deck and mode 0-3
	if randDeckAndMode == 0 {
		current_deckA = "init"
		randIndex := rand.Intn(len(fileOfCardsInitiate))
		aCard = fileOfCardsInitiate[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if randDeckAndMode == 1 {
		current_deckA = "nov"
		randIndex := rand.Intn(len(fileOfCardsNovice))
		aCard = fileOfCardsNovice[randIndex]
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if randDeckAndMode == 2 {
		current_deckA = "grad"
		randIndex := rand.Intn(len(fileOfCardsGraduate))
		aCard = fileOfCardsGraduate[randIndex]
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if randDeckAndMode == 3 {
		current_deckA = "mast"
		randIndex := rand.Intn(len(fileOfCardsMaster))
		aCard = fileOfCardsMaster[randIndex]
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if randDeckAndMode == 4 {
		//
		// Randomly access cards // deck and mode 4-7
		current_deckA = "inits"
		for {
			if indexInitS < len(fileOfCardsInitiate) {
				aCard = fileOfCardsInitiate[indexInitS]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
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
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
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
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
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
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
				indexMastS++
				break
			} else {
				indexMastS = 0
				continue
			}
		}
	}
	return
}

// recursion bool arg no longer needed or used
func pick_aCard_and_assign_fields() (promptField, objective, objective_kind, secondary_objective string) { // - -
	/*
		if cyclicArrayPulls.pulls[0] == "" {
			cyclicArrayPulls.InsertKChar("primedK0") // Prime the array
		}

	*/
	//
	if current_deck == "randAll" {
		// Get the first promptField value (and also the values that need to be returned with it)
		promptField, objective, objective_kind, secondary_objective = randomize_over_all_decks()
		// Now that we have a newly-acquired promptField, check to see if we have it stored in the cyclicArrayPulls slice ...
		// and, if it is so stored, obtain a replacement, and then look again through the entire slice. Repeat the entire
		// process as many times as may be required to finally obtain a value of promptField which is novel according to said slice.
		for {
			found := false
			for _, lastPull := range cyclicArrayPulls.pulls {
				if lastPull == promptField {
					// We also wish to store these duplicates in the map, to keep a tally of such events -- accessible via the rm Dir
					frequencyMapOfSeenChars[promptField]++ // The '++' increments the int value associated with promptField
					fmt.Printf("We've seen the pseudo-random char before lastPull: %s and promptField: %s\n", lastPull, promptField)
					found = true
					promptField, objective, objective_kind, secondary_objective = randomize_over_all_decks()
					break // Exit the inner loop, having a new and potentially novel promptField in hand
				}
			}

			if !found {
				// If promptField is not found in the slice ...
				break // All of our work here being done! We hereby exit the outer naked for loop
			} else {
				// A match WAS found (and a new promptField value has therefore been obtained; so, we need to restart the entire process)
				continue // explicitly continue, i.e., restart range with the replacement value of promptField
			}
		}

		// Exited the loops having obtained a newly-acquired, and verified as novel, value of promptField, so, ...
		// ... store that newly-acquired promptField in an array to be used as a memory of already-seen-and-used chars
		cyclicArrayPulls.InsertKChar(promptField) // Only the chars that are actually used are put into the array
		// Also, store that newly-acquired promptField in frequencyMapOfSeenChars
		frequencyMapOfSeenChars[promptField]++ // The '++' increments the int value associated with promptField
		return
	}
	/*
		*
		A stack overflow error occurs when the program's call stack becomes too deep, often due to recursive function calls without
		a proper base case or termination condition.
	*/
	//
	//
	//
	//
	// Iterate over decks, and then randomize each deck (using global values of the indices to maintain state)
	if current_deck == "rand" {
		if randDeck == 0 {
			current_deckA = "init"
			randIndex := rand.Intn(len(fileOfCardsInitiate))
			aCard = fileOfCardsInitiate[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
			promptField = aCard.Kanji
			objective = aCard.Meaning
			secondary_objective = aCard.Second_Meaning
			randDeck++
		} else if randDeck == 1 {
			current_deckA = "nov"
			randIndex := rand.Intn(len(fileOfCardsNovice))
			aCard = fileOfCardsNovice[randIndex]
			promptField = aCard.Kanji
			objective = aCard.Meaning
			secondary_objective = aCard.Second_Meaning
			randDeck++
		} else if randDeck == 2 {
			current_deckA = "grad"
			randIndex := rand.Intn(len(fileOfCardsGraduate))
			aCard = fileOfCardsGraduate[randIndex]
			promptField = aCard.Kanji
			objective = aCard.Meaning
			secondary_objective = aCard.Second_Meaning
			randDeck++
		} else if randDeck == 3 {
			current_deckA = "mast"
			randIndex := rand.Intn(len(fileOfCardsMaster))
			aCard = fileOfCardsMaster[randIndex]
			promptField = aCard.Kanji
			objective = aCard.Meaning
			secondary_objective = aCard.Second_Meaning
			// randDeck++
			randDeck = 0
		} else {
			/*
				current_deckA = "guru"
				randIndex := rand.Intn(len(fileOfCardsGuru))
				aCard = fileOfCardsGuru[randIndex]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning

			*/
			randDeck = 0
		}
	}

	// Iterate over decks, and then sequentially read all cards from each deck (using global values of the indices to maintain state)
	if current_deck == "all" {
		if forLoop == 0 {
			current_deckA = "inits"
			for {
				if indexInitS < len(fileOfCardsInitiate) {
					aCard = fileOfCardsInitiate[indexInitS]
					promptField = aCard.Kanji
					objective = aCard.Meaning
					secondary_objective = aCard.Second_Meaning
					indexInitS++
					forLoop++
					break // We only pull one card, and then exit
				} else {
					indexInitS = 0
					continue
				}
			}
		} else if forLoop == 1 {
			current_deckA = "novs"
			// Sequential prompting
			//
			for {
				if indexNovS < len(fileOfCardsNovice) {
					aCard = fileOfCardsNovice[indexNovS]
					promptField = aCard.Kanji
					objective = aCard.Meaning
					secondary_objective = aCard.Second_Meaning
					indexNovS++
					forLoop++
					break
				} else {
					indexNovS = 0
					continue
				}
			}
		} else if forLoop == 2 {
			current_deckA = "grads"
			// Sequential prompting
			//
			for {
				if indexGradS < len(fileOfCardsGraduate) {
					aCard = fileOfCardsGraduate[indexGradS]
					promptField = aCard.Kanji
					objective = aCard.Meaning
					secondary_objective = aCard.Second_Meaning
					indexGradS++
					forLoop++
					break
				} else {
					indexGradS = 0
					continue
				}
			}
		} else if forLoop == 3 {
			current_deckA = "masts"
			// Sequential prompting
			//
			for {
				if indexMastS < len(fileOfCardsMaster) {
					aCard = fileOfCardsMaster[indexMastS]
					promptField = aCard.Kanji
					objective = aCard.Meaning
					secondary_objective = aCard.Second_Meaning
					indexMastS++
					// forLoop++
					forLoop = 0
					break
				} else {
					indexMastS = 0
					continue
				}
			}
		} else {
			/*
				current_deckA = "gurus"
				// Sequential prompting
				//
				for {
					if indexGuruS < len(fileOfCardsGuru) {
						aCard = fileOfCardsGuru[indexGuruS]
						promptField = aCard.Kanji
						objective = aCard.Meaning
						secondary_objective = aCard.Second_Meaning
						indexGuruS++
						forLoop = 0
						break
					} else {
						indexGuruS = 0
						continue
					}
				}

			*/
		}
		objective_kind = "Romaji" // i.e., "Meaning"
	}

	// Random prompting from selected deck
	//
	if current_deck == "init" {
		randIndex := rand.Intn(len(fileOfCardsInitiate))
		aCard = fileOfCardsInitiate[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if current_deck == "nov" {
		randIndex := rand.Intn(len(fileOfCardsNovice))
		aCard = fileOfCardsNovice[randIndex]
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if current_deck == "grad" {
		randIndex := rand.Intn(len(fileOfCardsGraduate))
		aCard = fileOfCardsGraduate[randIndex]
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if current_deck == "mast" {
		randIndex := rand.Intn(len(fileOfCardsMaster))
		aCard = fileOfCardsMaster[randIndex]
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if current_deck == "guru" {
		randIndex := rand.Intn(len(fileOfCardsGuru))
		aCard = fileOfCardsGuru[randIndex]
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if current_deck == "init" {
		randIndex := rand.Intn(len(fileOfCardsInitiate))
		aCard = fileOfCardsInitiate[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if current_deck == "inits" {
		// Sequential prompting
		//
		for {
			if indexInitS < len(fileOfCardsInitiate) {
				aCard = fileOfCardsInitiate[indexInitS]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
				indexInitS++
				break
			} else {
				indexInitS = 0
				continue
			}
		}
	} else if current_deck == "novs" {
		// Sequential prompting
		//
		for {
			if indexNovS < len(fileOfCardsNovice) {
				aCard = fileOfCardsNovice[indexNovS]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
				indexNovS++
				break
			} else {
				indexNovS = 0
				continue
			}
		}
	} else if current_deck == "grads" {
		// Sequential prompting
		//
		for {
			if indexGradS < len(fileOfCardsGraduate) {
				aCard = fileOfCardsGraduate[indexGradS]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
				indexGradS++
				break
			} else {
				indexGradS = 0
				continue
			}
		}
	} else if current_deck == "masts" {
		// Sequential prompting
		//
		for {
			if indexMastS < len(fileOfCardsMaster) {
				aCard = fileOfCardsMaster[indexMastS]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
				indexMastS++
				break
			} else {
				indexMastS = 0
				continue
			}
		}
	} else if current_deck == "gurus" {
		// Sequential prompting
		//
		for {
			if indexGuruS < len(fileOfCardsGuru) {
				aCard = fileOfCardsGuru[indexGuruS]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
				indexGuruS++
				break
			} else {
				indexGuruS = 0
				continue
			}
		}
	}
	objective_kind = "Romaji" // i.e., "Meaning"

	/*
		//  Random prompting from all decks
		//
			randIndex := rand.Intn(len(fileOfCardsK))
			randIndexS := rand.Intn(len(fileOfCardsGuru2))
			randIndexD := rand.Intn(len(fileOfCardsGuru))

			randomFileOfCards = rand.Intn(3)

			// Kanji prompting, Meaning objective:
			if randomFileOfCards == 0 {
				aCard = fileOfCardsK[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
				promptField = aCard.Kanji
				objective = aCard.Meaning
			}
			if randomFileOfCards == 1 {
				aCard = fileOfCardsGuru2[randIndexS] // Randomly pick a 'card' from a 'deck' and store it in a global var
				promptField = aCard.Kanji
				objective = aCard.Meaning
			}
			if randomFileOfCards == 2 {
				aCard = fileOfCardsGuru[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
				promptField = aCard.Kanji
				objective = aCard.Meaning
			}

	*/

	return promptField, objective, objective_kind, secondary_objective

}
