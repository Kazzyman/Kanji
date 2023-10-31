package main

import "math/rand"

func pick_aCard_and_assign_fields() (promptField, objective, objective_kind, secondary_objective string) { // - -

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
	}

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
