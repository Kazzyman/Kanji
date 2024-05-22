package main

func directiveHandler(in, inew_prompt, inew_objective, inew_objective_kind, inew_secondary_objective string) (new_prompt, new_objective, new_objective_kind, new_secondary_objective string, returning_fr_dir bool) {
	var setcWasRun bool
	setcWasRun = false
	if in_list_of_Directives(in) {
		if in == "setc" {
			setcWasRun = true
			new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir = respond_to_UserSuppliedDirective(in)
		} else {
			setcWasRun = false
			_, _, _, _, returning_fr_dir = respond_to_UserSuppliedDirective(in)
		}
		// Either way, evaluateUsersGuess(...)
		evaluateUsersGuess(in, inew_prompt, inew_objective, inew_objective_kind, true, false, true, new_secondary_objective)
	} else {
		// if setc not run:
		evaluateUsersGuess(in, inew_prompt, inew_objective, inew_objective_kind, true, true, false, new_secondary_objective)
	}
	if setcWasRun {
		// return the forced new card PLUS the returning_fr_dir flag.
		return new_prompt, new_objective, new_objective_kind, new_secondary_objective, returning_fr_dir
	} else {
		// return originals, thereby keeping via not changing them. PLUS the returning_fr_dir flag.
		return inew_prompt, inew_objective, inew_objective_kind, inew_secondary_objective, returning_fr_dir // todo unsure of what to do here.
	}
}

func promptCase1(new_prompt string, field_to_prompt_from string) (in string) {
	// This prompt, deployed by new_objective_kind, takes new_prompt
	if field_to_prompt_from == "kanji" {
		// This prompt, deployed by objective_kind, takes promptField (rather than the new_prompt variant)
		in = promptWithDir(new_prompt) // Get user's input, from a randomly selected prompt
	}
	if field_to_prompt_from == "kunyomi" {
		in = promptWithDir(aCard.Kunyomi) // Get user's input, from a randomly selected prompt
	}
	return in
}

func freshCard(inew_prompt, inew_objective, inew_objective_kind, inew_secondary_objective string, already_used_map map[string]int) (new_prompt, new_objective, new_objective_kind, new_secondary_objective string) {
	var actuallyReturnNewValues bool
	actuallyReturnNewValues = false
	length := 0
	loopCounter := 0
	// fmt.Printf("length of map is %d\n", len(already_used_map))
	for length <= len(already_used_map) {
		length++
		loopCounter++
		if loopCounter > 9999 {
			// fmt.Printf("loopCounter:%d, so doing a clear of the map\n", loopCounter)
			clearMap()
			break
		}
		if is_pick_novel(inew_prompt) {
			already_used_map[inew_prompt]++
			break // keep it
		}
		// else [because of the double breaks above], pick another
		new_prompt, new_objective, new_objective_kind, new_secondary_objective = pick_RandomCard_Assign_fields()
		length = 0 // continue the loop ensuring that the entire map is read with this new pick
		actuallyReturnNewValues = true
	}
	if actuallyReturnNewValues {
		return new_prompt, new_objective, new_objective_kind, new_secondary_objective // return new pick.
	} else {
		return inew_prompt, inew_objective, inew_objective_kind, inew_secondary_objective // return old original values: keep first draw.
	}
}
