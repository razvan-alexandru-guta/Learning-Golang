package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	sports := make(map[int]string)

	sports[1] = "Tennis"
	sports[2] = "Football"
	sports[3] = "Basketball"
	sports[4] = "Baseball"
	sports[5] = "Handball"
	sports[6] = "Golf"

	fmt.Println("MENU")
	fmt.Println("-----------")
	fmt.Println("1 - Tennis")
	fmt.Println("2 - Football")
	fmt.Println("3 - Basketball")
	fmt.Println("4 - Baseball")
	fmt.Println("5 - Handball")
	fmt.Println("6 - Golf")
	fmt.Println("Q - QUIT")

	for {
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You pressed %s", sports[i]))
	}
	fmt.Println("Program exit ...")
}
