package main

import "fmt"

// LOGGERS:
func log_right(actual_prompt_string, usersSubmission string) { // - -
	recordGuess(actual_prompt_string, usersSubmission, aCard.Meaning, aCard.Second_Meaning)
	logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(actual_prompt_string)
	logHits_in_cyclicArrayHits("Right", actual_prompt_string)
	if theGameIsRunning {
		game_loop_counter++
		if game_loop_counter > game_duration_set_by_user {
			the_game_ends()
		}

		if weHadFailed_And_OnlyGotThisRightBecauseOfTheClue {
			// Then that fail has already been logged and we need to skip all logging.
			weHadFailed_And_OnlyGotThisRightBecauseOfTheClue = false
		} else {
			if guessLevelCounter == 2 {
				if gottenHonestly { // todo] do not accumulate if after an "error" or hint
					correctOnFirstAttemptAccumulator++ // ::: 1st
					gottenHonestly = false
				}
			} else if guessLevelCounter == 3 {
				correctOnSecondAttemptAccumulator++ // ::: 2nd
			} else {
				// ... then ... the guessLevelCounter was 4.
				fmt.Printf("here in log right, guessLevelCounter is:%d, and it should be 4\n", guessLevelCounter)
				correctOnThirdAttemptAccumulator++ // ::: 3rd
			}
			// ::: The other accumulator++  thang : failedOnThirdAttemptAccumulator ]todo[ ... gets handled in log_oops()
		}
	}
}
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
