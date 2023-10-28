package main

import "math/rand"

func pick_RandomCard_Assign_fields() (promptField, objective, objective_kind, secondary_objective string) { // - -
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
