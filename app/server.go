package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type jsonError struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

type jsonStat struct {
	Since string   `json:"since"`
	Total int      `json:"total"`
	Time  [101]int `json:"time"`
}

type stat struct {
	runningSince time.Time
	searchTotal  int
	searchTime   [101]int
	mutex        sync.Mutex
}

type statHandler struct {
	stat *stat
}

type searchHandler struct {
	records *rcd
	stat    *stat
}

func (h *statHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	j := h.stat.buildStatJSON()
	json.NewEncoder(w).Encode(j)
}

func (stats *stat) recordSearchTime(t time.Duration) {
	stats.mutex.Lock()
	ms := t.Milliseconds()
	log.Println("Search took", ms, "ms")
	ind := ms / 10
	if ind < 100 {
		stats.searchTime[ind]++
	} else {
		stats.searchTime[100]++
	}
	stats.searchTotal++
	stats.mutex.Unlock()
}

func (stats *stat) buildStatJSON() jsonStat {
	stats.mutex.Lock()
	j := jsonStat{
		stats.runningSince.UTC().Format("2006/01/02 15:04 MST"),
		stats.searchTotal,
		stats.searchTime,
	}
	stats.mutex.Unlock()
	return j
}

func (h *searchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	keywordArray := buildKeywordArray(r.FormValue("keyword"))
	if len(keywordArray) == 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(400)
		j := jsonError{}
		j.Error.Message = "keyword missing"
		json.NewEncoder(w).Encode(j)
		return
	}

	pref, err := strconv.Atoi(r.FormValue("pref"))
	if err != nil {
		pref = 0
	}
	if pref < 0 || pref > 47 {
		pref = 0
	}

	result := h.records.andSearch(keywordArray, pref)

	h.stat.recordSearchTime(time.Since(t))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)
}

func startServer() {
	var records rcd
	var stats stat

	records.readCSV("all.csv")

	searchHandler := &searchHandler{&records, &stats}

	statHandler := &statHandler{&stats}
	statHandler.stat.runningSince = time.Now()

	http.Handle("/search", searchHandler)
	http.Handle("/stat", statHandler)

	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":5555", nil))
}
