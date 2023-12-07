package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Coordinates struct {
	x, y int
}

type NumberCell struct {
	char rune
	visited bool
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)

	symbolsMap := make(map[Coordinates]rune)
	numbersMap := make(map[Coordinates]NumberCell)

	y := 0
	
	for scanner.Scan() {
		for x, char := range scanner.Text() {
			if isValidSymbol(char) {
				symbolsMap[Coordinates{x: x, y: y}] = char
			}
			if unicode.IsDigit(char) {
				numbersMap[Coordinates{x: x, y: y}] = NumberCell{ char: char, visited: false }
			}
		}
		y++
	}

	total := 0
	
	for key := range symbolsMap {
		adjecentCoordinates := getAdjacentCoordsForSymbol(key)
		for _, coord := range adjecentCoordinates {
			if numbersMap[coord].char != 0 && !numbersMap[coord].visited {
				numbersMap[coord] = markNumberCellAsVisited(numbersMap[coord])

				fullNumber := string(numbersMap[coord].char)

				indexToCheckFrom := 1
				for numbersMap[Coordinates{x: coord.x + indexToCheckFrom, y: coord.y}].char != 0 && !numbersMap[Coordinates{x: coord.x + indexToCheckFrom, y: coord.y}].visited {
					numbersMap[Coordinates{x: coord.x + indexToCheckFrom, y: coord.y}] = markNumberCellAsVisited(numbersMap[Coordinates{x: coord.x + indexToCheckFrom, y: coord.y}])
					
					fullNumber = fullNumber + string(numbersMap[Coordinates{x: coord.x + indexToCheckFrom, y: coord.y}].char)
					indexToCheckFrom++
				}

				indexToCheckFrom = -1
				for numbersMap[Coordinates{x: coord.x + indexToCheckFrom, y: coord.y}].char != 0 && !numbersMap[Coordinates{x: coord.x + indexToCheckFrom, y: coord.y}].visited {
					numbersMap[Coordinates{x: coord.x + indexToCheckFrom, y: coord.y}] = markNumberCellAsVisited(numbersMap[Coordinates{x: coord.x + indexToCheckFrom, y: coord.y}])
					
					fullNumber = string(numbersMap[Coordinates{x: coord.x + indexToCheckFrom, y: coord.y}].char) + fullNumber
					indexToCheckFrom--
				}

				partNumber, error := strconv.Atoi(fullNumber)

				if error != nil {
					fmt.Printf("Could not parse %s to int. Original Error: %s", fullNumber, error)
					os.Exit(0)
				}

				total += int(partNumber)

			}
		}
	}

	println("Sum of part numbers:", total)
	
}

func getAdjacentCoordsForSymbol(coord Coordinates) []Coordinates {
	return []Coordinates{
		{coord.x - 1, coord.y}, {coord.x + 1, coord.y}, // Left and Right
		{coord.x, coord.y - 1}, {coord.x, coord.y + 1}, // Up and Down
		{coord.x - 1, coord.y - 1}, {coord.x + 1, coord.y - 1}, // Upper diagonals
		{coord.x - 1, coord.y + 1}, {coord.x + 1, coord.y + 1}, // Lower diagonals
	}
}

func markNumberCellAsVisited(numberCell NumberCell) NumberCell {
	copy := numberCell
	copy.visited = true
	return copy
}

func convertStringToRune(s string) rune {
	return []rune(s)[0]
}

func isValidSymbol(char rune) bool {
	return char != '.' && !unicode.IsDigit(char)
}