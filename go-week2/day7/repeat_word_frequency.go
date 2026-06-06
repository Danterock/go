package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Print("input words: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	count := countWordFrequency(text)
	words := sortedWordKeys(count)
	for i := 0; len(words) > i; i++ {
		fmt.Printf("%s: %d\n", words[i], count[words[i]])
	}
}

func countWordFrequency(text string) map[string]int {
	text = strings.ToLower(text)
	textCopy := strings.Fields(text)
	count := make(map[string]int)
	for i := 0; i < len(textCopy); i++ {
		count[textCopy[i]]++
	}
	return count
}

func sortedWordKeys(freq map[string]int) []string {
	keys := make([]string, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
