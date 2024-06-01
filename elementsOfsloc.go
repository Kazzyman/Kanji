package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var fileExplored int

func countSLOC() {
	numberOfFilesExplored := 0

	filenameOfThisFile5 := "/Users/quasar/Kanji-new/elementsOfsloc.go"
	blankLines5, singleComments5, commentBlock15, commentBlock25, commentBlock35, runes15, runes25, runes35, totalLines5, nonEmptyLines5 := reportSLOCstats(filenameOfThisFile5)
	numberOfFilesExplored++

	filenameOfThisFile2 := "/Users/quasar/Kanji-new/findInFiles.go"
	blankLines2, singleComments2, commentBlock12, commentBlock22, commentBlock32, runes12, runes22, runes32, totalLines2, nonEmptyLines2 := reportSLOCstats(filenameOfThisFile2)
	numberOfFilesExplored++

	filenameOfThisFile6 := "/Users/quasar/Kanji-new/functions.go"
	blankLines6, singleComments6, commentBlock16, commentBlock26, commentBlock36, runes16, runes26, runes36, totalLines6, nonEmptyLines6 := reportSLOCstats(filenameOfThisFile6)
	numberOfFilesExplored++

	filenameOfThisFile7 := "/Users/quasar/Kanji-new/globalVariables.go"
	blankLines7, singleComments7, commentBlock17, commentBlock27, commentBlock37, runes17, runes27, runes37, totalLines7, nonEmptyLines7 := reportSLOCstats(filenameOfThisFile7)
	numberOfFilesExplored++

	filenameOfThisFile14 := "/Users/quasar/Kanji-new/line_parsing.go"
	blankLines14, singleComments14, commentBlock114, commentBlock214, commentBlock314, runes114, runes214, runes314, totalLines14, nonEmptyLines14 := reportSLOCstats(filenameOfThisFile14)
	numberOfFilesExplored++

	filenameOfThisFile8 := "/Users/quasar/Kanji-new/locateCard.go"
	blankLines8, singleComments8, commentBlock18, commentBlock28, commentBlock38, runes18, runes28, runes38, totalLines8, nonEmptyLines8 := reportSLOCstats(filenameOfThisFile8)
	numberOfFilesExplored++

	filenameOfThisFile1 := "/Users/quasar/Kanji-new/main.go"
	blankLines1, singleComments1, commentBlock11, commentBlock21, commentBlock31, runes11, runes21, runes31, totalLines1, nonEmptyLines1 := reportSLOCstats(filenameOfThisFile1)
	numberOfFilesExplored++

	filenameOfThisFile9 := "/Users/quasar/Kanji-new/memoryFunctions.go"
	blankLines9, singleComments9, commentBlock19, commentBlock29, commentBlock39, runes19, runes29, runes39, totalLines9, nonEmptyLines9 := reportSLOCstats(filenameOfThisFile9)
	numberOfFilesExplored++

	filenameOfThisFile10a := "/Users/quasar/Kanji-new/new_stats_functions.go"
	blankLines10a, singleComments10a, commentBlock110a, commentBlock210a, commentBlock310a, runes110a, runes210a, runes310a, totalLines10a, nonEmptyLines10a := reportSLOCstats(filenameOfThisFile10a)
	numberOfFilesExplored++

	filenameOfThisFile10 := "/Users/quasar/Kanji-new/objectsAndMethods.go"
	blankLines10, singleComments10, commentBlock110, commentBlock210, commentBlock310, runes110, runes210, runes310, totalLines10, nonEmptyLines10 := reportSLOCstats(filenameOfThisFile10)
	numberOfFilesExplored++
	filenameOfThisFile11 := "/Users/quasar/Kanji-new/pick_a_card.go"
	blankLines11, singleComments11, commentBlock111, commentBlock211, commentBlock311, runes111, runes211, runes311, totalLines11, nonEmptyLines11 := reportSLOCstats(filenameOfThisFile11)
	numberOfFilesExplored++
	filenameOfThisFile15 := "/Users/quasar/Kanji-new/prompts&directions.go"
	blankLines15, singleComments15, commentBlock115, commentBlock215, commentBlock315, runes115, runes215, runes315, totalLines15, nonEmptyLines15 := reportSLOCstats(filenameOfThisFile15)
	numberOfFilesExplored++
	filenameOfThisFile16 := "/Users/quasar/Kanji-new/statsFunctions.go"
	blankLines, singleComments, commentBlock01, commentBlock02, commentBlock03, runes01, runes02, runes03, totalLines16, nonEmptyLines16 := reportSLOCstats(filenameOfThisFile16)
	numberOfFilesExplored++
	filenameOfThisFile55 := "/Users/quasar/Kanji-new/functionsFromMain.go"
	blankLines55, singleComments55, commentBlock155, commentBlock255, commentBlock355, runes155, runes255, runes355, totalLines55, nonEmptyLines55 := reportSLOCstats(filenameOfThisFile55)
	numberOfFilesExplored++
	filenameOfThisFile18 := "/Users/quasar/Kanji-new/game_functions.go"
	blankLines18, singleComments18, commentBlock118, commentBlock218, commentBlock318, runes118, runes218, runes318, totalLines18, nonEmptyLines18 := reportSLOCstats(filenameOfThisFile18)
	numberOfFilesExplored++
	filenameOfThisFile17 := "/Users/quasar/Kanji-new/prompts.go"
	blankLines17, singleComments17, commentBlock117, commentBlock217, commentBlock317, runes117, runes217, runes317, totalLines17, nonEmptyLines17 := reportSLOCstats(filenameOfThisFile17)
	numberOfFilesExplored++
	totalLines := totalLines1 + totalLines2 + totalLines5 + totalLines6 + totalLines7 + totalLines8 + totalLines9 +
		totalLines10 + totalLines14 + totalLines15 + totalLines16 + totalLines11 + totalLines10a + totalLines55 + totalLines17 + totalLines18
	numberOfFilesExplored++
	nonEmptyLines := nonEmptyLines1 + nonEmptyLines2 + nonEmptyLines5 + nonEmptyLines6 + nonEmptyLines7 + nonEmptyLines8 + nonEmptyLines9 +
		nonEmptyLines10 + nonEmptyLines14 + nonEmptyLines15 + nonEmptyLines16 + nonEmptyLines11 + nonEmptyLines10a + nonEmptyLines55 + nonEmptyLines17 + nonEmptyLines18
	numberOfFilesExplored++
	fileExplored = numberOfFilesExplored

	// totalLines := totalLines1 + totalLines2 + totalLines5 + totalLines6 + totalLines7 + totalLines8 + totalLines9 + totalLines10 + totalLines14 + totalLines15 + totalLines16 + totalLines11

	// nonEmptyLines := nonEmptyLines1 + nonEmptyLines2 + nonEmptyLines5 + nonEmptyLines6 + nonEmptyLines7 + nonEmptyLines8 + nonEmptyLines9 + nonEmptyLines10 + nonEmptyLines14 + nonEmptyLines15 + nonEmptyLines16 + nonEmptyLines11

	blankLinesTotal := blankLines15 + blankLines + blankLines14 + blankLines11 + blankLines10 + blankLines9 + blankLines8 + blankLines7 + blankLines6 + blankLines5 + blankLines2 + blankLines1 + blankLines17 + blankLines18 +
		blankLines10a + blankLines55

	singleCommentsTotal := singleComments15 + singleComments + singleComments14 + singleComments11 + singleComments10 + singleComments9 + singleComments8 + singleComments7 + singleComments6 + singleComments5 + singleComments2 +
		singleComments1 + singleComments17 + singleComments18 + singleComments10a + singleComments55

	commentBlock1Total := commentBlock01 + commentBlock115 + commentBlock114 + commentBlock111 + commentBlock110 + commentBlock19 + commentBlock18 + commentBlock17 + commentBlock16 + commentBlock15 + commentBlock12 + commentBlock11 +
		commentBlock117 + commentBlock118 + commentBlock110a + commentBlock155
	commentBlock2Total := commentBlock215 + commentBlock02 + commentBlock214 + commentBlock211 + commentBlock210 + commentBlock29 + commentBlock28 + commentBlock27 + commentBlock26 + commentBlock25 + commentBlock22 + commentBlock21 +
		commentBlock217 + commentBlock218 + commentBlock210a + commentBlock255
	commentBlock3Total := commentBlock315 + commentBlock03 + commentBlock314 + commentBlock311 + commentBlock310 + commentBlock39 + commentBlock38 + commentBlock37 + commentBlock36 + commentBlock35 + commentBlock32 + commentBlock31 +
		commentBlock317 + commentBlock318 + commentBlock310a + commentBlock355

	runes1Total := runes115 + runes01 + runes114 + runes111 + runes110 + runes19 + runes18 + runes17 + runes16 + runes15 + runes12 + runes11 + runes117 + runes118 + runes110a + runes155
	runes2Total := runes215 + runes02 + runes214 + runes211 + runes210 + runes29 + runes28 + runes27 + runes26 + runes25 + runes22 + runes21 + runes217 + runes218 + runes210a + runes255
	runes3Total := runes315 + runes03 + runes314 + runes311 + runes310 + runes39 + runes38 + runes37 + runes36 + runes35 + runes32 + runes31 + runes317 + runes318 + runes310a + runes355

	grandTotal := blankLinesTotal + singleCommentsTotal + commentBlock2Total + runes2Total

	sumOfCodePlusNon := grandTotal + nonEmptyLines

	fmt.Printf("Total lines of Code (exclusive of data) = %d t-SLOC\n\n", totalLines)

	fmt.Printf("Total lines of executable Code = %d e-SLOC\n\n", nonEmptyLines)
	fmt.Printf("BlnkLns:%d + SnglCmLns:%d + ComBlks:%d + runes:%d = total of cmnts+spc = %d lines\n\n", blankLinesTotal, singleCommentsTotal, commentBlock2Total, runes2Total, grandTotal)

	fmt.Printf("Total of comments etc. + e-SLOC = %d = t-SLOC\n\n", sumOfCodePlusNon)

	if runes3Total > 0 || runes1Total > 0 || commentBlock3Total > 0 || commentBlock1Total > 0 { // if any of these was > 0
		fmt.Println("\n\n === hey we actually got something from where there should not have been anything === \n\n")
	}

}

