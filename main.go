package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/qsoulior/sudoku-solver/solver"
)

type Mode uint

const (
	Classic Mode = iota + 1
	X
	OddEven
	Jigsaw
	Asterix
	Window
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var body struct {
			Mode   Mode
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
		switch body.Mode {
		case Classic:
			s.AddConstraint(solver.RowConstraint{})
			s.AddConstraint(solver.ColumnConstraint{})
			s.AddConstraint(solver.SquareConstraint{})
		case X:
			s.AddConstraint(solver.RowConstraint{})
			s.AddConstraint(solver.ColumnConstraint{})
			s.AddConstraint(solver.SquareConstraint{})
			s.AddConstraint(solver.PrimaryConstraint{})
			s.AddConstraint(solver.SecondaryConstraint{})
		case OddEven:
			s.AddConstraint(solver.RowConstraint{})
			s.AddConstraint(solver.ColumnConstraint{})
			s.AddConstraint(solver.SquareConstraint{})
			s.AddConstraint(solver.OddEvenConstraint{&body.Layout})
		case Jigsaw:
			s.AddConstraint(solver.RowConstraint{})
			s.AddConstraint(solver.ColumnConstraint{})
			s.AddConstraint(solver.ShapeConstraint{&body.Layout})
		case Asterix:
			s.AddConstraint(solver.RowConstraint{})
			s.AddConstraint(solver.ColumnConstraint{})
			s.AddConstraint(solver.SquareConstraint{})
			s.AddConstraint(solver.AsterixConstraint{})
		case Window:
			s.AddConstraint(solver.RowConstraint{})
			s.AddConstraint(solver.ColumnConstraint{})
			s.AddConstraint(solver.SquareConstraint{})
			s.AddConstraint(solver.WindowConstraint{})
		}

		if s.Solve() {
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
