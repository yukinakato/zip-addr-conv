package main

import (
	"regexp"
	"strings"
)

const allowInputCharsMaxCount = 100

func substr(s string, n int) string {
	r := []rune(s)
	if len(r) > n && n > 0 {
		return string(r[:n])
	}
	return s
}

func buildKeywordArray(raw string) []string {
	regexToIgnore := regexp.MustCompile(`[-ｰ−－]+`)
	regexToSplit := regexp.MustCompile(`[\s　]+`)
	replaceNumWideToNarrow := strings.NewReplacer("０", "0", "１", "1", "２", "2", "３", "3", "４", "4", "５", "5", "６", "6", "７", "7", "８", "8", "９", "9")
	var keyword string

	// 文字数制限をかける
	// 無視したい文字（ハイフン等）を消去する
	// 連続したスペースや区切り文字を単独スペースに置き換える
	// 数字のみ全角から半角に変換する
	// 両端のスペースを削除する
	keyword = substr(raw, allowInputCharsMaxCount)
	keyword = regexToIgnore.ReplaceAllString(keyword, "")
	keyword = regexToSplit.ReplaceAllString(keyword, " ")
	keyword = replaceNumWideToNarrow.Replace(keyword)
	keyword = strings.TrimSpace(keyword)

	if len(keyword) == 0 {
		return []string{}
	}
	return strings.Split(keyword, " ")
}
