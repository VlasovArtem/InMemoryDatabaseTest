package domain

type InMemoryDatabase[K comparable, V any] interface {
	Set(key K, value V)
	Get(key K) V
	Delete(key K)
	StartTransaction()
	Commit()
	Rollback()
}
