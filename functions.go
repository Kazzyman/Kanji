package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// DIRECTIVES : --------------------------------------------------------------------------------------------
//
func game_on() (game string) { // - -
	game = "on"
	gameOn = true
	fmt.Println("The game is on")

	startBeforeCall = time.Now()
	currentTime := time.Now()
	TimeOfStartFromTop = time.Now()

	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	_, err1 := fmt.Fprintf(fileHandle,
		"\n The game began at: %s \n",
		currentTime.Format("15:04:05 on Monday 01-02-2006"))
	check_error(err1)
	return game
}
func game_off() (game string) { // - -
	game = "off"
	gameOn = false
	game_duration = 998

	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	currentTime := time.Now()

	// calculate elapsed time
	t_s2 := time.Now()
	elapsed := t_s2.Sub(TimeOfStartFromTop)

	// cast time durations to a String type for Fprintf "formatted print"
	TotalRun := elapsed.String()

	fmt.Printf("\nRun time was: %s, gameOn is: %t \n\n", TotalRun, gameOn)

	// End timer and report elapsed time
	_, err1 := fmt.Fprintf(fileHandle,
		"\n The game ended at: %s  Total prompts was: %d \n",
		currentTime.Format("15:04:05 on Monday 01-02-2006"), total_prompts)
	check_error(err1)
	_, err2 := fmt.Fprintf(fileHandle,
		"\n Elapsed time of game was: %s \n",
		TotalRun)
	check_error(err2)
	return game
}

// Creates a func named check_error which takes one parameter "e" of type error
func check_error(e error) { //    - -
	if e != nil {
		panic(e) // use panic() to display error code
	}
}

func testForDirective(in string) (result bool) { // - -
	if in == "?" || // <-- If it IS a directive
		in == "reset" ||
		in == "set" ||
		in == "stat" ||
		in == "notes" ||
		in == "st" ||
		in == "dir" ||
		in == "quit" ||
		in == "q" ||
		in == "exit" ||
		in == "ex" ||
		in == "stats" ||
		in == "format" ||
		in == "rm" ||
		in == "gameon" ||
		in == "gameoff" ||
		in == "gamed" {
		// Then:
		result = true
	}
	return result
}

func respond_to_UserSuppliedDirective(in string) (prompt, objective, kind string) { // - -
	var count int
	switch in {
	case "notes":
		fmt.Println("\nOnyomi (音読み, おにょみ): the reading based on the character's original Chinese pronunciation. \n" +
			"    Onyomi readings are typically used when kanji characters are combined to form compound words, especially in words with \n" +
			"a more formal or academic context.\n" +
			"    The Onyomi reading of a kanji character in Japanese was derived from the character's original Chinese pronunciation. \n" +
			"However, over time, the pronunciation of characters has evolved differently in the two languages; so, there are often \n" +
			"significant differences in how they are now pronounced in the two languages. There might be some similarities due to \n" +
			"their shared origin, but the sounds are usually no longer identical; largely because, the pronunciation in Japanese \n" +
			"has been influenced by the Japanese language's phonetic structure and historical linguistic developments.\n" +
			"    In addition, a kanji character may have a certain pronunciation in Mandarin Chinese, but its Onyomi reading in Japanese \n" +
			"may have undergone changes in consonants, vowels, or tones. And, the meanings associated with the characters \n" +
			"might also have diverged between the two languages.\n" +
			"    So, while there may be historical and etymological connections, the modern pronunciation of Onyomi in Japanese \n" +
			"does not necessarily closely resemble how a Mandarin Chinese speaker would pronounce the corresponding characters.\n\n" +
			"Kunyomi (訓読み, くにょみ): an alternative reading based on the character's native Japanese pronunciation. \n" +
			"    When Chinese characters were introduced, the Japanese adapted them to fit their language by assigning Kunyomi readings \n" +
			"based on existing Japanese words that had similar meanings. Kunyomi readings are commonly used when the kanji character \n" +
			"appears in isolation or is combined with Hiragana characters to form native Japanese words -- allowing for the creation of \n" +
			"new terms and expressions with extra nuance. \n" +
			"    Distinguishing Homophones: As more characters were borrowed from Chinese, it became essential to distinguish between \n" +
			"homophones and provide context-specific meanings. Kunyomi readings helped in achieving this, as they often carried specific \n" +
			"nuances and meanings in the Japanese language.\n" +
			"    Influence of Japanese Phonetics: The Japanese language has a different phonetic structure from Chinese. The Kunyomi \n" +
			"readings reflect this phonetic structure, including differences in syllables, vowel sounds, and pitch accent, making the \n" +
			"pronunciation more natural for native Japanese speakers.\n" +
			"    Development of Native Vocabulary: Over centuries, Japan developed its own vocabulary and linguistic expressions. \n" +
			"Kunyomi readings evolved to accommodate these developments and ensure that Chinese characters were effectively integrated \n" +
			"into the Japanese language.\n" +
			"    In summary, the creation and use of Kunyomi readings in Japanese kanji were motivated by a desire to adapt Chinese \n" +
			"characters to the linguistic, cultural, and semantic needs of the Japanese language. This process allowed for the \n" +
			"coexistence of both Onyomi and Kunyomi readings, contributing to the rich and nuanced nature of the Japanese writing system.\n")
	case "gamed":
		fmt.Println("Enter a number for how many prompts there will be in the game")
		_, _ = fmt.Scan(&count)
		game_duration = count - 2
	case "gameon":
		// game_loop_counter ++
		game_on()
	case "gameoff":
		game_off()
	case "reset":
		// Flush (clear) the old stats and hits arrays
		cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
		cyclicArrayHits = CyclicArrayHits{}
		// Also, flush (clear) the maps
		total_prompts = 0
		//
		//goland:noinspection ALL
		fmt.Println("\nArrays and maps flushed:\n")
		fmt.Println("    cyclicArrayOfTheJcharsGottenWrong")
		fmt.Println("    cyclicArrayHits")
		fmt.Println("    frequencyMapOf_IsFineOnChars")
		//goland:noinspection ALL
		fmt.Println("    frequencyMapOf_need_workOn\n")
		fmt.Println("  And, all Game values have also been reset")
	case "quit":
		os.Exit(1)
	case "q":
		os.Exit(1)
	case "exit":
		os.Exit(1)
	case "ex":
		os.Exit(1)
	case "?":
		fmt.Printf("\n%s\n%s\n%s\n%s\n%s\n%s\n\n", aCard.Kanji, aCard.Meaning, aCard.Long_Meaning, aCard.Onyomi, aCard.Kunyomi, aCard.Vocab)
	case "stat":
		hits()
	case "st":
		hits()
	case "stats":
		hits()
	case "format":
		formatter()
	case "dir": // reDisplay the DIRECTORY OF DIRECTIVES (and instructions):
		re_display_List_of_Directives()
	case "rm":
		read_map_of_fineOn()
		read_map_of_needWorkOn()
	case "set":
		prompt, objective, kind = reSet_aCard_andThereBy_reSet_thePromptString()
		if prompt == "" {
			fmt.Println("\n That string was not found, setting to \"west\" \n")
			prompt = "西"
			objective = "west"
			kind = "Romaji"
		}
	default:
		// fmt.Println("Directive not found") // Does not work because only existent cases are passed to the switch
	}
	/*
		if prompt == "" {
			fmt.Println("\n That string was not found, setting to \"west\" \n")
			prompt = "西"
			objective = "west"
			kind = "Romaji"
		}

	*/
	return prompt, objective, kind
}

