package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func countSLOC() {
	numberOfFilesExplored := 0
	// Create a regex to match the base file name
	re := regexp.MustCompile(`[^/]+$`)
	var match string

	fmt.Println("Expecting to process 18 files\n")

	filenameOfThisFile5 := "/Users/quasar/Kanji-new/elementsOfsloc.go"
	blankLines5, singleComments5, commentBlock15, commentBlock25, commentBlock35, runes15, runes25, runes35, totalLines5,
		nonEmptyLines5 := reportSLOCstats(filenameOfThisFile5)
	match = re.FindString(filenameOfThisFile5)
	fmt.Printf("--processed %s which was %d lines long\n ", match, totalLines5)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)
	//
	filenameOfThisFile2 := "/Users/quasar/Kanji-new/findInFiles.go"
	blankLines2, singleComments2, commentBlock12, commentBlock22, commentBlock32, runes12, runes22, runes32, totalLines2,
		nonEmptyLines2 := reportSLOCstats(filenameOfThisFile2)
	match = re.FindString(filenameOfThisFile2)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines2)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)
	//
	filenameOfThisFile42 := "/Users/quasar/Kanji-new/formatter.go"
	blankLines42, singleComments42, commentBlock412, commentBlock422, commentBlock432, runes412, runes422, runes432, totalLines42,
		nonEmptyLines42 := reportSLOCstats(filenameOfThisFile42)
	match = re.FindString(filenameOfThisFile42)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines42)
	//
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)
	//
	filenameOfThisFile6 := "/Users/quasar/Kanji-new/functions.go"
	blankLines6, singleComments6, commentBlock16, commentBlock26, commentBlock36, runes16, runes26, runes36, totalLines6,
		nonEmptyLines6 := reportSLOCstats(filenameOfThisFile6)
	match = re.FindString(filenameOfThisFile6)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines6)
	//
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)
	//
	filenameOfThisFile55 := "/Users/quasar/Kanji-new/functionsFromMain.go"
	blankLines55, singleComments55, commentBlock155, commentBlock255, commentBlock355, runes155, runes255, runes355, totalLines55,
		nonEmptyLines55 := reportSLOCstats(filenameOfThisFile55)
	match = re.FindString(filenameOfThisFile55)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines55)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)
	//
	filenameOfThisFile18 := "/Users/quasar/Kanji-new/game_functions.go"
	blankLines18, singleComments18, commentBlock118, commentBlock218, commentBlock318, runes118, runes218, runes318, totalLines18,
		nonEmptyLines18 := reportSLOCstats(filenameOfThisFile18)
	match = re.FindString(filenameOfThisFile18)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines18)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)
	//
	filenameOfThisFile7 := "/Users/quasar/Kanji-new/globalVariables.go"
	blankLines7, singleComments7, commentBlock17, commentBlock27, commentBlock37, runes17, runes27, runes37, totalLines7,
		nonEmptyLines7 := reportSLOCstats(filenameOfThisFile7)
	match = re.FindString(filenameOfThisFile7)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines7)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)
	//
	filenameOfThisFile14 := "/Users/quasar/Kanji-new/line_parsing.go"
	blankLines14, singleComments14, commentBlock114, commentBlock214, commentBlock314, runes114, runes214, runes314, totalLines14,
		nonEmptyLines14 := reportSLOCstats(filenameOfThisFile14)
	match = re.FindString(filenameOfThisFile14)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines14)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)
	//
	filenameOfThisFile514 := "/Users/quasar/Kanji-new/list_all_chars_lff.go"
	blankLines514, singleComments514, commentBlock5114, commentBlock5214, commentBlock5314, runes5114, runes5214, runes5314, totalLines514,
		nonEmptyLines514 := reportSLOCstats(filenameOfThisFile514)
	match = re.FindString(filenameOfThisFile514)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines514)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)
	//
	filenameOfThisFile8 := "/Users/quasar/Kanji-new/locateCard.go"
	blankLines8, singleComments8, commentBlock18, commentBlock28, commentBlock38, runes18, runes28, runes38, totalLines8,
		nonEmptyLines8 := reportSLOCstats(filenameOfThisFile8)
	match = re.FindString(filenameOfThisFile8)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines8)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)
	//
	filenameOfThisFile1 := "/Users/quasar/Kanji-new/main.go"
	blankLines1, singleComments1, commentBlock11, commentBlock21, commentBlock31, runes11, runes21, runes31, totalLines1,
		nonEmptyLines1 := reportSLOCstats(filenameOfThisFile1)
	match = re.FindString(filenameOfThisFile1)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines1)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)
	//
	filenameOfThisFile9 := "/Users/quasar/Kanji-new/memoryFunctions.go"
	blankLines9, singleComments9, commentBlock19, commentBlock29, commentBlock39, runes19, runes29, runes39, totalLines9,
		nonEmptyLines9 := reportSLOCstats(filenameOfThisFile9)
	match = re.FindString(filenameOfThisFile9)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines9)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)

	filenameOfThisFile10a := "/Users/quasar/Kanji-new/new_stats_functions.go"
	blankLines10a, singleComments10a, commentBlock110a, commentBlock210a, commentBlock310a, runes110a, runes210a, runes310a,
		totalLines10a, nonEmptyLines10a := reportSLOCstats(filenameOfThisFile10a)
	match = re.FindString(filenameOfThisFile10a)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines10a)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)

	filenameOfThisFile10 := "/Users/quasar/Kanji-new/objectsAndMethods.go"
	blankLines10, singleComments10, commentBlock110, commentBlock210, commentBlock310, runes110, runes210, runes310, totalLines10,
		nonEmptyLines10 := reportSLOCstats(filenameOfThisFile10)
	match = re.FindString(filenameOfThisFile10)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines10)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)

	filenameOfThisFile11 := "/Users/quasar/Kanji-new/pick_a_card.go"
	blankLines11, singleComments11, commentBlock111, commentBlock211, commentBlock311, runes111, runes211, runes311, totalLines11,
		nonEmptyLines11 := reportSLOCstats(filenameOfThisFile11)
	match = re.FindString(filenameOfThisFile11)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines11)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)

	filenameOfThisFile15 := "/Users/quasar/Kanji-new/prompts&directions.go"
	blankLines15, singleComments15, commentBlock115, commentBlock215, commentBlock315, runes115, runes215, runes315, totalLines15,
		nonEmptyLines15 := reportSLOCstats(filenameOfThisFile15)
	match = re.FindString(filenameOfThisFile15)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines15)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)

	filenameOfThisFile17 := "/Users/quasar/Kanji-new/prompts.go"
	blankLines17, singleComments17, commentBlock117, commentBlock217, commentBlock317, runes117, runes217, runes317, totalLines17,
		nonEmptyLines17 := reportSLOCstats(filenameOfThisFile17)
	match = re.FindString(filenameOfThisFile17)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines17)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)

	filenameOfThisFile16 := "/Users/quasar/Kanji-new/statsFunctions.go"
	blankLines, singleComments, commentBlock01, commentBlock02, commentBlock03, runes01, runes02, runes03, totalLines16,
		nonEmptyLines16 := reportSLOCstats(filenameOfThisFile16)
	match = re.FindString(filenameOfThisFile16)
	fmt.Printf("--processed %s which was %d lines long \n ", match, totalLines16)
	numberOfFilesExplored++
	fmt.Printf("%d files processed so far\n", numberOfFilesExplored)

	gt := totalLines16 + totalLines17 + totalLines15 + totalLines11 + totalLines10 + totalLines10a + totalLines9 + totalLines1 + totalLines8 +
		totalLines514 + totalLines14 + totalLines7 + totalLines18 + totalLines55 + totalLines6 + totalLines42 + totalLines2 + totalLines5

	fmt.Printf("double-checking, Total lines in all files: %d \n\n", gt)

	TotalNumberOfFilesExplored := numberOfFilesExplored

	totalLines := totalLines1 + totalLines2 + totalLines5 + totalLines6 + totalLines7 + totalLines8 + totalLines9 + totalLines42 +
		totalLines10 + totalLines14 + totalLines15 + totalLines16 + totalLines11 + totalLines10a + totalLines55 + totalLines17 + totalLines18 +
		totalLines514

	nonEmptyLines := nonEmptyLines1 + nonEmptyLines2 + nonEmptyLines5 + nonEmptyLines6 + nonEmptyLines7 + nonEmptyLines8 + nonEmptyLines9 +
		nonEmptyLines10 + nonEmptyLines14 + nonEmptyLines15 + nonEmptyLines16 + nonEmptyLines11 + nonEmptyLines10a + nonEmptyLines55 +
		nonEmptyLines17 + nonEmptyLines18 + nonEmptyLines42 + nonEmptyLines514

	blankLinesTotal := blankLines15 + blankLines + blankLines14 + blankLines11 + blankLines10 + blankLines9 + blankLines8 + blankLines7 +
		blankLines6 + blankLines5 + blankLines2 + blankLines1 + blankLines17 + blankLines18 +
		blankLines10a + blankLines55 + blankLines42 + blankLines514

	singleCommentsTotal := singleComments15 + singleComments + singleComments14 + singleComments11 + singleComments10 + singleComments9 +
		singleComments8 + singleComments7 + singleComments6 + singleComments5 + singleComments2 + singleComments42 +
		singleComments1 + singleComments17 + singleComments18 + singleComments10a + singleComments55 + singleComments514

	commentBlock1Total := commentBlock01 + commentBlock115 + commentBlock114 + commentBlock111 + commentBlock110 + commentBlock19 +
		commentBlock18 + commentBlock17 + commentBlock16 + commentBlock15 + commentBlock12 + commentBlock11 +
		commentBlock117 + commentBlock118 + commentBlock110a + commentBlock155 + commentBlock412 + commentBlock5114
	commentBlock2Total := commentBlock215 + commentBlock02 + commentBlock214 + commentBlock211 + commentBlock210 + commentBlock29 +
		commentBlock28 + commentBlock27 + commentBlock26 + commentBlock25 + commentBlock22 + commentBlock21 +
		commentBlock217 + commentBlock218 + commentBlock210a + commentBlock255 + commentBlock422 + commentBlock5214
	commentBlock3Total := commentBlock315 + commentBlock03 + commentBlock314 + commentBlock311 + commentBlock310 + commentBlock39 +
		commentBlock38 + commentBlock37 + commentBlock36 + commentBlock35 + commentBlock32 + commentBlock31 +
		commentBlock317 + commentBlock318 + commentBlock310a + commentBlock355 + commentBlock432 + commentBlock5314

	runes1Total := runes115 + runes01 + runes114 + runes111 + runes110 + runes19 + runes18 + runes17 + runes16 + runes15 + runes12 +
		runes11 + runes117 + runes118 + runes110a + runes155 + runes412 + runes5114
	runes2Total := runes215 + runes02 + runes214 + runes211 + runes210 + runes29 + runes28 + runes27 + runes26 + runes25 + runes22 +
		runes21 + runes217 + runes218 + runes210a + runes255 + runes422 + runes5214
	runes3Total := runes315 + runes03 + runes314 + runes311 + runes310 + runes39 + runes38 + runes37 + runes36 + runes35 + runes32 +
		runes31 + runes317 + runes318 + runes310a + runes355 + runes432 + runes5314

	grandTotal := blankLinesTotal + singleCommentsTotal + commentBlock2Total + runes2Total

	sumOfCodePlusNon := grandTotal + nonEmptyLines

	fmt.Printf("Total lines of Code (exclusive of data) = %d t-SLOC\n\n", totalLines)

	fmt.Printf("Total lines of executable Code = %d e-SLOC\n\n", nonEmptyLines)
	fmt.Printf("BlnkLns:%d + SnglCmLns:%d + ComBlks:%d + runes:%d = total of cmnts+spc = %d lines\n\n", blankLinesTotal,
		singleCommentsTotal, commentBlock2Total, runes2Total, grandTotal)

	fmt.Printf("Total of comments etc. + e-SLOC = %d = t-SLOC\n\n", sumOfCodePlusNon)
	fmt.Printf("")

	fmt.Printf("Total number of files: %d\n", TotalNumberOfFilesExplored)

	// elementsOfsloc.go, findInFiles.go, formatter.go, functions.go, functionsFromMain.go, game_functions.go, globalVariables.go,
	// line_parsing.go, list_all_chars_lff.go, locateCard.go, main.go, memoryFunctions.go, new_stats_functions.go, objectsAndMethods.go,
	// pick_a_card.go, prompts&directions.go, prompts.go, statsFunctions.go (18)

	if runes3Total > 0 || runes1Total > 0 || commentBlock3Total > 0 || commentBlock1Total > 0 { // if any of these was > 0
		fmt.Println("\n\n === hey we actually got something from where there should not have been anything === \n\n")
	}

}

