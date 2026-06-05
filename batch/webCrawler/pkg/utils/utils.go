package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"golang.org/x/text/width"
	"gopkg.in/natefinch/lumberjack.v2"
)


var _regPriceSpliter   = regexp.MustCompile(`[\s(（/／:、]`)
var _regUndefinedPrice = regexp.MustCompile(`(?i)(ask|open)`)
const _initPrice int   = 999999999
// 金額表記を数値に変換 "¥128,000" → 128000
func ParsePrice(price string) (int, error) {
	if _regUndefinedPrice.MatchString(price) {
		return -1, nil
	}
	var result int = 0
	var err error

	if _regPriceSpliter.MatchString(price) {
		result, err = parseMultiPrice(price)
	} else {
		result, err = parseSinglePrice(price)
	}
	if err != nil { return -1, err }
	if result == _initPrice { result = -1 }
	return result, nil
}

var _regNotNumber = regexp.MustCompile(`\D`)
// 金額（数値）へ変換
func parseSinglePrice(s string) (int, error) {
	price := width.Narrow.String(s)

	cleaned := _regNotNumber.ReplaceAllString(price, "")

	if cleaned == "" {
		return -1, fmt.Errorf("parseSinglePrice: 金額のパースに失敗しました。>>> %v\n", s)
	}
	result, _ := strconv.Atoi(cleaned)
	return result, nil
}

// 税込み金額（数値）へ変換
// 250,000円（税別） ／ 275,000円（税込）＞＞ 275000
func parseMultiPrice(s string) (int, error) {
	// 全分割
	prices := _regPriceSpliter.Split(s, -1)

	var errMessage string =
		"parseMultiPrice: 金額のパースに失敗しました。>>> %v\n"
	if len(prices) < 2 {
		return -1, fmt.Errorf(errMessage, s)
	}

	var minPrice int = _initPrice

	for _, price := range prices {
		p 		   := width.Narrow.String(price)
		trimed     := _regNotNumber.ReplaceAllString(p, "")

		if trimed == "" {
			continue
		}
		price, _ := strconv.Atoi(trimed)

		if price < minPrice {
			minPrice = price
		}
	}
	return minPrice, nil
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

// フレット数は１５～３０であると思われる
var _regFretStr   = regexp.MustCompile(`(1[5-9]|[2-3][0-9])\s*[Ff]+`)
var _refNotNumber = regexp.MustCompile(`\D`)
// ギターのフレット数を取得
func GetFretCount(s string) (int, error) {
	fretStr := _regFretStr.FindString(s)
	fret    := _refNotNumber.ReplaceAllString(fretStr, "")

	if len(fret) <= 0 {
		log.Println(s)
		return -1, fmt.Errorf("フレット数を取得できませんでした。>>> %v\n", s)
	}
	result, _ := strconv.Atoi(fret)
	return result, nil
}

var regScale = regexp.MustCompile(`\s*mm`)
// ギタースケールの単位を除去
func TrimScaleUnit(s string) int {
	halfed := width.Narrow.String(s)

	scale 	  := regScale.ReplaceAllString(halfed, "")
	result, _ := strconv.Atoi(scale)

	if result == 0 { return -1 }

	return result
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

// スレッドセーフなappend 注：mutexに直接&sync.Mutex{}を渡すのは禁止
func LockedAppend[T any](mutex *sync.Mutex, slice []T, elem ...T) []T {
	mutex.Lock()
	defer mutex.Unlock()
	slice = append(slice, elem...)

	return slice
}