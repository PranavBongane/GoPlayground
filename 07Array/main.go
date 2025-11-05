package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in Go")

	// Step 1: Declare an array of fixed size (4 elements)
	var fruitList [4]string

	// Step 2: Assign values by index
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	// fruitList[2] left empty (default zero value: "")
	fruitList[3] = "peach"

	fmt.Println("The fruit list is:", fruitList)
	fmt.Println("The fruit list length is:", len(fruitList))

	// Step 3: Accessing and updating elements
	fmt.Println("Second fruit is:", fruitList[1])
	fruitList[2] = "grapes" // updating value
	fmt.Println("Updated fruit list:", fruitList)

	// Step 4: Iterating through the array
	fmt.Println("\nIterating through fruitList:")
	for index, value := range fruitList {
		fmt.Printf("Index %d → %s\n", index, value)
	}

	// Step 5: Array literal initialization
	var vegList = [3]string{"potato", "beans", "brinjal"}
	fmt.Println("\nThe vegetable list is:", vegList)
	fmt.Println("The vegetable list length is:", len(vegList))

	// Step 6: Array type safety
	var a [3]int
	var b [4]int
	// a = b // ❌ Not allowed: arrays of different sizes are different types
	fmt.Printf("\nType of a: %T, Type of b: %T\n", a, b)

	// Step 7: Copy behavior (arrays are value types)
	copyArr := fruitList // Creates a full copy
	copyArr[0] = "mango" // Modify copied array
	fmt.Println("\nOriginal fruit list:", fruitList)
	fmt.Println("Copied fruit list  :", copyArr)

	// ----------------------------------------------------
	// BONUS SECTION: SLICES (real-world alternative to arrays)
	// ----------------------------------------------------
	fmt.Println("\n--- Slice Example ---")

	// Step 8: Slice declaration and initialization (no fixed size)
	sliceList := []string{"apple", "banana"}
	fmt.Println("Initial slice:", sliceList)
	fmt.Println("Length:", len(sliceList), "Capacity:", cap(sliceList))

	// Step 9: Append new elements (slices can grow dynamically)
	sliceList = append(sliceList, "peach", "grapes")
	fmt.Println("After appending:", sliceList)
	fmt.Println("Length:", len(sliceList), "Capacity:", cap(sliceList))

	// Step 10: Slice a portion of the array/slice
	subSlice := sliceList[1:3] // elements from index 1 to 2
	fmt.Println("Sub-slice (1:3):", subSlice)

	// Step 11: Demonstrate shared underlying array behavior
	subSlice[0] = "updated-banana"
	fmt.Println("Modified sub-slice:", subSlice)
	fmt.Println("Original slice after modification:", sliceList)
}
