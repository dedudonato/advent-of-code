package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	depth := 0
	horizontal := 0
	aim := 0

	for scanner.Scan() {
		instructions := strings.Split(scanner.Text(), " ")
		value, _ := strconv.Atoi(instructions[1])

		if instructions[0] == "forward" {
			horizontal += value
			depth += aim * value
		} else if instructions[0] == "up" {
			aim -= value
		} else {
			aim += value
		}
	}

	fmt.Println(depth * horizontal)
}
