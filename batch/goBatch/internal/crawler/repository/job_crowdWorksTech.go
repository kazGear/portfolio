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
    _jobFeaturesCrowdWorksTech map[string][]*model.JobFeature = make(map[string][]*model.JobFeature)
    _mutex                   *sync.Mutex                      = &sync.Mutex{}
)

type jobRepository struct {
    db *sqlx.DB
}

func NewJobRepositoryCrowdWorksTech(db *sqlx.DB) Repository[*model.Job] {
    return &jobRepository{ db: db }
}

func InjectionJobFeaturesCrowdWorksTech(features []*model.JobFeature, url string) {
    // 読み取りは同時アクセス可能
    _mutex.Lock()
    _jobFeaturesCrowdWorksTech[url] = features
    _mutex.Unlock()
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
        err := r.updates(job, savedFeatureJobIds)
        if err != nil {
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
    if titleLength < 1 || 100 < titleLength || job.Url == "" {
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
    // 生成されたjob_idを取得, 各Featureにセットして同期
    jobId, err := r.selectCurrentJobId(job, transaction)
    setJobId(jobId, job.Url)

    if err != nil {
        return err
    }
    // すでに案件の特徴が保存されてるか確認し、保存済みであれば後続処理は不要
    // Featuresが無くても同様
    _, exists := savedFeatureJobIds[jobId]
    _mutex.Lock()
    features  := _jobFeaturesCrowdWorksTech[job.Url]
    _mutex.Unlock()

    if exists || len(features) <= 0 {
        return transaction.Commit()
    }


    // まとめてインサート(features)
    sqlBulkInsertFeatures := createSqlBulkInsertFeatures(job.Url)
    _, err = transaction.Exec(sqlBulkInsertFeatures)

    if err != nil {
        return err
    }
    // 案件の特徴を保存したjobIdを記録
    _, err = transaction.Exec(sql.InsertJobId(), jobId)

    if err != nil {
        return err
    }

    err = transaction.Commit()

    if err != nil {
        return err
    }
    // 処理済のデータは削除
    removeJobFeaturesCrowdWorksTech(job.Url)
    return nil
}

// DB保存済の情報は不要なため削除、メモリ解法
func removeJobFeaturesCrowdWorksTech(url string) {
    _mutex.Lock()
    defer _mutex.Unlock()

    _, exists := _jobFeaturesCrowdWorksTech[url]

    if exists {
        delete(_jobFeaturesCrowdWorksTech, url)
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

    for _, feature := range _jobFeaturesCrowdWorksTech[url] {
        feature.JobId = jobId
    }
}

func createSqlBulkInsertFeatures(url string) string {
    var builder strings.Builder

    _mutex.Lock()
    features := _jobFeaturesCrowdWorksTech[url]
    _mutex.Unlock()

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