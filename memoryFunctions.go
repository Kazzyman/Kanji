package main

import (
	"fmt"
)

func read_map() {
	if len(frequencyMapOfSeenChars) == 0 {
		fmt.Printf(colorRed)
		fmt.Printf("\nThe seenMap is empty\n")
		fmt.Printf(colorReset)
	}
	for s, f := range frequencyMapOfSeenChars {
		if s != "" {
			if s != "primedK0" {
				fmt.Printf(" --- From seenMap: string is:")
				fmt.Printf(colorCyan)
				fmt.Printf("%s", s)
				fmt.Printf(colorReset)
				fmt.Printf(", freq is:")
				fmt.Printf(colorRed)
				fmt.Printf("%d", f)
				fmt.Printf(colorReset)
				fmt.Printf("\n")
			}
		}
	}
	fmt.Println("")
	var indexIntoArray int
	for indexIntoArray < len(cyclicArrayPulls.pulls) {
		if cyclicArrayPulls.pulls[indexIntoArray] != "" {
			if cyclicArrayPulls.pulls[indexIntoArray] != "primedK0" {
				fmt.Printf("Char stored in cyclicArrayPulls: %s \n", cyclicArrayPulls.pulls[indexIntoArray])
			}
		}
		indexIntoArray++
	}
}

func read_map_of_fineOn() { //     - -
	if len(frequencyMapOf_IsFineOnChars) == 0 {
		fmt.Printf(colorRed)
		fmt.Printf("\nThe FineOn Map is empty\n")
		fmt.Printf(colorReset)
	}
	for s, f := range frequencyMapOf_IsFineOnChars {
		fmt.Printf(" --- From MapOf_IsFineOn: string is:")
		fmt.Printf(colorCyan)
		fmt.Printf("%s", s)
		fmt.Printf(colorReset)
		fmt.Printf(", freq is:")
		fmt.Printf(colorRed)
		fmt.Printf("%d", f)
		fmt.Printf(colorReset)
		fmt.Printf(" ---\n")
	}
	fmt.Println("")
}

func read_map_of_needWorkOn() { //     - -
	if len(frequencyMapOf_need_workOn) == 0 {
		fmt.Printf(colorRed)
		fmt.Printf("\nThe need_workOn map is empty\n")
		fmt.Printf(colorReset)
	}
	for s, f := range frequencyMapOf_need_workOn {
		fmt.Printf(" --- From frequencyMapOf_need_workOn: string is:")
		fmt.Printf(colorCyan)
		fmt.Printf("%s", s)
		fmt.Printf(colorReset)
		fmt.Printf(", freq is:")
		fmt.Printf(colorRed)
		fmt.Printf("%d", f)
		fmt.Printf(colorReset)
		fmt.Printf(" ---\n")
	}
	fmt.Println("")
}
