package main

import (
	"fmt"
)

var board [3][3]string
var currentPlayer string

func main() {
	initializeBoard(3, 3)
	currentPlayer = "X"
	winner := ""

	for winner == "" {
		displayBoard()
		fmt.Printf("Player %s's turn\n", currentPlayer)
		row, col := getPlayerMove()
		if isValidMove(row, col) {
			board[row][col] = currentPlayer
			winner = checkWinner()
			if winner == "" {
				switchPlayer()
			}
		} else {
			fmt.Println("Invalid move! Try again.")
		}
	}

	displayBoard()
	if winner == "Draw" {
		fmt.Println("It's a draw!")
	} else {
		fmt.Printf("Player %s wins!\n", winner)
	}
}

/*
func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
}
*/

func initializeBoard(rows, cols int) [][]string {
	board := make([][]string, rows)
	for i := 0; i < rows; i++ {
		board[i] = make([]string, cols)
		for j := 0; j < cols; j++ {
			board[i][j] = " "
		}
	}
	return board
}

func displayBoard() {
	fmt.Println("  0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", board[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("  -+-+-")
		}
	}
	fmt.Println()
}

func getPlayerMove() (int, int) {
	var row, col int
	fmt.Print("Enter row (0, 1, or 2): ")
	fmt.Scan(&row)
	fmt.Print("Enter column (0, 1, or 2): ")
	fmt.Scan(&col)
	return row, col
}

func isValidMove(row, col int) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 || board[row][col] != " " {
		return false
	}
	return true
}

func checkWinner() string {
	for i := 0; i < 3; i++ {
		// Check rows
		if board[i][0] != " " && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return currentPlayer
		}
		// Check columns
		if board[0][i] != " " && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return currentPlayer
		}
	}

	// Check diagonals
	if board[0][0] != " " && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return currentPlayer
	}
	if board[0][2] != " " && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return currentPlayer
	}

	// Check for a draw
	if isBoardFull() {
		return "Draw"
	}

	return ""
}

func isBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}

func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}
