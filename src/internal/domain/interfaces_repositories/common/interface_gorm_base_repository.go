package interfaces_repositories_common


type IGormBaseRepository[T any] interface {
	Query() *T
	Insert(entity interface{}) error
	Update(entity interface{}) error
	Remove(entity interface{}) error
}