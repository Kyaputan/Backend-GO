package main

import "fmt"

func main() {
	for i := 0; i < 1000000000; i++ {
		if i%100000 == 0 {
			fmt.Println(i)}

}
fmt.Println(`Done`)
}
