# Sudoku Project

## Project Description

This project aims to solve a Sudoku puzzle using the **backtracking** algorithm, implemented in **Golang**. Backtracking is a depth-first search method often employed in solving puzzles like Sudoku, where each cell is tested for a valid value. This project demonstrates this technique by finding the solution for any valid Sudoku grid.

## Features

- Ability to solve standard 9x9 grids.
- Display of the Sudoku if a solution is found.
- Error handling.
- Verification of row, column, and sub-grid (3x3 regions) constraints.

## Algorithm Used

We use the **backtracking** algorithm, which follows this process:

1. Each empty cell __(an empty cell is represented by a ".")__ is visited one by one.
2. A number between 1 and 9 is placed in the cell, and the program checks if this number respects the Sudoku rules.
3. If a constraint is violated (row, column, or sub-grid), the program tries the next number.
4. If no number works for a cell, the program backtracks to the previous cell and increments its value.
5. This process continues until the entire grid is solved.

The algorithm guarantees to find a solution; otherwise, an error will be displayed on the console.

## Prerequisites

Before running the project, ensure that the following are installed on your machine:

- **Golang** (version 1.23 or later)
- **Git** to clone the repository

## Installation

To install and run the project, follow these steps:

1. Clone the project repository using Git:
   ```bash
   git clone https://ytrack.learn.ynov.com/git/gremy/raid-sudoku
   ```
## Testing the Project:
Here are some commands to test our project :

1. Solvable Sudoku

    ```bash
    go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
    ```
2. Unsolvable Sudoku
    ```bash
    go run . 1 2 3 4
    go run .
    go run . ".96.4...1" "1...6.1.4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
    ```

Please note that the commands should be run in the main project directory !

