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
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

type jobCrawlerService struct {
    repository repository.Repository[*model.Job]
}

func NewJobCrawlerService(repository repository.Repository[*model.Job]) CrawlerService {
    return &jobCrawlerService{ repository: repository }
}

type JobBoard struct {
    name     string
    scraper  scraper.Scraper[*model.Job]
    provider scraper.PageProvider
    parser   scraper.ModelParser[*model.Job]
    logger   *log.Logger
}

func NewJobBoard(name     string,
                 scraper  scraper.Scraper[*model.Job],
                 provider scraper.PageProvider,
                 parser   scraper.ModelParser[*model.Job],
                 logger   *log.Logger,
) *JobBoard {
    return &JobBoard{ name, scraper, provider, parser ,logger }
}

func (g *jobCrawlerService) RunCrawler() {
    envName            := "PARALLEL_COUNT_JOB"
    parallelCount, err := strconv.Atoi(os.Getenv(envName))

    if err != nil || parallelCount <= 0 {
        log.Panicf("%v must be a positive integer.", envName)
    }

    queue := make(chan struct{}, parallelCount) // 並列数制御

    jobBoards := jobBoardFactory()
    wg     := &sync.WaitGroup{}

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

            jobBoard.logger.Printf(C.DecoLabel, "Started job crawler " + jobBoard.name)

            startTime := time.Now() // 処理時間計測開始

            // クローラー起動
            jobBoard.scraper.CollectLinks(parentCtx)
            jobs := jobBoard.scraper.Scrape(jobBoard.provider, jobBoard.parser, parentCtx)
            okCnt, ngCnt, errs := g.repository.Save(jobs)

            // ログ
            jobBoard.logger.Printf("[Upsert result %v]: OK %v 件, NG %v 件", jobBoard.name, okCnt, ngCnt)
            log.Printf("[Upsert result %v]: OK %v 件, NG %v 件", jobBoard.name, okCnt, ngCnt) // コンソール用

            for _, err := range errs {
                jobBoard.logger.Println(err)
            }
            jobBoard.logger.Printf(C.DecoLabel, "Finished job crawler " + jobBoard.name)
            jobBoard.logger.Printf("Job crawler processing time: %v\n", time.Since(startTime))
        }(*jobBoard)
    }
    wg.Wait()
}

// 各種メーカー作成
func jobBoardFactory() map[string]*JobBoard {
    jobBoards := map[string]*JobBoard{}

    filepath := "./batch/goBatch/internal/crawler/logs/job/%v_%v.log"

    jobBoardName := C.Midworks
    logger       := utils.NewLogger(jobBoardName, filepath)
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperMidworks(logger),
        scraper.NewCallBacksMidworks(logger),
        scraper.NewCallBacksMidworks(logger),
        logger,
    )

    jobBoardName = C.TechReach
    logger       = utils.NewLogger(jobBoardName, filepath)
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperTechReach(logger),
        scraper.NewCallBacksTechReach(logger),
        scraper.NewCallBacksTechReach(logger),
        logger,
    )

    jobBoardName = C.EngineerFactory
    logger       = utils.NewLogger(jobBoardName, filepath)
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperEngineerFactory(logger),
        scraper.NewCallBacksEngineerFactory(logger),
        scraper.NewCallBacksEngineerFactory(logger),
        logger,
    )

    jobBoardName = C.FreelanceHub
    logger       = utils.NewLogger(jobBoardName, filepath)
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperFreelanceHub(logger),
        scraper.NewCallBacksFreelanceHub(logger),
        scraper.NewCallBacksFreelanceHub(logger),
        logger,
    )

    jobBoardName = C.CrowdWorksTech
    logger       = utils.NewLogger(jobBoardName, filepath)
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperCrowdworksTech(logger),
        scraper.NewCallBacksCrowdworksTech(logger),
        scraper.NewCallBacksCrowdworksTech(logger),
        logger,
    )

    jobBoardName = C.AGELESS
    logger       = utils.NewLogger(jobBoardName, filepath)
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperAgeless(logger),
        scraper.NewCallBacksAgeless(logger),
        scraper.NewCallBacksAgeless(logger),
        logger,
    )

    jobBoardName = C.SES_JOB_LINK
    logger       = utils.NewLogger(jobBoardName, filepath)
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperSesJobLink(logger),
        scraper.NewCallBacksSesJobLink(logger),
        scraper.NewCallBacksSesJobLink(logger),
        logger,
    )

    jobBoardName = C.FreelanceStart
    logger       = utils.NewLogger(jobBoardName, filepath)
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperFreelanceStart(logger),
        scraper.NewCallBacksFreelanceStart(logger),
        scraper.NewCallBacksFreelanceStart(logger),
        logger,
    )

    jobBoardName = C.FreelanceJob
    logger       = utils.NewLogger(jobBoardName, filepath)
    jobBoards[jobBoardName] = NewJobBoard(
        jobBoardName,
        scraper.NewScraperFreelanceJob(logger),
        scraper.NewCallBacksFreelanceJob(logger),
        scraper.NewCallBacksFreelanceJob(logger),
        logger,
    )

    return jobBoards
}