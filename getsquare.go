package sudoku

import "math"

// GetSquare retrieves the 3x3 subgrid (square) from the Sudoku grid that contains the cell at coordinates (y, x).
// It determines the starting row and column of the 3x3 subgrid based on the coordinates (y, x) and extracts
// all elements from that subgrid.
//
// Args:
//   grille: A 2D slice of strings representing the Sudoku grid.
//   y: The row index of the cell whose subgrid is to be retrieved.
//   x: The column index of the cell whose subgrid is to be retrieved.
//
// Returns:
//   A slice of strings containing all elements of the 3x3 subgrid that contains the cell at (y, x).
//
// Example:
//   grille := [9][9]string{
//     {"5", "3", ".",    ".", "7", ".",    ".", ".", "."},
//     {"6", ".", ".",    "1", "9", "5",    ".", ".", "."},
//     {".", "9", "8",    ".", ".", ".",    ".", "6", "."},
//     {"8", ".", ".",    ".", "6", ".",    ".", ".", "3"},
//     {"4", ".", ".",    "8", ".", "3",    ".", ".", "1"},
//     {"7", ".", ".",    ".", "2", ".",    ".", ".", "6"},
//     {".", "6", ".",    ".", ".", ".",    "2", "8", "."},
//     {".", ".", ".",    "4", "1", "9",    ".", ".", "5"},
//     {".", ".", ".",    ".", "8", ".",    ".", "7", "9"},
//   }
//   y := 4
//   x := 5
//   Result:
//   [".", "6", ".", "8", ".", "3", ".", "2", "."]
func GetSquare(grille [][]string, y int, x int) []string {
	// Determine the size of the square
	dimensionsquare := int(math.Sqrt(float64(len(grille))))
	// Determine the starting coordinates of the 3x3 square
	startRow := (y / dimensionsquare) * dimensionsquare
	startCol := (x / dimensionsquare) * dimensionsquare

	// Extract the 9 elements of the square
	square := []string{}
	for i := 0; i < dimensionsquare; i++ {
		for j := 0; j < dimensionsquare; j++ {
			square = append(square, grille[startRow+i][startCol+j])
		}
	}

	return square
}
