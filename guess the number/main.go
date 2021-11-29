package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const PromtWhenReady = "don't type the number, just press ENTER when you are ready."

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())
	//make random numbers
	var firstNumber int
	firstNumber = rand.Intn(8) + 2
	secondNumber := rand.Intn(8) + 2
	var thirdNumber = rand.Intn(8) + 2

	var answer int
	gameLogic(firstNumber, secondNumber, thirdNumber, answer)

}

func gameLogic(firstNumber int, secondNumber int, thirdNumber int, answer int) {
	reader := bufio.NewReader(os.Stdin)
	// display the instructions

	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")
	// take through the games
	fmt.Println("Think of a number between 1 and 10", PromtWhenReady)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, PromtWhenReady)
	answer = firstNumber
	reader.ReadString('\n')

	fmt.Println("Multiply the result by", secondNumber, PromtWhenReady)
	answer = answer * secondNumber
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally tought of", PromtWhenReady)
	answer = answer / firstNumber
	reader.ReadString('\n')

	fmt.Println("Now substract", thirdNumber, PromtWhenReady)
	answer = answer - thirdNumber
	reader.ReadString('\n')

	// give the answer

	fmt.Println("The answer is", answer)
}
