package sudoku

// ResolveTab checks if a number can be placed in a given slice of strings without conflicts.
// It determines whether the number is already present in the slice.
//
// Args:
//   tab: A slice of strings representing a row, column, or 3x3 subgrid of the Sudoku grid.
//   num: The number to be checked, represented as a string.
//
// Returns:
//   true if the number is not present in the slice (meaning it can be placed),
//   otherwise false if the number is already present.
func ResolveTab(tab []string, num string) bool {
	for i := 0; i < len(tab); i++ {
		if tab[i] == num {
			return false
		}
	}
	return true
}
