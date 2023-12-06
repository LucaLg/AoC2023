package daysix

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DaySix(scanner *bufio.Scanner, file *os.File) {
	defer file.Close()
	var inputArr []string
	var distance []int
	var time []int
	timeAdd := ""
	distanceAdd := ""
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		num := strings.Split(line, ":")
		inputArr = strings.Split(strings.TrimSpace(num[1]), " ")
		for _, char := range inputArr {
			if char != "" {
				c := strings.TrimSpace(char)
				number, err := strconv.Atoi(c)
				if err != nil {
					fmt.Println(err)
					return
				}
				if index == 0 {
					timeAdd += char
					time = append(time, number)
				} else {
					distanceAdd += char
					distance = append(distance, number)
				}
			}
		}
		index++
	}
	mult := 0
	for i := 0; i < len(distance); i++ {
		opts := calcWins(time[i], distance[i])
		if mult == 0 {
			mult = opts
		} else {
			mult *= opts
		}
	}
	fmt.Println(mult)
	//Part Two
	timeTwo, err := strconv.Atoi(timeAdd)
	distanceTwo, err := strconv.Atoi(distanceAdd)
	if err != nil {
		return
	}
	partTwo := calcWins(timeTwo, distanceTwo)
	fmt.Println("Part two", partTwo)
}
func calcWins(time int, distance int) int {
	winOptions := 0
	for ms := 1; ms < time; ms++ {
		possDistance := ms * (time - ms)
		if possDistance > distance {
			winOptions++
		}
	}
	return winOptions
}
