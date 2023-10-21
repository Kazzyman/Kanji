package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Use to format new elements for the various fileOfCardsX
func formatter() {
	fileHandle, err := os.OpenFile("formattedElements.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening or creating formattedElements.txt:", err)
		return
	}
	defer func(fileHandle *os.File) {
		_ = fileHandle.Close()
	}(fileHandle)

	// Open the input file
	file, err := os.Open("unformattedElements.txt")
	if err != nil {
		fmt.Println("Error opening file: unformattedElements: ", err)
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into its six (any number of) ';' delimited fields
		fields := strings.Split(line, ";")

		/*
				// Process the second field - was used on source file with only one meaning field (usually a capitalized sentence)
				secondFieldWords := strings.Fields(fields[1]) // Split the second field into its composite words
				// Take the first word of that second field, convert it to lowercase, and update that second field
				meaning := strings.ToLower(secondFieldWords[0]) // secondFieldWords[0] is the first word of the original sentence
			// Format the fields
				formattedLine := fmt.Sprintf("{\"%s\", \"%s\", \"%s\", \"%s\", \"%s\", \"%s\"},\n",
					fields[0], meaning, fields[1], fields[2], fields[3], fields[4]) // Generating six 六 fields
		*/

		// Check the length (the total number of the ';' delimited fields in the line) - should be 六
		if len(fields) > 6 {
			fmt.Printf("\n The line:%s has too many fields \n", line)
		} else if len(fields) < 6 {
			fmt.Printf("\n The line:%s has too few fields \n", line)
		} else {
			// Therefore the line had exactly 6 六 fields, and, is probably sans errors in its composition, so ...
			formattedLine := fmt.Sprintf("{\"%s\", \"%s\", \"%s\", \"%s\", \"%s\", \"%s\"},\n",
				fields[0], fields[1], fields[2], fields[3], fields[4], fields[5])

			// Print each formatted line to a file
			_, err := fmt.Fprintf(fileHandle, formattedLine)

			if err != nil {
				return
			}
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
