package main

// Contains two versions of lff dir. The first is for only comparing two data sets. The second func version is for
// ... comparing all nine data sets (claude-1, current-2, data-3, fresh-4, grad-5, guru-6, init-7, master-8, and novice-9)
// And, also the two yet to be used data sets:
// ... data_future (fileOf_Future) and:
// ... a_maybe_100 (data_file100_maybe)

// beauty_len = len(data_beauty)
//

import (
	"fmt"
	"os"
)

var one_card charSetStructKanji
var used_slice []string // To keep track of kanji chars already found and added to the output file.
var one_field string
var already_loged bool

// lff as a dir
// The purpose of this func is to parse all data sets (as apposed to files, as is done in fif), and
// ... compile a tally of kanji+meaning; while excluding duplicates.
// i.e., used to scan all data_category files and build a list of all unique cards (first line)
func list_from_files0() {
	used_slice = append(used_slice, "ricky") // Prime the slice with a string
	already_loged = false
	//
	// Open a file to be used to log the tally of unique cards
	output_file, err := os.OpenFile("ListOfUniqueCards.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening or creating ListOfUniqueCards.txt:", err)
		return
	} else {
		fmt.Fprintf(output_file, "This is a list of cards and where they were found\n\n")
	}
	defer func(output_file *os.File) {
		_ = output_file.Close()
	}(output_file)
	//
	// Open another file to be used to log duplicate cards found and the datasets involved
	output_file2, err := os.OpenFile("ListOfDuplicates.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening or creating ListOfDuplicates.txt:", err)
		return
	} else {
		fmt.Fprintf(output_file2, "This is a list of duplicates that were discovered \n\n")
	}
	defer func(output_file2 *os.File) {
		_ = output_file2.Close()
	}(output_file2)
	//
	//
	//
	for _, one_card = range claude {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s, found in dataset: claude\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s, from dataset: claude\n", one_card.Kanji, one_card.Meaning)
		}
	}
	//
	for _, one_card = range fileOfCardsGraduate {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s, found in dataset: grad\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s, from dataset: grad\n", one_card.Kanji, one_card.Meaning)
		}
	}

}

// ============================================================================================
// ============================================================================================

