package main

import (
	"fmt"
)

func test1() (result int) {

	defer func() {
		result++
	}()

	return 3
}

func main() {
	fmt.Println(test1())
}
