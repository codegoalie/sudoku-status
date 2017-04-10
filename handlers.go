package main

import (
	"encoding/json"
	"net/http"
)

type Stats struct {
	Count int `json:"count"`
}

func StatsIndex(repo repo, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	count := repo.GetPuzzleCount()
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Stats{Count: count}); err != nil {
		panic(err)
	}
}
