package main

import (
	"fmt"
)

func read_map_of_fineOn() { //     - -
	if len(frequencyMapOf_IsFineOnChars) == 0 {
		fmt.Print(colorRed, "\nThe FineOn Map is empty\n", colorReset)
	}
	for s, f := range frequencyMapOf_IsFineOnChars {
		fmt.Print(" --- From MapOf_IsFineOn: string is:", colorCyan, s, colorReset)
		fmt.Print(", freq is:", colorRed, f, colorReset, " ---\n")
	}
	fmt.Println("")
}

func read_already_used_map() { // Intentionally unused
	if len(already_used_map) == 0 {
		fmt.Printf(colorRed)
		fmt.Printf("\nalready_used_map is empty; probably it was reset after having already used every card in deck\n")
		fmt.Printf(colorReset)
	}
	for s, f := range already_used_map {
		fmt.Printf(" --- already_used_map: string is:")
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

func display_already_used_map() {
	fmt.Println("contents of already_used_map:")
	for str, f := range already_used_map {
		fmt.Printf("%s,%d \n", str, f)
	}
}

func clearMap() {
	already_used_map = make(map[string]int) // global
}

var stringFromMap string

var freqInMap int // Used also in: [is_pick_novel2(kanji string) bool] - -

func is_pick_novel(kanji string) bool {
	// for stringFromMap, freqInMap = range already_used_map {
	// for stringFromMap, _ = range already_used_map { ... cursor sugested 
	for stringFromMap = range already_used_map {
		// fmt.Printf("in is_pick_novel str:%s, f:%d \n", stringFromMap, freqInMap)

		if kanji == stringFromMap {
			// fmt.Printf("returning false, picked was:%s, map had:%s\n", kanji, stringFromMap)
			// Having either found the pick in the map prior to reading the entire map, or as the map's last element ...
			return false
		}
	} // having read the entire map, and not having found a previously used pick in said map ...
	// fmt.Printf("returning true, picked was:%s, map had:%s\n", kanji, stringFromMap)
	return true
}
