package storage

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