func reportSLOCstats(filepath string) (blankLines, singleComments, commentBlock1, commentBlock2, commentBlock3,
	runes1, runes2, runes3, totalLines, sloc int) { // ::: - -

	// Patterns to identify comments, blank lines, and strings
	singleLineCommentPattern := `^\s*//`
	multiLineCommentPattern := `(?s)/\*.*?\*/`
	blankLinePattern := `^\s*$`
	stringLiteralPattern := `(?s)"(?:\\.|[^\\"])*?"|` + "`(?:\\.|[^`])*?`"

	// Compile regular expressions
	singleLineCommentRe := regexp.MustCompile(singleLineCommentPattern)
	multiLineCommentRe := regexp.MustCompile(multiLineCommentPattern)
	blankLineRe := regexp.MustCompile(blankLinePattern)
	stringLiteralRe := regexp.MustCompile(stringLiteralPattern)

	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		// return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalLines = 0
	sloc = 0
	inMultiLineComment := false
	inMultiLineString := false

	for scanner.Scan() {
		line := scanner.Text()
		totalLines++

		// ::: Check for blank lines
		if blankLineRe.MatchString(line) {
			blankLines++
			continue
		}

		// ::: Check for single-line comments
		if singleLineCommentRe.MatchString(line) {
			singleComments++
			continue
		}

		// ::: Check for multi-line comment blocks
		if inMultiLineComment {
			if strings.Contains(line, "*/") {
				inMultiLineComment = false
				line = multiLineCommentRe.ReplaceAllString(line, "")
				if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
					commentBlock1++ // Does not normally accumulate anything.
					continue
				}
			} else {
				commentBlock2++ // This is where we find lines that match.
				continue
			}
		}
		if strings.Contains(line, "/*") {
			inMultiLineComment = true
			line = multiLineCommentRe.ReplaceAllString(line, "")
			if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
				commentBlock3++ // Does not normally accumulate anything.
				continue
			}
		}

		// ::: Check for multi-line strings // string literals // Runes
		if inMultiLineString {
			if strings.Count(line, "`")%2 != 0 || strings.Count(line, "\"")%2 != 0 {
				inMultiLineString = false
				line = stringLiteralRe.ReplaceAllString(line, "")
				if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
					runes1++ // Does not normally accumulate anything.
					continue
				}
			} else {
				runes2++ // This is where we find lines that match.
				continue
			}
		}
		if strings.Count(line, "`")%2 != 0 || strings.Count(line, "\"")%2 != 0 {
			inMultiLineString = true
			line = stringLiteralRe.ReplaceAllString(line, "")
			if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
				runes3++ // Does not normally accumulate anything.
				continue
			}
		}
		sloc++
	}

	if err := scanner.Err(); err != nil {
		// return 0, 0, err
	}

	return blankLines, singleComments, commentBlock1, commentBlock2, commentBlock3, runes1, runes2, runes3, totalLines, sloc
}
