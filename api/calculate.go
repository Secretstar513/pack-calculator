package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Secretstar513/pack-calculator/internal/calc"
)

type req struct{ Items int `json:"items"` }
type resp struct{ Result map[int]int `json:"result"` }

func Handler(w http.ResponseWriter, r *http.Request) {
	var q req
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "bad json", 400); return
	}
	out, err := calc.Calculate(q.Items, packSizes())
	if err != nil { http.Error(w, err.Error(), 400); return }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{Result: out})
}

func packSizes() []int {
	if env := os.Getenv("PACK_SIZES"); env != "" {
		parts := strings.Split(env, ",")
		var xs []int
		for _, p := range parts {
			if n, _ := strconv.Atoi(strings.TrimSpace(p)); n > 0 {
				xs = append(xs, n)
			}
		}
		return xs
	}
	return []int{250, 500, 1000, 2000, 5000}
}
