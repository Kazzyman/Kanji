package main

import (
	"strings"
	"unicode"
)

func check_for_match_in_other_fields(users_guess string) bool {
	// Look everywhere
	fields_from_aCard := []string{aCard.Meaning, aCard.Second_Meaning, aCard.Kunyomi, aCard.Onyomi, aCard.Vocab, aCard.Vocab2}

	for n, one_field_from_aCard := range fields_from_aCard { // Read through the slice of fields from aCard (see above)
		if customMatch(users_guess, one_field_from_aCard) {
			// found_one = true // Let the caller know that we found at least one match via our custom parsing algorithm
			// fmt.Printf("Tossed value is position number:%v\n", n) // If we want to know what the position was, print it
			if n != 99999 {
				n = 0
			} // Must use n for something, as we chose not to toss the returned parameter via a _,
			// break
			return true
		}
	}
	// else
	return false
}

// Older versions and segments of code:
/*
pattern := `(?i).*` + strings.Join(strings.Split(our_guess_in_lower_case, ""), `.*`) + `.*`
    reg := regexp.MustCompile(pattern)

    for _, str := range strings_from_card {
        if reg.MatchString(str) {
            found_one = true
            break
        }
    }
*/

/*
func customMatch(our_guess, str string) bool {
	ourGuessLower := strings.ToLower(our_guess)
	strLower := strings.ToLower(str)

	// Debugging info:
	fmt.Printf("strLower:%s, ourGuessLower:%s, %v\n", strLower, ourGuessLower, strings.Contains(strLower, ourGuessLower))
	fmt.Printf("ourGuessLower:%s, strLower:%s, %v\n\n", ourGuessLower, strLower, strings.Contains(ourGuessLower, strLower))

	return strings.Contains(strLower, ourGuessLower) ||
		strings.Contains(ourGuessLower, strLower)
	// Add more code here to parse strLower for any (non-alpha delimited) substrings in it that match any substring in ourGuessLower ...
	// ... such that if strLower is: "one, two, three, sam, sick, water", it will match if ourGuessLower is: "xxxxwaterxxxx" where x is any char
}

func check_for_match_in_other_fields(in string) (found_one bool) {
	our_guess_in_lower_case := strings.ToLower(in)
	strings_from_card := []string{aCard.Meaning, aCard.Second_Meaning, aCard.Kunyomi, aCard.Onyomi, aCard.Vocab, aCard.Vocab2}

	for _, str := range strings_from_card {
		if customMatch(our_guess_in_lower_case, str) {
			found_one = true
			break
		}
	}

		// Debugging: Print whether a match was found
		fmt.Printf("Search String: %s\n", our_guess_in_lower_case)
		fmt.Printf("Fields: %v\n", strings_from_card)
		fmt.Printf("Matches: %v\n", found_one)

		return found_one
}

*/

func bool_from_rune(r rune) bool {
	return !unicode.IsLetter(r) // r will be supplied by the caller, and is any
}

