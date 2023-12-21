package main

import (
	"fmt"
	"os"
)

var one_card charSetStructKanji
var used_slice []string // To keep track of kanji chars already found and added to the output file.
var one_field string
var already_loged bool

// lff as a dir
// The purpose of this func is to parse all data sets, and compile a tally of kanji+meaning which would exclude duplicates.
// Used to scan all data_category files and build a list of all unique cards (first line)
func list_from_files() {
	used_slice = append(used_slice, "ricky")
	already_loged = false
	//
	// Open a file to be used to log the tally of unique cards
	output_file, err := os.OpenFile("ListOfCards.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening or creating ListOfCards.txt:", err)
		return
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
	}
	defer func(output_file *os.File) {
		_ = output_file.Close()
	}(output_file)
	for _, one_card = range claude {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s claude\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s claude\n", one_card.Kanji, one_card.Meaning)
		}
	}

	for _, one_card = range fileOfCardsGraduate {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s grad\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s grad\n", one_card.Kanji, one_card.Meaning)
		}
	}

}

// lff as a dir
// The purpose of this func is to parse all data sets, and compile a tally of kanji+meaning which would exclude duplicates.
// Used to scan all data_category files and build a list of all unique cards (first line)
func list_from_filesO() {
	used_slice = append(used_slice, "ricky")
	already_loged = false
	//
	// Open a file to be used to log the tally of unique cards
	output_file, err := os.OpenFile("ListOfCards.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening or creating ListOfCards.txt:", err)
		return
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
	}
	defer func(output_file *os.File) {
		_ = output_file.Close()
	}(output_file)
	for _, one_card = range fileOfCardsMaster {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s master\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s master\n", one_card.Kanji, one_card.Meaning)
		}
	}
	for _, one_card = range fileOfCardsGuru {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s guru\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s guru\n", one_card.Kanji, one_card.Meaning)
		}
	}
	for _, one_card = range fileOf_Current {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s current\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s current\n", one_card.Kanji, one_card.Meaning)
		}
	}
	for _, one_card = range fileOfCardsGraduate {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s grad\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s grad\n", one_card.Kanji, one_card.Meaning)
		}
	}
	for _, one_card = range fileOfCardsInitiate {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s init\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s init\n", one_card.Kanji, one_card.Meaning)
		}
	}
	for _, one_card = range fileOfCardsNovice {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s nov\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s novice\n", one_card.Kanji, one_card.Meaning)
		}
	}
	for _, one_card = range fileOf_fresh {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s fresh\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s fresh\n", one_card.Kanji, one_card.Meaning)
		}
	}
	for _, one_card = range data_file {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s data\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s data\n", one_card.Kanji, one_card.Meaning)
		}
	}
	for _, one_card = range claude {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				fmt.Fprintf(output_file2, "%s : %s claude\n", one_card.Kanji, one_card.Meaning)
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s claude\n", one_card.Kanji, one_card.Meaning)
		}
	}
	//
}
