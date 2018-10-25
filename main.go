package main

import (
	"encoding/json"
	"net/http"
	"os"
	"log"
	"bytes"
)

type Record struct {
	id       int
	title    string
	subtitle string
	cost     int
	compare  int
	rate     int
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		return
	}

	http.HandleFunc("/", blankScreen)
	http.HandleFunc("/getJson", jsonWriter)
	http.ListenAndServe(":" + port, nil)
}

func jsonWriter(w http.ResponseWriter, r *http.Request) {
	s0 := []Record{Record{1, "本日（現時点まで）", "", 9, 8, 7}}
	s1 := append(s0, Record{2, "昨日", "先週の同じ曜日との比較", 344, 43243, 43})
	s2 := append(s1, Record{3, "今月（現時点まで）", "先週の同じ曜日との比較", 1, 0, 0})
	s3 := append(s2, Record{4, "先月", "先々月との比較", 93, 83, 72})
	record := append(s3, Record{5, "全期間", "", 20, 4, 434})

	js, err := json.Marshal(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bufBody := new(bytes.Buffer)
	bufBody.ReadFrom(r.Body)
	body := bufBody.String()

	log.Printf(body)
	log.Printf(string(js))

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func blankScreen(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "hello.html")
}
