package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lastDepth := 0
	increasingDepthCounter := 0

	for scanner.Scan() {
		currentDepth, _ := strconv.Atoi(scanner.Text())
		if currentDepth > lastDepth && lastDepth != 0 {
			increasingDepthCounter++
		}

		lastDepth = currentDepth
	}

	fmt.Println(increasingDepthCounter)
}
