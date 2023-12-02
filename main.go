package main

import (
	"AoC2023/day_two"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Hello from main")
	// dayone.DayOne()
	var file, scanner, err = readFileToScan("day_two/data.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	daytwo.DayTwo(file, scanner)
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
