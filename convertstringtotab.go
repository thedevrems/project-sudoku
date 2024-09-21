package sudoku

// ConvertStringToTab converts a string into a slice of strings.
// It iterates over each character in the input string and checks if it's either a "."
// or a digit between "0" and "9". If the character is valid, it adds it to the slice.
// If an invalid character is encountered, the function returns nil.
//
// Args:
//   str1: A string representing a row of the Sudoku grid.
//
// Returns:
//   A slice of strings where each element is a valid character from the input string,
//   or nil if the string contains invalid characters.
func ConvertStringToTab(str1 string) []string {
	var tab []string
	for _, ch := range str1 {
		// If the character is valid, add it to the slice
		if ch == '.' || (ch <= '9' && ch >= '1') {
			tab = append(tab, string(ch))
		} else {
			return nil
		}
	}
	return tab
}
