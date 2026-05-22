package main

import "fmt"

func main() {
	totalSeconds := 7384
	fmt.Printf("Hours: %d\nMinutes: %d\nSeconds: %d", totalSeconds/60/60, totalSeconds/60%60, totalSeconds%60)
}