func customMatch(our_guess, str string) bool {
	// Let's work only with lower-case versions of our data
	ourGuessLower := strings.ToLower(our_guess)
	strLower := strings.ToLower(str)

	// Check for a simple match of one within the other
	if strings.Contains(strLower, ourGuessLower) ||
		strings.Contains(ourGuessLower, strLower) { // i.e. :
		// if strLower contains ourGuessLower, OR ... // an exact match of our Guess anywhere within the card field
		// if ourGuessLower contains strLower        // an exact match of the entire card field within our Guess
		return true
	}

	/*
			// // Alternate, compact form, of the following single line of code:
				//  strParts := strings.FieldsFunc(strLower, bool_from_rune)
			// // i.e., sans the stand-alone named func bool_from_rune
		//
			// Split strLower into non-alpha delimited substrings and check if any match ourGuessLower
			strParts := strings.FieldsFunc(strLower, func(r rune) bool {
				return !unicode.IsLetter(r)
			})

	*/

	// // Version which uses a separate named function:
	// Split strLower into non-alpha delimited substrings; then check if any match ourGuessLower
	strParts := strings.FieldsFunc(strLower, bool_from_rune)
	/*
			strings.FieldsFunc expects a function as its second argument, one which takes a rune as its parameter,
			e.g., the bool_from_rune function defined immediately prior to this function.

			When we call strings.FieldsFunc(strLower, bool_from_rune), strings.FieldsFunc internally iterates through the
			characters in the strLower string, and for each character, it converts it to a rune before passing it as an argument
			to the bool_from_rune function.

			So, we don't need to explicitly supply the rune parameter when calling bool_from_rune within strings.FieldsFunc.
			The strings.FieldsFunc function takes care of this conversion for us.

			In summary, strings.FieldsFunc handles the conversion from characters in the input string to rune values and passes
			those rune values to the provided function (bool_from_rune) for processing. This is why we can call bool_from_rune
			without explicitly supplying a rune parameter to it.

		-- For additional foundational explanations of what is going on, please refer to the text which follows this func --
	*/

	for _, part := range strParts { // Read through the slice: strParts to obtain each part
		if strings.Contains(ourGuessLower, part) { // If ourGuessLower contains part, exit returning true
			return true
		}
	}
	// else:
	return false
}

// -- ADDITIONAL FOUNDATIONAL EXPLANATIONS OF WHAT IS GOING ON (in an earlier section): --
/*
The rune type in Go represents a Unicode code point, which is a numeric value that corresponds to a specific character
in the Unicode standard. In the code you provided, the function bool_from_rune is a custom function that takes a rune as
its argument and returns a boolean value based on a condition.

Here's how it works in the above code:

bool_from_rune is a function that accepts a rune (a Unicode code point) as its parameter.

Inside bool_from_rune, it checks if the given rune (Unicode code point) is not a letter using the unicode.IsLetter function.
The ! operator negates the result, so it returns true if the rune is not a letter and false if it is a letter.

The bool_from_rune function is used as an argument in strings.FieldsFunc

strings.FieldsFunc splits a string into fields based on a provided function, and in this case, it uses the bool_from_rune
function to determine where to split the string.

bool_from_rune checks if a rune is not a letter and is used in the splitting logic to split a string into non-alpha
delimited substrings. This allows us to break the input string into words or substrings based on non-letter characters
(e.g., spaces, punctuation, etc.).

It's a way to customize how strings are split when using strings.FieldsFunc

In the above code, it helps split strLower into non-alpha delimited substrings for further matching.

The r is simply a parameter name, and it's used to represent the rune value (Unicode code point) that you pass to the
function. It's a common convention in Go (and in many programming languages) to use descriptive parameter names that
make the code more readable and self-explanatory. In this case, r is used to indicate that the function operates on a
rune value.

The rune parameter represents a single Unicode code point, which can be any character from the Unicode standard.
The name r is just a choice made by the developer, yours truly, who wrote this code; it could have been named something
else, but r is a concise and commonly used name for such parameters, especially when dealing with Unicode code points.
It doesn't have any specific source or significance other than being a descriptive variable name.

When you call the strings.FieldsFunc function with bool_from_rune as its argument, you're not explicitly supplying a
parameter of type rune to it. Instead, the strings.FieldsFunc function itself will handle this.

Here's how it works:

strings.FieldsFunc expects a function that takes a rune as its parameter, e.g., bool_from_rune

When we call strings.FieldsFunc(strLower, bool_from_rune), strings.FieldsFunc internally iterates through the characters
in the strLower string, and for each character, it converts it to a rune before passing it as an argument to the
bool_from_rune function.

So, we don't need to explicitly supply the rune parameter when calling bool_from_rune within strings.FieldsFunc
The strings.FieldsFunc function takes care of this conversion for us.

In summary, strings.FieldsFunc handles the conversion from characters in the input string to rune values and passes
those rune values to the provided function (bool_from_rune) for processing. This is why we can call bool_from_rune
without explicitly supplying a rune parameter.

*/
