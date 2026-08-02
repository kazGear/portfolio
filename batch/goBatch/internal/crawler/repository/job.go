package repository

import (
	"fmt"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/jmoiron/sqlx"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/model"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/repository/sql"
)

var (
    _jobFeatures map[string][]*model.JobFeature = make(map[string][]*model.JobFeature)
    _jobOptions  map[string][]*model.JobOption  = make(map[string][]*model.JobOption)
    _mutex       *sync.Mutex                    = &sync.Mutex{}
)

type jobRepository struct {
    db *sqlx.DB
}

func NewJobRepository(db *sqlx.DB) Repository[*model.Job] {
    return &jobRepository{ db: db }
}

func InjectionJobFeatures(features []*model.JobFeature, url string) {
    _mutex.Lock()
    defer _mutex.Unlock()
    _jobFeatures[url] = features
}

func InjectionJobOptions(options []*model.JobOption, url string) {
    _mutex.Lock()
    defer _mutex.Unlock()
    _jobOptions[url] = options
}

// グローバル変数アクセス用（スレッドセーフ）
func getJobFeaturesCrowdWorksTech(url string) ([]*model.JobFeature, bool) {
    _mutex.Lock()
    defer _mutex.Unlock()

    features, exists := _jobFeatures[url]
    return features, exists
}

// グローバル変数アクセス用（スレッドセーフ）
func getJobOptionsCrowdWorksTech(url string) ([]*model.JobOption, bool) {
    _mutex.Lock()
    defer _mutex.Unlock()

    options, exists := _jobOptions[url]
    return options, exists
}

func (r *jobRepository) Save(jobs []*model.Job) (ok int, ng int, errors []error) {
    errs    := make([]error, 0, 200)
    okCount := 0
    ngCount := 0

    savedFeatureJobIds, err := r.selectSavedFeatureJobIds()

    if err != nil {
        return 0, len(jobs), []error{err}
    }

    for _, job := range jobs {
        if err := r.updates(job, savedFeatureJobIds); err != nil {
            errs = append(errs, err)
            ngCount++
            continue
        }
        okCount++
    }
    return okCount, ngCount, errs
}

func (r *jobRepository) updates(job *model.Job, savedFeatureJobIds map[int64]struct{}) error {
    // マルチバイトでも１文字としてカウント
    titleLength := utf8.RuneCountInString(job.Title)

    // 必須チェック
    if titleLength < 1 || 255 < titleLength || job.Url == "" {
        return fmt.Errorf("[Invalid must field]: Title=%v, Url=%v\n",
            job.Title,
            job.Url,
        )
    }
    // トランザクション開始
    transaction, err := r.db.Beginx()

    if err != nil {
        return err
    }
    defer transaction.Rollback()

    // 案件基本情報 update or insert
    if err := upsert(job, transaction); err != nil {
        return err
    }

    // 生成されたjob_idを取得, 各Featureにセットして同期
    jobId, err := r.selectCurrentJobId(job, transaction)
    setJobId(jobId, job.Url)

    if err != nil {
        return err
    }
    // すでに案件情報が保存されてるか確認し、保存済みであれば後続処理は不要
    // Features, optionsが共に無くても同様
    _, exists   := savedFeatureJobIds[jobId]
    features, _ := getJobFeaturesCrowdWorksTech(job.Url)
    options, _  := getJobOptionsCrowdWorksTech(job.Url)

    if exists || (len(features) <= 0 && len(options) <= 0) {
        return transaction.Commit()
    }
    // 付随情報をまとめてインサート(features ,options)
    sqlBulkInsertFeatures := createSqlBulkInsertFeatures(job.Url)
    _, err = transaction.Exec(sqlBulkInsertFeatures)

    if err != nil {
        return err
    }
    sqlBulkInsertOptions := createSqlBulkInsertOptions(job.Url)
    _, err = transaction.Exec(sqlBulkInsertOptions)

    if err != nil {
        return err
    }
    // 付随情報を保存したjobIdを記録
    _, err = transaction.Exec(sql.InsertJobId(), jobId)

    if err != nil {
        return err
    }

    if err := transaction.Commit(); err != nil {
        return err
    }

    // 処理済のデータは削除
    removeJobDataCrowdWorksTech(job.Url)
    removeJobOptionsCrowdWorksTech(job.Url)
    return nil
}

