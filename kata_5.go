package main

import (
	"fmt"
	"strings"
)

func reverseWords(stringInput string) string {
	// Get each Input word without spacers as array for strings
	stringArray := strings.Fields(stringInput)
	// Clear Input 'stringInput'
	stringInput = ""
	// Append each word from 'stringArray' backwards on 'stringInput'
	for i := len(stringArray); i > 0; i-- {
		stringInput += stringArray[i-1]
		if i <= len(stringArray) {
			stringInput += " "
		}
	}
	// Remove last spacer from 'stringInput'
	stringInput = strings.TrimRight(stringInput, " ")
	return stringInput
}

func main() {
	fmt.Printf("Output: %q\n", reverseWords("hello world"))
	fmt.Printf("Output: %q\n", reverseWords("  I love programming "))
	fmt.Printf("Output: %q\n", reverseWords("  the sky is blue  "))
}
