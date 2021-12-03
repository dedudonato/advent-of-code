package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type compare func(int, int) bool

func greaterOrEqual(a, b int) bool {
	return a >= b
}

func less(a, b int) bool {
	return a < b
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	binaryLength := 0
	inputLength := 0
	var binaryList []string

	for scanner.Scan() {
		binary := scanner.Text()
		binaryList = append(binaryList, binary)

		if binaryLength == 0 {
			binaryLength = len(binary)
		}

		inputLength++
	}

	oxygenBinary := filterList(binaryList, 0, greaterOrEqual)
	oxygenValue, _ := strconv.ParseInt(oxygenBinary, 2, 64)

	co2Binary := filterList(binaryList, 0, less)
	co2Value, _ := strconv.ParseInt(co2Binary, 2, 64)

	fmt.Println(oxygenValue * co2Value)
}

func filterList(binaryList []string, currentIndex int, compareFunction compare) string {
	length := len(binaryList)
	if length == 1 {
		return binaryList[0]
	}

	var onesList []string
	var zerosList []string
	for _, v := range binaryList {
		if rune(v[currentIndex]) == '1' {
			onesList = append(onesList, v)
		} else {
			zerosList = append(zerosList, v)
		}
	}

	if compareFunction(len(onesList), len(zerosList)) {
		return filterList(onesList, currentIndex+1, compareFunction)
	}

	return filterList(zerosList, currentIndex+1, compareFunction)
}
