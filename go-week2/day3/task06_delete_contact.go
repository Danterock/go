package main

import (
	"fmt"
	"sort"
)

func main() {
	book := map[string]string{
		"alex":  "111-11",
		"ivan":  "222-22",
		"maria": "333-33",
	}

	var name string
	fmt.Scan(&name)

	if !deleteContact(book, name) {
		fmt.Println("Not found")
		return
	}

	fmt.Println("Deleted:", name)

	keys := []string{}

	for key := range book {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	fmt.Println("Contacts:")

	for i := 0; i < len(keys); i++ {
		fmt.Printf("%s: %s\n", keys[i], book[keys[i]])
	}
}

func deleteContact(book map[string]string, name string) bool {
	_, ok := book[name]

	if !ok {
		return false
	}

	delete(book, name)
	return true
}
