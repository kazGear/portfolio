package scraper

import (
	"testing"

	C "github.com/kazGear/portfolio/goBatch/pkg/constants"
	"github.com/stretchr/testify/assert"
)

func TestGetJobPrice(t *testing.T) {
	prices := []struct {
		price    string
		wantMin  int
        wantMax  int
	}{
		{
			price: "55 ~ 60万",
            wantMin: 550000,
            wantMax: 600000,
		},
		{
			price: "55万 ~ 60万",
            wantMin: 550000,
            wantMax: 600000,
		},
		{
			price: "400,000 〜 550,000",
            wantMin: 400000,
            wantMax: 550000,
		},
		{
			price: "　1,500000　",
            wantMin: 1500000,
            wantMax: 1500000,
		},
		{
			price: " 150万円 ",
            wantMin: 1500000,
            wantMax: 1500000,
		},
        {
			price: "750000 -120万円 ",
            wantMin: 750000,
            wantMax: 1200000,
		},
        {
			price: "750000 - 850,000 ",
            wantMin: 750000,
            wantMax: 850000,
		},
		{
			price: "75万円～ 85 万円",
            wantMin: 750000,
            wantMax: 850000,
		},
        {
			price: "75 万円 ～ 85万円",
            wantMin: 750000,
            wantMax: 850000,
		},
        {
			price: " ~ 85 万円",
            wantMin: C.UndefinedPrice,
            wantMax: 850000,
		},
        {
			price: "75万円 ～",
            wantMin: 750000,
            wantMax: C.UndefinedPrice,
		},
	}

	for _, p := range prices {
		p := p // 並列テスト時の罠回避
		actualMin, actualMax := getJobPrice(p.price)
		assert.Equal(t, p.wantMin, actualMin)
        assert.Equal(t, p.wantMax, actualMax)
	}
}
