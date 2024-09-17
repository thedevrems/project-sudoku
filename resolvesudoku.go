package sudoku

import "strconv"

// ResolveSudoku solves a Sudoku grid using the backtracking method.
// The function returns true if the Sudoku is successfully solved, otherwise false.
//
// The function iterates over the grid and, for each empty cell (represented by "."),
// it attempts to place a number from 1 to 9. If a valid number can be placed,
// the function recursively calls itself to solve the remaining grid. If a conflict arises
// later in the grid, it backtracks by resetting the cell and trying another number.
//
// Args:
//   grille: A 2D slice of strings representing the Sudoku grid, where "." represents an empty cell.
//
// Returns:
//   true if the Sudoku grid is successfully solved, otherwise false.
func ResolveSudoku(grille [][]string) bool {
	// Iterate through each cell of the grid
	for y := 0; y < len(grille); y++ {
		for x := 0; x < len(grille[y]); x++ {
			// If the cell is empty (represented by ".")
			if grille[y][x] == "." {
				// Try each number from 1 to 9
				for num := 1; num <= 9; num++ {
					// Convert the number to a string
					char := strconv.Itoa(num)
					// Check if the number can be placed at position (y, x)
					if isValid(grille, y, x, char) {
						// Place the number in the cell
						grille[y][x] = char
						// Recursively attempt to solve the rest of the grid
						if ResolveSudoku(grille) {
							return true // If solved, return true
						}
						// If the solution does not work, reset the cell (backtracking)
						grille[y][x] = "."
					}
				}
				// If no number works, return false to backtrack
				return false
			}
		}
	}
	// If all cells are filled correctly, return true
	return true
}
