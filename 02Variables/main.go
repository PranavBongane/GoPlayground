package main

import "fmt"

// Exported constant — name starts with uppercase so it's exported (public).
// If this constant were in a package other than `main`, other packages could access it as <package>.Globle.
// Note: in `package main` export doesn't matter much, but the naming rule still applies.
const Global = "test globle var...."

func main() {
	// Explicit variable declaration with type
	var userName string = "pranav"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n", userName)

	// Boolean type variable
	var isAlive bool = true
	fmt.Println(isAlive)
	fmt.Printf("Variable is of type: %T \n", isAlive)

	// Unsigned integer (uint8) — stores only positive numbers (0–255)
	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	// Integer variable (can store both positive and negative numbers)
	var iValue int = 256
	fmt.Println(iValue)
	fmt.Printf("Variable is of type: %T \n", iValue)

	// Float32 type — stores floating-point numbers (limited precision)
	var smallFloat float32 = 255.4152632563255412525
	fmt.Println(smallFloat) // value is rounded due to float32 precision
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// Float64 type — higher precision floating-point number
	var Float float64 = 255.4152632563255412525
	fmt.Println(Float)
	fmt.Printf("Variable is of type: %T \n", Float)

	// Default (zero) values in Go
	var newVar int // defaults to 0
	fmt.Println(newVar)
	fmt.Printf("Variable is of type: %T \n", newVar)

	var strVar string // defaults to empty string ""
	fmt.Println(strVar)
	fmt.Printf("Variable is of type: %T \n", strVar)

	// Type inference — Go automatically detects the variable type
	var implicitVar = "this is a string...."
	fmt.Println(implicitVar)
	fmt.Printf("Variable is of type: %T \n", implicitVar)

	// Walrus operator (:=) — short variable declaration (only allowed inside functions)
	testnum := 300000
	fmt.Println(testnum)
	fmt.Printf("Variable is of type: %T \n", testnum)

	// Reassigning new value (same type — valid)
	testnum = 366666
	fmt.Println(testnum)
	fmt.Printf("Variable is of type: %T \n", testnum)

	// Another walrus example with string
	testString := "test walrus"
	fmt.Println(testString)
	fmt.Printf("Variable is of type: %T \n", testString)

	// Printing exported global constant
	fmt.Println(Global)
	fmt.Printf("Variable is of type: %T \n", Global)
}

// Suggestion: if you prefer correct spelling and clearer intent, rename to:
// const Global = "test global var...."
