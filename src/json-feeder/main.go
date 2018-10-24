package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Record struct {
	Id       int
	Title    string
	Subtitle string
	Cost     int
	Compare  int
	Rate     int
}

func main() {
	port := os.GetEnv("PORT")

	if port == "" {
		return
	}

	http.HandleFunc("/", jsonWriter)
	http.ListenAndServe(":" + port, nil)
}

func jsonWriter(w http.ResponseWriter, r *http.Request) {
	record := Record{1, "本日（現時点まで）", "先週の同じ曜日との比較", 9, 8, 7}

	js, err := json.Marshal(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
