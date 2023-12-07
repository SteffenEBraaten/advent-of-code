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
	char    rune
	visited bool
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	symbolsMap, numbersMap, err := processFile(file)
	if err != nil {
		log.Fatal(err)
	}

	total, err := calculatePartNumbersSum(symbolsMap, numbersMap)
	if err != nil {
		log.Fatal(err)
	}

	println("Sum of part numbers:", total)
}

func processFile(file *os.File) (map[Coordinates]rune, map[Coordinates]NumberCell, error) {
	scanner := bufio.NewScanner(file)
	symbolsMap := make(map[Coordinates]rune)
	numbersMap := make(map[Coordinates]NumberCell)
	y := 0

	for scanner.Scan() {
		processLine(scanner.Text(), y, symbolsMap, numbersMap)
		y++
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return symbolsMap, numbersMap, nil
}

func processLine(line string, y int, symbolsMap map[Coordinates]rune, numbersMap map[Coordinates]NumberCell) {
	for x, char := range line {
		coord := Coordinates{x: x, y: y}
		if isValidSymbol(char) {
			symbolsMap[coord] = char
		}
		if unicode.IsDigit(char) {
			numbersMap[coord] = NumberCell{char: char, visited: false}
		}
	}
}

func calculatePartNumbersSum(symbolsMap map[Coordinates]rune, numbersMap map[Coordinates]NumberCell) (int, error) {
	total := 0

	for coord := range symbolsMap {
		sum, err := sumAdjacentPartNumbers(coord, numbersMap)
		if err != nil {
			return 0, err
		}
		total += sum
	}

	return total, nil
}

func sumAdjacentPartNumbers(coord Coordinates, numbersMap map[Coordinates]NumberCell) (int, error) {
	sum := 0
	adjacentCoords := getAdjacentCoordsForSymbol(coord)

	for _, adjacentCoord := range adjacentCoords {
		number, err := parseFullNumberAtCoord(adjacentCoord, numbersMap)
		if err != nil {
			return 0, err
		}
		sum += number
	}

	return sum, nil
}

func parseFullNumberAtCoord(coord Coordinates, numbersMap map[Coordinates]NumberCell) (int, error) {
	if cell, exists := numbersMap[coord]; exists && !cell.visited {
		fullNumber := string(cell.char)
		numbersMap[coord] = markNumberCellAsVisited(cell)

		// Extend the number to the left and right
		fullNumber = extendNumber(coord, -1, fullNumber, numbersMap) // Left
		fullNumber = extendNumber(coord, 1, fullNumber, numbersMap)  // Right

		partNumber, err := strconv.Atoi(fullNumber)
		if err != nil {
			return 0, fmt.Errorf("could not parse %s to int: %v", fullNumber, err)
		}

		return partNumber, nil
	}
	return 0, nil
}

func extendNumber(startCoord Coordinates, xAxisFromStartPoint int, currentNumber string, numbersMap map[Coordinates]NumberCell) string {
	for {
		nextCoord := Coordinates{x: startCoord.x + xAxisFromStartPoint, y: startCoord.y}
		cell, exists := numbersMap[nextCoord]
		if !exists || cell.visited || cell.char == 0 {
			break
		}
		numbersMap[nextCoord] = markNumberCellAsVisited(cell)
		if xAxisFromStartPoint < 0 {
			currentNumber = string(cell.char) + currentNumber
		} else {
			currentNumber += string(cell.char)
		}
		startCoord = nextCoord
	}
	return currentNumber
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