package main

import (
	"fmt"
)

// Used in respond_to_UserSuppliedDirective() in case: "setc"
// Also used in the 'setc' directive via: reSet_aCard_andThereBy_reSet_thePromptString()
// ... to reSet the card, i.e., the Kanji/prompt(and all other aCard fields) via its Meaning field
//
// The sole function of this func is to set the global: foundElement
func silentlyLocateCard(setKeyRequest string) { //  - -
	foundElement = nil // Prime the global foundElement, a pointer thus: var foundElement *charSetStructKanji
	//
	// Firstly, look in fileOfCardsInitiate
	for _, card := range fileOfCardsInitiate { // The new local variable: card will be an object defined by a structure
		if card.Meaning == setKeyRequest {
			// v v v if we find a 'card' in the range of 'fileOfCardsInitiate',
			// ... we set the foundElement global var
			foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card
			// i.e., it is a pointer thus: var foundElement *charSetStructKanji
			break
		}
	}
	//
	if foundElement == nil { // If we did not locate the card we seek: look in ALL other decks
		// fmt.Println("Element not found in fileOfCardsInitiate: func silentlyLocateCard()")
		//
		// As a SECONDARY possible location, look in fileOfCardsNovice
		for _, card := range fileOfCardsNovice {
			if card.Meaning == setKeyRequest {
				foundElement = &card
				break
			}
		}
		if foundElement == nil { // If we STILL have not yet located the card we seek:
			// fmt.Println("Element not found in fileOfCardsNovice: func silentlyLocateCard()")
			//
			// As a TERTIARY possible location, look in fileOfCardsGraduate
			for _, card := range fileOfCardsGraduate {
				if card.Meaning == setKeyRequest {
					foundElement = &card
					break
				}
			}
		}
		if foundElement == nil {
			// fmt.Println("Element not found in fileOfCardsGraduate: func silentlyLocateCard()")
			//
			// As a QUATERNARY possible location, look in fileOfCardsMaster
			for _, card := range fileOfCardsMaster {
				if card.Meaning == setKeyRequest {
					foundElement = &card
					break
				}
			}
		}
		if foundElement == nil {
			// fmt.Println("Element not found in fileOfCardsMaster: func silentlyLocateCard()")
			//
			// As THE FINAL possible location, look in fileOfCardsGuru
			for _, card := range fileOfCardsGuru {
				if card.Meaning == setKeyRequest {
					foundElement = &card
					break
				}
			}
			if foundElement == nil {

				fmt.Printf(string(colorRed))
				fmt.Printf("\nElement ")
				fmt.Printf(string(colorReset))
				fmt.Printf("\"%s\"", setKeyRequest)
				fmt.Printf(string(colorRed))
				fmt.Printf(" was not found in any deck : silentlyLocateCard()\n")
				fmt.Printf(string(colorReset))
			}
		}
	}
}
