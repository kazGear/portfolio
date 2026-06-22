package utils

import (
	"crypto/sha1"
	"encoding/json"
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
	"github.com/chromedp/cdproto/cdp"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	C "github.com/kazGear/portfolio/webCrawler/pkg/constants"
	"github.com/shopspring/decimal"

	"golang.org/x/text/width"
	"gopkg.in/natefinch/lumberjack.v2"
)


var _regPriceSpliter   = regexp.MustCompile(`[()（）/／:、]`)
var _regUndefinedPrice = regexp.MustCompile(`(?i)(ask|open|オープン)`)
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
// 250,000円（税別） ／ 275,000円（税込）＞＞ 250000
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

var regWood = regexp.MustCompile(`\s+`)
// 木材コードを探しだす
func SearchWoodCode(s string) int {
	trimed := regWood.ReplaceAllString(s, "")

	for _, wood := range C.GetWoods() {
		if strings.Contains(strings.ToLower(trimed), strings.ToLower(wood.Name)) {
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

var regScale = regexp.MustCompile(`(\..*|\s)*(mm|”)`)
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
	date 	 := time.Now().Format(C.DateTime)
	filename := fmt.Sprintf("logs/%v_%v.log", maker, date)

	os.MkdirAll("logs", 0755)

	// ローテーション設定
	log.SetOutput(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    5,   // 5MBでローテーション
		MaxBackups: 7,   // 最大7ファイル保持
		MaxAge:     10,  // 30日で削除
		Compress:   false,
	})
}

// ログインスタンスを作成
func NewLogger(makerName string) *log.Logger {
    // 日付入りのログファイル名
    date := time.Now().Format(C.DateTime)
    filename := fmt.Sprintf("logs/%v_%v.log", makerName, date)

    // ディレクトリがなければ作成
    os.MkdirAll("logs", 0755)

    // メーカーごとにローテーション設定
    writer := &lumberjack.Logger{
        Filename:   filename,
        MaxSize:    5,   // 5MBでローテーション
        MaxBackups: 7,   // 最大7ファイル保持
        MaxAge:     10,  // 10日で削除
        Compress:   false,
    }
    return log.New(writer, "", log.LstdFlags)
}

// 取得したリンクを表示
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
        Cd: C.Red,
        Keywords: []string{
            "red", "cherry", "apple", "fiesta", "burgundy", "cranberry",
            "garnet", "cardinal", "tomato", "ember", "lava",
            "vermillion", "rose",
        },
    },
    {
        Cd: C.Pink,
        Keywords: []string{
            "pink", "coral", "sakura", "rose", "twinkle",
        },
    },
    {
        Cd: C.Orange,
        Keywords: []string{
            "orange", "sunset", "sunrise", "autumn", "coral",
            "tangerine", "poppy",
        },
    },
    {
        Cd: C.Yellow,
        Keywords: []string{
            "yellow", "honey", "amber", "mustard", "lemon", "blond",
        },
    },
    {
        Cd: C.Green,
        Keywords: []string{
            "green", "citron", "ivy", "forest", "olive", "mint",
            "snake", "iguana", "malachite",
        },
    },
    {
        Cd: C.SkyBlue,
        Keywords: []string{
            "skyblue", "sky", "frost",
        },
    },
    {
        Cd: C.Blue,
        Keywords: []string{
            "blue", "marine", "supreme", "nebula", "peacock", "mercury",
            "aqua", "turquoise", "azure", "navy", "bonnet",
        },
    },
    {
        Cd: C.Purple,
        Keywords: []string{
            "purple", "indigo", "violet", "lavender", "plum", "amethyst", "sugilite", "tanzanite",
        },
    },
    {
        Cd: C.Gray,
        Keywords: []string{
            "gray", "granite", "pewter", "slate", "ash", "graphite",
            "charcoal", "stone", "meteorite", "rusty", "iron",
        },
    },
	{
		Cd: C.Brown,
		Keywords: []string{
			"brown", "walnut", "mahogany", "chocolate", "bourbon", "tobacco",
		},
	},
	{
		Cd: C.Natural,
		Keywords: []string{
			"natural", "raw", "naked", "plain", "wood", "driftwood", "burnt",
		},
	},
	{
		Cd: C.Gold,
		Keywords: []string{
			"gold", "champagne", "brass",
		},
	},
	{
		Cd: C.Silver,
		Keywords: []string{
			"silver", "chrome",
		},
	},
	{
		Cd: C.Black,
		Keywords: []string{
			"black", "obsidian", "onyx", "ebony", "jet", "pitch",
		},
	},
	{
		Cd: C.White,
		Keywords: []string{
			"white", "snow", "ivory", "cream", "pearl", "fox",
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
	return C.OthersColor
}

// URLを使用できる形式に変換（next.jsの謎パス等
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

// 相対パスを絶対パスに変換
func CreateImagePath(absUrl string, src string) (string, error) {
	var err error
	base, err := url.Parse(absUrl)
	ref, err  := url.Parse(src)

	if err != nil {
		return "", fmt.Errorf(`[Url parse failed]: %v %w`, src, err)
	}
	result := base.ResolveReference(ref)

	return result.String(), nil
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

func TrimSpace() func(string) string {
	return func(s string) string {
		return strings.TrimSpace(s)
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
func ConvertLabel(label string, fieldMap map[string]string) (string, bool) {
    val, exist := fieldMap[label]
	return val, exist
}

type exchange struct {
	Amount int                `json:"amount"`
    Base   string             `json:"base"`
    Date   string             `json:"date"`
    Rates  map[string]float64 `json:"rates"`
}

// 為替レート取得
func GetExchangeUSDtoJPY() float64 {
	res, err := http.Get(`https://api.frankfurter.dev/v1/latest?from=USD&to=JPY`)
	if err != nil || res == nil {
		return 1 // ×1で＄表記そのままにする
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	var exchange exchange
	json.Unmarshal(body, &exchange)

	if err != nil {
		return 1
	}
	return exchange.Rates["JPY"]
}

// 国外価格 * rate >> 日本価格
func CalcExchangedPrice(foreignPrice string, rate float64) string {
	parsed, _ := ParsePrice(foreignPrice)
	foreignP  := decimal.NewFromInt(int64(parsed))
	exchange  := decimal.NewFromFloat(rate)
	// 小数点以下は切り捨て
	return foreignP.Mul(exchange).Truncate(0).String()
}

// リンクの重複を排除する
func GetDistinctLinks(links []string) []string {
	removed := map[string]struct{}{}

	for _, link := range links {
		removed[link] = struct{}{}
	}
	distinctLinks := []string{}

	for link := range removed {
		distinctLinks = append(distinctLinks, link)
	}
	return distinctLinks
}

// 重複なしのURL配列を取得
func MapToSliceUrl(visited map[string]struct{}) []string {
    urls  := make([]string, 0, 500)
    mutex := &sync.Mutex{}

    for k, _ := range visited {
        urls = LockedAppend(mutex, urls, k)
    }
    return urls
}

// 必要なリンクだけ取得
func GetNeedLinks(links []string, needPattern *regexp.Regexp, cap int) []string {
    needLinks := make([]string, 0, cap)

    for _, link := range links {
        if needPattern.MatchString(link) {
            needLinks = append(needLinks, link)
        }
    }
    return needLinks
}

// 相対パスから絶対パスへ変換
func ToAbsLinks(links []string, prefix string, cap int) []string {
    absLinks := make([]string, 0, cap)

    for _, link := range links {
        absLinks = append(absLinks, prefix + link)
    }
    return absLinks
}

// link収集
func CollectLinks(eachSelector string, doc *goquery.Document, cap int) []string {
    var links = make([]string, 0, cap)

    // 複数リンクを収集
    doc.Find(eachSelector).Each(func(idx int, selector *goquery.Selection) {
        link, _ := selector.Attr("href")
        if link != "" {
            links = append(links, link)
        }
    })
    return links
}

// cdp.Node.Attributes の構造 []string{ id", "frame1", "src", "https://example.com", "class", "foo", ...}
func GetAttr(node *cdp.Node, attrName string) string {
    for i := 0; i < len(node.Attributes)-1; i += 2 {
        if node.Attributes[i] == attrName {
            // 次の要素が属性の内容
            return node.Attributes[i+1]
        }
    }
    return ""
}

// 文字列の様式を統一する
func NormalizeString(str string) string {
    normalized := width.Narrow.String(str)
    normalized  = strings.ToLower(normalized)
    normalized  = strings.TrimSpace(normalized)
    normalized  = strings.ReplaceAll(normalized, " ", "")
    normalized  = strings.ReplaceAll(normalized, "\r\n", "")
    normalized  = strings.ReplaceAll(normalized, "\n", "")
    normalized  = strings.ReplaceAll(normalized, "\"", "")
    normalized  = strings.ReplaceAll(normalized, "“", "")
    normalized  = strings.ReplaceAll(normalized, "”", "")
    normalized  = strings.ReplaceAll(normalized, "'", "")
    normalized  = strings.ReplaceAll(normalized, "-", "")
    return normalized
}