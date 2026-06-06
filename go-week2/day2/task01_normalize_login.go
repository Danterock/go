package main

import (
	"fmt"
	"strings"
)

func main() {
	var login string
	fmt.Print("Enter your login: ")
	fmt.Scan(&login)
	fmt.Print("Login: ", normalizeLogin(login))
}

func normalizeLogin(login string) string {
	login = strings.TrimSpace(login)
	if login == "" {
		return "Invalid login"
	}
	login = strings.ToLower(login)
	return login
}
