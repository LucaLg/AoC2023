package dayfour

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func DayFour(scanner *bufio.Scanner, file *os.File) {
	defer file.Close()
	var sum = 0
	for scanner.Scan() {
		line := scanner.Text()
		points := readLine(line)
		if points != -1 {
			sum += points
		}
	}
	fmt.Println(sum)
}
func readLine(line string) int {
	game := strings.Split(line, ":")
	splitNumbers := strings.Split(game[1], "|")
	var sum = 0
	winningNums := []int{}
	winningString := strings.Split(strings.TrimSpace(splitNumbers[0]), " ")
	ownString := strings.Split(strings.TrimSpace(splitNumbers[1]), " ")
	fmt.Println(winningString)
	fmt.Println(ownString)
	for _, nums := range winningString {
		// fmt.Println(nums)
		if nums != "" {
			num, err := strconv.Atoi(nums)
			if err != nil {
				fmt.Println(err)
			}

			winningNums = append(winningNums, num)
		}
	}
	for _, nums := range ownString {
		if nums != "" {
			num, err := strconv.Atoi(nums)
			if err != nil {
				fmt.Println(err)
			}
			if slices.Contains(winningNums, num) {
				if sum == 0 {
					sum += 1
				} else {
					sum *= 2
				}
			}
		}
	}
	// fmt.Println(sum)
	return sum

}