func upsert(job *model.Job, transaction *sqlx.Tx) error {
    // UPDATE（存在すれば更新）
    res, err := transaction.NamedExec(sql.UpdateJob(), job)

    if err != nil {
        return err
    }
    // UPDATE で更新された行数を確認
    updateRows, err := res.RowsAffected()

    if err != nil {
        return err
    }
    // UPDATE されてないなら INSERT
    if updateRows == 0 {
        _, err := transaction.NamedExec(sql.InsertJob(), job)

        if err != nil {
            return err
        }
    }
    return nil
}

// DB保存済の情報は不要なため削除、メモリ解法
func removeJobDataCrowdWorksTech(url string) {
    _, exists := getJobFeaturesCrowdWorksTech(url)

    _mutex.Lock()
    defer _mutex.Unlock()

    if exists {
        delete(_jobFeatures, url)
    }
}

// DB保存済の情報は不要なため削除、メモリ解法
func removeJobOptionsCrowdWorksTech(url string) {
    _, exists := getJobOptionsCrowdWorksTech(url)

    _mutex.Lock()
    defer _mutex.Unlock()

    if exists {
        delete(_jobOptions, url)
    }
}

func (r *jobRepository) selectSavedFeatureJobIds() (map[int64]struct{}, error) {
    savedFeatureJobIds := make(map[int64]struct{})

    var jobIds []int64
    err := r.db.Select(&jobIds, sql.SelectCreatedFeatures())

    if err != nil {
        return map[int64]struct{}{}, err
    }

    for _, id := range jobIds {
        savedFeatureJobIds[id] = struct{}{}
    }
    return savedFeatureJobIds, nil
}

func (r *jobRepository) selectCurrentJobId(job *model.Job, transaction *sqlx.Tx) (int64, error) {
    rows, err := transaction.NamedQuery(sql.SelectCurrentJobId(), job)

    if err != nil {
        return -1, err
    }
    defer rows.Close()

    if !rows.Next() {
        return -1, fmt.Errorf("Error no rows.")
    }
    var jobId int64

    err = rows.Scan(&jobId)

    if err != nil {
        return -1, err
    }

    return jobId, nil
}

func setJobId(jobId int64, url string) {
    _mutex.Lock()
    defer _mutex.Unlock()

    for _, feature := range _jobFeatures[url] {
        feature.JobId = jobId
    }

    for _, option := range _jobOptions[url] {
        option.JobId = jobId
    }
}

func createSqlBulkInsertFeatures(url string) string {
    var builder strings.Builder

    features, _ := getJobFeaturesCrowdWorksTech(url)

    if len(features) == 0 {
        return ""
    }
    builder.WriteString(" INSERT INTO t_job_features VALUES ")

    for i := 0; i < len(features); i++ {
        // 初回以外カンマで区切る
        if i > 0 {
            builder.WriteString(",")
        }

        builder.WriteString(fmt.Sprintf(" (%v, '%v', '%v', '%v') ",
            features[i].JobId,
            features[i].FeatureName,
            features[i].Category,
            features[i].RequirementType),
        )
    }
    builder.WriteString(";")

    return builder.String()
}

func createSqlBulkInsertOptions(url string) string {
    var builder strings.Builder

    options, _ := getJobOptionsCrowdWorksTech(url)

    if len(options) == 0 {
        return ""
    }
    builder.WriteString(" INSERT INTO t_job_options VALUES ")

    for i := 0; i < len(options); i++ {
        // 初回以外カンマで区切る
        if i > 0 {
            builder.WriteString(",")
        }

        builder.WriteString(fmt.Sprintf(" (%v, '%v') ",
            options[i].JobId,
            options[i].Option),
        )
    }
    builder.WriteString(";")

    return builder.String()
}