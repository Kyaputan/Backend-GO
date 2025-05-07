package main

import "fmt"

func main() {
	a := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
	}
	var b map[string]string

	fmt.Println(a)
	fmt.Println(b == nil)

	for index, val := range a {
		fmt.Println(index, val)
	}

}
