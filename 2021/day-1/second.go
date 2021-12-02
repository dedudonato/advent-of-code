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

	var depths []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentDepth, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, currentDepth)
	}

	increasingDepthCounter := 0
	for i, v := range depths {
		if i > 2 && v > depths[i-3] {
			increasingDepthCounter++
		}
	}

	fmt.Println(increasingDepthCounter)
}
