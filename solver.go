package main

type SudokuGrid [9][9]int

func isValid(grid *SudokuGrid, row int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if grid[i][col] == num {
			return false
		}
	}

	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

func solve(grid *SudokuGrid) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(grid, i, j, num) {
						grid[i][j] = num

						if solve(grid) {
							return true
						} else {
							grid[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}
