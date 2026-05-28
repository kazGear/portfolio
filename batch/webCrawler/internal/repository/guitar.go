package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kazGear/portfolio/webCrawler/internal/model"
)

func SaveGuitars(db *sqlx.DB, list []model.Guitar) error {
    query := `
        INSERT INTO guitars (model, maker, price, url, image, shop)
        VALUES (:model, :maker, :price, :url, :image, :shop)
        ON CONFLICT (url) DO UPDATE SET
            model = EXCLUDED.model,
            maker = EXCLUDED.maker,
            price = EXCLUDED.price,
            image = EXCLUDED.image,
            shop  = EXCLUDED.shop;
    `
    for _, g := range list {
        _, err := db.NamedExec(query, g)
        if err != nil {
            return err
        }
    }
    return nil
}
