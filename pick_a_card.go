package main

import (
	"fmt"
	"math/rand"
)

func pick_RandomCard_Assign_fields() { // - -
	if current_deck == "fresh" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(fileOf_fresh))
			aCard = fileOf_fresh[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}
	if current_deck == "init" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(fileOfCardsInitiate))
			aCard = fileOfCardsInitiate[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}
	if current_deck == "grad" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(fileOfCardsGraduate))
			aCard = fileOfCardsGraduate[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}
	if current_deck == "nov" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(fileOfCardsNovice))
			aCard = fileOfCardsNovice[randIndex]
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}
	if current_deck == "claude" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(claude))
			aCard = claude[randIndex]
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}
	if current_deck == "mast" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(fileOfCardsMaster))
			aCard = fileOfCardsMaster[randIndex]
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}
	if current_deck == "guru" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(fileOfCardsGuru))
			aCard = fileOfCardsGuru[randIndex]
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}
	if current_deck == "current" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(fileOf_Current))
			aCard = fileOf_Current[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}

	// data deck: data_data_184cards.go
	if current_deck == "data" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(data_file))
			aCard = data_file[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}

	// advanced
	if current_deck == "beauty" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(data_beauty))
			aCard = data_beauty[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}

	// alternate "word problems"
	if current_deck == "words" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(fileOf_Words))
			aCard = fileOf_Words[randIndex]    // Randomly pick a 'card' from a 'deck' and store it in a global var
			actual_prompt_string = aCard.Kanji // English meaning
			primary_meaning = aCard.Meaning    // the Japanese word
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}

	// unused deck
	if current_deck == "2098_lines_of_cards" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(claudeOld))
			aCard = claudeOld[randIndex]
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}

	// fileOf__Quiz 1
	if current_deck == "quiz_novice" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(fileOf_Quiz_novice))
			aCard = fileOf_Quiz_novice[randIndex]
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}
	// fileOf__Quiz 2
	if current_deck == "quiz_comp" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(fileOf_Quiz_comp))
			aCard = fileOf_Quiz_comp[randIndex]
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	} // fileOf__Quiz 3
	if current_deck == "quiz_master" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(fileOf_Quiz_master))
			aCard = fileOf_Quiz_master[randIndex]
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	} // fileOf__Quiz 4
	if current_deck == "quiz_guru" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(fileOf_Quiz_guru))
			aCard = fileOf_Quiz_guru[randIndex]
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
				continue
			} else {
				break // break out of local for loop and naturally fall-through to the return
			}
		}
	}

	// fileOf__Quiz 5
	if current_deck == "quiz_japan" {
		for { // This for loop assigns fields, and is needed to handle empty cards
			randIndex := rand.Intn(len(fileOf_Quiz_japan))
			aCard = fileOf_Quiz_japan[randIndex]
			actual_prompt_string = aCard.Kanji
			primary_meaning = aCard.Meaning
			secondary_meaning = aCard.Second_Meaning
			if actual_prompt_string == "" || actual_prompt_string == " " {
				emptyCardCounter++
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
			randDeckAndMode = rand.Intn(8) // if an empty card was picked, maybe try a different deck, or not
			if randDeckAndMode == 0 {
				current_deck_B = "init"
				deck_len = init_len
				randIndex := rand.Intn(len(fileOfCardsInitiate))
				aCard = fileOfCardsInitiate[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
				actual_prompt_string = aCard.Kanji
				primary_meaning = aCard.Meaning
				secondary_meaning = aCard.Second_Meaning
			} else if randDeckAndMode == 1 {
				current_deck_B = "nov"
				deck_len = novice_len
				randIndex := rand.Intn(len(fileOfCardsNovice))
				aCard = fileOfCardsNovice[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
				actual_prompt_string = aCard.Kanji
				primary_meaning = aCard.Meaning
				secondary_meaning = aCard.Second_Meaning
			} else if randDeckAndMode == 2 {
				current_deck_B = "claude"
				deck_len = claude_len
				randIndex := rand.Intn(len(fileOfCardsGraduate))
				aCard = fileOfCardsGraduate[randIndex]
				actual_prompt_string = aCard.Kanji
				primary_meaning = aCard.Meaning
				secondary_meaning = aCard.Second_Meaning
			} else if randDeckAndMode == 3 {
				current_deck_B = "mast"
				deck_len = master_len
				randIndex := rand.Intn(len(fileOfCardsMaster))
				aCard = fileOfCardsMaster[randIndex]
				actual_prompt_string = aCard.Kanji
				primary_meaning = aCard.Meaning
				secondary_meaning = aCard.Second_Meaning
			} else if randDeckAndMode == 4 {
				current_deck_B = "current"
				deck_len = current_len
				randIndex := rand.Intn(len(fileOf_Current))
				aCard = fileOf_Current[randIndex]
				actual_prompt_string = aCard.Kanji
				primary_meaning = aCard.Meaning
				secondary_meaning = aCard.Second_Meaning
			} else if randDeckAndMode == 5 {
				current_deck_B = "fresh"
				deck_len = fresh_len
				randIndex := rand.Intn(len(fileOf_fresh))
				aCard = fileOf_fresh[randIndex]
				actual_prompt_string = aCard.Kanji
				primary_meaning = aCard.Meaning
				secondary_meaning = aCard.Second_Meaning
			} else if randDeckAndMode == 6 {
				current_deck_B = "guru"
				deck_len = guru_len
				randIndex := rand.Intn(len(fileOfCardsGuru))
				aCard = fileOfCardsGuru[randIndex]
				actual_prompt_string = aCard.Kanji
				primary_meaning = aCard.Meaning
				secondary_meaning = aCard.Second_Meaning
			} else if randDeckAndMode == 7 {
				current_deck_B = "grad"
				deck_len = guru_len
				randIndex := rand.Intn(len(fileOfCardsGraduate))
				aCard = fileOfCardsGraduate[randIndex]
				actual_prompt_string = aCard.Kanji
				primary_meaning = aCard.Meaning
				secondary_meaning = aCard.Second_Meaning
			}
			if actual_prompt_string == "" || actual_prompt_string == " " { // if the prior conditional picks an empty card, pick another
				fmt.Printf("%s--- An empty kanji card was skipped in %s:%s !!!!%s\n",
					colorRed, current_deck, current_deck_B, colorReset) // Verified
				continue
			} else {
				break // break out of local for loop (in 'all') and naturally fall-through to the return
			}
		}
	}
}
