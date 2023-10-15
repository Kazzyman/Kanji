package main

import (
	"fmt"
)

// used ONLY in the 'set' directive to reSet the prompt & "aCard." fields
func silentlyLocateCard(setKeyRequest string) { //  - -
	foundElement = nil
	for _, card := range fileOfCardsK {
		if card.Meaning == setKeyRequest {
			// v v v if we find a 'card' in the range of 'fileOfCardsK',
			// ... we set the foundElement global var, which is used in reSet_aCard_andThereBy_reSet_thePromptString()
			foundElement = &card // foundElement is a global var and contains all the fields of a card or element
			break
		}
	}
	if foundElement == nil {
		// fmt.Println("Element not found in fileOfCardsK: func silentlyLocateCard()")
	}
	if foundElement == nil {

		for _, card := range fileOfCardsGuru2 {
			if card.Meaning == setKeyRequest {
				// v v v if we find a 'card' in the range of 'fileOfCardsGuru2',
				// ... we set the foundElement global var, which is used in reSet_aCard_andThereBy_reSet_thePromptString()
				foundElement = &card // foundElement is a global var and contains all the fields of a card or element
				break
			}
		}
		// if it is still nil ...
		if foundElement == nil {
			// fmt.Println("Element not found in fileOfCardsGuru2: func silentlyLocateCard()")

			for _, card := range fileOfCardsGuru {
				if card.Meaning == setKeyRequest {
					// v v v if we find a 'card' in the range of 'fileOfCardsGuru',
					// ... we set the foundElement global var, which is used in reSet_aCard_andThereBy_reSet_thePromptString()
					foundElement = &card // foundElement is a global var and contains all the fields of a card or element
					break
				}
			}
			if foundElement == nil {
				fmt.Printf("\nElement %s not found\n", setKeyRequest)
			}
		}
	}

}
