package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type game struct {
	currentBoard, goalBoard []int
	totalMoves              int
	nilssonIn               bool
	manhattanScore          int
	nilssonScore            int
	f                       int
	moves                   []string
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

func yesNo() bool {
	var raw string

	if _, err := fmt.Scan(&raw); err != nil {
		log.Print(" Scan for row failed, due to ", err)
	}

	return string(raw[0]) == "y" || string(raw[0]) == "Y"
}

func getUserInput() ([]int, []int, bool) {
	NilssonIn := true

	currentBoard := []int{1, 2, 5, 6, 3, 4, 7, 8, 0}
	goalBoard := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// currentBoard := []int{1, 2, 3, 7, 4, 5, 6, 8, 0}
	// goalBoard := []int{1, 2, 3, 8, 6, 4, 7, 5, 0}

	fmt.Println("Use defaults? (y/n)")
	if !yesNo() {
		fmt.Println("Would you like to use Nilsson scoring as the heuristic? (y/n)")
		NilssonIn = yesNo()
		fmt.Println("Add your Boards, separate with commas, use 0 as the free space")
		fmt.Println("Enter the starting state:")
		currentBoard = getBoard()
		fmt.Println("Enter the goal state:")
		goalBoard = getBoard()

	}

	return currentBoard, goalBoard, NilssonIn
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

func searchList(list [][]int, thing []int) bool {
	for _, i := range list {
		for j := range i {
			if i[j] != thing[j] {
				return false
			}

		}
	}
	return true
}

func getNilsson(g game) int {
	val := 0

	for idx := range order {
		if idx == 8 {
			val += 3
			break
		}
		if idx == 0 {
			continue
		}

		cur := pair{g.currentBoard[order[idx%8]], g.currentBoard[order[(idx+1)%8]]}
		goal := pair{g.goalBoard[order[idx%8]], g.goalBoard[order[(idx+1)%8]]}

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

func (g game) updateScores() game {
	if g.nilssonIn {
		g.nilssonScore = getNilsson(g)
		g.f = g.nilssonScore + g.totalMoves
	} else {
		g.manhattanScore = getManhattan(g)
		g.f = g.manhattanScore + g.totalMoves
	}

	return g

}

func moveBoard(g game, idx int, dif int, move string) game {

	newG := g

	newG.currentBoard = make([]int, len(g.currentBoard))
	copy(newG.currentBoard, g.currentBoard)

	newG.moves = make([]string, len(g.moves))
	copy(newG.moves, g.moves)

	blank := newG.currentBoard[idx]
	if blank != 0 {
		log.Println("blank is ", blank)
	}

	newG.currentBoard[idx] = newG.currentBoard[idx+dif]

	newG.currentBoard[idx+dif] = blank

	newG.totalMoves++

	newG = newG.updateScores()

	newG.moves = append(newG.moves, move)

	return newG

}
func printBoard(g game) {
	space := "- - - - - - -  - - - - - - -"
	fmt.Println(space)

	for i := 0; i < length; i += 3 {
		fmt.Println("|", g.currentBoard[i], "|", g.currentBoard[i+1], "|", g.currentBoard[i+2], "|->|", g.goalBoard[i], "|", g.goalBoard[i+1], "|", g.goalBoard[i+2], "|")
		fmt.Println(space)
	}

}

func (g game) getMoves(past [][]int) []game {

	moves := []game{}

	idx := searchVals(g.currentBoard, 0)

	if idx > 2 {
		move := moveBoard(g, idx, -3, "up")
		if !searchList(past, move.currentBoard) {
			moves = append(moves, move)
		}

		//up
	}
	if idx < 6 {

		move := moveBoard(g, idx, 3, "down")
		if !searchList(past, move.currentBoard) {
			moves = append(moves, move)
		}
		//down
	}
	if idx%size != 2 {

		move := moveBoard(g, idx, 1, "right")
		if !searchList(past, move.currentBoard) {
			moves = append(moves, move)
		}
		//left
	}
	if idx%size != 0 {

		move := moveBoard(g, idx, -1, "left")
		if !searchList(past, move.currentBoard) {
			moves = append(moves, move)
		}
		//right
	}

	return moves

}
func getLowest(moves []game) (game, []game) {

	sort.Slice(moves, func(i, j int) bool {
		return moves[i].f > moves[j].f
	})

	if len(moves) < 2 {
		return game{}, []game{}
	}

	return moves[len(moves)-1], moves[:len(moves)-1]

}

func main() {

	currentBoard, goalBoard, nilssonIn := getUserInput()

	g := game{
		currentBoard: currentBoard,
		goalBoard:    goalBoard,
		totalMoves:   0,
		nilssonIn:    nilssonIn,
		moves:        []string{}}

	g = g.updateScores()

	printBoard(g)

	var movesQueue []game
	var pastStates [][]int
	nodes := 0

	for (3 < g.nilssonScore && g.nilssonIn) || (g.manhattanScore > 0 && !g.nilssonIn) {

		for _, i := range g.getMoves(pastStates) {
			movesQueue = append(movesQueue, i)

		}
		pastStates = append(pastStates, g.currentBoard)
		g, movesQueue = getLowest(append(movesQueue, g.getMoves(pastStates)...))
		nodes++
		if len(movesQueue) < 1 {
			fmt.Println("Unsolvable")
			break
		}

	}

	fmt.Println("Sequence:", g.moves, "\nNodes expanded:", nodes, "\nNodes visited:", g.totalMoves)

}
