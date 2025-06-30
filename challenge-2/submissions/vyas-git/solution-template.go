package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input from standard input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// Call the ReverseString function
		output := ReverseString(input)

		// Print the result
		fmt.Println(output)
	}
}

// ReverseString returns the reversed string of s.
func ReverseString(s string) string {
        i,j:= 0, len(s)-1
        r := []rune(s)
        for i<j{
            r[i],r[j] = r[j],r[i]
            i++
            j--
        }
	return string(r)
}
