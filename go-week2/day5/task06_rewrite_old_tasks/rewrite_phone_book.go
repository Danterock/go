package main

import (
	"fmt"
	"strings"
)

var name string
var book = map[string]string{
	"alex":  "111-11",
	"ivan":  "222-22",
	"maria": "333-33",
}

func getPhone(book map[string]string, name string) (string, bool) {
	phone, ok := book[name]
	return phone, ok
}

func printPhone(name string) string {
	name = strings.ToLower(name)
	phone, ok := getPhone(book, name)
	if ok {
		return "Phone: " + phone
	}
	return "Not Found"
}

func main() {
	fmt.Print("input your name: ")
	fmt.Scanln(&name)
	fmt.Println(printPhone(name))
}
