package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const maxRed = 12
const maxGreen = 13
const maxBlue = 14

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	gameIdTotal := 0
	powerTotal := 0

	for scanner.Scan() {
		gameId, power := parseGame(scanner.Text())
		gameIdTotal += gameId
		powerTotal += power
	}

	println("Possible Game IDs summed up:", gameIdTotal)
	println("Sum of power:", powerTotal)
}

func cubeOverLimit(color string, numberOfCubes int) bool {
	switch color {
		case "red":
			return numberOfCubes > maxRed
		case "green":
			return numberOfCubes > maxGreen
		case "blue":
			return numberOfCubes > maxBlue
	default:
		return false
	}
}

func parseGame(line string) (int, int) {
	regex := regexp.MustCompile(`\d+`)
	gameId, err := strconv.Atoi(regex.FindString(line))

	if err != nil {
		log.Fatal("could not extract value for gameId ", err)
	}

	game := strings.Split(line, ": ")[1]
	sets := strings.Split(game, "; ")

	leastAmountOfCubes := map[string]int {
		"green" : 0,
		"red" : 0,
		"blue" : 0,
	}

	for _, cubesLine := range sets {
		cubes := strings.Split(cubesLine, ", ")

		for i := range cubes {
			cubeUsage := strings.Split(cubes[i], " ")
			numberOfCubes, err := strconv.Atoi(cubeUsage[0])
			color := cubeUsage[1]

			if err != nil {
				log.Fatal(err, " failed to convert number of cubes")
			}

			if leastAmountOfCubes[color] < numberOfCubes {
				leastAmountOfCubes[color] = numberOfCubes
			}

			if cubeOverLimit(color, numberOfCubes) {
				gameId = 0
			}
		}
	}

	power := leastAmountOfCubes["green"] * leastAmountOfCubes["red"] * leastAmountOfCubes["blue"]

	return gameId, power
}