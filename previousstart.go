package sudoku

// PreviousStart takes a list of command-line arguments and returns a Sudoku grid
// as a 2D array. It expects exactly 9 strings, each containing 9 characters corresponding
// to the rows of the Sudoku grid. Additionally, each of these 9 characters must be either
// a "." or a value between "1" and "9". If the input does not meet these conditions,
// an error message is printed and the function returns nil.
//
// Args:
//   args: A slice of strings where the first element is the program name (ignored)
//         and the remaining 9 elements are strings representing the rows of the Sudoku grid.
//
// Returns:
//   A 2D slice of strings representing the Sudoku grid, or nil if the input is invalid.
func PreviousStart(args []string) [][]string {
	// Error handling if the number of arguments is incorrect
	if len(args) != 10 {
		PrintError("Incorrect number of arguments")
		return nil
	}

	// Getting grid dimensions
	tailleGrilleSudoku := len(args) - 1 // Exclude the program name
	grilleSudoku := make([][]string, tailleGrilleSudoku)

	// Adding each converted line in the grid
	for i := 1; i <= tailleGrilleSudoku; i++ {
		// If a string doesn't contain exactly 9 characters, return an error
		if len(args[i]) != tailleGrilleSudoku {
			PrintError("One or more rows do not have the correct length")
			return nil
		} else {
			// Convert the string into a row of the grid
			if ConvertStringToTab(args[i]) != nil {
				grilleSudoku[i-1] = ConvertStringToTab(args[i])
			} else {
				PrintError("Invalid input values for the Sudoku grid")
				return nil
			}
		}
	}
	return grilleSudoku
}
