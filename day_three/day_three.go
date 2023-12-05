package daythree

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

type LineToCheck struct {
	Y    int
	Line string
}
type Coord struct {
	x int
	y int
}

func DayThree(scanner *bufio.Scanner, file *os.File) {
	defer file.Close()
	var sum int
	var inputArr []string
	for scanner.Scan() {
		line := scanner.Text()
		inputArr = append(inputArr, line)
	}
	sum, posMap := lineScan(inputArr)
	partTwo(inputArr, posMap)
	fmt.Println(sum)
}
func partTwo(inputArr []string, coordMap map[Coord]int) {
	var sum int
	for y, line := range inputArr {
		for x, char := range line {
			if char == '*' {
				coords := getNeighbourNumbers(inputArr, x, y, len(line))
				var vals []int
				for _, coord := range coords {
					if !slices.Contains(vals, coordMap[coord]) {
						vals = append(vals, coordMap[coord])
					}
				}
				if len(vals) == 2 {
					mult := vals[0] * vals[1]
					sum += mult
				}
			}
		}
		fmt.Println("Part two", sum)
	}
}
func lineScan(inputArr []string) (int, map[Coord]int) {
	var posMap = make(map[Coord]int)
	var sum = 0
	for y, line := range inputArr {
		var num = ""
		var isValid = false
		var coords []Coord
		for x, char := range line {
			if unicode.IsDigit(char) {
				if !isValid {
					isValid = checkIfNumValid(inputArr, x, y, len(line))
				}
				coords = append(coords, Coord{x: x, y: y})
				num = num + string(char)
			}
			if (!unicode.IsDigit(char) && num != "") || (x == len(line)-1 && num != "") {
				if isValid {
					number, err := strconv.Atoi(num)
					if err != nil {
						fmt.Println(err)
						return -1, nil
					}
					sum += number
					for _, c := range coords {
						posMap[c] = number
					}
				}
				clear(coords)
				isValid = false
				num = ""
			}
		}
	}
	return sum, posMap
}
func getNeighbourNumbers(inputArr []string, x int, y int, lineLength int) []Coord {
	neighbours := [][]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {0, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	var coords []Coord
	for _, cell := range neighbours {
		currY := cell[1] + y
		currX := cell[0] + x
		if 0 <= currX && currX < lineLength && 0 <= currY && currY < len(inputArr) {
			value := inputArr[currY][currX]
			if unicode.IsDigit(rune(value)) {

				coords = append(coords, Coord{x: currX, y: currY})
			}
		}
	}
	return coords
}
func checkIfNumValid(inputArr []string, x int, y int, lineLength int) bool {
	neighbours := [][]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {0, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	for _, cell := range neighbours {
		currY := cell[1] + y
		currX := cell[0] + x
		if 0 <= currX && currX < lineLength && 0 <= currY && currY < len(inputArr) {
			value := inputArr[currY][currX]
			if value != '.' && !unicode.IsDigit(rune(value)) {
				return true
			}
		}
	}
	return false
}
