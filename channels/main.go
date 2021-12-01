package main

import (
	"fmt"
	"myapp/loopTest"
)

func mainFunc(ch chan int) {
	for {
		fmt.Println("main", <-ch)
	}
}

func main() {
	chKey := make(chan int)
	chValue := make(chan int)
	myMap := make(map[int]int)

	myMap[1] = 1
	myMap[2] = 2
	myMap[3] = 3

	go loopTest.TestFunc(chValue)

	go mainFunc(chKey)

	for key, value := range myMap {
		fmt.Println(key)
		select {
		case chKey <- key:
			fmt.Print("4")
		case chValue <- value:
			fmt.Print("5")
		}
	}
	fmt.Println("done")
}
