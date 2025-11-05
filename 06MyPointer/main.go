package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointer class")

	// Step 1: Declare a pointer variable without assigning anything
	// By default, it holds the zero value for pointers — which is <nil>
	var ptr *int
	fmt.Println(ptr) // Output: <nil>

	// Step 2: Declare a normal integer variable
	myVar := 90

	// Step 3: Create a pointer that stores the memory address of 'myVar'
	var ptr2 = &myVar

	// Step 4: Print the memory address stored in ptr2
	fmt.Println("The memory address:", ptr2)

	// Step 5: Dereference the pointer to access the actual value stored at that memory address
	fmt.Println("Actual value at the memory address:", *ptr2)

	// Step 6: Modify the value at that address using dereferencing
	// This change will reflect in 'myVar' too, since both refer to the same memory
	*ptr2 = *ptr2 * 2

	// Step 7: Print the updated value of 'myVar'
	fmt.Println("The changed value is:", myVar)
}
