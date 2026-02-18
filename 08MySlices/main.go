package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice study in Go")

	// 1️⃣ Creating a basic slice
	var fruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// 2️⃣ Appending new items to the slice
	fruitList = append(fruitList, "banana", "mango")
	fmt.Println("After appending fruits:", fruitList)

	// 3️⃣ Slicing operation (from index 1 up to but not including 3)
	fruitList = fruitList[1:3]
	fmt.Println("After slicing [1:3]:", fruitList)

	// 4️⃣ Creating an int slice with make()
	highScores := make([]int, 4)
	highScores[0] = 324
	highScores[1] = 249
	highScores[2] = 478
	highScores[3] = 546
	// highScores[4] = 777 // ❌ would throw error — index out of range

	fmt.Println("Initial high scores:", highScores)

	// 5️⃣ Appending values dynamically (Go auto-expands capacity)
	highScores = append(highScores, 456, 676, 878)
	fmt.Println("After appending scores:", highScores)

	// 6️⃣ Sorting the slice
	sort.Ints(highScores)
	fmt.Println("Sorted high scores:", highScores)
	fmt.Println("Is slice sorted?", sort.IntsAreSorted(highScores))

	// 7️⃣ Removing an element by index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("\nCourses before removal:", courses)

	var index int = 2                                       // remove "swift"
	courses = append(courses[:index], courses[index+1:]...) // combine before and after parts
	fmt.Println("Courses after removal:", courses)
}
