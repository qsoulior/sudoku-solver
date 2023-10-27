package solver

type Grid [9][9]int

type Solver struct {
	grid        *Grid
	constraints []Constraint
}

func New(grid *Grid) *Solver {
	return &Solver{
		grid:        grid,
		constraints: make([]Constraint, 0),
	}
}

func (s *Solver) AddConstraint(c Constraint) {
	s.constraints = append(s.constraints, c)
}

func (s *Solver) Valid(row int, col int, num int) bool {
	valid := true
	for _, c := range s.constraints {
		valid = valid && c.Valid(s.grid, row, col, num)
	}
	return valid
}

func (s *Solver) Solve() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.grid[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if s.Valid(i, j, num) {
						s.grid[i][j] = num

						if s.Solve() {
							return true
						}

						s.grid[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}
