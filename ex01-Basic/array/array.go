package main

import (
	"fmt"
)

// var array_name = [length]datatype{values}
// array_name := [length]datatype{values}
func main() {
	var array1 = [4]string{"1", "2", "3", "4"}
	var array2 = []int{1, 2, 3, 4}
	for i := 0; i < len(array1); i++ {
		fmt.Printf("Type: %T, Value: %v\n", array1[i], array1[i])
	}
	for i := 0; i < len(array2); i++ {
		fmt.Printf("Type: %T, Value: %v\n", array2[i], array2[i])
	}
}
