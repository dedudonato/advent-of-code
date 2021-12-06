package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	first  Coordinate
	second Coordinate
}

type Coordinate struct {
	x int
	y int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var pairs []Pair

	for scanner.Scan() {
		line := scanner.Text()
		first, second := getCoordinates(line)

		if first.x == second.x || first.y == second.y {
			pair := Pair{
				first:  first,
				second: second,
			}

			pairs = append(pairs, pair)
		}
	}

	fmt.Println(getMoreThanTwoIntersections(pairs))
}

func getMoreThanTwoIntersections(pairs []Pair) int {
	intersectionMap := make(map[Coordinate]int)

	for _, p := range pairs {
		if p.first.x == p.second.x {
			min := math.Min(float64(p.first.y), float64(p.second.y))
			max := math.Max(float64(p.first.y), float64(p.second.y))

			for i := int(min); i <= int(max); i++ {
				coord := Coordinate{x: p.first.x, y: i}
				if v, ok := intersectionMap[coord]; ok {
					intersectionMap[coord] = v + 1
				} else {
					intersectionMap[coord] = 1
				}
			}
		} else {
			min := math.Min(float64(p.first.x), float64(p.second.x))
			max := math.Max(float64(p.first.x), float64(p.second.x))

			for i := int(min); i <= int(max); i++ {
				coord := Coordinate{x: i, y: p.first.y}
				if v, ok := intersectionMap[coord]; ok {
					intersectionMap[coord] = v + 1
				} else {
					intersectionMap[coord] = 1
				}
			}
		}
	}

	sum := 0
	for _, v := range intersectionMap {
		if v > 1 {
			sum++
		}
	}

	return sum
}

func getCoordinates(line string) (first Coordinate, second Coordinate) {
	coordinates := strings.Split(line, " -> ")

	firstValues := strings.Split(coordinates[0], ",")
	xFirst, _ := strconv.Atoi(firstValues[0])
	yFirst, _ := strconv.Atoi(firstValues[1])

	first = Coordinate{x: xFirst, y: yFirst}

	secondValues := strings.Split(coordinates[1], ",")
	xSecond, _ := strconv.Atoi(secondValues[0])
	ySecond, _ := strconv.Atoi(secondValues[1])

	second = Coordinate{x: xSecond, y: ySecond}
	return
}
