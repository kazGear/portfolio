package repository

import (
	"fmt"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
	"github.com/kazGear/portfolio/webCrawler/pkg/utils"
)

type Repository interface {
    Upsert(g model.Guitar) error
    UpsertAll(guitars []model.Guitar)
}

type guitarRepository struct {
    db *sqlx.DB
}

func NewGuitarRepository(db *sqlx.DB) Repository {
    return &guitarRepository{ db: db }
}

func (r *guitarRepository) Upsert(guitar model.Guitar) error {
    log.Printf("Try upsert >>> %v", guitar.String())

    // pkチェック
    if guitar.Maker <= 0 || len(guitar.Name) <= 0 || len(guitar.Color) <= 0 {
        return fmt.Errorf("invalid primary key: maker=%v, name=%v, color=%v\n",
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
    res, err := r.db.NamedExec(`
        UPDATE t_guitars
           SET body_finish         = :body_finish,
               body_material       = :body_material,
               body_material_front = :body_material_front,
               body_material_back  = :body_material_back,
               bridge              = :bridge,
               color               = :color,
               color_cd            = :color_cd,
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
           AND color = :color
    `, guitar)

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
                color_cd,
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
                :color_cd,
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

func (r *guitarRepository) UpsertAll(guitars []model.Guitar) {
    errs  := make([]error, 0, 300)
    mutex := &sync.Mutex{}

    for _, guitar := range guitars {
        err := r.Upsert(guitar)
        errs = utils.LockedAppend(mutex, errs, err)
    }
    for _, err := range errs {
        log.Printf("[upsert error]: %v", err)
    }
}