package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
        for s := range scanner.Text()
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
