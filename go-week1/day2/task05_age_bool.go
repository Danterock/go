package main

import "fmt"

func main() {
	var name string
	var age int
	var isAdult bool
	fmt.Print("Enter name and age: \n")
	fmt.Scan(&name, &age)
	isAdult = age >= 18
	fmt.Printf("Name: %s\nAge: %d\nAdult: %t", name, age, isAdult)
}
