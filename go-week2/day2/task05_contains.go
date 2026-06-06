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
	textO := strings.Fields(text)
	query = strings.ToLower(query)
	queryO := strings.Fields(query)
	for i := 0; i < len(text); i++ {
		if queryO[0] == textO[i] {
			return true
		}
		return false
	}
	return false
}
