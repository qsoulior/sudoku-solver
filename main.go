package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		grid := new(SudokuGrid)
		d := json.NewDecoder(r.Body)
		d.Decode(grid)

		solved := solve(grid)
		if solved {
			w.Header().Set("Content-Type", "application/json")
			e := json.NewEncoder(w)
			e.Encode(*grid)
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
