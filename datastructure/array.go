package datastructure

import "fmt"

func GetArray() [3]int {
	var array [3]int // static 3 index only
	array[0] = 10
	array[1] = 20
	array[2] = 30
	return array
}

func ConvertArrayToSlice() {
	var array [3]int
	array[0] = 10
	array[1] = 20
	array[2] = 30

	slice := array[:]
	slice = append(slice, 40, 50, 60)

	fmt.Println("ConvertArrayToSlice.array=", array)
	fmt.Println("ConvertArrayToSlice.slice=", slice)
}
