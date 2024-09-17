package main

import (
	"os"
	"sudoku"
)

// main is the entry point of the program. It initializes the Sudoku grid from
// the program's command-line arguments, checks if the grid can be solved,
// and then prints the solved grid or an error message if it cannot be solved.
func main() {
	// Doing some tests before trying to find a solution
	grilleSudoku := sudoku.PreviousStart(os.Args)
	//  Checks if the Sudoku grid can be solved
	if sudoku.ResolveSudoku(grilleSudoku) {
		// Prints the solved Sudoku grid
		sudoku.PrintSudoku(grilleSudoku)
	} else {
		// Prints an error message if the grid cannot be solved
		sudoku.PrintError("Unable to resolve the grid.")
	}
}
