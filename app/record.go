package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"sync"
)

const resultsMaxCount = 50

type jsonAddress struct {
	Zipcode      string `json:"zipcode"`
	AddressKana  string `json:"kana"`
	AddressKanji string `json:"kanji"`
}

type rcd struct {
	data  [][]string
	index map[int]map[string]int
	mutex sync.Mutex
}

func (records *rcd) readCSV(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	records.data, err = reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	records.createIndex()
}

func (records *rcd) createIndex() {
	prefCode := map[string]int{
		"北海道":  1,
		"青森県":  2,
		"岩手県":  3,
		"宮城県":  4,
		"秋田県":  5,
		"山形県":  6,
		"福島県":  7,
		"茨城県":  8,
		"栃木県":  9,
		"群馬県":  10,
		"埼玉県":  11,
		"千葉県":  12,
		"東京都":  13,
		"神奈川県": 14,
		"新潟県":  15,
		"富山県":  16,
		"石川県":  17,
		"福井県":  18,
		"山梨県":  19,
		"長野県":  20,
		"岐阜県":  21,
		"静岡県":  22,
		"愛知県":  23,
		"三重県":  24,
		"滋賀県":  25,
		"京都府":  26,
		"大阪府":  27,
		"兵庫県":  28,
		"奈良県":  29,
		"和歌山県": 30,
		"鳥取県":  31,
		"島根県":  32,
		"岡山県":  33,
		"広島県":  34,
		"山口県":  35,
		"徳島県":  36,
		"香川県":  37,
		"愛媛県":  38,
		"高知県":  39,
		"福岡県":  40,
		"佐賀県":  41,
		"長崎県":  42,
		"熊本県":  43,
		"大分県":  44,
		"宮崎県":  45,
		"鹿児島県": 46,
		"沖縄県":  47,
	}
	records.index = make(map[int]map[string]int)
	var start, end int
	for pref, code := range prefCode {
		start, end = -1, -1
		for i, r := range records.data {
			if strings.HasPrefix(r[2], pref) {
				if start == -1 {
					start = i
				}
				if end < i {
					end = i
				}
			} else {
				if end != -1 {
					break
				}
			}
		}
		if start == -1 || end == -1 {
			records.index[code] = map[string]int{"start": 0, "end": 0}
		} else {
			records.index[code] = map[string]int{"start": start, "end": end + 1}
		}
	}
	// validation
	cum := 0
	for _, v := range records.index {
		cum += v["end"] - v["start"]
	}
	if cum != len(records.data) {
		log.Println("num records doesn't match")
		log.Fatalln("cum:", cum, "all", len(records.data))
	}
}

func (records *rcd) andSearch(keywordArray []string, pref int) []jsonAddress {
	records.mutex.Lock()
	if len(keywordArray) == 0 {
		return []jsonAddress{}
	}
	var matched bool
	var searchRange [][]string
	var matchedIds []int
	if pref > 0 && pref < 48 {
		searchRange = records.data[records.index[pref]["start"]:records.index[pref]["end"]]
	} else {
		searchRange = records.data[:]
	}
L_ALL:
	for id, record := range searchRange {
		for _, kwd := range keywordArray {
			matched = false
			for _, elm := range record {
				matched = matched || strings.Contains(elm, kwd)
			}
			if !matched {
				break
			}
		}
		if matched {
			matchedIds = append(matchedIds, id)
			if len(matchedIds) == resultsMaxCount {
				break L_ALL
			}
		}
	}
	result := []jsonAddress{}
	for _, i := range matchedIds {
		result = append(result,
			jsonAddress{
				searchRange[i][0],
				searchRange[i][1],
				searchRange[i][2],
			})
	}
	records.mutex.Unlock()
	return result
}
