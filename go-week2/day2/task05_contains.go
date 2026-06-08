package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var sentens, word string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your sentence: ")
	sentens, _ = reader.ReadString('\n')
	sentens = strings.TrimSpace(sentens)
	fmt.Scan(&word)
	word = strings.TrimSpace(word)
	if containsIgnoreCase(sentens, word) == true {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
func containsIgnoreCase(text string, query string) bool {
	text = strings.ToLower(text)
	query = strings.ToLower(query)
	text0 := strings.Fields(text)
	for i := 0; i < len(text0); i++ {
		if text0[i] == query {
			return true
		}
	}
	return false
}
