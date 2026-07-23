package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/model"
	"github.com/kazGear/portfolio/goBatch/internal/crawler/repository/sql"
)

type Repository interface {
    Upsert(g *model.Guitar) error
    UpsertAll(guitars []*model.Guitar) (ok int, ng int, errs []error)
}

type guitarRepository struct {
    db *sqlx.DB
}

func NewGuitarRepository(db *sqlx.DB) Repository {
    return &guitarRepository{ db: db }
}

func (r *guitarRepository) Upsert(guitar *model.Guitar) error {
    // pkチェック
    if guitar.Maker <= 0 || len(guitar.Name) <= 0 || len(guitar.Color) <= 0 {
        return fmt.Errorf("[Invalid primary key]: maker=%v, name=%v, color=%v\n",
            guitar.Maker,
            guitar.Name,
            guitar.Color,
        )
    }
    // 画像確認
    if len(guitar.Src) <= 0 {
        return fmt.Errorf("画像URLは必須項目です。")
    }

    // 1. UPDATE（存在すれば更新）
    res, err := r.db.NamedExec(sql.UpdateGuitar(), guitar)

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
        _, err := r.db.NamedExec(sql.InsertGuitar(), guitar)
        return err
    }
    return nil
}

func (r *guitarRepository) UpsertAll(guitars []*model.Guitar) (int, int, []error) {
    errs    := make([]error, 0, 300)
    okCount := 0
    ngCount := 0

    for _, guitar := range guitars {
        err := r.Upsert(guitar)
        if err != nil {
            errs = append(errs, err)
            ngCount++
            continue
        }
        okCount++
    }
    return okCount, ngCount, errs
}