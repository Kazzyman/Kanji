package main

// LOGGERS:
func log_right(prompt_it_was, in string) { // - -
	recordGuess(prompt_it_was, in, aCard.Meaning, aCard.Second_Meaning)
	logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(prompt_it_was)
	logHits_in_cyclicArrayHits("Right", prompt_it_was)
}
func log_oops(prompt_it_was, field_it_was, guess string) { // - -
	logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(prompt_it_was)
	logHits_in_cyclicArrayHits("Oops", prompt_it_was)
	logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(prompt_it_was +
		":it was:" + field_it_was + ":but you had guessed:" + guess)
}

//
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

//
// A special Universal logger|Inserter: so we could drill the user more on chars he has missed, deprecated?
// ... now only used in the new hits function ?
func logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(Jchar string) { // - -
	cyclicArrayOfTheJcharsGottenWrong.InsertCharsWrong(Jchar)
}
