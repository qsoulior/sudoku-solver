package main

import "github.com/qsoulior/sudoku-solver/solver"

type Grid = solver.Grid

// base
func isValid(values *Grid, row int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if values[row][i] == num {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if values[i][col] == num {
			return false
		}
	}

	return true
}

// classic
func isValidClassic(values *Grid, row int, col int, num int) bool {
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if values[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return isValid(values, row, col, num)
}

// X
func isValidX(values *Grid, row int, col int, num int) bool {
	if row == col {
		for i := 0; i < 9; i++ {
			if values[i][i] == num {
				return false
			}
		}
	}

	if row == 8-col {
		for i := 0; i < 9; i++ {
			if values[i][8-i] == num {
				return false
			}
		}
	}

	return isValidClassic(values, row, col, num)
}

// figured
func isValidFigured(values *Grid, layout *Grid, row int, col int, num int) bool {
	color := layout[row][col]
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if color == layout[i][j] && values[i][j] == num {
				return false
			}
		}
	}

	return isValid(values, row, col, num)
}

// classic
func solveClassic(values *Grid) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if values[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if isValidClassic(values, i, j, num) {
						values[i][j] = num

						if solveClassic(values) {
							return true
						} else {
							values[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}

// X
func solveX(values *Grid) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if values[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if isValidX(values, i, j, num) {
						values[i][j] = num

						if solveX(values) {
							return true
						} else {
							values[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}

// odd-even
func solveOddEven(values *Grid, layout *Grid) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if values[i][j] == 0 {
				for num := layout[i][j]; num <= 9; num += 2 {
					if isValidClassic(values, i, j, num) {
						values[i][j] = num

						if solveOddEven(values, layout) {
							return true
						} else {
							values[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}

// figured
func solveFigured(values *Grid, layout *Grid) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if values[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if isValidFigured(values, layout, i, j, num) {
						values[i][j] = num

						if solveFigured(values, layout) {
							return true
						} else {
							values[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}
