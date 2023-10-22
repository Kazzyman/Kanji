package main

import "math/rand"

func pick_RandomCard_Assign_fields() (promptField, objective, objective_kind string) { // - -
	// Random prompting from selected deck
	//
	if current_deck == "init" {
		randIndex := rand.Intn(len(fileOfCardsInitiate))
		aCard = fileOfCardsInitiate[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kanji
		objective = aCard.Meaning
	} else if current_deck == "nov" {
		randIndex := rand.Intn(len(fileOfCardsNovice))
		aCard = fileOfCardsNovice[randIndex]
		promptField = aCard.Kanji
		objective = aCard.Meaning
	} else if current_deck == "grad" {
		randIndex := rand.Intn(len(fileOfCardsGraduate))
		aCard = fileOfCardsGraduate[randIndex]
		promptField = aCard.Kanji
		objective = aCard.Meaning
	} else if current_deck == "mast" {
		randIndex := rand.Intn(len(fileOfCardsMaster))
		aCard = fileOfCardsMaster[randIndex]
		promptField = aCard.Kanji
		objective = aCard.Meaning
	} else if current_deck == "guru" {
		randIndex := rand.Intn(len(fileOfCardsGuru))
		aCard = fileOfCardsGuru[randIndex]
		promptField = aCard.Kanji
		objective = aCard.Meaning
	} else if current_deck == "init" {
		randIndex := rand.Intn(len(fileOfCardsInitiate))
		aCard = fileOfCardsInitiate[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kanji
		objective = aCard.Meaning
	} else if current_deck == "novs" {
		// Sequential prompting
		//
		for {
			if index < len(fileOfCardsNovice) {
				aCard = fileOfCardsNovice[index]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				index++
				break
			} else {
				index = 0
				continue
			}
		}
	} else if current_deck == "grads" {
		// Sequential prompting
		//
		for {
			if index < len(fileOfCardsGraduate) {
				aCard = fileOfCardsGraduate[index]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				index++
				break
			} else {
				index = 0
				continue
			}
		}
	} else if current_deck == "masts" {
		// Sequential prompting
		//
		for {
			if index < len(fileOfCardsMaster) {
				aCard = fileOfCardsMaster[index]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				index++
				break
			} else {
				index = 0
				continue
			}
		}
	} else if current_deck == "gurus" {
		// Sequential prompting
		//
		for {
			if index < len(fileOfCardsGuru) {
				aCard = fileOfCardsGuru[index]
				promptField = aCard.Kanji
				objective = aCard.Meaning
				index++
				break
			} else {
				index = 0
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

	return promptField, objective, objective_kind

}
