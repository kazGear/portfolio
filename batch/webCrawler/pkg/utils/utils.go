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
	var result int = _initPrice
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

	if len(prices) < 2 {
		return -1, fmt.Errorf("[ParseMultiPrice parse error]: %v\n", s)
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
	try, err := strconv.Atoi(s)
	if err == nil { return try, nil }

	fretStr := _regFretStr.FindString(s)
	fret    := _refNotNumber.ReplaceAllString(fretStr, "")

	if len(fret) <= 0 {
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

// ログ設定(グローバル設定)
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

// ログインスタンスを作成
func NewLogger(makerName string) *log.Logger {
    // 日付入りのログファイル名
    date := time.Now().Format("2006-01-02")
    filename := fmt.Sprintf("logs/%v_%v.log", makerName, date)

    // ディレクトリがなければ作成
    os.MkdirAll("logs", 0755)

    // メーカーごとにローテーション設定
    writer := &lumberjack.Logger{
        Filename:   filename,
        MaxSize:    5,   // 5MBでローテーション
        MaxBackups: 7,   // 最大7ファイル保持
        MaxAge:     10,  // 10日で削除
        Compress:   true,
    }
    return log.New(writer, "", log.LstdFlags)
}

// 取得したログを表示
func LogCollectedLinks(links []string, logger *log.Logger) {
	for _, link := range links {
		logger.Printf("[Collected link]: %v\n", link)
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
type colorKeyword struct {
    Cd       int
    Keywords []string
}

var colorKeywords = []colorKeyword{
    {
        Cd: constants.Red,
        Keywords: []string{
            "red", "cherry", "apple", "fiesta", "burgundy", "cranberry",
            "garnet", "cardinal", "tomato", "ember", "lava",
            "vermillion", "rose",
        },
    },
    {
        Cd: constants.Pink,
        Keywords: []string{
            "pink", "coral", "sakura", "rose", "twinkle",
        },
    },
    {
        Cd: constants.Orange,
        Keywords: []string{
            "orange", "sunset", "sunrise", "autumn", "coral",
            "tangerine", "poppy",
        },
    },
    {
        Cd: constants.Yellow,
        Keywords: []string{
            "yellow", "honey", "amber", "mustard", "lemon", "blond",
        },
    },
    {
        Cd: constants.Green,
        Keywords: []string{
            "green", "citron", "ivy", "forest", "olive", "mint",
            "snake", "iguana", "malachite",
        },
    },
    {
        Cd: constants.SkyBlue,
        Keywords: []string{
            "skyblue", "sky", "frost",
        },
    },
    {
        Cd: constants.Blue,
        Keywords: []string{
            "blue", "marine", "supreme", "nebula", "peacock", "mercury",
            "aqua", "turquoise", "azure", "navy", "bonnet",
        },
    },
    {
        Cd: constants.Purple,
        Keywords: []string{
            "purple", "indigo", "violet", "lavender", "plum", "amethyst", "sugilite", "tanzanite",
        },
    },
    {
        Cd: constants.Gray,
        Keywords: []string{
            "gray", "granite", "pewter", "slate", "ash", "graphite",
            "charcoal", "stone", "meteorite", "rusty", "iron",
        },
    },
    {
        Cd: constants.Black,
        Keywords: []string{
            "black", "obsidian", "onyx", "ebony", "jet", "pitch",
        },
    },
    {
        Cd: constants.White,
        Keywords: []string{
            "white", "snow", "ivory", "cream", "pearl", "fox",
        },
    },
    {
        Cd: constants.Brown,
        Keywords: []string{
            "brown", "walnut", "mahogany", "chocolate", "bourbon", "tobacco",
        },
    },
    {
        Cd: constants.Gold,
        Keywords: []string{
            "gold", "champagne", "brass",
        },
    },
    {
        Cd: constants.Silver,
        Keywords: []string{
            "silver", "chrome",
        },
    },
    {
        Cd: constants.Natural,
        Keywords: []string{
            "natural", "raw", "naked", "plain", "wood", "driftwood", "burnt",
        },
    },
}

// カラー名を抽象化してカラーコードに変換する
func ConvertColorCd(colorName string) int {
	lowerColorName := strings.ToLower(colorName)

	for _, colors := range colorKeywords {
		cd := colors.Cd
		for _, color := range colors.Keywords {
			if strings.Contains(lowerColorName, strings.ToLower(color)) {
				return cd
			}
		}
	}
	return constants.OthersColor
}

// URLを使用できる形式に変換
func ConvertRealUrl(proxyUrl string) (string, error) {
	u, err := url.Parse(proxyUrl)

	if err != nil || len(u.String()) == 0 {
		return proxyUrl, fmt.Errorf("[URL parse error]: %v %w\n", proxyUrl, err)
	}
	encodedUrl   := u.Query().Get("url")
	realUrl, err := url.QueryUnescape(encodedUrl)

	if err != nil || len(realUrl) == 0 {
		return proxyUrl, fmt.Errorf("[URL convert error]: %v %w\n", proxyUrl, err)
	}
	return realUrl, nil
}

// 画像保存するためのパスを作成。dir名＋ファイル名
func CreateImageSavePath(saveDirName string, url string) string {
	hash 	 := sha1.Sum([]byte(url)) // urlをハッシュ化
	filename := fmt.Sprintf("%x.png", hash)
	savePath := filepath.Join(saveDirName, filename)
	return savePath
}

// 画像取得 savePathはファイル名まで含める
func DownloadImage(url, savePath string) error {
    if err := os.MkdirAll(filepath.Dir(savePath), 0755); err != nil {
        return fmt.Errorf("[Make dir failed]: %w\n", err)
    }
    resp, err := http.Get(url)

	if resp == nil {
		return fmt.Errorf("[Response error]: %v\n", "res == nil")
	}
    if err != nil {
        return fmt.Errorf("[Response error]: %w\n", err)
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)

	if err != nil {
        return fmt.Errorf("[Resource read error]: %w\n", err)
    }
	// atomicな保存
	tmp := savePath + ".tmp"
	os.WriteFile(tmp, data, 0644)  // 一時ファイルに書く
	os.Rename(tmp, savePath)
	return nil
}

// まとめてリソースをDL。不要ならDLしない
func AutoDownLoader(guitars []*model.Guitar, saveDirName string) []error {
 	var wg 	  = &sync.WaitGroup{}
	var mutex = &sync.Mutex{}
	queue 	 := make(chan struct{}, 7) // 並列数制御
	errs  	 := make([]error, 0, 300)

	for _, g := range guitars {
		g := g // 並列バグ対策
		queue <- struct{}{}
		wg.Add(1)

		go func(guitar *model.Guitar) {
			defer wg.Done()
			defer func() { <-queue }() // 次のワーカーへ

			if strings.HasPrefix(guitar.Src, "http") {
				return
			}
			url, err := ConvertRealUrl(guitar.Src)

			if !strings.HasPrefix(guitar.Src, "http") {
				return // url convert failed
			}
			if err != nil { LockedAppend(mutex, errs, err) }

			savePath := CreateImageSavePath(saveDirName, url)

			if _, err := os.Stat(savePath); err == nil {
				return // キャッシュ、画像は保存済
			}
			DownloadImage(url, savePath) // atomicな保存
			guitar.Src = savePath
		}(g)
	}
	wg.Wait()

	if len(errs) == 0 {
		return nil
	} else {
		return errs
	}
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
func ParseWight(weight string) (float64, error) {
	w := width.Narrow.String(weight)
	w  = regWight.FindString(w)

	result, err := strconv.ParseFloat(w, 64)

	if err != nil {
		return -1, fmt.Errorf("[Weight parse error] %v %w\n", weight, err)
	}
	return result, nil
}

// サイトの項目名をフィールド名に変換
func ConvertLabel(label string, fieldMap map[string]string) string {
    return fieldMap[label]
}