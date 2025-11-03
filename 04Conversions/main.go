package main

import (
	"bufio"   // For buffered input from the console
	"fmt"     // For printing to console
	"os"      // To access system input/output (like os.Stdin)
	"strconv" // For converting string to numeric types
	"strings" // For trimming and processing string input
)

func main() {
	// Introductory messages
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	// Create a reader to take input from user (reads from standard input)
	reader := bufio.NewReader(os.Stdin)

	// Read the user input until newline ('\n')
	input, _ := reader.ReadString('\n')

	// Display what the user entered
	fmt.Println("Thanks for rating,", input)

	// Clean up the input using strings.TrimSpace to remove newline and extra spaces
	// Then convert it from string to float64 using strconv.ParseFloat
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// Error handling
	if err != nil {
		// If the conversion fails (e.g., user enters "abc"), print the error
		fmt.Println("Error:", err)
	} else {
		// If conversion succeeds, perform a simple arithmetic operation
		fmt.Println("Added 1 to your rating:", numRating+1)
	}
}
