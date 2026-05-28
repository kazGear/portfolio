package utils

import (
	"regexp"
	"strconv"
)

// "¥128,000" → 128000
func ParsePrice(s string) int {
    // 数字以外を全部削除
    re := regexp.MustCompile(`\D`)
    cleaned := re.ReplaceAllString(s, "")

    if cleaned == "" {
        return 0
    }

    price, err := strconv.Atoi(cleaned)
    if err != nil {
        return 0
    }

    return price
}
