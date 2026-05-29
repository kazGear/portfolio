package utils

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// 金額表記を数値に変換 "¥128,000" → 128000
func ParsePrice(price string) (int, error) {
	var result int = 0
	var err error

	if strings.Contains(price, "／") {
		result, err = parseDoublePrice(price)
	} else {
		result, err = parseSinglePrice(price)
	}

	if (err != nil) { return -1, err }
	return result, nil
}

// 金額（数値）へ変換
func parseSinglePrice(s string) (int, error) {
	price := ToHankakuNumber(s)

	re := regexp.MustCompile(`\D`)
	cleaned := re.ReplaceAllString(price, "")

	if cleaned == "" {
		log.Fatal("parseSinglePrice: 金額のパースに失敗しました。")
	}
	result, err := strconv.Atoi(cleaned)

	if err != nil { return 0, err }
	return result, nil
}

// 税込み金額（数値）へ変換
// 250,000円（税別） ／ 275,000円（税込）＞＞ 275000
func parseDoublePrice(s string) (int, error) {
	prices := strings.Split(s, "／")

	var errMessage string =
		"parseDoublePrice: 金額のパースに失敗しました。"
	if len(prices) != 2 {
		log.Fatal(errMessage)
	}

	p1 := ToHankakuNumber(prices[0])
	p2 := ToHankakuNumber(prices[1])

	re := regexp.MustCompile(`\D`)
	cleaned1 := re.ReplaceAllString(p1, "")
	cleaned2 := re.ReplaceAllString(p2, "")

	if cleaned1 == "" || cleaned2 == "" {
		log.Fatal(errMessage)
	}
	price1, err1 := strconv.Atoi(cleaned1)
	price2, err2 := strconv.Atoi(cleaned2)
	if err1 != nil { return 0, err1 }
	if err2 != nil { return 0, err2 }

	var result int = 0
	// 税込み表示を選択
	if (price1 < price2) {
		result = price2
	} else {
		result = price1
	}
	return result, nil
}

// 変換 全角数字 >> 半角数字
func ToHankakuNumber(num string) string {
	hankakuNum := strings.Map(func(r rune) rune {
		if '０' <= r && r <= '９' {
			// 全角数字は Unicode 上で 半角数字 + 0xFEE0 の差
			return r - 0xFEE0
		}
		return r
	}, num)
	return hankakuNum
}