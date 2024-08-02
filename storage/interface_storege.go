package storage

// Storage adalah interface untuk menyimpan data
type Storage interface {
	Save(data string) error
	Load() (string, error)
}
