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
	port := os.Getenv("PORT")

	if port == "" {
		return
	}

	http.HandleFunc("/getJson", jsonWriter)
	http.ListenAndServe(":" + port, nil)
}

func jsonWriter(w http.ResponseWriter, r *http.Request) {
	record := []Record{}
	record.append(record, Record{1, "本日（現時点まで）", "", 9, 8, 7})
	record.append(record, Record{2, "昨日", "先週の同じ曜日との比較", 344, 43243, 43})
	record.append(record, Record{3, "今月（現時点まで）", "先週の同じ曜日との比較", 1, 0, 0})
	record.append(record, Record{4, "先月", "先々月との比較", 93, 83, 72})
	record.append(record, Record{5, "全期間", "", 20, 4, 434})

	js, err := json.Marshal(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
