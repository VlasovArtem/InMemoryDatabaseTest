package repository

import (
	"test-inmemory-database/internal/domain"
)

type inMemoryDatabase[K comparable, V any] struct {
	data        map[K]V
	transaction *inMemoryDatabase[K, V]
}

func (database *inMemoryDatabase[K, V]) Set(key K, value V) {
	if database.transaction != nil {
		database.transaction.Set(key, value)
		return
	}
	database.data[key] = value
}

func (database *inMemoryDatabase[K, V]) Get(key K) V {
	if database.transaction != nil {
		return database.transaction.Get(key)
	}
	return database.data[key]
}

func (database *inMemoryDatabase[K, V]) Delete(key K) {
	if database.transaction != nil {
		database.transaction.Delete(key)
		return
	}
	delete(database.data, key)
}

func (database *inMemoryDatabase[K, V]) StartTransaction() {
	if database.transaction != nil {
		database.transaction.StartTransaction()
		return
	}

	databaseSnapshot := database.createDatabaseSnapshot()

	database.transaction = &inMemoryDatabase[K, V]{
		data: databaseSnapshot,
	}
}

func (database *inMemoryDatabase[K, V]) createDatabaseSnapshot() map[K]V {
	snapshot := make(map[K]V, len(database.data))
	for key, value := range database.data {
		snapshot[key] = value
	}
	return snapshot
}

func (database *inMemoryDatabase[K, V]) Commit() {
	if database.transaction == nil {
		return
	}
	if database.transaction.transaction != nil {
		database.transaction.Commit()
		return
	}
	database.data = database.transaction.data

	database.transaction = nil
}

func (database *inMemoryDatabase[K, V]) Rollback() {
	if database.transaction == nil {
		return
	}
	if database.transaction.transaction != nil {
		database.transaction.Rollback()
		return
	}
	database.transaction = nil
}

func NewInMemoryDatabase[K comparable, V any]() domain.InMemoryDatabase[K, V] {
	return &inMemoryDatabase[K, V]{data: make(map[K]V)}
}
