package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestRecordSearchTime(t *testing.T) {
	var stats stat

	stats.recordSearchTime(1_000_000)
	if stats.searchTotal != 1 {
		t.Errorf("search total expected: 1, actual: %d", stats.searchTotal)
	}
	if stats.searchTime[0] != 1 {
		t.Errorf("searchTime[0-10 ms] expected: 1, actual: %d", stats.searchTime[0])
	}

	stats.recordSearchTime(25_000_000)
	if stats.searchTotal != 2 {
		t.Errorf("search total expected: 2, actual: %d", stats.searchTotal)
	}
	if stats.searchTime[2] != 1 {
		t.Errorf("searchTime[20-30 ms] expected: 1, actual: %d", stats.searchTime[2])
	}

	stats.recordSearchTime(2500_000_000)
	if stats.searchTotal != 3 {
		t.Errorf("search total expected: 3, actual: %d", stats.searchTotal)
	}
	if stats.searchTime[100] != 1 {
		t.Errorf("searchTime[over 1000 ms] expected: 1, actual: %d", stats.searchTime[100])
	}
}

func TestBuildStatJSON(t *testing.T) {
	var stats stat
	stats.runningSince = time.Date(2020, 8, 1, 15, 14, 13, 0, time.UTC)
	stats.searchTotal = 1
	stats.searchTime = [101]int{1}

	expected := jsonStat{"2020/08/01 15:14 UTC", 1, stats.searchTime}
	actual := stats.buildStatJSON()
	if expected != actual {
		t.Errorf("jsonStat differs expected: %v, actual %v", expected, actual)
	}
}

func TestSearch(t *testing.T) {
	var records rcd
	var stats stat
	var ja []jsonAddress
	var je jsonError
	records.readCSV("testdata/test.csv")
	h := searchHandler{&records, &stats}

	jerr := jsonError{}
	jerr.Error.Message = "keyword missing"

	cases := []struct {
		url      string
		status   int
		expected interface{}
	}{
		{
			"/search?keyword=4747472",
			200,
			[]jsonAddress{
				{"4747472", "おきなわけん56789", "沖縄県56789"},
			},
		},
		{
			"/search?keyword=" + url.PathEscape("き") + "&pref=1",
			200,
			[]jsonAddress{
				{"0101011", "ほっかいどうかきくけこ", "北海道かきくけこ"},
			},
		},
		{
			"/search?keyword=2&pref=100",
			200,
			[]jsonAddress{
				{"0101012", "ほっかいどうさしすせそ", "北海道さしすせそ"},
				{"2828280", "ひょうごけんたちつてと", "兵庫県たちつてと"},
				{"2828281", "ひょうごけんなにぬねの", "兵庫県なにぬねの"},
				{"2828282", "ひょうごけんはひふへほ", "兵庫県はひふへほ"},
				{"4747471", "おきなわけん012345", "沖縄県01234"},
				{"4747472", "おきなわけん56789", "沖縄県56789"},
			},
		},
		{
			"/search?keyword=" + url.PathEscape("ひょ　１-"),
			200,
			[]jsonAddress{
				{"2828281", "ひょうごけんなにぬねの", "兵庫県なにぬねの"},
			},
		},
		{
			"/search?keyword=" + url.PathEscape("   な　１　"),
			200,
			[]jsonAddress{
				{"2828281", "ひょうごけんなにぬねの", "兵庫県なにぬねの"},
				{"4747471", "おきなわけん012345", "沖縄県01234"},
			},
		},
		{
			"/search?",
			400,
			jerr,
		},
		{
			"/search?keyword=",
			400,
			jerr,
		},
		{
			"/search?keyword=%20",
			400,
			jerr,
		},
	}

	for _, tt := range cases {
		req := httptest.NewRequest("GET", tt.url, nil)
		res := httptest.NewRecorder()
		h.ServeHTTP(res, req)

		if res.Code != tt.status {
			t.Errorf("status code: expected %d, actual %d", tt.status, res.Code)
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}
		switch tt.expected.(type) {
		case []jsonAddress:
			err = json.Unmarshal(b, &ja)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(tt.expected, ja) {
				t.Errorf("expected: %s actual: %s", tt.expected, ja)
			}
		case jsonError:
			err = json.Unmarshal(b, &je)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(tt.expected, je) {
				t.Errorf("expected: %s actual: %s", tt.expected, je)
			}
		default:
			t.Errorf("no type matched")
		}
	}
}
