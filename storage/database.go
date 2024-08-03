package storage

import "fmt"

// DatabaseStorage adalah implementasi dari Storage yang menyimpan data dalam database
type DatabaseStorage struct {
	data string
}

func (s *DatabaseStorage) Save(data string) error {
	s.data = data
	return nil
}

func (s *DatabaseStorage) Load() (string, error) {
	return s.data, nil
}

// implementation struct field provider
type Database struct {
	Name string
}

func NewDatabase(dbName string) *Database {
	return &Database{Name: dbName}
}

func NewDatabaseWithCleanUp(name string) (*Database, func(), error) {
	db := &Database{Name: name}

	// Cleanup function to close the database connection
	cleanup := func() {
		fmt.Println("Closing database connection")
	}

	return db, cleanup, nil
}
