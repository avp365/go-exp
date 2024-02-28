package main

import "fmt"

func test1() (result int) {

	defer func() {
		result++
	}()

	return 3
}

func main() {
	p := make([]int, 3)
	s := []int{1}
	fmt.Println(p)
	fmt.Println(s)
	p = s
	fmt.Println(p)
}
