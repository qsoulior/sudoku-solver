package solver

type Grid [9][9]int

type Solver struct {
	Grid        *Grid
	Constraints []Constraint
}

func New(grid *Grid) *Solver {
	return &Solver{
		Grid:        grid,
		Constraints: make([]Constraint, 0),
	}
}

func (s *Solver) AddConstraint(c Constraint) {
	s.Constraints = append(s.Constraints, c)
}

func (s *Solver) Valid(row int, col int, num int) bool {
	valid := true
	for _, c := range s.Constraints {
		valid = valid && c.Valid(s.Grid, row, col, num)
	}
	return valid
}

func (s *Solver) Solve() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.Grid[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if s.Valid(i, j, num) {
						s.Grid[i][j] = num

						if s.Solve() {
							return true
						} else {
							s.Grid[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}
