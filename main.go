package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/qsoulior/sudoku-solver/solver"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var body struct {
			Mode   int
			Values solver.Grid
			Layout solver.Grid
		}
		d := json.NewDecoder(r.Body)
		err := d.Decode(&body)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		s := solver.New(&body.Values)

		var solved bool
		switch body.Mode {
		case 1:
			s.AddConstraint(&solver.RowConstraint{})
			s.AddConstraint(&solver.ColumnConstraint{})
			s.AddConstraint(&solver.SquareConstraint{})
			// solved = solveClassic(&body.Values)
		case 2:
			s.AddConstraint(&solver.RowConstraint{})
			s.AddConstraint(&solver.ColumnConstraint{})
			s.AddConstraint(&solver.SquareConstraint{})
			s.AddConstraint(&solver.PrimaryConstraint{})
			s.AddConstraint(&solver.SecondaryConstraint{})
			// solved = solveX(&body.Values)
		case 3:
			s.AddConstraint(&solver.RowConstraint{})
			s.AddConstraint(&solver.ColumnConstraint{})
			s.AddConstraint(&solver.SquareConstraint{})
			s.AddConstraint(&solver.OddEvenConstraint{&body.Layout})
			// solved = solveOddEven(&body.Values, &body.Layout)
		case 4:
			s.AddConstraint(&solver.RowConstraint{})
			s.AddConstraint(&solver.ColumnConstraint{})
			s.AddConstraint(&solver.FigureConstraint{&body.Layout})
			// solved = solveFigured(&body.Values, &body.Layout)
		}
		solved = s.Solve()

		if solved {
			w.Header().Set("Content-Type", "application/json")
			e := json.NewEncoder(w)
			e.Encode(body.Values)
			return
		}
		w.WriteHeader(400)
		return
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "failed to load html template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleIndex)

	http.ListenAndServe("localhost:3000", mux)
}
