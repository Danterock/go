package main

import (
	"fmt"
	"strings"
)

func main() {
	book := map[string]string{
		"alex":  "111-11",
		"ivan":  "222-22",
		"maria": "333-33",
	}
	var name string
	fmt.Print("input your name: ")
	fmt.Scanln(&name)
	name = strings.ToLower(name)

	phone, ok := getPhone(book, name)
	if ok {
		fmt.Println("Phone:", phone)
	} else {
		fmt.Println("Not Found")
	}
}

func getPhone(book map[string]string, name string) (string, bool) {
	phone, ok := book[name]
	return phone, ok
}
