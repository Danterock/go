package main

import "fmt"

func main() {
	var Min int
	fmt.Print("Input minutes:")
	fmt.Scan(&Min)
	fmt.Printf("Hours: %d\nMinutes: %d", Min/60, Min%60)
}