func reportSLOCstats(filepath string) (blankLines, singleComments, commentBlock1, commentBlock2, commentBlock3, runes1, runes2, runes3, totalLines, sloc int) { // ::: - -
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
			if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) { // blankLines, singleComments, commentBlock1, commentBlock2, commentBlock3, runes1, runes2, runes3
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

/*
func countSLOC() {

	filenameOfThisFile55 := "/Users/quasar/Kanji-new/functionsFromMain.go"
	totalLines55, nonEmptyLines55 := reportSLOCstats(filenameOfThisFile55)

	filenameOfThisFile5 := "/Users/quasar/Kanji-new/elementsOfsloc.go"
	totalLines5, nonEmptyLines5 := reportSLOCstats(filenameOfThisFile5)

	filenameOfThisFile2 := "/Users/quasar/Kanji-new/findInFiles.go"
	totalLines2, nonEmptyLines2 := reportSLOCstats(filenameOfThisFile2)

	filenameOfThisFile6 := "/Users/quasar/Kanji-new/functions.go"
	totalLines6, nonEmptyLines6 := reportSLOCstats(filenameOfThisFile6)

	filenameOfThisFile7 := "/Users/quasar/Kanji-new/globalVariables.go"
	totalLines7, nonEmptyLines7 := reportSLOCstats(filenameOfThisFile7)

	filenameOfThisFile14 := "/Users/quasar/Kanji-new/line_parsing.go"
	totalLines14, nonEmptyLines14 := reportSLOCstats(filenameOfThisFile14)

	filenameOfThisFile8 := "/Users/quasar/Kanji-new/locateCard.go"
	totalLines8, nonEmptyLines8 := reportSLOCstats(filenameOfThisFile8)

	filenameOfThisFile1 := "/Users/quasar/Kanji-new/main.go"
	totalLines1, nonEmptyLines1 := reportSLOCstats(filenameOfThisFile1)

	filenameOfThisFile9 := "/Users/quasar/Kanji-new/memoryFunctions.go"
	totalLines9, nonEmptyLines9 := reportSLOCstats(filenameOfThisFile9)

	filenameOfThisFile10a := "/Users/quasar/Kanji-new/new_stats_functions.go"
	totalLines10a, nonEmptyLines10a := reportSLOCstats(filenameOfThisFile10a)

	filenameOfThisFile10 := "/Users/quasar/Kanji-new/objectsAndMethods.go"
	totalLines10, nonEmptyLines10 := reportSLOCstats(filenameOfThisFile10)

	filenameOfThisFile11 := "/Users/quasar/Kanji-new/pick_a_card.go"
	totalLines11, nonEmptyLines11 := reportSLOCstats(filenameOfThisFile11)

	filenameOfThisFile15 := "/Users/quasar/Kanji-new/prompts&directions.go"
	totalLines15, nonEmptyLines15 := reportSLOCstats(filenameOfThisFile15)

	filenameOfThisFile16 := "/Users/quasar/Kanji-new/statsFunctions.go"
	totalLines16, nonEmptyLines16 := reportSLOCstats(filenameOfThisFile16)

	filenameOfThisFile18 := "/Users/quasar/Kanji-new/game_functions.go"
	totalLines18, nonEmptyLines18 := reportSLOCstats(filenameOfThisFile18)

	filenameOfThisFile17 := "/Users/quasar/Kanji-new/prompts.go"
	totalLines17, nonEmptyLines17 := reportSLOCstats(filenameOfThisFile17)

	totalLines := totalLines1 + totalLines2 + totalLines5 + totalLines6 + totalLines7 + totalLines8 + totalLines9 +
		totalLines10 + totalLines14 + totalLines15 + totalLines16 + totalLines11 + totalLines10a + totalLines55 + totalLines17 + totalLines18

	nonEmptyLines := nonEmptyLines1 + nonEmptyLines2 + nonEmptyLines5 + nonEmptyLines6 + nonEmptyLines7 + nonEmptyLines8 + nonEmptyLines9 +
		nonEmptyLines10 + nonEmptyLines14 + nonEmptyLines15 + nonEmptyLines16 + nonEmptyLines11 + nonEmptyLines10a + nonEmptyLines55 + nonEmptyLines17 + nonEmptyLines18

	fmt.Printf("Total lines of Code: %d\n", totalLines)
	fmt.Printf("Total lines of executable Code: %d\n", nonEmptyLines)

}

func reportSLOCstats(filepath string) (int, int) {
	// Patterns to identify comments, blank lines, and strings
	singleLineCommentPattern :=
	multiLineCommentPattern :=
	blankLinePattern :=
	stringLiteralPattern :=

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
	totalLines := 0
	sloc := 0
	inMultiLineComment := false
	inMultiLineString := false

	for scanner.Scan() {
		line := scanner.Text()
		totalLines++

		// Check for blank lines
		if blankLineRe.MatchString(line) {
			continue
		}

		// Check for single-line comments
		if singleLineCommentRe.MatchString(line) {
			continue
		}

		// Check for multi-line comments
		if inMultiLineComment {
			if strings.Contains(line, ...") {
				inMultiLineComment = false
				line = multiLineCommentRe.ReplaceAllString(line, "")
				if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
					continue
				}
			} else {
				continue
			}
		}

		if strings.Contains(line, "/*") {
			inMultiLineComment = true
			line = multiLineCommentRe.ReplaceAllString(line, "")
			if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
				continue
			}
		}

		// Check for multi-line strings
		if inMultiLineString {
			if strings.Count(line, "`")%2 != 0 || strings.Count(line, "\"")%2 != 0 {
inMultiLineString = false
line = stringLiteralRe.ReplaceAllString(line, "")
if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
continue
}
} else {
continue
}
}

if strings.Count(line, "`")%2 != 0 || strings.Count(line, "\"")%2 != 0 {
inMultiLineString = true
line = stringLiteralRe.ReplaceAllString(line, "")
if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
continue
}
}

sloc++
}

if err := scanner.Err(); err != nil {
// return 0, 0, err
}

return totalLines, sloc
}

*/
