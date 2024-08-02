package service

import "golang-dependensi-inject/storage"

// NewInMemoryStorage menyediakan instance dari InMemoryStorage
func NewCachingData() *storage.CachingData {
	return &storage.CachingData{}
}

// NewDatabaseStorage menyediakan instance dari DatabaseStorage
func NewDatabaseStorage() *storage.DatabaseStorage {
	return &storage.DatabaseStorage{}
}
