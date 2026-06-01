package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
)

type Repository interface {
    Upsert(g model.Guitar) error
    UpsertAll(guitars []model.Guitar) error
}

type guitarRepository struct {
    db *sqlx.DB
}

func NewGuitarRepository(db *sqlx.DB) Repository {
    return &guitarRepository{ db: db }
}

func (r *guitarRepository) Upsert(guitar model.Guitar) error {
    // pkチェック
    if guitar.Maker <= 0 || len(guitar.Name) <= 0 {
        return fmt.Errorf("invalid primary key: maker=%d, name=%s", guitar.Maker, guitar.Name)
    }
    // 1. UPDATE（存在すれば更新）
    res, err := r.db.NamedExec(`
        UPDATE t_guitars
           SET body_finish         = :body_finish,
               body_material       = :body_material,
               body_material_front = :body_material_front,
               body_material_back  = :body_material_back,
               bridge              = :bridge,
               color               = :color,
               controls            = :controls,
               comment             = :comment,
               fingerboard         = :fingerboard,
               fret_count          = :fret_count,
               inlays              = :inlays,
               joint               = :joint,
               neck_material       = :neck_material,
               pickups             = :pickups,
               price               = :price,
               scale_length_mm     = :scale_length_mm,
               series              = :series,
               src                 = :src,
               weight              = :weight
         WHERE maker = :maker
           AND name  = :name
    `, guitar)

    if err != nil {
        return err
    }
    // 2. UPDATE で更新された行数を確認
    rows, err := res.RowsAffected()
    if err != nil {
        return err
    }
    // 3. 行がなければ INSERT
    if rows == 0 {
        _, err := r.db.NamedExec(`
            INSERT INTO t_guitars
            (
                maker,
                name,
                body_finish,
                body_material,
                body_material_front,
                body_material_back,
                bridge,
                color,
                controls,
                comment,
                fingerboard,
                fret_count,
                inlays,
                joint,
                neck_material,
                pickups,
                price,
                scale_length_mm,
                series,
                src,
                weight
            )
                VALUES
            (
                :maker,
                :name,
                :body_finish,
                :body_material,
                :body_material_front,
                :body_material_back,
                :bridge,
                :color,
                :controls,
                :comment,
                :fingerboard,
                :fret_count,
                :inlays,
                :joint,
                :neck_material,
                :pickups,
                :price,
                :scale_length_mm,
                :series,
                :src,
                :weight
            )
        `, guitar)
        return err
    }
    return nil
}

func (r *guitarRepository) UpsertAll(guitars []model.Guitar) error {
    for _, guitar := range guitars {
        if err := r.Upsert(guitar); err != nil {
            log.Printf("DB保存失敗: maker: %v, name: %v (%v)", guitar.Maker, guitar.Name, err)
        }
    }
    return nil
}