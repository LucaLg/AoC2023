package main

import (
	daythree "AoC2023/day_three"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var file, scanner, err = readFileToScan("day_three/data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// dayfour.DayFour(scanner, file)
	daythree.DayThree(scanner, file)
}
func readFileToScan(path string) (*os.File, *bufio.Scanner, error) {

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error file reading", err)
		return nil, nil, err
	}
	scanner := bufio.NewScanner(file)
	return file, scanner, nil
}
