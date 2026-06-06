package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your full name: ")
	name, _ = reader.ReadString('\n')
	fmt.Print("Initials: ", getInitials(name))
}

func getInitials(fullName string) string {
	fname := strings.Fields(fullName)
	initials := ""
	for i := 0; i < len(fname); i++ {
		firstl := []rune(fname[i])[0]
		initials += string(firstl) + "."
	}
	return initials
}
