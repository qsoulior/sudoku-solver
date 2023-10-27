package solver

type Constraint interface {
	Valid(values *Grid, row int, col int, num int) bool
}

type RowConstraint struct {
}

func (c *RowConstraint) Valid(values *Grid, row int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if values[row][i] == num {
			return false
		}
	}
	return true
}

type ColumnConstraint struct {
}

func (c *ColumnConstraint) Valid(values *Grid, row int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if values[i][col] == num {
			return false
		}
	}
	return true
}

type SquareConstraint struct {
}

func (c *SquareConstraint) Valid(values *Grid, row int, col int, num int) bool {
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

func (c *PrimaryConstraint) Valid(values *Grid, row int, col int, num int) bool {
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

func (c *SecondaryConstraint) Valid(values *Grid, row int, col int, num int) bool {
	if row == 8-col {
		for i := 0; i < 9; i++ {
			if values[i][8-i] == num {
				return false
			}
		}
	}
	return true
}

type FigureConstraint struct {
	Layout *Grid
}

func (c *FigureConstraint) Valid(values *Grid, row int, col int, num int) bool {
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

func (c *OddEvenConstraint) Valid(values *Grid, row int, col int, num int) bool {
	if c.Layout[row][col] == 2 {
		return num%2 == 0
	}
	return num%2 != 0
}