func pick_RandomCard_Assign_fields() (promptField, objective, objective_kind string) { // - -
	/*
		Kanji   string
		Meaning string
		Onyomi  string aCard.Onyomi done
		Kunyomi string
		Vocab   string

		Kata   string
		Hira   string
		Romaji string
		HiraHint string
		KataHint string
	*/

	// Random prompting
	//
	randIndex := rand.Intn(len(fileOfCardsK))
	aCard = fileOfCardsK[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
	promptField = aCard.Kanji
	objective = aCard.Meaning
	objective_kind = "Romaji"

	/*
		// Sequential prompting from the manually semi-randomized deck
		//
		for {
			if index < len(fileOfCards) {
				aCard = fileOfCards[index] // pick a 'card' from a 'deck' and store it in a global var
				promptField = aCard.Kanji
				objective = aCard.Meaning
				index++
				break
			} else {
				index = 0
				continue
			}
		}

	*/

	objective_kind = "Romaji" // i.e., "Meaning"
	/*
		//  Random prompting from all decks
		//
			randIndex := rand.Intn(len(fileOfCardsK))
			randIndexS := rand.Intn(len(fileOfCardsGuru2))
			randIndexD := rand.Intn(len(fileOfCardsGuru))

			randomFileOfCards = rand.Intn(3)

			// Kanji prompting, Meaning objective:
			if randomFileOfCards == 0 {
				aCard = fileOfCardsK[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
				promptField = aCard.Kanji
				objective = aCard.Meaning
			}
			if randomFileOfCards == 1 {
				aCard = fileOfCardsGuru2[randIndexS] // Randomly pick a 'card' from a 'deck' and store it in a global var
				promptField = aCard.Kanji
				objective = aCard.Meaning
			}
			if randomFileOfCards == 2 {
				aCard = fileOfCardsGuru[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
				promptField = aCard.Kanji
				objective = aCard.Meaning
			}

	*/

	return promptField, objective, objective_kind

}

// Handles the Directive 'set'
func reSet_aCard_andThereBy_reSet_thePromptString() (prompt, objective, objective_kind string) { //  - -
	var theMeaningOfCardToSilentlyLocate string

	fmt.Printf("\nEnter a Meaning to")
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" reSet the prompt via the corresponding Kanji card :> ")
	fmt.Printf("%s", colorReset) //
	_, _ = fmt.Scan(&theMeaningOfCardToSilentlyLocate)

	// Tentatively, prepare to Scan for user's input and attempt locating a matching 'aCard'

	// Confidently, go-looking for user's input: locate matching 'aCard'
	silentlyLocateCard(theMeaningOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
	if foundElement != nil {
		aCard = *foundElement // Set the global var-object 'aCard'
		prompt = aCard.Kanji
		objective = aCard.Meaning
	} else {
		fmt.Printf("\n %s not found \n", theMeaningOfCardToSilentlyLocate)
	}
	objective_kind = "Romaji"
	fmt.Println("")

	return prompt, objective, objective_kind
}

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
