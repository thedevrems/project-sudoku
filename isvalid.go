package sudoku

// isValid checks if placing a number in a specific cell of the Sudoku grid is valid.
// It verifies that the number does not violate Sudoku rules for the row, column, and 3x3 subgrid (square).
//
// Args:w
//   grille: A 2D slice of strings representing the Sudoku grid.
//   row: The row index of the cell being checked.
//   col: The column index of the cell being checked.
//   num: The number to be placed in the cell, represented as a string.
//
// Returns:
//   true if placing the number in the specified cell is valid according to Sudoku rules,
//   otherwise false.
func isValid(grille [][]string, row int, col int, num string) bool {
	return ResolveTab(grille[row], num) && ResolveTab(GetColumn(grille, col), num) && ResolveTab(GetSquare(grille, row, col), num)
}
