package main

import (
	"fmt"
	"regexp"
)

// used ONLY in the 'set' directive to reSet the prompt & "aCard." fields
func silentlyLocateCard(setKeyRequest string) { //  - -

	var isAlphanumeric bool

	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true {
	case findAlphasIn.MatchString(setKeyRequest):
		isAlphanumeric = true
	default:
		isAlphanumeric = false
	}

	if isAlphanumeric == false { // ... then we should be safe to proceed with what will be a Hiragana char
		for _, card := range fileOfCardsK {
			if card.Kanji == setKeyRequest {
				// v v v if we find a 'card' in the range of 'fileOfCards',
				// ... we set the foundElement global var, which is used in reSet_aCard_andThereBy_reSet_thePromptString()
				foundElement = &card // foundElement is a global var and contains all the fields of a card or element
				break
			}
		}
		if foundElement == nil {
			fmt.Println("Element not found in: silentlyLocateCard(setKeyRequest string)")
		}
	} else {
		fmt.Printf(colorRed)
		fmt.Println("\nYou bastard!")
		fmt.Printf("\nYou have killed me with an Alpha string instead of a Hiragana\n\n")
		fmt.Printf(colorReset)
	}
}
