package sudoku

import (
	"fmt"
)

const (
	green = "\033[32m" // Green color
	blue  = "\033[34m" // Blue color
	reset = "\033[0m"  // Reset color
)

// PrintSudoku displays the Sudoku grid with colors for better visibility.
// It prints the grid where each number is displayed in green, with vertical and horizontal separators
// in blue to demarcate 3x3 subgrids.
//
// Args:
//
//	grille: A 2D slice of strings representing the Sudoku grid to be printed.
//
// The grid is formatted as follows:
// - Each number is displayed in green.
// - Vertical separators (|) in blue are added after every 3 cells within a row.
// - Horizontal separators (---) in blue are added after every 3 rows, separating 3x3 subgrids.
func PrintSudoku(grille [][]string) {
	for y := 0; y < len(grille); y++ {
		for x := 0; x < len(grille[y]); x++ {
			// Display the character in green
			char := rune(grille[y][x][0])
			fmt.Printf("%s%c%s", green, char, reset)

			if x != len(grille[y])-1 {
				fmt.Print(" ") // Use fmt.Print for space
			}
			// Add a vertical separator in blue after every 3x3 cell
			if (x+1)%3 == 0 && x != len(grille[y])-1 {
				fmt.Printf("%s| %s", blue, reset) // Use fmt for separator
			}
		}
		// New line after each row
		fmt.Println() // Use fmt.Println for new line

		// Add a horizontal separator in blue after every group of 3 rows
		if (y+1)%3 == 0 && y != len(grille)-1 {
			fmt.Printf("%s", blue)    // Start blue color
			for i := 0; i < 21; i++ { // 21 dashes to separate groups of 3 rows
				fmt.Print("-")
			}
			fmt.Printf("%s\n", reset) // Reset color and new line
		}
	}
}
