package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	fmt.Printf("max")
// 	time.Sleep(3 * time.Second)
// 	// fmt.Printf("\033[0;0H")
// 	fmt.Printf("\033[A\r")
// 	fmt.Printf("lol")
// }

var DEBUG bool = false

func log(format string, a ...interface{}) {
	trueFormat := fmt.Sprintf("%s\n", format)
	fmt.Printf(trueFormat, a...)
}

func debug(format string, a ...interface{}) {
	if DEBUG {
		log(format, a...)
	}
}

func getLiveCount(board [][]int, xpos int, ypos int) int {
	count := 0
	debug("Around [%d][%d]", xpos, ypos)

	for _, xval := range []int{xpos - 1, xpos, xpos + 1} {
		for _, yval := range []int{ypos - 1, ypos, ypos + 1} {
			if !(xval == xpos && yval == ypos) &&
				xval >= 0 && yval >= 0 &&
				xval < len(board) && yval < len(board[xval]) &&
				board[xval][yval] == 1 {
				debug("[%d][%d] is live", xval, yval)
				count++
			}
		}
	}
	return count
}

func printBoard(board [][]int) {
	// for i := 0; i <= len(board); i++ {
	// 	fmt.Printf("\033[A")
	// }
	// fmt.Printf("\r")
	fmt.Print("Board: \r\n")
	for xpos := range board {
		// print(xpos)
		for ypos := range board[xpos] {
			// print(ypos)
			// fmt.Printf("[%d][%d][%d]\n", xpos, ypos, board[xpos][ypos])
			if board[xpos][ypos] == 1 {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\r\n")
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	board := make([][]int, 10)
	// set the board

	for xpos := range board {
		board[xpos] = make([]int, 10)
	}

	for xpos := range board {
		for ypos := range board[xpos] {
			if rand.Intn(10) < 4 {
				board[xpos][ypos] = 1
			} else {
				board[xpos][ypos] = 0
			}
		}
	}

	printBoard(board)
	// run
	for {
		newBoard := board
		for xpos := range board {
			for ypos := range board[xpos] {
				liveCount := getLiveCount(board, xpos, ypos)
				debug("[%d][%d] = [%d][%d]\n", xpos, ypos, board[xpos][ypos], liveCount)
				if board[xpos][ypos] == 1 && (liveCount < 2 || liveCount > 3) {
					newBoard[xpos][ypos] = 0
				} else if board[xpos][ypos] == 0 && liveCount == 3 {
					newBoard[xpos][ypos] = 1
				}
			}
		}
		board = newBoard

		printBoard(board)
		time.Sleep(time.Second)
	}

}
