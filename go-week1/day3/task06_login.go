package main

import "fmt"

func main() {
	var login string
	var password string
	fmt.Print("Enter your login and password: ")
	fmt.Scan(&login, &password)
	if login == "student" && password == "go12345" {
		fmt.Println("Access granted")
	} else if login != "student" {
		fmt.Println("Wrong login")
	} else if password != "go12345" {
		fmt.Println("Wrong password")
	}
}
