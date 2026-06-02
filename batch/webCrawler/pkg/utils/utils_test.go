package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var errMessage = "Failed >>> want: %v, actual: %v"

func TestParsePrice(t *testing.T) {
	prices := []string{
		"10000",
		"10,000",
		"￥10,000",
		"１００００",
		"１０，０００円",
		"250,000円（税別） ／ 275,000円（税込）",
		"BK, SW 726,000 yen (without tax: 660,000 yen) DCAR 748,000 yen (without tax: 680,000 yen)",
	}

	for _, price := range prices {
		price := price // 並列テスト時の罠回避
		result, _ := ParsePrice(price)
		assert.True(t, 1 <= result && result <= 100000000)
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
	}

	for _, s := range scales {
		s := s
		actual := TrimScaleUnit(s.scale)
		assert.Equal(t, s.want, actual)
	}
}
