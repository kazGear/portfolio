package repository

import "github.com/kazGear/portfolio/goBatch/internal/crawler/model"

type GuitarRepository interface {
	Saver[*model.Guitar]
}

type JobRepository interface {
	Saver[*model.Job]
	Selector[string, map[int64]struct{}]
}

type Saver[T any] interface {
	Save(models []T) (ok int, ng int, errors []error)
}

type Selector[T1 any, T2 any] interface {
	Select(T1) T2
}
// Deleter / Updater も追加する？