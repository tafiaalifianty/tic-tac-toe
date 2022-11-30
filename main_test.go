package main

import "testing"

func TestTicTacToe(t *testing.T) {
	value_xwin := []string{
		"X", "O", "X", "X", "O", "O", "X", "X", "O",
	}
	result_xwin := CheckBoard(value_xwin)
	if result_xwin != "X wins!" {
		// unit test failed
		panic("Result is not X wins")

	}

	value_owin := []string{
		"X", "O", "O", "X", "O", "X", "O", "X", "O",
	}
	result_owin := CheckBoard(value_owin)
	if result_owin != "O wins!" {
		// unit test failed
		panic("Result is not O wins")

	}

	value_draw := []string{
		"O", "X", "O", "X", "O", "X", "X", "O", "X",
	}
	result_draw := CheckBoard(value_draw)
	if result_draw != "It's a draw!" {
		// unit test failed
		panic("Result is not draw")

	}

	value_progress := []string{
		"X", "O", "X", "X", "-", "-", "O", "-", "-",
	}
	result_progress := CheckBoard(value_progress)
	if result_progress != "Game still in progress" {
		// unit test failed
		panic("Result is not in progress")

	}

	value_invalid := []string{
		"X", "X", "X", "O", "O", "O", "X", "X", "O",
	}
	result_invalid := CheckBoard(value_invalid)
	if result_invalid != "Invalid Game Board" {
		// unit test failed
		panic("Result is not invalid")

	}
}
