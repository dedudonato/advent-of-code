package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DAYS = 256

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lanternfish := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ",")
		for _, v := range values {
			intValue, _ := strconv.Atoi(v)
			lanternfish = append(lanternfish, intValue)
		}
	}

	for i := 0; i < DAYS; i++ {
		size := len(lanternfish)
		for j := 0; j < size; j++ {
			lanternfishValue := lanternfish[j]
			if lanternfishValue == 0 {
				lanternfish[j] = 6
				lanternfish = append(lanternfish, 8)
			} else {
				lanternfish[j] = lanternfishValue - 1
			}
		}
	}

	fmt.Println(len(lanternfish))
}
