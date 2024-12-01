package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"unicode"
)

func main() {
	flag.Parse()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		file.Close()
	}

	scanner := bufio.NewScanner(file)

	var rightList []int
	var leftList []int

	for scanner.Scan() {
		line := scanner.Text()
		left, right, err := getLeftAndRightNumber(line)

		if err != nil {
			log.Println("Error converting to number: ", err)
			os.Exit(1)
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})

	totalDistance := 0
	similarityScore := 0

	similaritiesMap := make(map[int]int)

	for i := range rightList {
		similaritiesMap[leftList[i]] += 1
		totalDistance += diff(rightList[i], leftList[i])
	}

	for i, _ := range rightList {
		similarityScore += rightList[i] * similaritiesMap[rightList[i]]
	}

	fmt.Printf("Total distance: %d\n", totalDistance)
	fmt.Printf("Similarity score: %d\n", similarityScore)
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func getLeftAndRightNumber(line string) (int, int, error) {
	leftNumberAsString := ""
	rightNumberAsString := ""
	hasHitWhiteSpace := false
	for _, char := range line {
		if unicode.IsSpace(char) {
			hasHitWhiteSpace = true
		}

		if unicode.IsDigit(char) {
			if hasHitWhiteSpace == false {
				leftNumberAsString += string(char)
			} else {
				rightNumberAsString += string(char)
			}
		}
	}

	leftNumber, leftNumberErr := strconv.Atoi(leftNumberAsString)
	if leftNumberErr != nil {
		log.Fatal("Could not convert let number from string to int")
	}

	rightNumber, rightNumberErr := strconv.Atoi(rightNumberAsString)
	if rightNumberErr != nil {
		log.Fatal("Could not convert right number from string to int")
	}

	return leftNumber, rightNumber, nil
}
