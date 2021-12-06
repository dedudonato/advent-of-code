package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Board struct {
	grid   [][]int
	values map[int]bool
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int
	var currentBoard []string
	var boards []Board

	for scanner.Scan() {
		line := scanner.Text()
		if len(numbers) == 0 {
			numbers = generateNumbers(line)
		} else if line == "" && len(currentBoard) == 5 {
			boards = append(boards, *newBoard(currentBoard))
			currentBoard = make([]string, 0)
		} else if line != "" {
			currentBoard = append(currentBoard, line)
		}
	}

	boards = append(boards, *newBoard(currentBoard))

	lastNumber, sum := playLastBingo(boards, numbers)

	fmt.Println(lastNumber * sum)
}

func generateNumbers(line string) []int {
	var result []int
	for _, v := range strings.Split(line, ",") {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}

	return result
}

func newBoard(rows []string) *Board {
	b := &Board{
		grid:   make([][]int, 5),
		values: make(map[int]bool, 25),
	}

	for i, row := range rows {
		for j, v := range strings.Fields(row) {
			integerValue, _ := strconv.Atoi(v)

			if j == 0 {
				b.grid[i] = make([]int, 5)
			}

			b.grid[i][j] = integerValue
			b.values[integerValue] = true
		}
	}

	return b
}

func playLastBingo(boards []Board, numbers []int) (lastNumber int, sum int) {
	numbersDrawn := make(map[int]bool)

	for _, v := range numbers {
		numbersDrawn[v] = true
		for _, b := range boards {
			if b.values[v] {
				if hasWon(b, v, numbersDrawn) {
					boards = remove(boards, b)
					if len(boards) == 0 {
						return v, sumOfUndrawn(b, numbersDrawn)
					}
				}
			}
		}
	}

	return -1, -1
}

func remove(boards []Board, boardToRemove Board) []Board {
	index := -1
	for i, b := range boards {
		if reflect.DeepEqual(b, boardToRemove) {
			index = i
		}
	}
	boards[index] = boards[len(boards)-1]
	return boards[:len(boards)-1]
}

func hasWon(board Board, value int, numbersDrawn map[int]bool) bool {
	x := 0
	y := 0

	for i := range board.grid {
		for j, v := range board.grid[i] {
			if v == value {
				x = i
				y = j
			}
		}
	}

	columnResult := true

	for i := 0; i < 5; i++ {
		if ok := numbersDrawn[board.grid[i][y]]; !ok {
			columnResult = false
		}
	}

	if columnResult {
		return true
	}

	for j := 0; j < 5; j++ {
		if ok := numbersDrawn[board.grid[x][j]]; !ok {
			return false
		}
	}

	return true
}

func sumOfUndrawn(board Board, numbersDrawn map[int]bool) int {
	sum := 0

	for i := range board.grid {
		for _, v := range board.grid[i] {
			if ok := numbersDrawn[v]; !ok {
				sum += v
			}
		}
	}

	return sum
}
