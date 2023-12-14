package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ScratchCard struct {
	id int
	numberOfAquiredWinningNumbers int
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

	println("Total points:", totalPoints)

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

	winningNumbers := parseNumbers(splitNumbers[0])
	numbers := parseNumbers(splitNumbers[1])

	numberOfAquiredWinningNumbers := getNumberOfAquiredWinningNumbers(winningNumbers, numbers)

	scratchcard := ScratchCard{
		id: id,
		numberOfAquiredWinningNumbers: numberOfAquiredWinningNumbers,
	}

	return scratchcard
}

func getNumberOfAquiredWinningNumbers(winningNumbers []int, numbers []int) int {
	var numberOfAquiredWinningNumbers int

	for _, winningNumber := range winningNumbers {
		for _, number := range numbers {
			if winningNumber == number {
				numberOfAquiredWinningNumbers++
			}
		}
	}

	return numberOfAquiredWinningNumbers
}

func parseNumbers(numbers string) []int {
	numberStrings := strings.Split(numbers, space)

	var numberInts []int

	for _, numberString := range numberStrings {
		if numberString != "" {
			number, err := strconv.Atoi(numberString)

			if err != nil {
				log.Fatal(err)
			}

			numberInts = append(numberInts, number)
		}
	}

	return numberInts

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