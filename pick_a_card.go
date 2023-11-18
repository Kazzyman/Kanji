package main

import (
	"fmt"
	"math/rand"
)

func pick_RandomCard_Assign_fields() (promptField, objective, objective_kind, secondary_objective string) { // - -
	if current_deck == "fresh" {
		for { // This for loop is only needed to check for empty cards
			randIndex := rand.Intn(len(fileOf_fresh))
			aCard = fileOf_fresh[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
			promptField = aCard.Kanji
			objective = aCard.Meaning
			secondary_objective = aCard.Second_Meaning
			if promptField == "" || promptField == " " {
				fmt.Printf("%s--- An empty kanji card was skipped in fresh!!!!%s\n", colorRed, colorReset) // Verified
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}

	if current_deck == "init" {
		for { // This for loop is only needed to check for empty cards
			randIndex := rand.Intn(len(fileOfCardsInitiate))
			aCard = fileOfCardsInitiate[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
			promptField = aCard.Kanji
			objective = aCard.Meaning
			secondary_objective = aCard.Second_Meaning
			if promptField == "" || promptField == " " {
				fmt.Printf("%s--- An empty kanji card was skipped in init!!!!%s\n", colorRed, colorReset) // Verified
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}

	if current_deck == "nov" {
		for { // This for loop is only needed to check for empty cards
			randIndex := rand.Intn(len(fileOfCardsNovice))
			aCard = fileOfCardsNovice[randIndex]
			promptField = aCard.Kanji
			objective = aCard.Meaning
			secondary_objective = aCard.Second_Meaning
			if promptField == "" || promptField == " " {
				fmt.Printf("%s--- An empty kanji card was skipped i nov!!!!%s\n", colorRed, colorReset) // Verified
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}
	/*
		if current_deck == "grad" {
			for { // This for loop is only needed to check for empty cards
				randIndex := rand.Intn(len(fileOfCardsGraduate))
				aCard = fileOfCardsGraduate[randIndex]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
				if promptField == "" || promptField == " " {
					fmt.Printf("%s--- An empty kanji card was skipped in grad!!!!%s\n", colorRed, colorReset) // Verified
					continue
				} else {
					break // break out of local for loop and naturally fall-through to the return
				}
			}
		}

	*/
	if current_deck == "grad" {
		for { // This for loop is only needed to check for empty cards
			randIndex := rand.Intn(len(data_file100_maybe))
			aCard = data_file100_maybe[randIndex]
			promptField = aCard.Kanji
			objective = aCard.Meaning
			secondary_objective = aCard.Second_Meaning
			if promptField == "" || promptField == " " {
				fmt.Printf("%s--- An empty kanji card was skipped in grad!!!!%s\n", colorRed, colorReset) // Verified
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}

	if current_deck == "mast" {
		for { // This for loop is only needed to check for empty cards
			randIndex := rand.Intn(len(fileOfCardsMaster))
			aCard = fileOfCardsMaster[randIndex]
			promptField = aCard.Kanji
			objective = aCard.Meaning
			secondary_objective = aCard.Second_Meaning
			if promptField == "" || promptField == " " {
				fmt.Printf("%s--- An empty kanji card was skipped in mast!!!!%s\n", colorRed, colorReset) // Verified
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}

	if current_deck == "guru" {
		for { // This for loop is only needed to check for empty cards
			randIndex := rand.Intn(len(fileOfCardsGuru))
			aCard = fileOfCardsGuru[randIndex]
			promptField = aCard.Kanji
			objective = aCard.Meaning
			secondary_objective = aCard.Second_Meaning
			if promptField == "" || promptField == " " {
				fmt.Printf("%s--- An empty kanji card was skipped in guru!!!!%s\n", colorRed, colorReset) // Verified
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}

	if current_deck == "current" {
		for { // This for loop is only needed to check for empty cards
			randIndex := rand.Intn(len(fileOf_Current))
			aCard = fileOf_Current[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
			promptField = aCard.Kanji
			objective = aCard.Meaning
			secondary_objective = aCard.Second_Meaning
			if promptField == "" || promptField == " " {
				fmt.Printf("%s--- An empty kanji card was skipped in current!!!!%s\n", colorRed, colorReset) // Verified
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}

	if current_deck == "data" {
		for { // This for loop is only needed to check for empty cards
			randIndex := rand.Intn(len(data_file))
			aCard = data_file[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
			promptField = aCard.Kanji
			objective = aCard.Meaning
			secondary_objective = aCard.Second_Meaning
			if promptField == "" || promptField == " " {
				fmt.Printf("%s--- An empty kanji card was skipped in data!!!!%s\n", colorRed, colorReset) // Verified
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}

	if current_deck == "all" {
		// Randomize over all decks, i.e., Randomly chose a deck
		// fmt.Printf("\n random number is:%d, and must include 0-6 inclusive\n", randDeckAndMode)
		//
		for { // This for loop is only needed to check for empty cards picked from any random deck
			randDeckAndMode = rand.Intn(7) // if an empty card was picked, maybe try a different deck, or not
			if randDeckAndMode == 0 {
				current_deck_B = "init"
				randIndex := rand.Intn(len(fileOfCardsInitiate))
				aCard = fileOfCardsInitiate[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
			} else if randDeckAndMode == 1 {
				current_deck_B = "nov"
				randIndex := rand.Intn(len(fileOfCardsNovice))
				aCard = fileOfCardsNovice[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
			} else if randDeckAndMode == 2 {
				current_deck_B = "grad"
				randIndex := rand.Intn(len(fileOfCardsGraduate))
				aCard = fileOfCardsGraduate[randIndex]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
			} else if randDeckAndMode == 3 {
				current_deck_B = "mast"
				randIndex := rand.Intn(len(fileOfCardsMaster))
				aCard = fileOfCardsMaster[randIndex]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
			} else if randDeckAndMode == 4 {
				current_deck_B = "current"
				randIndex := rand.Intn(len(fileOf_Current))
				aCard = fileOf_Current[randIndex]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
			} else if randDeckAndMode == 5 {
				current_deck_B = "fresh"
				randIndex := rand.Intn(len(fileOf_fresh))
				aCard = fileOf_fresh[randIndex]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
			} else if randDeckAndMode == 6 {
				current_deck_B = "guru"
				randIndex := rand.Intn(len(fileOfCardsGuru))
				aCard = fileOfCardsGuru[randIndex]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				secondary_objective = aCard.Second_Meaning
			}
			if promptField == "" || promptField == " " { // if the prior conditional picks an empty card, pick another
				fmt.Printf("%s--- An empty kanji card was skipped in %s:%s !!!!%s\n",
					colorRed, current_deck, current_deck_B, colorReset) // Verified
				continue
			} else {
				break // break out of local for loop (in 'all') and naturally fall-through to the return
			}
		}
	}
	objective_kind = "Romaji"                                          // i.e., "Meaning"
	return promptField, objective, objective_kind, secondary_objective // returns the values of the global vars
}
