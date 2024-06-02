package main

import (
	"fmt"
	"os"
)

// LOGGERS:
func log_right(actual_prompt_string, usersSubmission string) { // - -
	recordGuess(actual_prompt_string, usersSubmission, aCard.Meaning, aCard.Second_Meaning)
	logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(actual_prompt_string)
	logHits_in_cyclicArrayHits("Right", actual_prompt_string)
	if aGameIsRunning {
		game_loop_counter++
		if game_loop_counter > game_duration_set_by_user {
			game_loop_counter = 0
			game_off()
			postGame_wrap_up()
		}

		if weHadFailed_And_OnlyGotThisRightBecauseOfTheClue {
			// Then that fail has already been logged and we need to skip tallying the score.
			weHadFailed_And_OnlyGotThisRightBecauseOfTheClue = false
		} else {
			if guessLevelCounter == 2 {
				if gottenHonestly {
					correctOnFirstAttemptAccumulator++ // ::: 1st
					gottenHonestly = false
				}
			} else if guessLevelCounter == 3 {
				correctOnSecondAttemptAccumulator++ // ::: 2nd
			} else {
				correctOnThirdAttemptAccumulator++ // ::: 3rd
			}
			// ::: The other accumulator++  thang : failedOnThirdAttemptAccumulator ]todo[ ... gets handled in log_oops()
		}
	}
	if guessLevelCounter == 3 {
		fileHandle, err := os.OpenFile("Kanji-newLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		check_error(err)
		_, err1 := fmt.Fprintf(fileHandle,
			"\n%s may have mistyped==%s:%s", nameOfPlayer, aCard.Meaning, aCard.Kanji) // mistyped is a word?
		check_error(err1)
	} else if guessLevelCounter == 4 {
		fileHandle, err := os.OpenFile("Kanji-newLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		check_error(err)
		_, err1 := fmt.Fprintf(fileHandle,
			"\n%s had a some difficulty with==%s:%s", nameOfPlayer, aCard.Meaning, aCard.Kanji)
		check_error(err1)
	}
}

/*

 */

func log_oops(actual_prompt_string, primary_meaning, usersSubmission string) { // - -
	logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(actual_prompt_string)
	logHits_in_cyclicArrayHits("Oops", actual_prompt_string)
	logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(actual_prompt_string +
		":it was:" + primary_meaning + ":but you had guessed:" + usersSubmission)
}

// 'Reinforce-or-Skip' loggers|Inserters:
func logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(promptToSkip string) { // - -
	frequencyMapOf_IsFineOnChars[promptToSkip]++
}
func logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(promptToWorkMoreOn string) { // - -
	frequencyMapOf_need_workOn[promptToWorkMoreOn]++
}

// Universal hits logger|Inserter:
func logHits_in_cyclicArrayHits(RightOrOops, JChar string) { // - -
	cyclicArrayHits.InsertRightOrOops(RightOrOops)
	cyclicArrayHits.InsertChar(JChar)
}

// A special Universal logger|Inserter: so we could drill the user more on chars he has missed, deprecated?
// ... now only used in the new hits function ?
func logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(Jchar string) { // - -
	cyclicArrayOfTheJcharsGottenWrong.InsertCharsWrong(Jchar)
}

/*
.
*/

func resetAllLogs() {
	cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
	cyclicArrayHits = CyclicArrayHits{}
	// Also, flush (clear) the maps
	total_prompts = 0
	kanjiHitMap = make(map[string]CardInfo)
	frequencyMapOf_IsFineOnChars = make(map[string]int)
	frequencyMapOf_need_workOn = make(map[string]int)
	already_used_map = make(map[string]int)
	//
	//goland:noinspection ALL
	fmt.Println("\nArrays and maps flushed:\n")
	fmt.Println("    cyclicArrayOfTheJcharsGottenWrong")
	fmt.Println("    cyclicArrayHits")
	fmt.Println("    frequencyMapOf_IsFineOnChars")
	fmt.Println("    frequencyMapOf_need_workOn")
	fmt.Println("    already_used_map")

	fmt.Println("    kanjiHitMap")
	//goland:noinspection ALL
	fmt.Println("    frequencyMapOf_need_workOn\n")
	if weHaveBeenHereBefore {
		fmt.Println("  All Game values have also been reset; answer this last one from the prior deck")
	} else {
		fmt.Println("  Ready to begin.")
	}
	weHaveBeenHereBefore = true
}
