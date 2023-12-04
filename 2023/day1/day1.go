package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var numbersMap = map[string]string{
	"one": "1",
	"two": "2",
	"three": "3",
	"four": "4",
	"five": "5",
	"six": "6",
	"seven": "7",
	"eight": "8",
	"nine": "9",
}

func main() {
	includeVerbal := flag.Bool("includeVerbal", false, "Include verbal expressions of numbers")
	flag.Parse()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		first, err1 := findFirstDigitOrVerbal(line, *includeVerbal)
		last, err2 := findLastDigitOrVerbal(line, *includeVerbal)

		if err1 != nil || err2 != nil {
			log.Println("Error parsing line: ", err1, err2)
			continue
		}

		number, err := strconv.Atoi(first + last)
		if err != nil {
			log.Println("Error converting to number: ", err)
			continue
		}

		total += number
	}

	println(total)
}

func findFirstDigitOrVerbal(line string, includeVerbal bool) (string, error) {
	for i, char := range line {
		if includeVerbal && unicode.IsLetter(char) {
			for key, value := range numbersMap {
				if len(line)-i >= len(key) && line[i:i+len(key)] == key {
					return value, nil
				}
			}
		}
		if unicode.IsDigit(char) {
			return string(char), nil
		}
	}
	return "", fmt.Errorf("no digit or verbal number found")
}

func findLastDigitOrVerbal(line string, includeVerbal bool) (string, error) {
	for i := len(line) - 1; i >= 0; i-- {
		char := rune(line[i])
		if includeVerbal && unicode.IsLetter(char) {
			for key, value := range numbersMap {
				if i+len(key) <= len(line) && line[i:i+len(key)] == key {
					return value, nil
				}
			}
		}
		if unicode.IsDigit(char) {
			return string(char), nil
		}
	}
	return "", fmt.Errorf("no digit or verbal number found")
}
