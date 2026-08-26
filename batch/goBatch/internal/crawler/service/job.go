package service

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/kazGear/portfolio/goBatch/internal/crawler/model"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/repository"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/scraper"
	C "github.com/kazGear/portfolio/goBatch/pkg/constants"
)

type jobCrawlerService struct {
    repository repository.JobRepository
}

func NewJobCrawlerService(repository repository.JobRepository) CrawlerService {
    return &jobCrawlerService{ repository: repository }
}

type JobBoard struct {
    name     string
    scraper  scraper.Scraper[*model.Job]
    provider scraper.PageProvider
    parser   scraper.ModelParser[*model.Job]
}

func NewJobBoard(name     string,
                 scraper  scraper.Scraper[*model.Job],
                 provider scraper.PageProvider,
                 parser   scraper.ModelParser[*model.Job],
) *JobBoard {
    return &JobBoard{ name, scraper, provider, parser }
}

func (j *jobCrawlerService) RunCrawler() {
    envName            := "PARALLEL_COUNT_JOB"
    parallelCount, err := strconv.Atoi(os.Getenv(envName))

    if err != nil || parallelCount <= 0 {
        log.Panicf("%v must be a positive integer.", envName)
    }

    queue := make(chan struct{}, parallelCount) // 並列数制御

    jobBoards := jobBoardFactory()
    wg        := &sync.WaitGroup{}

    // クロール + スクレイピング + DB保存
    for _, jobBoard := range jobBoards {
        jobBoard := jobBoard
        wg.Add(1)

        go func(jobBoard JobBoard) {
            queue <- struct{}{}
            defer wg.Done()
            defer func() { <- queue }() // 次のワーカーへ

            // chromedpコンテキスト構築
            cancelAlloc, cancelParent, parentCtx := createChromedpCtx()
            defer cancelAlloc()
            defer cancelParent()

            log.Printf(C.DecoLabel, "Started job crawler " + jobBoard.name)

            startTime := time.Now() // 処理時間計測開始

            // クローラー起動
            jobBoard.scraper.CollectLinks(parentCtx)
            jobs := jobBoard.scraper.Scrape(jobBoard.provider, jobBoard.parser, parentCtx)
            okCnt, ngCnt, errs := j.repository.Save(jobs)

            // ログ
            log.Printf("[Upsert result %v]: OK %v 件, NG %v 件", jobBoard.name, okCnt, ngCnt)

            for _, err := range errs {
                log.Println(err)
            }
            log.Printf(C.DecoLabel, "Finished job crawler " + jobBoard.name)
            log.Printf("%v crawler processing time: %v\n", jobBoard.name, time.Since(startTime))
        }(*jobBoard)
    }
    wg.Wait()
}

// 各種メーカー作成
func jobBoardFactory() map[string]*JobBoard {
    jobBoards := map[string]*JobBoard{}

    jobBoardName := C.Midworks
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperMidworks(),
        scraper.NewCallBacksMidworks(),
        scraper.NewCallBacksMidworks(),
    )

    jobBoardName = C.TechReach
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperTechReach(),
        scraper.NewCallBacksTechReach(),
        scraper.NewCallBacksTechReach(),
    )

    jobBoardName = C.EngineerFactory
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperEngineerFactory(),
        scraper.NewCallBacksEngineerFactory(),
        scraper.NewCallBacksEngineerFactory(),
    )

    jobBoardName = C.FreelanceHub
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperFreelanceHub(),
        scraper.NewCallBacksFreelanceHub(),
        scraper.NewCallBacksFreelanceHub(),
    )

    jobBoardName = C.AGELESS
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperAgeless(),
        scraper.NewCallBacksAgeless(),
        scraper.NewCallBacksAgeless(),
    )

    jobBoardName = C.SES_JOB_LINK
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperSesJobLink(),
        scraper.NewCallBacksSesJobLink(),
        scraper.NewCallBacksSesJobLink(),
    )

    jobBoardName = C.FreelanceStart
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperFreelanceStart(),
        scraper.NewCallBacksFreelanceStart(),
        scraper.NewCallBacksFreelanceStart(),
    )

    jobBoardName = C.FreelanceJob
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperFreelanceJob(),
        scraper.NewCallBacksFreelanceJob(),
        scraper.NewCallBacksFreelanceJob(),
    )

    // jobBoardName = C.CrowdWorksTech
    // jobBoards[jobBoardName] = NewJobBoard(
    //     jobBoardName,
    //     scraper.NewScraperCrowdworksTech(),
    //     scraper.NewCallBacksCrowdworksTech(),
    //     scraper.NewCallBacksCrowdworksTech(),
    // )

    return jobBoards
}