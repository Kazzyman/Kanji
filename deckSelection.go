package main

import "math/rand"

func pick_RandomCard_Assign_fields() (promptField, objective, objective_kind string) { // - -
	// Random prompting from initiate deck
	//
	/*
		randIndex := rand.Intn(len(fileOfCardsInitiate))
		aCard = fileOfCardsInitiate[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kanji
		objective = aCard.Meaning
		objective_kind = "Romaji"
	*/

	// Random prompting from fileOfAllCardsSequentially deck
	//
	randIndex := rand.Intn(len(fileOfAllCardsSequentially))
	aCard = fileOfAllCardsSequentially[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
	promptField = aCard.Kanji
	objective = aCard.Meaning
	objective_kind = "Romaji"

	/*
		// Sequential prompting from the manually semi-randomized deck
		//
		for {
			if index < len(fileOfCards) {
				aCard = fileOfCards[index] // pick a 'card' from a 'deck' and store it in a global var
				promptField = aCard.Kanji
				objective = aCard.Meaning
				index++
				break
			} else {
				index = 0
				continue
			}
		}

	*/

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
