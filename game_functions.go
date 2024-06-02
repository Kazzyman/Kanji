package main

import (
	"fmt"
	"os"
	"time"
)

func initialize_game() { // ::: - -
	aGameIsRunning = true
	game_loop_counter = 0
	correctOnFirstAttemptAccumulator = 0
	correctOnSecondAttemptAccumulator = 0
	correctOnThirdAttemptAccumulator = 0
	failedOnThirdAttemptAccumulator = 0
	//
	// ::: file-writing and time-stamping is all that follows ==========================================
	currentTime := time.Now()
	TimeOfStartFromInceptionOfGame = time.Now()

	fileHandle, err := os.OpenFile("Kanji-newLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	_, err1 := fmt.Fprintf(fileHandle,
		"\n\n-- A game began at: %s",
		currentTime.Format("15:04:05 on Monday 01-02-2006"))
	check_error(err1)
}

/*
.
*/

func postGame_wrap_up() { // ::: - -
	aGameIsRunning = false
	now_using_game_duration_set_by_user = false

	fileHandle, err := os.OpenFile("Kanji-newLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	currentTime := time.Now()

	// calculate elapsed time
	t_s2 := time.Now()
	elapsed := t_s2.Sub(TimeOfStartFromInceptionOfGame)

	// Format the elapsed time to display minutes and whole seconds
	minutes := int(elapsed.Minutes())
	seconds := int(elapsed.Seconds()) % 60

	// Create the formatted string
	TotalRun := fmt.Sprintf("%02d:%02d", minutes, seconds)

	// ::: ------------------------------------------------------------------- v v v v
	Process_users_input_as_a_guess() // ::: trying this ---- and these ------- v v v v
	if guessLevelCounter == 2 {
		if gottenHonestly {
			correctOnFirstAttemptAccumulator++
			gottenHonestly = false
		}
	} else if guessLevelCounter == 3 {
		correctOnSecondAttemptAccumulator++
	} else { // :::  this only runs if user gets one right on his last guess (while at the highest guessLevelCounter)
		correctOnThirdAttemptAccumulator++
		// ::: fail/error accumulator gets incremented (or at least it gets displayed [below]) during the oops message
	} // - - - - - - - - - - - - - - - - - - - - - - - - - ::: -----------------------------------------------------------------

	if correctOnFirstAttemptAccumulator > 0 && correctOnSecondAttemptAccumulator > 0 && correctOnThirdAttemptAccumulator > 0 && failedOnThirdAttemptAccumulator == 0 { // ::: done
		fmt.Println(colorRed)
		fmt.Printf("\nYour Game run-time was:%s,  you got %s%d%s correct on your first try,  %s%d%s right on your second try,\n"+
			"... and you got %s%d right on your third try. \n\n", TotalRun, colorReset, correctOnFirstAttemptAccumulator, colorRed, colorReset, correctOnSecondAttemptAccumulator,
			colorRed, colorReset, correctOnThirdAttemptAccumulator)
	} else if correctOnFirstAttemptAccumulator > 0 && correctOnSecondAttemptAccumulator > 0 && correctOnThirdAttemptAccumulator == 0 && failedOnThirdAttemptAccumulator == 0 { // ::: done
		fmt.Println(colorRed)
		fmt.Printf("\nYour Game run-time was:%s,  you got %s%d%s correct on your first try,  %s%d right on your second try. \n\n", TotalRun, colorReset, correctOnFirstAttemptAccumulator,
			colorRed, colorReset, correctOnSecondAttemptAccumulator)
	} else if correctOnFirstAttemptAccumulator > 0 && correctOnSecondAttemptAccumulator == 0 && correctOnThirdAttemptAccumulator == 0 && failedOnThirdAttemptAccumulator == 0 { // ::: done
		fmt.Println(colorRed)
		fmt.Printf("\nYour Game run-time was:%s,  Gongratulations! you got %s%d correct on your first try. \n\n", TotalRun, colorReset, correctOnFirstAttemptAccumulator)
	} else {
		fmt.Println(colorRed)
		fmt.Printf("\nYour Game run-time was:%s,  you got %s%d%s correct on your first try,  %s%d%s right on your second try,\n"+
			"... and you got %s%d%s right on your third try, and were unable to answer correctly without a hint "+
			"%s%d times. \n\n", TotalRun, colorReset, correctOnFirstAttemptAccumulator, colorRed, colorReset, correctOnSecondAttemptAccumulator,
			colorRed, colorReset, correctOnThirdAttemptAccumulator, colorRed, colorReset, failedOnThirdAttemptAccumulator)
	}

	// End timer and report elapsed time and other stats to a file.
	_, err1 := fmt.Fprintf(fileHandle,
		"\n-- A game ended at: %s  Total prompts was: %d \n",
		currentTime.Format("15:04:05 on Monday 01-02-2006"), game_loop_counter)
	check_error(err1)

	_, err3 := fmt.Fprintf(fileHandle, "%s's results were as follows: Right on first attempt:%d, on 2nd attempt:%d, 3rd attempt:%d, even a hint was ineffective:%d, %d/%d\n",
		nameOfPlayer, correctOnFirstAttemptAccumulator,
		correctOnSecondAttemptAccumulator, correctOnThirdAttemptAccumulator,
		failedOnThirdAttemptAccumulator, game_loop_counter, game_duration_set_by_user)
	check_error(err3)

	_, err2 := fmt.Fprintf(fileHandle,
		"The Elapsed time of the game was: %s \n\n\n",
		TotalRun)
	check_error(err2)
}

/*
.
*/
