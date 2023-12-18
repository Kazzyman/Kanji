package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var one_card charSetStructKanji
var used_slice []string
var one_field string
var already_loged bool

// lff as a dir
func list_from_files() {
	used_slice = append(used_slice, "ricky")
	already_loged = false
	// Open a file to be used to log the cards
	output_file, err := os.OpenFile("ListOfCards.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening or creating ListOfCards.txt:", err)
		return
	}
	defer func(output_file *os.File) {
		_ = output_file.Close()
	}(output_file)
	//

	/*
		for
		read from a deck
		if card not in slice, put in slice, write to file
	*/
	for _, one_card = range fileOfCardsMaster {
		already_loged = false
		// fmt.Println(one_card.Kanji, " is one kanji from the deck")
		// make one pass through used_slice with each one_card
		for _, one_field = range used_slice {
			// fmt.Println(one_field, " is from the slice")
			if one_card.Kanji == one_field {
				already_loged = true
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s \n", one_card.Kanji, one_card.Meaning)
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
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s \n", one_card.Kanji, one_card.Meaning)
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
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s \n", one_card.Kanji, one_card.Meaning)
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
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s \n", one_card.Kanji, one_card.Meaning)
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
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s \n", one_card.Kanji, one_card.Meaning)
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
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s \n", one_card.Kanji, one_card.Meaning)
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
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s \n", one_card.Kanji, one_card.Meaning)
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
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s \n", one_card.Kanji, one_card.Meaning)
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
				break
			}
		}
		if already_loged == false {
			used_slice = append(used_slice, one_card.Kanji)
			// fmt.Printf("%s is one_card.Kanji from inner loop\n", one_card.Kanji)
			fmt.Fprintf(output_file, "%s : %s \n", one_card.Kanji, one_card.Meaning)
		}
	}
	//
}

