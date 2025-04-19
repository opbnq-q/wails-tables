package services

type Service[T any] interface {
	GetAll() ([]*T, error)
	GetById(id uint) (*T, error)
	Create(item T) (T, error)
	Update(item T) (T, error)
	Delete(id uint) error
	Count() (int64, error)
}
