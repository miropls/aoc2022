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

func convertToRPS(s string) string {

	if s == "A" || s == "X" {
		s = "R"
	} else if s == "B" || s == "Y" {
		s = "P"
	} else if s == "C" || s == "Z" {
		s = "S"
	}
	return s
}

func calculateRoundScore(o, p string) int {
	var roundScore int

	if o == p {
		roundScore = pointsForChoice[p] + 3
	} else if o == "R" && p == "S" || o == "S" && p == "P" || o == "P" && p == "R" {
		roundScore = pointsForChoice[p]
	} else {
		roundScore = pointsForChoice[p] + 6
	}

	return roundScore
}

func main() {
	playerTotalPoints := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		playerTotalPoints += calculateRoundScore(convertToRPS(s[0]), convertToRPS(s[1]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(playerTotalPoints)
}
