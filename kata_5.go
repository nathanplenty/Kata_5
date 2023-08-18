package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// 1. Get String
	// 1.1. If String is empty -> return it (Guard)
	if len(s) < 1 {
		return s
	}
	// 2. Split after each spacer
	rev := strings.Split(s, " ")
	in := ""
	for i := len(rev); i > 0; i-- {
		if rev[i-1] == "" || rev[i-1] == " " {
			// 3. Remove split with only spacers
			continue
		}
		// 4. Append each split from the stack
		in = in + rev[i-1] + " "
	}
	in = strings.TrimRight(in, " ")
	// 5. Return reverse
	// fmt.Printf("%q\n", in)
	return in
}

func main() {
	fmt.Printf("Output: %q\n", reverseWords("hello world"))
	fmt.Printf("Output: %q\n", reverseWords("  I love programming "))
	fmt.Printf("Output: %q\n", reverseWords("  the sky is blue  "))
}
