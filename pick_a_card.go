package main

import "math/rand"

func randomize_over_all_decks() (promptField, objective, objective_kind, secondary_objective string) {
	// Randomize over all decks (except Guru)
	// Randomly chose a deck
	randDeckAndMode = rand.Intn(6)
	// fmt.Printf("\n random number is:%d, and must include 0-5 inclusive\n", randDeckAndMode)
	//
	if randDeckAndMode == 0 {
		current_deck = "init"
		randIndex := rand.Intn(len(fileOfCardsInitiate))
		aCard = fileOfCardsInitiate[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if randDeckAndMode == 1 {
		current_deck = "nov"
		randIndex := rand.Intn(len(fileOfCardsNovice))
		aCard = fileOfCardsNovice[randIndex]
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if randDeckAndMode == 2 {
		current_deck = "grad"
		randIndex := rand.Intn(len(fileOfCardsGraduate))
		aCard = fileOfCardsGraduate[randIndex]
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if randDeckAndMode == 3 {
		current_deck = "mast"
		randIndex := rand.Intn(len(fileOfCardsMaster))
		aCard = fileOfCardsMaster[randIndex]
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if randDeckAndMode == 4 {
		current_deck = "current"
		randIndex := rand.Intn(len(fileOf_Current))
		aCard = fileOf_Current[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	} else if randDeckAndMode == 5 {
		current_deck = "fresh"
		randIndex := rand.Intn(len(fileOf_fresh))
		aCard = fileOf_fresh[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kanji
		objective = aCard.Meaning
		secondary_objective = aCard.Second_Meaning
	}
	return
}

func randomize_over_current() (promptField, objective, objective_kind, secondary_objective string) {
	current_deck = "current"
	randIndex := rand.Intn(len(fileOf_Current))
	aCard = fileOf_Current[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
	promptField = aCard.Kanji
	objective = aCard.Meaning
	secondary_objective = aCard.Second_Meaning
	return
}

func randomize_over_fresh() (promptField, objective, objective_kind, secondary_objective string) {
	current_deck = "fresh"
	// for {
	randIndex := rand.Intn(len(fileOf_fresh))
	aCard = fileOf_fresh[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
	promptField = aCard.Kanji
	objective = aCard.Meaning
	secondary_objective = aCard.Second_Meaning
	/*
		if promptField == "" {
			// // fmt.Println("an empty kanji card was found") // Verified
			continue
		} else {
			break
		}

	*/
	// }
	return
}

func pick_RandomCard_Assign_fields() (promptField, objective, objective_kind, secondary_objective string) { // - -

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
	} else if current_deck == "fresh" {
		promptField, objective, objective_kind, secondary_objective = randomize_over_fresh()
	} else if current_deck == "current" {
		promptField, objective, objective_kind, secondary_objective = randomize_over_current()
	} else if current_deck == "all" {
		promptField, objective, objective_kind, secondary_objective = randomize_over_all_decks()
	}
	objective_kind = "Romaji" // i.e., "Meaning"
	return promptField, objective, objective_kind, secondary_objective
}
