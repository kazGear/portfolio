package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"golang.org/x/text/width"
	"gopkg.in/natefinch/lumberjack.v2"
)

const splitPattern = `[\s(（/／,、]`

// 金額表記を数値に変換 "¥128,000" → 128000
func ParsePrice(price string) (int, error) {
	var result int = 0
	var err error

	if regexp.MustCompile(splitPattern).MatchString(price) {
		result, err = parseDoublePrice(price)
	} else {
		result, err = parseSinglePrice(price)
	}

	if (err != nil) { return -1, err }
	return result, nil
}

// 金額（数値）へ変換
func parseSinglePrice(s string) (int, error) {
	price := width.Narrow.String(s)

	re := regexp.MustCompile(`\D`)
	cleaned := re.ReplaceAllString(price, "")

	if cleaned == "" {
		return -1, fmt.Errorf("parseSinglePrice: 金額のパースに失敗しました。>>> %v\n", s)
	}
	result, err := strconv.Atoi(cleaned)

	if err != nil { return -1, err }
	return result, nil
}

// 税込み金額（数値）へ変換
// 250,000円（税別） ／ 275,000円（税込）＞＞ 275000
func parseDoublePrice(s string) (int, error) {
	prices := regexp.MustCompile(splitPattern).Split(s, 2)

	var errMessage string =
		"parseDoublePrice: 金額のパースに失敗しました。>>> %v\n"
	if len(prices) != 2 {
		return -1, fmt.Errorf(errMessage, s)
	}
	p1 := width.Narrow.String(prices[0])
	p2 := width.Narrow.String(prices[1])

	re := regexp.MustCompile(`\D`)
	cleaned1 := re.ReplaceAllString(p1, "")
	cleaned2 := re.ReplaceAllString(p2, "")

	if cleaned1 == "" || cleaned2 == "" {
		return -1, fmt.Errorf(errMessage, "price1: " + cleaned1 + ", price2: " + cleaned2 + " >>> " + s)
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

// 木材コードを探しだす
func SearchWoodCode(s string) int {
	for _, wood := range constants.GetWoods() {
		if strings.Contains(strings.ToLower(s), strings.ToLower(wood.Name)) {
			return wood.Code
		}
	}
	return 0 // 該当なし
}

// ギターのフレット数を取得
func GetFretCount(s string) (int, error) {
	// フレット数は１５～３０であると思われる
	reg  := regexp.MustCompile(`(1[5-9]|[2-3][0-9])`)
	fret := reg.FindString(s)

	if len(fret) <= 0 {
		log.Println(s)
		return -1, fmt.Errorf("フレット数を取得できませんでした。>>> %v\n", s)
	}
	result, err := strconv.Atoi(fret)

	if err != nil {
		return -1, err
	}
	return result, nil
}

// ギタースケールの単位を除去
func TrimScaleUnit(s string) (int, error) {
	halfed := width.Narrow.String(s)

	reg   := regexp.MustCompile(`\s*mm`)
	scale := reg.ReplaceAllString(halfed, "")

	result, err := strconv.Atoi(scale)

	if err != nil {
		return -1, nil
	}
	return result, nil
}

// ログ設定
func LoggerInit(maker string) {
	date 	 := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("logs/%v_%v.log", maker, date)

	os.MkdirAll("logs", 0755)

	// ローテーション設定
	log.SetOutput(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    5,   // 5MBでローテーション
		MaxBackups: 7,   // 最大7ファイル保持
		MaxAge:     30,  // 30日で削除
		Compress:   true,
	})
}