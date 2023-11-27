package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var firstElement1 string
var firstElement2 string

// Used to scan all data_category files to see if cards in a source file are already in one or more category files
func find_in_files() {
	// Open a file to be used to log the cards that match in any category file
	output_file, err := os.OpenFile("CardsAlreadyInCategories.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening or creating CardsAlreadyInCategories.txt:", err)
		return
	}
	defer func(output_file *os.File) {
		_ = output_file.Close()
	}(output_file)

	// Open the master source file for reading
	source, err := os.Open("2098_lines_of_cards.go") // The input file
	if err != nil {
		fmt.Println("Error opening file: a_maybe_100.go", err)
		return
	}
	defer func(source *os.File) {
		_ = source.Close()
	}(source)

	//
	// Open the category files for reading
	grad, err := os.Open("dataGraduate_165.go")
	if err != nil {
		fmt.Println("Error opening file: dataGraduate.go", err)
		return
	}
	defer func(grad *os.File) {
		_ = grad.Close()
	}(grad)
	//
	guru, err := os.Open("dataGuru_28.go")
	if err != nil {
		fmt.Println("Error opening file: dataGuru.go", err)
		return
	}
	defer func(guru *os.File) {
		_ = guru.Close()
	}(guru)
	//
	initiate, err := os.Open("dataInitiate_194_of_4124.go")
	if err != nil {
		fmt.Println("Error opening file: dataInitiate.go", err)
		return
	}
	defer func(initiate *os.File) {
		_ = initiate.Close()
	}(initiate)
	//
	master, err := os.Open("dataMaster_111_of_147.go")
	if err != nil {
		fmt.Println("Error opening file: dataMaster.go", err)
		return
	}
	defer func(master *os.File) {
		_ = master.Close()
	}(master)
	//
	novice, err := os.Open("dataNovice_88.go")
	if err != nil {
		fmt.Println("Error opening file: dataNovice.go", err)
		return
	}
	defer func(novice *os.File) {
		_ = novice.Close()
	}(novice)

	// Create scanners to read the files line by line:
	scan_source := bufio.NewScanner(source) // "input" file being checked against older files
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
	for scan_source.Scan() { // *********************************   ************************************* [1]
		entire_line1 := scan_source.Text() // read a line1 ********************************************** [2a]

		// Split line1 at the opening double quote
		all_fields1 := strings.Split(entire_line1, "\"") // ************************************** [2b]

		// Extract the first element within double quotes (e.g., the character "耳")
		if len(all_fields1) >= 3 { // for every line read from 1, read a line from 2 // *************** [2c]
			firstElement1 = all_fields1[1] // why the second field ???, dont care, it works!

			if len(firstElement1) == 3 { // if it is a kanji char (occupies 3 bytes ?) // ************* [2d]
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
			}
		}
	} // 1 // end of loop 1 ****************************************************************************** [7]
}
