package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var firstElement1 string
var firstElement2 string

func find_in_files() {
	// Open a file to be used to log the cards that match any category file
	fileHandle, err := os.OpenFile("CardsAlreadyInCategories.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening or creating CardsAlreadyInCategories.txt:", err)
		return
	}
	defer func(fileHandle *os.File) {
		_ = fileHandle.Close()
	}(fileHandle)

	// Open the master source file for reading
	file1, err := os.Open("a_prob_done.txt")
	if err != nil {
		fmt.Println("Error opening file: a_prob_done.txt", err)
		return
	}
	defer func(file1 *os.File) {
		_ = file1.Close()
	}(file1)

	//
	// Open the category files for reading
	file2, err := os.Open("dataGraduate.go")
	if err != nil {
		fmt.Println("Error opening file: dataGraduate.go", err)
		return
	}
	defer func(file2 *os.File) {
		_ = file2.Close()
	}(file2)
	//
	//
	file3, err := os.Open("dataGuru.go")
	if err != nil {
		fmt.Println("Error opening file: dataGuru.go", err)
		return
	}
	defer func(file3 *os.File) {
		_ = file3.Close()
	}(file3)
	//
	file4, err := os.Open("dataInitiate.go")
	if err != nil {
		fmt.Println("Error opening file: dataInitiate.go", err)
		return
	}
	defer func(file4 *os.File) {
		_ = file4.Close()
	}(file4)
	//
	file5, err := os.Open("dataMaster.go")
	if err != nil {
		fmt.Println("Error opening file: dataMaster.go", err)
		return
	}
	defer func(file5 *os.File) {
		_ = file5.Close()
	}(file5)
	//
	file6, err := os.Open("dataNovice.go")
	if err != nil {
		fmt.Println("Error opening file: dataNovice.go", err)
		return
	}
	defer func(file6 *os.File) {
		_ = file6.Close()
	}(file6)

	// Create scanners to read the files line by line
	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)
	scanner3 := bufio.NewScanner(file3)
	scanner4 := bufio.NewScanner(file4)
	scanner5 := bufio.NewScanner(file5)
	scanner6 := bufio.NewScanner(file6)

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
	// Iterate through each line in file1
	for scanner1.Scan() { // *********************************   ************************************* [1]
		entire_line1 := scanner1.Text() // read a line1 ********************************************** [2a]

		// Split line1 at the opening double quote
		all_fields1 := strings.Split(entire_line1, "\"") // ************************************** [2b]

		// Extract the first element within double quotes (e.g., the character "耳")
		if len(all_fields1) >= 3 { // for every line read from 1, read a line from 2 // *************** [2c]
			firstElement1 = all_fields1[1] // why the second field ???, dont care, it works!

			if len(firstElement1) == 3 { // if it is a kanji char (occupies 3 bytes ?) // ************* [2d]
				// We have now read and trimmed our first element from file1 as: firstElement1
				//
				// Iterate through each line in file2
				// Reset scanner2 to the beginning of file2 before the next iteration
				file2.Seek(0, 0)
				for scanner2.Scan() { // **************************************************** [3]
					entire_line2 := scanner2.Text() // ************************************* [4a]

					// Split the 2 line at the opening double quote
					all_fields2 := strings.Split(entire_line2, "\"") // *****************[4b]

					// Extract the first element within double quotes (the character "耳")
					if len(all_fields2) >= 3 { // ******************************************* [4c]
						firstElement2 = all_fields2[1] // why the second field ???, dont care!
					}
					if firstElement1 == firstElement2 { // ********************************************** [5a]
						// then the card is not unique to the master source file // ********************* [5b]
						fmt.Fprintf(fileHandle, "match, all_fields1: %s \n", all_fields1)
						fmt.Fprintf(fileHandle, "match, all_fields2: %s \n", all_fields2)
						fmt.Fprintf(fileHandle, "match, firstElement1 is %s \n", firstElement1)
						fmt.Fprintf(fileHandle, "match, firstElement2 is %s \\n\\n \n\n", firstElement2)
					}
				} // 2 // end of loop 2 ******************************************************************** [6]
				scanner2 = bufio.NewScanner(file2) // This is actually needed, but why?? (dont care, it works!)
				//
				//
				file3.Seek(0, 0)
				for scanner3.Scan() {
					entire_line2 := scanner3.Text()
					all_fields2 := strings.Split(entire_line2, "\"")
					if len(all_fields2) >= 3 {
						firstElement2 = all_fields2[1]
					}
					if firstElement1 == firstElement2 {
						fmt.Fprintf(fileHandle, "match, all_fields1: %s \n", all_fields1)
						fmt.Fprintf(fileHandle, "match, all_fields2: %s \n", all_fields2)
						fmt.Fprintf(fileHandle, "match, firstElement1 is %s \n", firstElement1)
						fmt.Fprintf(fileHandle, "match, firstElement2 is %s \\n\\n \n\n", firstElement2)
					}
				} // 3 // end of loop 3
				scanner3 = bufio.NewScanner(file3) // These are actually needed, but why? (dont care, it works!)
				//
				//
				file4.Seek(0, 0)
				for scanner4.Scan() {
					entire_line2 := scanner4.Text()
					all_fields2 := strings.Split(entire_line2, "\"")
					if len(all_fields2) >= 3 {
						firstElement2 = all_fields2[1]
					}
					if firstElement1 == firstElement2 {
						fmt.Fprintf(fileHandle, "match, all_fields1: %s \n", all_fields1)
						fmt.Fprintf(fileHandle, "match, all_fields2: %s \n", all_fields2)
						fmt.Fprintf(fileHandle, "match, firstElement1 is %s \n", firstElement1)
						fmt.Fprintf(fileHandle, "match, firstElement2 is %s \\n\\n \n\n", firstElement2)
					}
				} // 4 // end of loop 4
				scanner4 = bufio.NewScanner(file4)
				//
				//
				file5.Seek(0, 0)
				for scanner5.Scan() {
					entire_line2 := scanner5.Text()
					all_fields2 := strings.Split(entire_line2, "\"")
					if len(all_fields2) >= 3 {
						firstElement2 = all_fields2[1]
					}
					if firstElement1 == firstElement2 {
						fmt.Fprintf(fileHandle, "match, all_fields1: %s \n", all_fields1)
						fmt.Fprintf(fileHandle, "match, all_fields2: %s \n", all_fields2)
						fmt.Fprintf(fileHandle, "match, firstElement1 is %s \n", firstElement1)
						fmt.Fprintf(fileHandle, "match, firstElement2 is %s \\n\\n \n\n", firstElement2)
					}
				} // 5 // end of loop 5
				scanner5 = bufio.NewScanner(file5)
				//
				//
				file6.Seek(0, 0)
				for scanner6.Scan() {
					entire_line2 := scanner6.Text()
					all_fields2 := strings.Split(entire_line2, "\"")
					if len(all_fields2) >= 3 {
						firstElement2 = all_fields2[1]
					}
					if firstElement1 == firstElement2 {
						fmt.Fprintf(fileHandle, "match, all_fields1: %s \n", all_fields1)
						fmt.Fprintf(fileHandle, "match, all_fields2: %s \n", all_fields2)
						fmt.Fprintf(fileHandle, "match, firstElement1 is %s \n", firstElement1)
						fmt.Fprintf(fileHandle, "match, firstElement2 is %s \\n\\n \n\n", firstElement2)
					}
				} // 6 // end of loop 6
				scanner6 = bufio.NewScanner(file6)
			}
		}
	} // 1 // end of loop 1 ****************************************************************************** [7]
}
