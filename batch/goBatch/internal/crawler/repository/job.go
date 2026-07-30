package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/model"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/repository/sql"
)

type jobRepository struct {
    db *sqlx.DB
}

func NewJobRepository(db *sqlx.DB) Repository[*model.Job] {
    return &jobRepository{ db: db }
}

func (r *jobRepository) Save(jobs []*model.Job) (ok int, ng int, errors []error) {
    errs    := make([]error, 0, 300)
    okCount := 0
    ngCount := 0

    for _, job := range jobs {
        err := r.upsert(job)
        if err != nil {
            errs = append(errs, err)
            ngCount++
            continue
        }
        okCount++
    }
    return okCount, ngCount, errs
}

func (r *jobRepository) upsert(job *model.Job) error {
    // 必須チェック
    if (1 <= len(job.Title) && len(job.Title) <= 100) || job.Url == "" {
        return fmt.Errorf("[Invalid must field]: Title=%v, Url=%v\n",
            job.Title,
            job.Url,
        )
    }

    // 1. UPDATE（存在すれば更新）
    res, err := r.db.NamedExec(sql.UpdateJob(), job)

    if err != nil {
        return err
    }
    // 2. UPDATE で更新された行数を確認
    rows, err := res.RowsAffected()

    if err != nil {
        return err
    }
    // 3. UPDATE されてないなら INSERT
    if rows == 0 {
        _, err := r.db.NamedExec(sql.InsertJob(), job)
        return err
    }
    return nil
}
