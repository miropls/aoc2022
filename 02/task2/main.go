package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var pointsForChoice = map[string]int{
	"R": 1,
	"P": 2,
	"S": 3,
}

var RPS = map[string]string{
	"A": "R",
	"B": "P",
	"C": "S",
	"X": "R",
	"Y": "P",
	"Z": "S",
}

var winLoseCondition = map[string][2]string{
	"R": {"S", "P"},
	"P": {"R", "S"},
	"S": {"P", "R"},
}

func calculateRoundScore(o, p string) int {
	roundScore := 0
	opponentChoice := RPS[o]

	switch p {
	case "X":
		roundScore += 0 + pointsForChoice[winLoseCondition[opponentChoice][0]]
	case "Y":
		roundScore += 3 + pointsForChoice[opponentChoice]
	case "Z":
		roundScore += 6 + pointsForChoice[winLoseCondition[opponentChoice][1]]
	}

	return roundScore
}

func main() {
	playerTotalPoints := 0

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		players := strings.Split(scanner.Text(), " ")
		playerTotalPoints += calculateRoundScore(players[0], players[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("total: ", playerTotalPoints)
}
