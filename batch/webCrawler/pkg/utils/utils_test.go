package utils

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errMessage = "Failed >>> want: %v, actual: %v"

func TestParsePrice(t *testing.T) {
	prices := []struct {
		price string
		want  int
	}{
		{
			price: "10000", want: 10000,
		},
		{
			price: "10,000", want: 10000,
		},
		{
			price: "￥10,000", want: 10000,
		},
		{
			price: "１０，０００", want: 10000,
		},
		{
			price: "１０，０００円", want: 10000,
		},
		{
			price: "250,000円（税別） ／ 275,000円（税込）",
			want: 250000,
		},
		{
			price: "BK, SW 726,000 yen (without tax: 660,000 yen) DCAR 748,000 yen (without tax: 680,000 yen)",
			want: 660000,
		},
		{
			price: "ASK", want: -1,
		},
		{
			price: "open", want: -1,
		},
		{
			price: "999999999", want: -1,
		},
		{
			price: "$ 1 149", want: 1149,
		},
	}

	for _, p := range prices {
		p := p // 並列テスト時の罠回避
		actual, _ := ParsePrice(p.price)
		assert.Equal(t, p.want, actual)
	}
}

func TestSearchWoodCode(t *testing.T) {
	var maple int = 6
	var hardMaple int = 1

	woods := []string{
		"hardMaple Paduak 7P",
		"HardMaple, Walnut, Paduak 7P",
		"Maple, Paduak 7P HardMaple", // 具体名が拾われる（マスタ順サーチ）
		"HardMaple Maple",
		"Hard Maple Maple",
	}

	assert.Equal(t, hardMaple, SearchWoodCode(woods[0]))
	assert.Equal(t, hardMaple, SearchWoodCode(woods[1]))
	assert.Equal(t, hardMaple, SearchWoodCode(woods[2]))
	assert.Equal(t, hardMaple, SearchWoodCode(woods[3]))
	assert.Equal(t, maple, SearchWoodCode(woods[4]))
}

func TestGetFretCount(t *testing.T) {
	frets := []struct {
		input string
		want  int
	}{
		{
			input: "JESCAR FW57110-NS, 24frets",
			want:  24,
		},
		{
			input: "JESCAR FW57110-NS, 14frets",
			want:  -1,
		},
		{
			input: "JESCAR FW57110-NS15frets",
			want:  15,
		},
		{
			input: "JESCAR FW57110-NS39frets",
			want:  39,
		},
		{
			input: "JESCAR FW57110-NS40frets",
			want:  -1,
		},
		{
			input: "JESCAR FW58118-NS, 22frets",
			want:  22,
		},
		{
			input: "JESCAR FW58118-NS, 22frets",
			want:  22,
		},
		{
			input: "Jescar FW57110, 24FRET",
			want:  24,
		},
		{
			input: "Jescar FW57110, 24F",
			want:  24,
		},
		{
			input: "22FRET Stainless",
			want:  22,
		},
		{
			input: "22 frets (Stainless)",
			want:  22,
		},
		{
			input: "22FRET/24FRET（比較表記）",
			want:  22,
		},
		{
			input: "24",
			want:  24,
		},
	}

	for _, fret := range frets {
		fret := fret

		actual, err := GetFretCount(fret.input)

		if err != nil || actual == -1 {
			assert.Error(t, err)
		} else {
			if fret.want != actual {
				t.Fatalf(errMessage, fret.want, actual)
			}
		}
	}
}

func TestTrimScaleUnit(t *testing.T) {
	scales := []struct {
		scale string
		want  int
	}{
		{
			scale: "648mm",
			want:  648,
		},
		{
			scale: "648 mm",
			want:  648,
		},
		{
			scale: "648  mm",
			want:  648,
		},
		{
			scale: "６４８ｍｍ",
			want:  648,
		},
		{
			scale: "６４８　ｍｍ",
			want:  648,
		},
		{
			scale: "",
			want:  -1,
		},
		{
			scale: "820.00 mm",
			want:  820,
		},
		{
			scale: `24.75" / 628.65mm`,
			want:  24,
		},
	}

	for _, s := range scales {
		s := s
		actual := TrimScaleUnit(s.scale)
		assert.Equal(t, s.want, actual)
	}
}

