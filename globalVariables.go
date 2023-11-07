package main

import "time"

// This file contains many of the global vars NOT located in the objectsAndMethods.go file
//

// All of the decks will draw cards per this aCard var:
var aCard = charSetStructKanji{}

//
var foundElement *charSetStructKanji

// All cards in all of the decks are of this Structure:
type charSetStructKanji struct {
	Kanji          string
	Meaning        string
	Second_Meaning string
	Onyomi         string
	Kunyomi        string
	Vocab          string
	Vocab2         string
}

var returning_from_a_wrong_guess = false
var current_deck string
var current_deckA string
var total_prompts int

var forLoop int
var randDeck int
var randDeckAndMode int

// Indexes for the sequential files, used in pick_a_card.go
var indexNovS int
var indexInitS int
var indexGuruS int
var indexMastS int
var indexGradS int
var indexFresh int

var TimeOfStartFromTop = time.Now()

var game_loop_counter int
var game_duration = 998
var count int
var gameOn bool

// Constants:
const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorCyan = "\033[36m"
const colorPurple = "\033[35m"
const colorYellow = "\033[33m"

//
// Global Maps:
//
// Used in : func read_map_of_fineOn()
// ... and in : func logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(promptToSkip string)
var frequencyMapOf_IsFineOnChars = make(map[string]int) // - -
//
// Used in : func read_map_of_needWorkOn()
// ... and in : func logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(promptToWorkMoreOn string)
var frequencyMapOf_need_workOn = make(map[string]int) // - -
