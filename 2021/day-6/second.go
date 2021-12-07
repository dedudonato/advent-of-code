package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DAYS = 256
const SIZE = 9

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lanternfishsCount := make([]int, SIZE)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ",")
		for _, v := range values {
			intValue, _ := strconv.Atoi(v)
			lanternfishsCount[intValue]++
		}
	}

	for i := 0; i < DAYS; i++ {
		lastValue := 0
		for j := SIZE - 1; j >= 0; j-- {
			if j == 0 {
				lanternfishsCount[6] += lanternfishsCount[0]
				lanternfishsCount[8] += lanternfishsCount[0]
				lanternfishsCount[0] = lastValue
			} else if j == 8 {
				lastValue = lanternfishsCount[8]
				lanternfishsCount[8] = 0
			} else {
				temp := lanternfishsCount[j]
				lanternfishsCount[j] = lastValue
				lastValue = temp
			}
		}
	}

	sum := 0
	for _, v := range lanternfishsCount {
		sum += v
	}

	fmt.Println(sum)
}
