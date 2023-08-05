package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	listOfElves := []int{}
	totalCalories := 0
	top3CaloriesSummed := 0

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			listOfElves = append(listOfElves, totalCalories)
			totalCalories = 0
			scanner.Scan()
		}

		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			break
		}

		totalCalories += num
	}

	sort.SliceStable(listOfElves, func(i, j int) bool {
		return listOfElves[i] > listOfElves[j]
	})

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for v := range listOfElves[:3] {
		top3CaloriesSummed += listOfElves[v]
	}

	fmt.Println(top3CaloriesSummed)
}
