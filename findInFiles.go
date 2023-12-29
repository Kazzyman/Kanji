package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var firstElement1 string
var firstElement2 string

// Dir : fif
// Used to scan all of the data_category files to see if cards in the "source" file (claude) are already in one or more 
// ... of the other category files (current, data, fresh, graduate, guru, initiate, master, or novice) 
func find_in_files() {
	// Open a new output file to be used to log the cards that match in any category file
	output_file, err := os.OpenFile("fif_AlreadyInOthers.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening or creating fif_AlreadyInOthers.txt:", err)
		return
	} else {
		//
		//
		fmt.Fprintf(output_file, "This is just a list from comparing claude with each of the other files\n\n")
	}
	defer func(output_file *os.File) {
		_ = output_file.Close()
	}(output_file)
	//
	// Claude will be acting as the source file
	// Open the "source" file (claude) for reading
	source, err := os.Open("data_claude_298_cards.go") // The input file
	if err != nil {
		fmt.Println("Error opening file: data_claude_298_cards.go", err)
		return
	}
	defer func(source *os.File) {
		_ = source.Close()
	}(source)
	//
	// Open all the other data category files which will be compared against the source (claude) 
	// Open the category files for reading
	//
	// Current 
	current, err := os.Open("data_current_24cards.go")
	if err != nil {
		fmt.Println("Error opening file: data_current_24cards.go", err)
		return
	}
	defer func(current *os.File) {
		_ = current.Close()
	}(current)
	//
	// Data 
	data, err := os.Open("data_data_184cards.go")
	if err != nil {
		fmt.Println("Error opening file: data_data_184cards.go", err)
		return
	}
	defer func(data *os.File) {
		_ = data.Close()
	}(data)
	//
	// Fresh 
	fresh, err := os.Open("data_fresh_85cards.go")
	if err != nil {
		fmt.Println("Error opening file: data_fresh_85cards.go", err)
		return
	}
	defer func(fresh *os.File) {
		_ = fresh.Close()
	}(fresh)
	//
	// Graduate 
	grad, err := os.Open("data_Graduate_33cards.go")
	if err != nil {
		fmt.Println("Error opening file: data_Graduate_33cards.go", err)
		return
	}
	defer func(grad *os.File) {
		_ = grad.Close()
	}(grad)
	//
	// Guru
	guru, err := os.Open("data_Guru_6cards.go")
	if err != nil {
		fmt.Println("Error opening file: data_Guru_6cards.go", err)
		return
	}
	defer func(guru *os.File) {
		_ = guru.Close()
	}(guru)
	//
	// Initiate 
	initiate, err := os.Open("data_Initiate_119cards.go")
	if err != nil {
		fmt.Println("Error opening file: data_Initiate_119cards.go", err)
		return
	}
	defer func(initiate *os.File) {
		_ = initiate.Close()
	}(initiate)
	//
	// Master
	master, err := os.Open("data_Master_46cards.go")
	if err != nil {
		fmt.Println("Error opening file: data_Master_46cards.go", err)
		return
	}
	defer func(master *os.File) {
		_ = master.Close()
	}(master)
	//
	// Novice
	novice, err := os.Open("data_Novice_23cards.go")
	if err != nil {
		fmt.Println("Error opening file: data_Novice_23cards.go", err)
		return
	}
	defer func(novice *os.File) {
		_ = novice.Close()
	}(novice)

	// Create scanners to read the files line by line:
	scan_source := bufio.NewScanner(source) // "input" file (claude) is being checked against the other files
	// The other files:
	scan_current := bufio.NewScanner(current)
	scan_data := bufio.NewScanner(data)
	scan_fresh := bufio.NewScanner(fresh)
	scan_grad := bufio.NewScanner(grad)
	scan_guru := bufio.NewScanner(guru)
	scan_initiate := bufio.NewScanner(initiate)
	scan_master := bufio.NewScanner(master)
	scan_novice := bufio.NewScanner(novice)

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
	// Iterate through each line in the source file (claude)
	for scan_source.Scan() { // *********************************   ************************************* [1]
		entire_line1 := scan_source.Text() // read a line1 ********************************************** [2a]

		// Split line1 at the opening double quote
		all_fields1 := strings.Split(entire_line1, "\"") // ************************************** [2b]

		// Extract the first element within double quotes (e.g., the character "耳")
		if len(all_fields1) >= 3 { // for every line read from 1, read a line from 2 // *************** [2c]
			firstElement1 = all_fields1[1] // why the second field ???, don't care, it works!

			if len(firstElement1) == 3 { // if it is a kanji char (occupies 3 bytes ?) // ************* [2d]
				// We have now read and trimmed our first element from the source file (claude) as: firstElement1
				//
				// Current
				// Iterate through each line in current
				// .Seek resets scan_current to the beginning of current before the next iteration
				current.Seek(0, 0)
				for scan_current.Scan() { // **************************************************** [3]
					entire_line2 := scan_current.Text() // ************************************* [4a]

					// Split the 2 line at the opening double quote
					all_fields2 := strings.Split(entire_line2, "\"") // *****************[4b]

					// Extract the first element within double quotes (the character, i.e. "耳")
					if len(all_fields2) >= 3 { // ******************************************* [4c]
						firstElement2 = all_fields2[1]
					}
					if firstElement1 == firstElement2 { // ********************************************** [5a]
						// then the card is not unique to the source file // ********************* [5b]
						fmt.Fprintf(output_file, "%s is also in the Current category file\n", firstElement1)
					}
				} // 2 // end of loop 2 ? ******************************************************************** [6]
				scan_current = bufio.NewScanner(current) // These are actually needed, but why? (don't care, it works!)
				//
				// Data
				data.Seek(0, 0)
				for scan_data.Scan() {
					entire_line2 := scan_data.Text()
					all_fields2 := strings.Split(entire_line2, "\"")
					if len(all_fields2) >= 3 {
						firstElement2 = all_fields2[1]
					}
					if firstElement1 == firstElement2 {
						fmt.Fprintf(output_file, "%s is also in the Data category file\n", firstElement1)
					}
				}
				scan_data = bufio.NewScanner(data) // These are actually needed, but why? (don't care, it works!)
				//
				// Fresh
				fresh.Seek(0, 0)
				for scan_fresh.Scan() {
					entire_line2 := scan_fresh.Text()
					all_fields2 := strings.Split(entire_line2, "\"")
					if len(all_fields2) >= 3 {
						firstElement2 = all_fields2[1]
					}
					if firstElement1 == firstElement2 {
						fmt.Fprintf(output_file, "%s is also in the Fresh category file\n", firstElement1)
					}
				}
				scan_fresh = bufio.NewScanner(fresh) // These are actually needed, but why? (don't care, it works!)
				//
				// Grad
				grad.Seek(0, 0)
				for scan_grad.Scan() {
					entire_line2 := scan_grad.Text()
					all_fields2 := strings.Split(entire_line2, "\"")
					if len(all_fields2) >= 3 {
						firstElement2 = all_fields2[1] // why the second field ???, don't care!
					}
					if firstElement1 == firstElement2 {
						fmt.Fprintf(output_file, "%s is also in the Graduate category file\n", firstElement1)
					}
				}
				scan_grad = bufio.NewScanner(grad) // This is actually needed, but why?? (don't care, it works!)
				//
				// Guru
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
				}
				scan_guru = bufio.NewScanner(guru) // These are actually needed, but why? (don't care, it works!)
				//
				// Initiate
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
				}
				scan_initiate = bufio.NewScanner(initiate)
				//
				// Master
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
				}
				scan_master = bufio.NewScanner(master)
				//
				// Novice
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
				}
				scan_novice = bufio.NewScanner(novice)
			}
		}
	} // 1 // end of loop 1 ****************************************************************************** [7]
}
