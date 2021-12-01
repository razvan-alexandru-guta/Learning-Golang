package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Winner         string `json:"winner"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	message := ""
	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}
	if playerValue == computerValue {
		drawMap := make(map[int]string)
		drawMap[1] = "Tie!"
		drawMap[2] = "Great minds!"
		drawMap[3] = "You are both so good"
		rand.Seed(time.Now().UnixNano())
		roundResult = "It's a draw"
		message = drawMap[rand.Intn(3)+1]
	} else if playerValue == (computerValue+1)%3 {
		winnerMap := make(map[int]string)
		winnerMap[1] = "Yuppy!"
		winnerMap[2] = "Nicee!!"
		winnerMap[3] = "Congrats!"
		rand.Seed(time.Now().UnixNano())
		roundResult = "Player wins"
		message = winnerMap[rand.Intn(3)+1]
	} else {
		loserMap := make(map[int]string)
		loserMap[1] = "Try again!"
		loserMap[2] = "No problem!!"
		loserMap[3] = "Off!"
		rand.Seed(time.Now().UnixNano())
		roundResult = "Computer wins"
		message = loserMap[rand.Intn(3)+1]
	}
	var result Round
	result.Winner = message
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result
}