// lff as a dir
// The purpose of this func is to parse all data sets (as apposed to files, as is done in fif), and
// ... compile a tally of kanji+meaning; while excluding duplicates.
// i.e., used to scan all data_category files and build a list of all unique cards (first line)
func list_from_files() {
	used_slice = append(used_slice, "ricky") // Prime the slice with a string
	already_loged = false
	//
	// Open a file to be used to log the tally of unique cards
	output_file, err := os.OpenFile("ListOfUniqueCards.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening or creating ListOfUniqueCards.txt:", err)
		return
	} else {
		fmt.Fprintf(output_file, "This is a list of cards and where they were found\n\n")
	}
	defer func(output_file *os.File) {
		_ = output_file.Close()
	}(output_file)
	//
	// Open another file to be used to log duplicate cards found and the datasets involved
	output_file2, err := os.OpenFile("ListOfDuplicates.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening or creating ListOfDuplicates.txt:", err)
		return
	} else {
		fmt.Fprintf(output_file2, "This is a list of duplicates that were discovered \n\n")
	}
	defer func(output_file2 *os.File) {
		_ = output_file2.Close()
	}(output_file2)
	//
	//
	//
	// 1
	counter_claude := 0
	for _, one_card = range claude {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s, found in dataset: claude\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			counter_claude++
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s, from dataset: claude\n", one_card.Kanji, one_card.Meaning)
		}
	}
	fmt.Fprintf(output_file, "So that is %d cards from the claude dataset:%d\n\n", counter_claude, claude_len)

	// 2
	counter_current := 0
	for _, one_card = range fileOf_Current {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s, found in dataset: current\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			counter_current++
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s, from dataset: current\n", one_card.Kanji, one_card.Meaning)
		}
	}
	fmt.Fprintf(output_file, "So that is %d cards from the current dataset:%d\n\n", counter_current, current_len)

	// 3
	counter_data := 0
	for _, one_card = range data_file {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s, found in dataset: data\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			counter_data++
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s, from dataset: data\n", one_card.Kanji, one_card.Meaning)
		}
	}
	fmt.Fprintf(output_file, "So that is %d cards from the data dataset:%d\n\n", counter_data, current_len)

	// 4
	counter_fresh := 0
	for _, one_card = range fileOf_fresh {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s, found in dataset: fresh\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			counter_fresh++
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s, from dataset: fresh\n", one_card.Kanji, one_card.Meaning)
		}
	}
	fmt.Fprintf(output_file, "So that is %d cards from the fresh dataset:%d\n\n", counter_fresh, fresh_len)

	// 5
	counter_grad := 0
	for _, one_card = range fileOfCardsGraduate {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s, found in dataset: grad\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			counter_grad++
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s, from dataset: grad\n", one_card.Kanji, one_card.Meaning)
		}
	}
	fmt.Fprintf(output_file, "So that is %d cards from the grad dataset:%d\n\n", counter_grad, grad_len)

	// 6
	counter_guru := 0
	for _, one_card = range fileOfCardsGuru {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s, found in dataset: guru\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			counter_guru++
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s, from dataset: guru\n", one_card.Kanji, one_card.Meaning)
		}
	}
	fmt.Fprintf(output_file, "So that is %d cards from the guru dataset:%d\n\n", counter_guru, guru_len)

	// 7
	counter_init := 0
	for _, one_card = range fileOfCardsInitiate {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s, found in dataset: init\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			counter_init++
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s, from dataset: init\n", one_card.Kanji, one_card.Meaning)
		}
	}
	fmt.Fprintf(output_file, "So that is %d cards from the init dataset:%d\n\n", counter_init, init_len)

	// 8
	counter_master := 0
	for _, one_card = range fileOfCardsMaster {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s, found in dataset: master\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			counter_master++
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s, from dataset: master\n", one_card.Kanji, one_card.Meaning)
		}
	}
	fmt.Fprintf(output_file, "So that is %d cards from the master dataset:%d\n\n", counter_master, master_len)

	// 9
	counter_novice := 0
	for _, one_card = range fileOfCardsNovice {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s, found in dataset: novice\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			counter_novice++
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s, from dataset: novice\n", one_card.Kanji, one_card.Meaning)
		}
	}
	fmt.Fprintf(output_file, "So that is %d cards from the novice dataset:%d\n\n", counter_novice, novice_len)

	total_of_all_cards_in_all_datasets_excluding_duplicates := counter_claude + counter_current + counter_data +
		counter_fresh + counter_grad + counter_guru + counter_init + counter_master + counter_novice

	fmt.Fprintf(output_file, "So that is %d cards total, excluding the %d cards in data_future\n\n",
		total_of_all_cards_in_all_datasets_excluding_duplicates, future_len)

	fmt.Fprintf(output_file, "And then there are also the cards in data_file100_maybe:%d \n\n\n", data_file100_maybe_len)

	// 10
	counter_future := 0
	for _, one_card = range fileOf_Future {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s, found in dataset: future\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			counter_future++
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s, from dataset: future\n", one_card.Kanji, one_card.Meaning)
		}
	}
	fmt.Fprintf(output_file, "So that is %d cards from the future dataset:%d\n\n", counter_future, future_len)

	//
	//
	// 11
	counter_maybe := 0
	for _, one_card = range data_file100_maybe {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s, found in dataset: maybe\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			counter_maybe++
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s, from dataset: maybe\n", one_card.Kanji, one_card.Meaning)
		}
	}
	fmt.Fprintf(output_file, "So that is %d cards from the maybe dataset:%d\n\n",
		counter_maybe, data_file100_maybe_len)

	// and then there is :
	// beauty_len = len(data_beauty)

	//
	//
	// 11
	counter_beauty := 0
	for _, one_card = range data_beauty {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s, found in dataset: beauty\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			counter_beauty++
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s, from dataset: beauty\n", one_card.Kanji, one_card.Meaning)
		}
	}
	fmt.Fprintf(output_file, "So that is %d cards from the beauty dataset:%d\n\n",
		counter_beauty, beauty_len)

}
