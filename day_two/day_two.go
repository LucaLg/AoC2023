package daytwo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DayTwo(file *os.File, scanner *bufio.Scanner) {
	defer file.Close()
	var sum = 0
	var maxSum = 0
	for scanner.Scan() {
		line := scanner.Text()
		id, maxMult, err := checkLine(line)
		if err != nil {

			return
		}
		if id != -1 {
			fmt.Println(id)
			sum += id
		}
		maxSum += maxMult
	}
	fmt.Println(maxSum)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error during scanning:", err)
	}
}
func checkLine(line string) (int, int, error) {
	var m = make(map[string]int)
	var maxMap = make(map[string]int)
	s := strings.Split(line, ":")
	id, err := strconv.Atoi(strings.Split(s[0], " ")[1])
	if err != nil {
		return -1, -1, err
	}
	colorOperations := strings.Split(s[1], ";")
	for _, op := range colorOperations {
		roundOp := strings.Split(op, ",")
		for _, f := range roundOp {
			trimmed := strings.TrimSpace(f)
			splitOp := strings.Split(trimmed, " ")
			color := splitOp[1]
			amount, err := strconv.Atoi(splitOp[0])
			if err != nil {
				return -1, -1, err
			}
			m[color] = m[color] + amount
		}
		for color, val := range m {
			if val > maxMap[color] {
				maxMap[color] = val
			}
		}
		for k := range m {
			delete(m, k)
		}

	}
	fmt.Println(maxMap)
	var maxMult = 1
	for _, val := range maxMap {
		maxMult *= val
	}
	if m["green"] <= 13 && m["red"] <= 12 && m["blue"] <= 14 {

		return id, maxMult, nil
	}
	return -1, maxMult, nil

}
