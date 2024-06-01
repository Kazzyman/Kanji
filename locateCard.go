package main

import (
	"fmt"
	"regexp"
)

// Used in respond_to_UserSupplied_Directive() in case: "setc"
// Also used in the 'setc' directive via: reSet_aCard_andThereBy_reSet_thePromptString()
// ... to reSet the card, i.e., the Kanji/prompt(and all other aCard fields) via its Meaning field
//
// The sole function of this func is to set the global: foundElement
func silentlyLocateCard(setKeyRequest string) { //  - -
	// card.Meaning will (sometimes) be either a string such as "red-1" or "red-2"
	// ... likewise, setKeyRequest will (only sometimes) be either a string such as "red-1" or "red-2"
	foundElement = nil                        // Prime the global foundElement, a pointer thus: var foundElement *charSetStructKanji
	regex := regexp.MustCompile(`^([a-z]+)-`) // Used to strip the -int suffix from the card.Meaning string
	regexH := regexp.MustCompile(`-`)         // Used to determine if the user entered a hyphenated val for setKeyRequest
	var meaning string                        // Used to stand-in for card.Meaning in cases where setKeyRequest is sans hyphen

	//
	// Firstly, look in the "claude" deck
	for _, card := range claude { // This new local variable: card, will be an object defined by a structure
		if len(regexH.FindStringSubmatch(setKeyRequest)) > 0 { // regex-H scans setKeyRequest to see if that string is hyphenated ...
			if card.Meaning == setKeyRequest { // ... go and look for the literal (complete) hyphenated meaning (special case).
				// v v v if we find a 'card' in the range of 'claude',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card.
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"claude\" deck, part 1\n\n", setKeyRequest)
				break
			}
		} else { // The user has entered a value for setKeyRequest sans hyphen (the usual case) ...
			// ... we therefore check it against card.Meaning which may OR MAY NOT be hyphenated, itself.
			matches := regex.FindStringSubmatch(card.Meaning)
			if len(matches) > 0 { // setKeyRequest is sans hyphen, and card.Meaning IS hyphenated; so (maybe) truncate card.Meaning
				// todo:
				// Only if card.Meaning is of the form meaning-int rather than meaning-secondWordOfMeaning ...
				meaning = regex.FindStringSubmatch(card.Meaning)[1] // index of 1 grabs the chars left of the hyphen
			} else { // setKeyRequest is hyphenated, so we will be looking for the complete setKeyRequest in each card.Meaning
				// setKeyRequest string was not hyphenated, and card.Meaning has a non-int val right of any hyphen ...
				// ... so grab all of card.Meaning (usual case)
				meaning = card.Meaning
			}

			// if card.Meaning == setKeyRequest { // Was the prior condition in need of extra attention ...
			if meaning == setKeyRequest { // The revised condition that will also handle compound strings suffixed by -int
				// v v v if we find a 'card' in the range of 'claude',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"claude\" deck, part 2\n\n", setKeyRequest)
				break
			}
		}
	}
	//
	// When no match is found above, look in the "fileOf_Current" deck
	for _, card := range fileOf_Current { // This new local variable: card, will be an object defined by a structure
		if len(regexH.FindStringSubmatch(setKeyRequest)) > 0 { // regex-H scans setKeyRequest to see if that string is hyphenated ...
			if card.Meaning == setKeyRequest { // ... go and look for the literal (complete) hyphenated meaning (special case).
				// v v v if we find a 'card' in the range of 'fileOf_Current',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card.
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"current\" deck, part 1\n\n", setKeyRequest)
				break
			}
		} else { // The user has entered a value for setKeyRequest sans hyphen (the usual case) ...
			// ... we therefore check it against card.Meaning which may OR MAY NOT be hyphenated, itself.
			matches := regex.FindStringSubmatch(card.Meaning)
			if len(matches) > 0 { // setKeyRequest is sans hyphen, and card.Meaning IS hyphenated; so truncate card.Meaning
				meaning = regex.FindStringSubmatch(card.Meaning)[1] // index of 1 grabs the chars left of the hyphen
			} else { // setKeyRequest is hyphenated, so we will be looking for the complete setKeyRequest in each card.Meaning
				meaning = card.Meaning // setKeyRequest string was not hyphenated, so grab all of card.Meaning (usual case)
			}

			// if card.Meaning == setKeyRequest { // Was the prior condition in need of extra attention ...
			if meaning == setKeyRequest { // The revised condition that will also handle compound strings suffixed by -int
				// v v v if we find a 'card' in the range of 'fileOf_Current',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"current\" deck, part 2\n\n", setKeyRequest)
				break
			}
		}
	}
	//
	// When no match is found above, look in the "data_file" deck
	for _, card := range data_file { // This new local variable: card, will be an object defined by a structure
		if len(regexH.FindStringSubmatch(setKeyRequest)) > 0 { // regex-H scans setKeyRequest to see if that string is hyphenated ...
			if card.Meaning == setKeyRequest { // ... go and look for the literal (complete) hyphenated meaning (special case).
				// v v v if we find a 'card' in the range of 'data_file',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card.
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"data\" deck, part 1\n\n", setKeyRequest)
				break
			}
		} else { // The user has entered a value for setKeyRequest sans hyphen (the usual case) ...
			// ... we therefore check it against card.Meaning which may OR MAY NOT be hyphenated, itself.
			matches := regex.FindStringSubmatch(card.Meaning)
			if len(matches) > 0 { // setKeyRequest is sans hyphen, and card.Meaning IS hyphenated; so truncate card.Meaning
				meaning = regex.FindStringSubmatch(card.Meaning)[1] // index of 1 grabs the chars left of the hyphen
			} else { // setKeyRequest is hyphenated, so we will be looking for the complete setKeyRequest in each card.Meaning
				meaning = card.Meaning // setKeyRequest string was not hyphenated, so grab all of card.Meaning (usual case)
			}

			// if card.Meaning == setKeyRequest { // Was the prior condition in need of extra attention ...
			if meaning == setKeyRequest { // The revised condition that will also handle compound strings suffixed by -int
				// v v v if we find a 'card' in the range of 'data_file',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"data\" deck, part 2\n\n", setKeyRequest)
				break
			}
		}
	}

	// When no match is found above, look in the "fileOf_fresh" deck
	for _, card := range fileOf_fresh { // This new local variable: card, will be an object defined by a structure
		if len(regexH.FindStringSubmatch(setKeyRequest)) > 0 { // regex-H scans setKeyRequest to see if that string is hyphenated ...
			if card.Meaning == setKeyRequest { // ... go and look for the literal (complete) hyphenated meaning (special case).
				// v v v if we find a 'card' in the range of 'fileOf_fresh',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card.
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"fresh\" deck, part 1\n\n", setKeyRequest)
				break
			}
		} else { // The user has entered a value for setKeyRequest sans hyphen (the usual case) ...
			// ... we therefore check it against card.Meaning which may OR MAY NOT be hyphenated, itself.
			matches := regex.FindStringSubmatch(card.Meaning)
			if len(matches) > 0 { // setKeyRequest is sans hyphen, and card.Meaning IS hyphenated; so truncate card.Meaning
				meaning = regex.FindStringSubmatch(card.Meaning)[1] // index of 1 grabs the chars left of the hyphen
			} else { // setKeyRequest is hyphenated, so we will be looking for the complete setKeyRequest in each card.Meaning
				meaning = card.Meaning // setKeyRequest string was not hyphenated, so grab all of card.Meaning (usual case)
			}

			// if card.Meaning == setKeyRequest { // Was the prior condition in need of extra attention ...
			if meaning == setKeyRequest { // The revised condition that will also handle compound strings suffixed by -int
				// v v v if we find a 'card' in the range of 'fileOf_fresh',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"fresh\" deck, part 2\n\n", setKeyRequest)
				break
			}
		}
	}
	//

	// When no match is found above, look in the "fileOfCardsGraduate" deck
	for _, card := range fileOfCardsGraduate { // This new local variable: card, will be an object defined by a structure
		if len(regexH.FindStringSubmatch(setKeyRequest)) > 0 { // regex-H scans setKeyRequest to see if that string is hyphenated ...
			if card.Meaning == setKeyRequest { // ... go and look for the literal (complete) hyphenated meaning (special case).
				// v v v if we find a 'card' in the range of 'fileOf_fresh',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card.
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"Graduate\" deck, part 1\n\n", setKeyRequest)
				break
			}
		} else { // The user has entered a value for setKeyRequest sans hyphen (the usual case) ...
			// ... we therefore check it against card.Meaning which may OR MAY NOT be hyphenated, itself.
			matches := regex.FindStringSubmatch(card.Meaning)
			if len(matches) > 0 { // setKeyRequest is sans hyphen, and card.Meaning IS hyphenated; so truncate card.Meaning
				meaning = regex.FindStringSubmatch(card.Meaning)[1] // index of 1 grabs the chars left of the hyphen
			} else { // setKeyRequest is hyphenated, so we will be looking for the complete setKeyRequest in each card.Meaning
				meaning = card.Meaning // setKeyRequest string was not hyphenated, so grab all of card.Meaning (usual case)
			}

			// if card.Meaning == setKeyRequest { // Was the prior condition in need of extra attention ...
			if meaning == setKeyRequest { // The revised condition that will also handle compound strings suffixed by -int
				// v v v if we find a 'card' in the range of 'fileOf_fresh',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"Graduate\" deck, part 2\n\n", setKeyRequest)
				break
			}
		}
	}
	//

	// When no match is found above, look in the "fileOfCardsInitiate" deck
	for _, card := range fileOfCardsInitiate { // This new local variable: card, will be an object defined by a structure
		if len(regexH.FindStringSubmatch(setKeyRequest)) > 0 { // regex-H scans setKeyRequest to see if that string is hyphenated ...
			if card.Meaning == setKeyRequest { // ... go and look for the literal (complete) hyphenated meaning (special case).
				// v v v if we find a 'card' in the range of 'fileOf_fresh',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card.
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"Initiate\" deck, part 1\n\n", setKeyRequest)
				break
			}
		} else { // The user has entered a value for setKeyRequest sans hyphen (the usual case) ...
			// ... we therefore check it against card.Meaning which may OR MAY NOT be hyphenated, itself.
			matches := regex.FindStringSubmatch(card.Meaning)
			if len(matches) > 0 { // setKeyRequest is sans hyphen, and card.Meaning IS hyphenated; so truncate card.Meaning
				meaning = regex.FindStringSubmatch(card.Meaning)[1] // index of 1 grabs the chars left of the hyphen
			} else { // setKeyRequest is hyphenated, so we will be looking for the complete setKeyRequest in each card.Meaning
				meaning = card.Meaning // setKeyRequest string was not hyphenated, so grab all of card.Meaning (usual case)
			}

			// if card.Meaning == setKeyRequest { // Was the prior condition in need of extra attention ...
			if meaning == setKeyRequest { // The revised condition that will also handle compound strings suffixed by -int
				// v v v if we find a 'card' in the range of 'fileOf_fresh',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"Initiate\" deck, part 2\n\n", setKeyRequest)
				break
			}
		}
	}
	//

	// When no match is found above, look in the "fileOfCardsMaster" deck
	for _, card := range fileOfCardsMaster { // This new local variable: card, will be an object defined by a structure
		if len(regexH.FindStringSubmatch(setKeyRequest)) > 0 { // regex-H scans setKeyRequest to see if that string is hyphenated ...
			if card.Meaning == setKeyRequest { // ... go and look for the literal (complete) hyphenated meaning (special case).
				// v v v if we find a 'card' in the range of 'fileOf_fresh',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card.
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"Master\" deck, part 1\n\n", setKeyRequest)
				break
			}
		} else { // The user has entered a value for setKeyRequest sans hyphen (the usual case) ...
			// ... we therefore check it against card.Meaning which may OR MAY NOT be hyphenated, itself.
			matches := regex.FindStringSubmatch(card.Meaning)
			if len(matches) > 0 { // setKeyRequest is sans hyphen, and card.Meaning IS hyphenated; so truncate card.Meaning
				meaning = regex.FindStringSubmatch(card.Meaning)[1] // index of 1 grabs the chars left of the hyphen
			} else { // setKeyRequest is hyphenated, so we will be looking for the complete setKeyRequest in each card.Meaning
				meaning = card.Meaning // setKeyRequest string was not hyphenated, so grab all of card.Meaning (usual case)
			}

			// if card.Meaning == setKeyRequest { // Was the prior condition in need of extra attention ...
			if meaning == setKeyRequest { // The revised condition that will also handle compound strings suffixed by -int
				// v v v if we find a 'card' in the range of 'fileOf_fresh',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"Master\" deck, part 2\n\n", setKeyRequest)
				break
			}
		}
	}
	//

	// When no match is found above, look in the "fileOfCardsNovice" deck
	for _, card := range fileOfCardsNovice { // This new local variable: card, will be an object defined by a structure
		if len(regexH.FindStringSubmatch(setKeyRequest)) > 0 { // regex-H scans setKeyRequest to see if that string is hyphenated ...
			if card.Meaning == setKeyRequest { // ... go and look for the literal (complete) hyphenated meaning (special case).
				// v v v if we find a 'card' in the range of 'fileOf_fresh',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card.
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"Novice\" deck, part 1\n\n", setKeyRequest)
				break
			}
		} else { // The user has entered a value for setKeyRequest sans hyphen (the usual case) ...
			// ... we therefore check it against card.Meaning which may OR MAY NOT be hyphenated, itself.
			matches := regex.FindStringSubmatch(card.Meaning)
			if len(matches) > 0 { // setKeyRequest is sans hyphen, and card.Meaning IS hyphenated; so truncate card.Meaning
				meaning = regex.FindStringSubmatch(card.Meaning)[1] // index of 1 grabs the chars left of the hyphen
			} else { // setKeyRequest is hyphenated, so we will be looking for the complete setKeyRequest in each card.Meaning
				meaning = card.Meaning // setKeyRequest string was not hyphenated, so grab all of card.Meaning (usual case)
			}

			// if card.Meaning == setKeyRequest { // Was the prior condition in need of extra attention ...
			if meaning == setKeyRequest { // The revised condition that will also handle compound strings suffixed by -int
				// v v v if we find a 'card' in the range of 'fileOf_fresh',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"Novice\" deck, part 2\n\n", setKeyRequest)
				break
			}
		}
	}
	//

	// When no match is found above, look in the "fileOfCardsGuru" deck
	for _, card := range fileOfCardsGuru { // This new local variable: card, will be an object defined by a structure
		if len(regexH.FindStringSubmatch(setKeyRequest)) > 0 { // regex-H scans setKeyRequest to see if that string is hyphenated ...
			if card.Meaning == setKeyRequest { // ... go and look for the literal (complete) hyphenated meaning (special case).
				// v v v if we find a 'card' in the range of 'fileOf_fresh',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card.
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"Guru\" deck, part 1\n\n", setKeyRequest)
				break
			}
		} else { // The user has entered a value for setKeyRequest sans hyphen (the usual case) ...
			// ... we therefore check it against card.Meaning which may OR MAY NOT be hyphenated, itself.
			matches := regex.FindStringSubmatch(card.Meaning)
			if len(matches) > 0 { // setKeyRequest is sans hyphen, and card.Meaning IS hyphenated; so truncate card.Meaning
				meaning = regex.FindStringSubmatch(card.Meaning)[1] // index of 1 grabs the chars left of the hyphen
			} else { // setKeyRequest is hyphenated, so we will be looking for the complete setKeyRequest in each card.Meaning
				meaning = card.Meaning // setKeyRequest string was not hyphenated, so grab all of card.Meaning (usual case)
			}

			// if card.Meaning == setKeyRequest { // Was the prior condition in need of extra attention ...
			if meaning == setKeyRequest { // The revised condition that will also handle compound strings suffixed by -int
				// v v v if we find a 'card' in the range of 'fileOf_fresh',
				// ... we set the foundElement global var
				foundElement = &card // foundElement is a global var which contains(refers to) all the fields of a card
				// i.e., it is a pointer thus: var foundElement *charSetStructKanji
				fmt.Printf("setKeyRequest: %s was found in the \"Guru\" deck, part 2\n\n", setKeyRequest)
				break
			}
		}
	}

	// todo: add the remaining decks
}
