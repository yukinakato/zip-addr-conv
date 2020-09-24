package main

import (
	"reflect"
	"testing"
)

func TestCreateIndex(t *testing.T) {
	var records rcd
	records.readCSV("testdata/test.csv")
	records.createIndex()

	expected := []struct {
		prefcode int
		start    int
		end      int
	}{
		{1, 0, 3},
		{28, 3, 6},
		{47, 6, 9},
	}

	for _, ex := range expected {
		as := records.index[ex.prefcode]["start"]
		if as != ex.start {
			t.Errorf("at code %d: expected start: %d actual: %d", ex.prefcode, ex.start, as)
		}
		ae := records.index[ex.prefcode]["end"]
		if ae != ex.end {
			t.Errorf("at code %d: expected end: %d actual: %d", ex.prefcode, ex.end, ae)
		}
	}
}

func TestAndSearch(t *testing.T) {
	var records rcd
	records.readCSV("testdata/test.csv")
	records.createIndex()

	cases := []struct {
		keywordArray []string
		expected     []jsonAddress
	}{
		{
			[]string{"2828282"},
			[]jsonAddress{
				{"2828282", "ひょうごけんはひふへほ", "兵庫県はひふへほ"},
			},
		},
		{
			[]string{"2"},
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
			[]string{"ひょ", "1"},
			[]jsonAddress{
				{"2828281", "ひょうごけんなにぬねの", "兵庫県なにぬねの"},
			},
		},
		{
			[]string{"な", "1"},
			[]jsonAddress{
				{"2828281", "ひょうごけんなにぬねの", "兵庫県なにぬねの"},
				{"4747471", "おきなわけん012345", "沖縄県01234"},
			},
		},
	}

	for _, tt := range cases {
		actual := records.andSearch(tt.keywordArray, 0)
		if !reflect.DeepEqual(tt.expected, actual) {
			t.Errorf("expected: %s actual: %s", tt.expected, actual)
		}
	}
}
