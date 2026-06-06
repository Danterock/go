package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var words string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your full name: ")
	words, _ = reader.ReadString('\n')
	fmt.Print("Words:", countWords(words))
}

func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}
