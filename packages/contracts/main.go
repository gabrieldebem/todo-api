package contracts

type RepositoryInterface[T any] interface {
	FindAll() ([]T, error)
	FindWhere(query interface{}, args ...interface{}) ([]T, error)
	First(u *T) (T, error)
	Create(u *T) (T, error)
	Update(u *T) (T, error)
	Delete(u *T) (T, error)
}
