package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type game struct {
	size                    int
	currentBoard, goalBoard []int
}

func getBoard(size int) []int {

	var raw string
	len := size * size

	board := make([]int, len)

	for i := 0; i < size; i++ {

		fmt.Println("	Row", i, ":")
		if _, err := fmt.Scan(&raw); err != nil {
			log.Print(" Scan for row failed, due to ", err)
		}

		split := strings.Split(raw, ",")

		for index, element := range split {
			val, err := strconv.ParseInt(element, 10, 0)
			if err != nil {
				log.Print(" Conversion failed, due to ", err)
			}
			board[index+i] = int(val)

		}

	}
	return board
}

func getUserInput() (int, []int, []int) {
	var size int

	fmt.Println("Enter a Board Size:")

	if _, err := fmt.Scan(&size); err != nil {
		log.Print(" Scan for size failed, due to ", err)
	}

	fmt.Println("Enter the starting state:")
	currentBoard := getBoard(size)
	fmt.Println("Enter the goal state:")
	goalBoard := getBoard(size)

	return size, currentBoard, goalBoard
}

func main() {

	size, currentBoard, goalBoard := getUserInput()

	g := game{
		size:         size,
		currentBoard: currentBoard,
		goalBoard:    goalBoard}

	fmt.Println(g.currentBoard)

}
