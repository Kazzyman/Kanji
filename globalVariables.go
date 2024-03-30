package main

import "time"

// This file contains All global vars NOT located in the objectsAndMethods.go file
var already_used_map = make(map[string]int)

// All of the decks will draw cards per this aCard var:
var aCard = charSetStructKanji{}

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

var weHaveBeenHereBefore bool
var future_len int
var data_file100_maybe_len int

var beauty_len int
var init_len int
var guru_len int
var master_len int
var novice_len int
var grad_len int
var fresh_len int
var current_len int
var data_len int
var claude_len int
var words_len int

var deck_len int

var randDeckAndMode int
var current_deck string
var current_deck_B string
var total_prompts int

var TimeOfStartFromTop = time.Now()

var game_loop_counter int
var game_duration = 998

var gameOn bool

var runeOfCode = `
    silentlyLocateCard("west") // Set the Convenience-global: foundElement
    if foundElement != nil {
        aCard = *foundElement // Set the global var-object 'aCard'
        prompt = aCard.Kanji
        objective = aCard.Meaning
    }
`

// Constants:
const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorCyan = "\033[36m"
const colorPurple = "\033[35m"

// Global Maps:
//
// Used in : func read_map_of_fineOn()
// ... and in : func logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(promptToSkip string)
var frequencyMapOf_IsFineOnChars = make(map[string]int) // - -
// Used in : func read_map_of_needWorkOn()
// ... and in : func logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(promptToWorkMoreOn string)
var frequencyMapOf_need_workOn = make(map[string]int) // - -
