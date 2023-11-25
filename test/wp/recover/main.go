package main

import (
	"fmt"
	"time"
)

func worker() {
	defer fmt.Println("End function")

	defer func() {
		if r := recover().(error); r != nil {

			fmt.Println("Recovered. Error:\n", r)
			go worker()
		}

	}()

	fmt.Println("Worker start:")

	time.Sleep(2 * time.Second)

	var ptr *int
	fmt.Println(*ptr)

}
func main() {

	go worker()

	time.Sleep(50 * time.Second)

}
