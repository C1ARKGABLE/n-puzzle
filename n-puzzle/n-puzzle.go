package main

// Imports the necessary packages
import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

// Search order for the Nilsson Index. Search the grid in a clockwise manner
// 0,1,2
// 3,4,5
// 6,7,8
var order = []int{0, 1, 2, 5, 8, 7, 6, 3, 4}

// Default sizes of the board... Slice must be 9 ints long [8]int. Width of board is 3
const length = 9
const width = 3

// Main structure for keeping track of the game board and moves
type game struct {
	currentBoard, goalBoard []int
	totalMoves              int
	nilssonIn               bool
	manhattanScore          int
	nilssonScore            int
	f                       int
	moves                   []string
}

// Pair or tuple type group... Python brain wants to use Python data types
type pair struct {
	a, b interface{}
}

func abs(x int) int {
	// Abolute value of two ints... Idk why Go doesn't have this
	if x < 0 {
		return -x
	}
	return x
}

func getBoard() []int {
	// This gets a board input from the user
	var raw string

	board := make([]int, 0)

	for i := 0; i < width; i++ {
		// For each row

		fmt.Println("	Row", i, ":")
		// Grab string values. seperated by commas.
		if _, err := fmt.Scan(&raw); err != nil {
			log.Print(" Scan for row failed, due to ", err)
		}

		// Split the string into a slice of strings
		split := strings.Split(raw, ",")

		for _, element := range split {
			// For each element in slice of strings convert to int
			val, err := strconv.ParseInt(element, 10, 0)
			if err != nil {
				log.Print(" Conversion failed, due to ", err)
			}
			// Add the int to a board (aka slice of ints)
			board = append(board, int(val))

		}

	}
	// Check if the user is inputing more or less values than they should
	if len(board) != length {
		log.Print(" Wrong length of board")
	}
	// Yay, this is now a board
	return board
}

func yesNo() bool {
	// This asks the user a yes or no question and returns a boolean answer.
	var raw string

	if _, err := fmt.Scan(&raw); err != nil {
		log.Print(" Scanning", err)
	}

	return string(raw[0]) == "y" || string(raw[0]) == "Y"
}

func getUserInput() ([]int, []int, bool) {

	// Default values:
	NilssonIn := false

	// This is a test

	currentBoard := []int{8, 6, 7, 2, 5, 4, 3, 0, 1}
	goalBoard := []int{1, 2, 3, 4, 5, 6, 7, 8, 0}

	fmt.Println("Use defaults? (y/n)")
	if !yesNo() {
		// If user wants to enter non-default values, ask these questions:
		fmt.Println("Would you like to use Nilsson scoring as the heuristic? (y/n)")
		NilssonIn = yesNo()
		fmt.Println("Add your Boards, separate with commas, use 0 as the free space")
		fmt.Println("Enter the starting state:")
		currentBoard = getBoard()
		fmt.Println("Enter the goal state:")
		goalBoard = getBoard()

	}
	// Return stuff
	return currentBoard, goalBoard, NilssonIn
}

func searchSlice(s1 []int, thing int) int {
	// Find index of int in a slice
	idx := -1

	for i := range s1 {
		if s1[i] == thing {
			idx = i
			break
		}
	}

	return idx
}

func compareSlice(s1 []int, s2 []int) bool {
	// Compare two slices...
	// If they are the same at all indices:
	// return true
	// else return false

	for idx := range s1 {
		if s1[idx] != s2[idx] {
			return false
		}
	}
	return true
}

func searchList(s1 [][]int, s2 []int) bool {
	// Is a slice in a list of slices?
	found := false
	for _, i := range s1 {
		found = found || compareSlice(i, s2)
	}
	return found
}

func getNilsson(g game) int {
	// Search clockwise around the board.
	// Pair current and next state clockwise together
	// If the pair is the same in the current state as in the goal state?
	// If not, score increases by six
	// If not, no increase
	// Add score to the manhattan distance, return
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

		goalIdx := searchSlice(g.goalBoard, element)

		val += (abs(idx/width-goalIdx/width) + abs(idx%width-goalIdx%width))
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

func printBoard(g game, md bool) {

	if md {
		fmt.Println(".| **S** |.||.| **G** |.")
		fmt.Println("-|-|-|-|-|-|-")
		for i := 0; i < length; i += 3 {
			fmt.Println(g.currentBoard[i], "|", g.currentBoard[i+1], "|", g.currentBoard[i+2], "|->|", g.goalBoard[i], "|", g.goalBoard[i+1], "|", g.goalBoard[i+2])
		}
	} else {
		space := "- - - - - - -  - - - - - - -"
		fmt.Println(space)

		for i := 0; i < length; i += 3 {
			fmt.Println("|", g.currentBoard[i], "|", g.currentBoard[i+1], "|", g.currentBoard[i+2], "|->|", g.goalBoard[i], "|", g.goalBoard[i+1], "|", g.goalBoard[i+2], "|")
			fmt.Println(space)
		}
	}
}

func (g game) getMoves(past [][]int) []game {

	moves := []game{}

	idx := searchSlice(g.currentBoard, 0)

	if idx > 2 {
		move := moveBoard(g, idx, -3, "down")
		if !searchList(past, move.currentBoard) {
			moves = append(moves, move)
		}

		//up
	}
	if idx < 6 {

		move := moveBoard(g, idx, 3, "up")
		if !searchList(past, move.currentBoard) {
			moves = append(moves, move)
		}
		//down
	}
	if idx%width != 2 {

		move := moveBoard(g, idx, 1, "left")
		if !searchList(past, move.currentBoard) {
			moves = append(moves, move)
		}
		//left
	}
	if idx%width != 0 {

		move := moveBoard(g, idx, -1, "right")
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

	if len(moves) < 1 {
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

	printBoard(g, true)

	var movesQueue []game
	var pastStates [][]int
	nodes := 0

	for !compareSlice(g.currentBoard, g.goalBoard) {

		for _, i := range g.getMoves(pastStates) {
			movesQueue = append(movesQueue, i)

		}
		pastStates = append(pastStates, g.currentBoard)

		g, movesQueue = getLowest(append(movesQueue, g.getMoves(pastStates)...))

		nodes++

		if len(movesQueue) < 2 {
			fmt.Println("Unsolvable")
			break
		}

	}

	fmt.Println("Sequence:", g.moves, "\nNodes expanded:", nodes, "\nNodes visited:", g.totalMoves)
}
