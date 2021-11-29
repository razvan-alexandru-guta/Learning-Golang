package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	talk := doctor.Intro()

	fmt.Println(talk)

	for {
		fmt.Print("User input -> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "quit" {
			break
		} else {
			fmt.Println(userInput)
		}

	}

}
