package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type game struct {
	currentBoard, goalBoard []int
	scores                  map[int]int
	totalMoves              int
	sumManhattan            int
	NilssonScore            int
}

type pair struct {
	a, b interface{}
}

const length = 9
const size = 3

var order = []int{0, 1, 2, 5, 8, 7, 6, 3, 4}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func getBoard() []int {

	var raw string

	board := make([]int, 0)

	for i := 0; i < size; i++ {

		fmt.Println("	Row", i, ":")
		if _, err := fmt.Scan(&raw); err != nil {
			log.Print(" Scan for row failed, due to ", err)
		}

		split := strings.Split(raw, ",")

		for _, element := range split {
			val, err := strconv.ParseInt(element, 10, 0)
			if err != nil {
				log.Print(" Conversion failed, due to ", err)
			}
			board = append(board, int(val))

		}

	}
	if len(board) != length {
		log.Print(" Wrong length of board")
	}
	return board
}

func getUserInput() ([]int, []int) {

	fmt.Println("Add your Boards, separate with commas, use 0 as the free space")

	fmt.Println("Enter the starting state:")
	currentBoard := []int{0, 1, 3, 8, 2, 4, 7, 6, 5} //getBoard()
	fmt.Println("Enter the goal state:")
	goalBoard := []int{1, 2, 3, 8, 0, 4, 7, 6, 5} //getBoard()

	return currentBoard, goalBoard
}

func searchVals(list []int, thing int) int {
	idx := -1

	for i := range list {
		if list[i] == thing {
			idx = i
			break
		}
	}

	return idx
}

func getNilsson(g game) int {
	val := 0

	for idx, element := range order {
		if idx == 0 {
			continue
		}
		if element == 4 {
			val += 3
			break
		}

		cur := pair{g.currentBoard[order[idx]], g.currentBoard[order[idx+1]]}
		goal := pair{g.goalBoard[order[idx]], g.goalBoard[order[idx+1]]}

		if cur != goal {
			val += 6

		}

	}

	return getManhattan(g) + val
}

func getManhattan(g game) int {
	val := 0

	for idx, element := range g.currentBoard {

		if element == 0 {
			continue
		}

		goalIdx := searchVals(g.goalBoard, element)

		val += (abs(idx/size-goalIdx/size) + abs(idx%size-goalIdx%size))
	}

	return val
}

func getMoves(g game) {

	idx := searchVals(g.currentBoard, 0)

	fmt.Println(idx)

	if idx%size > 1 {
		fmt.Println("Up")

	} else {

	}

	if idx%size < size {
		fmt.Println("Down")

	} else {

	}

	if idx/size > 1 {
		fmt.Println("Left")

	} else {

	}

	if idx/size < size {
		fmt.Println("Right")

	} else {

	}

}

func main() {

	currentBoard, goalBoard := getUserInput()

	g := game{
		currentBoard: currentBoard,
		goalBoard:    goalBoard,
		totalMoves:   0,
		sumManhattan: int(math.Inf(1)),
		NilssonScore: int(math.Inf(1))}

	fmt.Println(getNilsson(g))
	getMoves(g)

}
