package utils

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	"github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"golang.org/x/text/width"
	"gopkg.in/natefinch/lumberjack.v2"
)


var _regPriceSpliter   = regexp.MustCompile(`[()（）/／:、]`)
var _regUndefinedPrice = regexp.MustCompile(`(?i)(ask|open)`)
const _initPrice int   = 999999999
// 金額表記を数値に変換 "¥128,000" → 128000
func ParsePrice(price string) (int, error) {
	if _regUndefinedPrice.MatchString(price) {
		return -1, nil
	}
	var result int = 0
	var err error

	price = strings.ReplaceAll(price, " ", "")

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
		return -1, fmt.Errorf("[ParseSinglePrice parse error]: %v\n", s)
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
		"[ParseMultiPrice parse error]: %v\n"
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
	try, err :=strconv.Atoi(s)
	if err == nil { return try, nil }

	fretStr := _regFretStr.FindString(s)
	fret    := _refNotNumber.ReplaceAllString(fretStr, "")

	if len(fret) <= 0 {
		log.Println(s)
		return -1, fmt.Errorf("[Get fretOfNumber error]: %v\n", s)
	}
	result, _ := strconv.Atoi(fret)
	return result, nil
}

var regScale = regexp.MustCompile(`(\..*|\s)*mm`)
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
		MaxAge:     10,  // 30日で削除
		Compress:   true,
	})
}

// 取得したログを表示
func LogCollectedLinks(links []string) {
	for _, link := range links {
		log.Printf("[Collected link]: %v\n", link)
	}
}

// スレッドセーフなappend 注：mutexに直接&sync.Mutex{}を渡すのは禁止
func LockedAppend[T any](mutex *sync.Mutex, slice []T, elem ...T) []T {
	mutex.Lock()
	defer mutex.Unlock()
	slice = append(slice, elem...)

	return slice
}

// color > colorCd 変換用
var colorRegexMap = map[int]*regexp.Regexp{
    constants.Red:     regexp.MustCompile(`(?i)\bRed\b`),
    constants.Pink:    regexp.MustCompile(`(?i)\bPink\b`),
    constants.Orange:  regexp.MustCompile(`(?i)\bOrange\b`),
    constants.Yellow:  regexp.MustCompile(`(?i)\bYellow\b`),
    constants.Green:   regexp.MustCompile(`(?i)\bGreen\b`),
    constants.SkyBlue: regexp.MustCompile(`(?i)\bSkyBlue\b`),
    constants.Blue:    regexp.MustCompile(`(?i)\bBlue\b`),
    constants.Purple:  regexp.MustCompile(`(?i)\bPurple\b`),
    constants.Gray:    regexp.MustCompile(`(?i)\bGray\b`),
    constants.Black:   regexp.MustCompile(`(?i)\bBlack\b`),
    constants.White:   regexp.MustCompile(`(?i)\bWhite\b`),
    constants.Natural: regexp.MustCompile(`(?i)\bNatural\b`),
    constants.Brown:   regexp.MustCompile(`(?i)\bBrown\b`),
    constants.Gold:    regexp.MustCompile(`(?i)\bGold\b`),
    constants.Silver:  regexp.MustCompile(`(?i)\bSilver\b`),
}
// カラー名を抽象化してカラーコードに変換する
func ConvertColorCd(colorName string) int {
	for cd, regex := range colorRegexMap {
		if regex.MatchString(colorName) {
			return cd
		}
	}
	return constants.OthersColor
}

// URLを使用できる形式に変換
func ConvertRealUrl(proxyUrl string) string {
	u, err 	     := url.Parse(proxyUrl)

	if err != nil || len(u.String()) == 0 {
		log.Printf("[URL parse error]: %v\n", proxyUrl)
		return proxyUrl
	}
	encodedUrl   := u.Query().Get("url")
	realUrl, err := url.QueryUnescape(encodedUrl)

	if err != nil || len(realUrl) == 0 {
		log.Printf("[URL convert error]: %v\n", proxyUrl)
		return proxyUrl
	}
	return realUrl
}

// 画像保存するためのパスを作成。dir名＋ファイル名
func CreateImageSavePath(saveDirName string, url string) string {
	hash 	 := sha1.Sum([]byte(url)) // urlをハッシュ化
	filename := fmt.Sprintf("%x.png", hash)
	savePath := filepath.Join(saveDirName, filename)
	return savePath
}

// 画像取得 savePathはファイル名まで含める
func DownloadImage(url, savePath string) {
    if err := os.MkdirAll(filepath.Dir(savePath), 0755); err != nil {
        log.Printf("[Make dir failed]: %v\n", err)
    }
    resp, err := http.Get(url)

	if resp == nil {
		log.Printf("[Response error]: %v\n", "res == nil")
		return
	}
    if err != nil {
        log.Printf("[Response error]: %v\n", err)
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)

	if err != nil {
        log.Printf("[Resource read error]: %v\n", err)
    }
	// atomicな保存
	tmp := savePath + ".tmp"
	os.WriteFile(tmp, data, 0644)  // 一時ファイルに書く
	os.Rename(tmp, savePath)
}

// まとめてリソースをDL。不要ならDLしない
func AutoDownLoader(guitars []*model.Guitar, saveDirName string) {
	var wg = &sync.WaitGroup{}
	queue := make(chan struct{}, 5) // 並列数制御

	for _, g := range guitars {
		g := g // 並列バグ対策
		queue <- struct{}{}
		wg.Add(1)

		go func(guitar *model.Guitar) {
			defer wg.Done()
			defer func() { <-queue }() // 次のワーカーへ

			if strings.HasPrefix(guitar.Src, "https://") || strings.HasPrefix(guitar.Src, "http://") {
				return
			}
			url := ConvertRealUrl(guitar.Src)
			savePath := CreateImageSavePath(saveDirName, url)

			if _, err := os.Stat(savePath); err == nil {
				return // キャッシュ、画像は保存済
			}
			DownloadImage(url, savePath) // atomicな保存
			guitar.Src = savePath
		}(g)
	}
	wg.Wait()
}

// 指定したラベルの要素を取得
func GetElem(doc *goquery.Document) func(selector string) string {
	return func(selector string) string {
		return strings.TrimSpace(doc.Find(selector).Text())
	}
}

// 指定したラベルの次（兄弟要素）の要素を取得
func GetElemNextToLabel(doc *goquery.Document) func(selector string) string {
	return func(selector string) string {
		return strings.TrimSpace(doc.Find(selector).Next().Text())
	}
}

var regWight = regexp.MustCompile(`\d\.\d{1,2}`)
// 重量を抽出する（Kg単位）
func ParseWight(weight string) float64 {
	w := width.Narrow.String(weight)
	w  = regWight.FindString(w)

	result, err := strconv.ParseFloat(w, 64)

	if err != nil {
		log.Printf("[Weight parse error] %v\n", weight)
		return -1
	}
	return result
}

// サイトの項目名をフィールド名に変換
func ConvertLabel(label string, fieldMap map[string]string) string {
    return fieldMap[label]
}