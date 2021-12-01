package loopTest

import "fmt"

func TestFunc(chValue chan int) {
	for {
		auxValue := <-chValue
		fmt.Println("test", auxValue)
	}
}
