package main

import (
	"fmt"
)

func main() {
	var phoneNumber string
	fmt.Print("Enter your phone number: ")
	fmt.Scan(&phoneNumber)
	fmt.Print("Masked: ", maskPhone(phoneNumber))
}

func maskPhone(phone string) string {
	masked := ""
	runes := []rune(phone)
	if len(phone) <= 4 {
		return "Invalid phone"
	}
	for i := 0; i < len(runes)-4; i++ {
		masked += "*"
	}
	for i := len(runes) - 4; i < len(runes); i++ {
		masked += string(runes[i])
	}
	return masked
}
