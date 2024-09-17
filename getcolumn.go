package sudoku

// GetColumn returns the column at index `c` from a Sudoku grid.
// It extracts and returns the entire column as a slice of strings.
//
// Args:
//   grille: A 2D slice of strings representing the Sudoku grid.
//   c: The index of the column to be retrieved.
//
// Returns:
//   A slice of strings representing the column at index `c` from the Sudoku grid.
func GetColumn(grille [][]string, c int) []string {
	column := make([]string, 9)
	for i := 0; i < 9; i++ {
		column[i] = grille[i][c]
	}
	return column
}
    