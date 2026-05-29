package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePrice(t *testing.T) {
	prices := []string{
		"10000",
		"10,000",
		"￥10,000",
		"１００００",
		"１０，０００円",
		"250,000円（税別） ／ 275,000円（税込）",
	}

	for _, price := range prices {
		price := price // ← これ重要（並列テスト時の罠回避）
		result, _ := ParsePrice(price)
		assert.True(t, 0 <= result && result <= 100000000)
	}
}

func TestToHankakuNumber(t *testing.T) {
	prices := []string{
		"10000",
		"10,000",
		"￥10,000",
		"１００００",
		"１０，０００円",
	}

	for _, v := range prices {
		price := ToHankakuNumber(v)
		assert.Contains(t, price, "000")
	}
}