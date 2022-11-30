package main

import "fmt"

var win = [8][3]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

func isCWin(board []string, c string) bool {
	for i := 0; i < 8; i++ {
		if board[win[i][0]] == c && board[win[i][1]] == c && board[win[i][2]] == c {
			return true
		}
	}
	return false
}

func isValid(board []string) bool {
	xCount := 0
	oCount := 0

	for i := 0; i < 9; i++ {
		if board[i] == "X" {
			xCount++
		} else if board[i] == "O" {
			oCount++
		}
	}

	if xCount == oCount+1 || oCount == xCount+1 {
		if isCWin(board, "O") {
			if isCWin(board, "X") {
				return false
			}
		}
		return true
	}
	return false
}

func inProgress(board []string) bool {
	for _, v := range board {
		if v == "-" {
			return true
		}
	}
	return false
}

func CheckBoard(board []string) string {
	if isValid(board) {
		if inProgress(board) {
			return "Game still in progress"
		} else if isCWin(board, "X") {
			return "X wins!"
		} else if isCWin(board, "O") {
			return "O wins!"
		} else {
			return "It's a draw!"
		}
	}
	for _, v := range board {
		if v != "X" && v != "O" && v != "-" {
			return "Invalid Game Board"
		}
	}
	return "Invalid Game Board"
}

func main() {
	var board = make([]string, 9)
	for i := 0; i < 9; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%s", &board[i])
	}

	fmt.Println(board)

	fmt.Println(CheckBoard(board))
}
