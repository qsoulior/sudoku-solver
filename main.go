package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var body struct {
			Mode   int
			Values Grid
			Layout Grid
		}
		d := json.NewDecoder(r.Body)
		err := d.Decode(&body)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		var solved bool
		switch body.Mode {
		case 1:
			solved = solveClassic(&body.Values)
		case 2:
			solved = solveX(&body.Values)
		case 3:
			solved = solveOddEven(&body.Values, &body.Layout)
		case 4:
			solved = solveFigured(&body.Values, &body.Layout)
		}

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
