package pointer

import "fmt"

func Explanation() {
	// Declare an integer variable
	x := 10

	// Declare a pointer to an integer and assign it the address of x
	var p *int = &x

	// Print the value of x and the value at the pointer p
	fmt.Println("Value of x:", x)             // Output: Value of x: 10
	fmt.Println("Value at p:", *p, " -> ", p) // Output: Value at p: 10

	// Modify the value at the pointer p
	*p = 20

	// x is modified since p points to x
	fmt.Println("New value of x:", x) // Output: New value of x: 20
}
