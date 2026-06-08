package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLine(message string) string {
	fmt.Println(message)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func readInt(message string) int {
	text := readLine(message)

	number, err := strconv.Atoi(text)
	if err != nil {
		return -1
	}

	return number
}

func normalizePriority(priority string) (string, bool) {
	priority = strings.ToLower(priority)
	priority = strings.TrimSpace(priority)
	if priority == "low" || priority == "high" || priority == "medium" {
		return priority, true
	}
	return "Invalid priority", false
}
