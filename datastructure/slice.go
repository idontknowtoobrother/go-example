package datastructure

import "fmt"

func GetSlice() []int {
	slice := []int{ // dynamic
		10,
		20,
		30,
		40,
	}

	fmt.Println("slice=", slice)
	fmt.Println("len(slice)=", len(slice))
	fmt.Println("cap(slice)=", cap(slice))

	subSlice := slice[0:2]

	fmt.Println("subSlice of slice=", subSlice)
	fmt.Println("len(subSlice)=", len(subSlice))
	fmt.Println("cap(subSlice)=", cap(subSlice))

	return slice
}
