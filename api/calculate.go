package handler                                    // ✅ required by Vercel

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Secretstar513/pack-calculator/internal/calc"
)

type req struct{ Items int `json:"items"` }
type resp struct{ Result map[int]int `json:"result"` }

func Handler(w http.ResponseWriter, r *http.Request) { // ✅ exported, exact name
	var q req
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	out, err := calc.Calculate(q.Items, packSizes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp{Result: out})
}

// optional: read PACK_SIZES env var or fall back to defaults
func packSizes() []int {
	if env := strings.TrimSpace(strings.Trim(os.Getenv("PACK_SIZES"), "[]")) ; env != "" {
		var xs []int
		for _, s := range strings.Split(env, ",") {
			if n, _ := strconv.Atoi(strings.TrimSpace(s)); n > 0 { xs = append(xs, n) }
		}
		if len(xs) > 0 { return xs }
	}
	return []int{250, 500, 1000, 2000, 5000}
}
