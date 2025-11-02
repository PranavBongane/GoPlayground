package main

import (
	"bufio" // Provides buffered input/output utilities
	"fmt"   // For formatted I/O (Println, Printf, etc.)
	"os"    // Gives access to OS-level functions (e.g., Stdin)
)

func main() {
	// Simple welcome message
	welcome := "Hi User"
	fmt.Println(welcome)

	// Create a new reader to read input from the console (standard input)
	reader := bufio.NewReader(os.Stdin)

	// Prompt user for input
	fmt.Println("Enter the ratings for the match: ")

	// Read user input until newline ('\n')
	// The second return value is an error (ignored here using _)
	rating, _ := reader.ReadString('\n')

	// Print back the user input
	fmt.Println("Thanks for rating,", rating)

	// Print the data type of the input
	fmt.Printf("Type of rating is %T ", rating)
}
