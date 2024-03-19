package solver

import (
	"slices"
)

type Constraint interface {
	Valid(values *Grid, row int, col int, num uint) bool
}

type RowConstraint struct {
}

func (c RowConstraint) Valid(values *Grid, row int, col int, num uint) bool {
	for i := 0; i < 9; i++ {
		if values[row][i] == num {
			return false
		}
	}
	return true
}

type ColumnConstraint struct {
}

func (c ColumnConstraint) Valid(values *Grid, row int, col int, num uint) bool {
	for i := 0; i < 9; i++ {
		if values[i][col] == num {
			return false
		}
	}
	return true
}

type SquareConstraint struct {
}

func (c SquareConstraint) Valid(values *Grid, row int, col int, num uint) bool {
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if values[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

type PrimaryConstraint struct {
}

func (c PrimaryConstraint) Valid(values *Grid, row int, col int, num uint) bool {
	if row == col {
		for i := 0; i < 9; i++ {
			if values[i][i] == num {
				return false
			}
		}
	}
	return true
}

type SecondaryConstraint struct {
}

func (c SecondaryConstraint) Valid(values *Grid, row int, col int, num uint) bool {
	if row == 8-col {
		for i := 0; i < 9; i++ {
			if values[i][8-i] == num {
				return false
			}
		}
	}
	return true
}

type ShapeConstraint struct {
	Layout *Grid
}

func (c ShapeConstraint) Valid(values *Grid, row int, col int, num uint) bool {
	color := c.Layout[row][col]
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if color == c.Layout[i][j] && values[i][j] == num {
				return false
			}
		}
	}
	return true
}

type OddEvenConstraint struct {
	Layout *Grid
}

func (c OddEvenConstraint) Valid(values *Grid, row int, col int, num uint) bool {
	if c.Layout[row][col] == 2 {
		return num%2 == 0
	}
	return num%2 != 0
}

type AsterixConstraint struct {
}

var indexes = [][2]int{{4, 1}, {2, 2}, {6, 2}, {1, 4}, {4, 4}, {7, 4}, {2, 6}, {6, 6}, {4, 7}}

func (c AsterixConstraint) Valid(values *Grid, row int, col int, num uint) bool {
	nums := make([]uint, 9)
	for i := 0; i < 9; i++ {
		index := indexes[i]
		nums[i] = values[index[0]][index[1]]
	}

	if slices.Contains(indexes, [2]int{row, col}) {
		return !slices.Contains(nums, num)
	}

	return true
}

type WindowConstraint struct {
}

var s1 = [][2]int{{1, 1}, {1, 2}, {1, 3}, {2, 1}, {2, 2}, {2, 3}, {3, 1}, {3, 2}, {3, 3}}
var s2 = [][2]int{{5, 1}, {5, 2}, {5, 3}, {6, 1}, {6, 2}, {6, 3}, {7, 1}, {7, 2}, {7, 3}}
var s3 = [][2]int{{1, 5}, {1, 6}, {1, 7}, {2, 5}, {2, 6}, {2, 7}, {3, 5}, {3, 6}, {3, 7}}
var s4 = [][2]int{{5, 5}, {5, 6}, {5, 7}, {6, 5}, {6, 6}, {6, 7}, {7, 5}, {7, 6}, {7, 7}}

func (w WindowConstraint) Valid(values *Grid, row int, col int, num uint) bool {
	nums1 := make([]uint, 9)
	nums2 := make([]uint, 9)
	nums3 := make([]uint, 9)
	nums4 := make([]uint, 9)

	for i := 0; i < 9; i++ {
		index := s1[i]
		nums1[i] = values[index[0]][index[1]]

		index = s2[i]
		nums2[i] = values[index[0]][index[1]]

		index = s3[i]
		nums3[i] = values[index[0]][index[1]]

		index = s4[i]
		nums4[i] = values[index[0]][index[1]]
	}

	rc := [2]int{row, col}

	if slices.Contains(s1, rc) {
		return !slices.Contains(nums1, num)
	}

	if slices.Contains(s2, rc) {
		return !slices.Contains(nums2, num)
	}

	if slices.Contains(s3, rc) {
		return !slices.Contains(nums3, num)
	}

	if slices.Contains(s4, rc) {
		return !slices.Contains(nums4, num)
	}

	return true
}