func TestConvertColorCd(t *testing.T) {
	colors := []struct {
		color string
		want  int
	}{
		{
			color: "190 Red",
			want: 1,
		},
		{
			color: "Pearl Pink / SAKURA Pink",
			want: 2,
		},
		{
			color: "ORANGE DEEP",
			want: 3,
		},
		{
			color: "Mustard Yellow",
			want: 4,
		},
		{
			color: "Neon Green w/Kamikaze Graphic",
			want: 5,
		},
		{
			color: "skyBlue",
			want: 6,
		},
		{
			color: "Driftwood Blue w/Bla Filler",
			want: 7,
		},
		{
			color: "Indigo Purple w/Purple Pearl Dark",
			want: 8,
		},
		{
			color: "gray",
			want: 9,
		},
		{
			color: "See Thru Black Sunburst",
			want: 10,
		},
		{
			color: "Snow White",
			want: 11,
		},
		{
			color: "Driftwood Natural w/Dark Filler",
			want: 12,
		},
		{
			color: "Brown Burst",
			want: 13,
		},
		{
			color: "Metallic Gold",
			want: 14,
		},
		{
			color: "Royal Silver",
			want: 15,
		},
		{
			color: "extraGold",
			want: 14,
		},
		{
			color: "3 Tone Sunburst",
			want: 99,
		},
		{
			color: "blueeeeee",
			want: 7,
		},
		{
			color: "aqua",
			want: 7,
		},
	}
	for _, c := range colors {
		actual := ConvertColorCd(c.color)
		assert.Equal(t, c.want, actual)
	}
}

func TestParseWight(t *testing.T) {
	weights := []struct {
		weight string
		want   float64
	}{
		{
			weight: "2.30 +/- 10% Kg", want: 2.3,
		},
		{
			weight: "２.３０ +/- 10% Kg", want: 2.3,
		},
		{
			weight: "2 +/- 10% Kg", want: -1,
		},
		{
			weight: "2.30 +/- 10%", want: 2.3,
		},
	}

	for _, w := range weights {
		w := w
		actual, _ := ParseWight(w.weight)
		assert.Equal(t, w.want, actual)
	}
}

func TestConvertRealUrl(t *testing.T) {
	urls := []string {
		`/_next/image?url=https%3A%2F%2Fstrandbergs.cdn-norce.tech%2F81de294c-bd81-40ae-ad76-e8b983e5b528&w=3840&q=85`,
		`/_next/image?url=https%3A%2F%2Fstrandbergs.cdn-norce.tech%2Fe3b3ad85-57b9-449f-9d72-6dfc03de6c11&w=3840&q=85`,
		// `/product-images/Custom/CUSGJ5555/Ebony2-Pickup/front-banner-1600_900.png`,
	}

	for _, url := range urls {
		url := url
		converted, _ := ConvertRealUrl(url)
		log.Println(converted)
		assert.True(t, strings.HasPrefix(converted, "http"))
	}
}

func TestGetExchangeUSDtoJPY(t *testing.T) {
	rate := GetExchangeUSDtoJPY()
	assert.GreaterOrEqual(t, rate, 1.0)
	assert.LessOrEqual(t, rate, 360.0)
}
func TestCalcExchangedPrice(t *testing.T) {
	tests := []struct {
		dollar string
		rate   float64
		want   string
	}{
		{
			dollar: "99", rate: 120.0, want: "11880",
		},
		{
			dollar: "99", rate: 120.1, want: "11889",
		},
		{
			dollar: "99", rate: 120.11, want: "11890",
		},
		{
			dollar: "$99", rate: 120.11, want: "11890",
		},
	}

	for _, test := range tests {
		test := test
		actual := CalcExchangedPrice(test.dollar, test.rate)
		assert.Equal(t, test.want, actual)
	}
}

func TestConvertRelToAbsUrl(t *testing.T) {
	result  := `https://www.ibanez.com/common/product_artist_file/file/p_region_AZ2407F_BSR_00_01.png`
	baseUrl := `https://www.ibanez.com/jp/products/detail/az2407f_00_01.html`
	src 	:= `/common/product_artist_file/file/p_region_AZ2407F_BSR_00_01.png`
	try, _  := ConvertRelToAbsUrl(baseUrl, src)

	assert.Equal(t, try, result)
}