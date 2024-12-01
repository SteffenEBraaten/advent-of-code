package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ScratchCard struct {
	id int
	numberOfAquiredWinningNumbers int
	copies int
}

const numberSeperator = "|"
const numberStart = ": "
const space = " "

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scratchCards := parseScratchcardFile(file)

	totalPoints := 0

	for _, scratchCard := range scratchCards {
		totalPoints += calculatePoints(scratchCard.numberOfAquiredWinningNumbers)
	}

	for _, scratchCard := range scratchCards {
		if scratchCard.copies == 0 && scratchCard.numberOfAquiredWinningNumbers > 0 {
			for i := 1; scratchCard.numberOfAquiredWinningNumbers > i; i++ {
				scratchCards[scratchCard.id - 1 + i].copies += 1
			}
		}
		if scratchCard.copies > 0 && scratchCard.numberOfAquiredWinningNumbers > 0 {
			for i := 1; scratchCard.numberOfAquiredWinningNumbers > i; i++ {
				scratchCards[scratchCard.id - 1 + i].copies += 1 * scratchCard.copies
			}
		}
		println(scratchCard.copies)
	}

	numberOfCopyCards := 0
	for _, scratchCard := range scratchCards {
		numberOfCopyCards += scratchCard.copies
	}

	println("Total points:", totalPoints)
	println("Total number of cards:", len(scratchCards) + numberOfCopyCards)
}

func parseScratchcardFile(file *os.File) []ScratchCard {
	scanner := bufio.NewScanner(file)

	var scratchcards []ScratchCard

	for scanner.Scan() {
		line := scanner.Text()
		scratchcards = append(scratchcards, parseScratchCardLine(line))
	}

	return scratchcards
}

func parseScratchCardLine(line string) ScratchCard {
	id := getId(line)

	noPrefix := removePrefix(line)

	splitNumbers := strings.SplitN(noPrefix, numberSeperator, 2)

	winningNumbers, err := parseNumbers(splitNumbers[0])
	if err != nil {
		log.Fatal(err, "Could not parse winning numbers from line: ", line)
	}

	numbers, err := parseNumbers(splitNumbers[1])
	if err != nil {
		log.Fatal(err, "Could not parse numbers from line: ", line)
	}

	numberOfAquiredWinningNumbers := getNumberOfAquiredWinningNumbers(winningNumbers, numbers)

	scratchcard := ScratchCard{
		id: id,
		numberOfAquiredWinningNumbers: numberOfAquiredWinningNumbers,
	}

	return scratchcard
}

func getNumberOfAquiredWinningNumbers(winningNumbers, numbers []int) int {
	winNumsMap := make(map[int]bool)
	for _, num := range winningNumbers {
		winNumsMap[num] = true
	}

	count := 0
	for _, num := range numbers {
		if winNumsMap[num] {
			count++
		}
	}
	return count
}

func parseNumbers(numbers string) ([]int, error) {
	numberStrings := strings.Split(numbers, space)
	var numberInts []int

	for _, numberString := range numberStrings {
		if numberString != "" {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				return nil, fmt.Errorf("failed to parse number: %v", err)
			}
			numberInts = append(numberInts, number)
		}
	}
	return numberInts, nil
}

func removePrefix(line string) string {
	pos := strings.Index(line, numberStart)
	if pos != -1 {
		return line[pos+2:]
	}
	return ""
}

func getId(line string) int {
	regex := regexp.MustCompile(`\d+`)

	id, err := strconv.Atoi(regex.FindString(line))

	if err != nil {
		log.Fatal(err, "Could not parse id from line: ", line)
	}

	return id
}

func calculatePoints(numberOfAquiredWinningNums int) int {
	if numberOfAquiredWinningNums == 0 {
		return 0
	}

	if numberOfAquiredWinningNums == 1 {
		return 1
	}

	points := 1
	for i := 1; i < numberOfAquiredWinningNums; i++ {
		points = points * 2
	}

	return points
}