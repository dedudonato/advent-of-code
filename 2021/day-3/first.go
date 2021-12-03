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
	binaryLength := 0
	inputLength := 0
	var onesCount []int

	for scanner.Scan() {
		binary := scanner.Text()

		if binaryLength == 0 {
			binaryLength = len(binary)
			onesCount = make([]int, binaryLength)
		}

		for i, v := range binary {
			if v == '1' {
				onesCount[i]++
			}
		}

		inputLength++
	}

	var binaryGammaRate string
	var binaryEpsilonRate string

	for _, v := range onesCount {
		if v > inputLength/2 {
			binaryGammaRate += "1"
			binaryEpsilonRate += "0"
		} else {
			binaryGammaRate += "0"
			binaryEpsilonRate += "1"
		}
	}

	gammaRate, _ := strconv.ParseInt(binaryGammaRate, 2, 64)
	epsilonRate, _ := strconv.ParseInt(binaryEpsilonRate, 2, 64)

	fmt.Println(gammaRate * epsilonRate)
}