// Used to scan all data_category files and build a list of all unique cards (first line)
func list_from_files2() {
	// Open a file to be used to log the cards
	output_file, err := os.OpenFile("ListOfCards.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening or creating ListOfCards.txt:", err)
		return
	}
	defer func(output_file *os.File) {
		_ = output_file.Close()
	}(output_file)

	/*
		// Open the master source file for reading
		source, err := os.Open("2098_lines_of_cards.go") // The input file
		if err != nil {
			fmt.Println("Error opening file: a_maybe_100.go", err)
			return
		}
		defer func(source *os.File) {
			_ = source.Close()
		}(source)

	*/

	//
	// Open the category files for reading
	grad, err := os.Open("data_Graduate_33cards.go")
	if err != nil {
		fmt.Println("Error opening file: data_Graduate_33cards.go", err)
		return
	}
	defer func(grad *os.File) {
		_ = grad.Close()
	}(grad)
	//
	guru, err := os.Open("data_Guru_6cards.go")
	if err != nil {
		fmt.Println("Error opening file: data_Guru_6cards.go", err)
		return
	}
	defer func(guru *os.File) {
		_ = guru.Close()
	}(guru)
	//
	initiate, err := os.Open("data_Initiate_120cards.go")
	if err != nil {
		fmt.Println("Error opening file: data_Initiate_120cards.go", err)
		return
	}
	defer func(initiate *os.File) {
		_ = initiate.Close()
	}(initiate)
	//
	master, err := os.Open("data_Master_41cards.go")
	if err != nil {
		fmt.Println("Error opening file: data_Master_41cards.go", err)
		return
	}
	defer func(master *os.File) {
		_ = master.Close()
	}(master)
	//
	novice, err := os.Open("data_Novice_20cards.go")
	if err != nil {
		fmt.Println("Error opening file: data_Novice_20cards.go", err)
		return
	}
	defer func(novice *os.File) {
		_ = novice.Close()
	}(novice)

	// Create scanners to read the files line by line:
	// scan_source := bufio.NewScanner(source) // "input" file being checked against older files
	// Older files:
	scan_grad := bufio.NewScanner(grad)
	scan_guru := bufio.NewScanner(guru)
	scan_initiate := bufio.NewScanner(initiate)
	scan_master := bufio.NewScanner(master)
	scan_novice := bufio.NewScanner(novice)
	// Add the other files in collection (todo here)

	/* pseudo code:
	scan(loop1:) [1] - 7
	    read a line from 1 as first-Element-1 [2]
	        scan(loop2:) [3]
	            read a line from 2 as first-Element-2 [4]
	            compare first-Element-1 to first-Element-2 [5a]
				log any match [5b]
	        loop2: exits after EOF-2 [6]
	loop1: exits after EOF-1 [7]
	*/
	// Iterate through each line in source
	for { // *********************************   ************************************* [1]
		// entire_line1 := scan_source.Text() // read a line1 ********************************************** [2a]

		// Split line1 at the opening double quote
		// all_fields1 := strings.Split(entire_line1, "\"") // ************************************** [2b]

		// We have now read and trimmed our first element from source as: firstElement1
		//
		// Iterate through each line in guru
		// .Seek resets scan_grad to the beginning of guru before the next iteration
		grad.Seek(0, 0)
		for scan_grad.Scan() { // **************************************************** [3]
			entire_line2 := scan_grad.Text() // ************************************* [4a]

			// Split the 2 line at the opening double quote
			all_fields2 := strings.Split(entire_line2, "\"") // *****************[4b]

			// Extract the first element within double quotes (the character "耳")
			if len(all_fields2) >= 3 { // ******************************************* [4c]
				firstElement2 = all_fields2[1] // why the second field ???, dont care!
			}
			if firstElement1 == firstElement2 { // ********************************************** [5a]
				// then the card is not unique to the source file // ********************* [5b]
				fmt.Fprintf(output_file, "%s is also in the Graduate category file\n", firstElement1)
			}
		} // 2 // end of loop 2 ******************************************************************** [6]
		scan_grad = bufio.NewScanner(grad) // This is actually needed, but why?? (dont care, it works!)
		//
		// .Seek resets, as was done for grad; refer to that comment
		guru.Seek(0, 0)
		for scan_guru.Scan() {
			entire_line2 := scan_guru.Text()
			all_fields2 := strings.Split(entire_line2, "\"")
			if len(all_fields2) >= 3 {
				firstElement2 = all_fields2[1]
			}
			if firstElement1 == firstElement2 {
				fmt.Fprintf(output_file, "%s is also in the Guru category file\n", firstElement1)
			}
		} // 3 // end of loop 3
		scan_guru = bufio.NewScanner(guru) // These are actually needed, but why? (dont care, it works!)
		//
		// .Seek resets, as was done for grad; refer to that comment
		initiate.Seek(0, 0)
		for scan_initiate.Scan() {
			entire_line2 := scan_initiate.Text()
			all_fields2 := strings.Split(entire_line2, "\"")
			if len(all_fields2) >= 3 {
				firstElement2 = all_fields2[1]
			}
			if firstElement1 == firstElement2 {
				fmt.Fprintf(output_file, "%s is also in the Initiate category file\n", firstElement1)
			}
		} // 4 // end of loop 4
		scan_initiate = bufio.NewScanner(initiate)

		//
		// .Seek resets, as was done for grad; refer to that comment
		master.Seek(0, 0)
		for scan_master.Scan() {
			entire_line2 := scan_master.Text()
			all_fields2 := strings.Split(entire_line2, "\"")
			if len(all_fields2) >= 3 {
				firstElement2 = all_fields2[1]
			}
			if firstElement1 == firstElement2 {
				fmt.Fprintf(output_file, "%s is also in the Master category file\n", firstElement1)
			}
		} // 5 // end of loop 5
		scan_master = bufio.NewScanner(master)
		//
		// .Seek resets, as was done for grad; refer to that comment
		novice.Seek(0, 0)
		for scan_novice.Scan() {
			entire_line2 := scan_novice.Text()
			all_fields2 := strings.Split(entire_line2, "\"")
			if len(all_fields2) >= 3 {
				firstElement2 = all_fields2[1]
			}
			if firstElement1 == firstElement2 {
				fmt.Fprintf(output_file, "%s is also in the Novice category file\n", firstElement1)
			}
		} // 6 // end of loop 6
		scan_novice = bufio.NewScanner(novice)

	} // 1 // end of loop 1 ****************************************************************************** [7]
}
