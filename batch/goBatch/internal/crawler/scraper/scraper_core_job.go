package scraper

import (
	"log"

	"github.com/kazGear/portfolio/goBatch/internal/crawler/model"
	C "github.com/kazGear/portfolio/goBatch/pkg/constants"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

// ギター構造体の構築フレームワーク
func buildJobFrame(data map[string]string, url string, logger *log.Logger) (*model.Job) {
	job  := model.Job{}
    trim := utils.TrimSpace()

	job.Title = trim(data[C.Title])
	job.Url = trim(url)
	job.Description = data[C.Description]

	return &job
}