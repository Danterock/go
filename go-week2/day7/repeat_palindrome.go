package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var sentens string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your sentence: ")
	sentens, _ = reader.ReadString('\n')
	fmt.Print("Palindrome: ", isPalindrome(sentens))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.TrimSpace(s)
	runes := []rune(s)
	left := 0
	right := len(runes) - 1
	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true

}
