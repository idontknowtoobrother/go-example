package funcs

import "fmt"

// @ Syntax
//
// func functionName(parameter1 type1, parameter2 type2) returnType {
//     // Function body
//     // Return statement (if the function has a return type)
// }
//

func PlusTwoNumber(left int, right int) int { // return int
	return left + right
}

func Hello(name string) { // void
	fmt.Println("Hello, ", name)
}
