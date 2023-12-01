package dayone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func DayOne() {
	fmt.Println("Hello from dayOne")
	file, err := os.Open("day_one/data2.txt")

	if err != nil {
		fmt.Println("Error file reading", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var sum int = 0
	for scanner.Scan() {
		line := scanner.Text()
		line = replaceNumbers(line)
		number := getNumber(line)
		sum += number
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	println(sum)
}
func replaceNumbers(line string) string {
	newLine := strings.ReplaceAll(line, "one", "one1one")
	newLine = strings.ReplaceAll(newLine, "two", "two2two")
	newLine = strings.ReplaceAll(newLine, "three", "three3three")
	newLine = strings.ReplaceAll(newLine, "four", "four4four")
	newLine = strings.ReplaceAll(newLine, "five", "five5five")
	newLine = strings.ReplaceAll(newLine, "six", "six6six")
	newLine = strings.ReplaceAll(newLine, "seven", "seven7seven")
	newLine = strings.ReplaceAll(newLine, "eight", "eight8eight")
	newLine = strings.ReplaceAll(newLine, "nine", "nine9nine")
	return newLine
}
func getNumber(line string) int {
	var firstDigit, lastDigit rune

	for _, char := range line {
		if unicode.IsDigit(char) {
			if firstDigit == 0 {
				firstDigit, lastDigit = char, char
			} else {
				lastDigit = char
			}
		}
	}
	addedNumber := string(firstDigit) + string(lastDigit)
	if res, err := strconv.Atoi(addedNumber); err == nil {
		return res
	}
	return -1
}
