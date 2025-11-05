package main

import (
	"fmt"
	"time" // The 'time' package provides functionality for measuring and displaying time.
)

func main() {
	fmt.Println("Welcome to time study of GoLang")

	// Step 1: Get the current local time
	presentTime := time.Now()

	// Prints the current date and time in default format
	fmt.Println(presentTime)

	// Step 2: Format the current time
	// In Go, formatting uses a reference date: "Mon Jan 2 15:04:05 MST 2006"
	// You must use this exact layout pattern to define your format.
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// Explanation of format parts:
	// 01 -> month (MM)
	// 02 -> day (DD)
	// 2006 -> year (YYYY)
	// 15 -> hour (24-hour format)
	// 04 -> minute
	// 05 -> second
	// Monday -> full day name

	// Step 3: Create a custom date/time value manually using time.Date()
	createdDate := time.Date(2020, time.May, 10, 23, 23, 0, 0, time.UTC)

	// Step 4: Format and print the custom date in a readable way
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
