package repository

type Repository[T any] interface {
	Save(models []T) (ok int, ng int, errors []error)
}