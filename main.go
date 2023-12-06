package main

import (
	daysix "AoC2023/day_six"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var file, scanner, err = readFileToScan("day_six/data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	daysix.DaySix(scanner, file)
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
