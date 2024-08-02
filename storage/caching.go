package storage

// Caching adalah implementasi dari Storage yang menyimpan data dalam memori
type CachingData struct {
	data string
}

func (s *CachingData) Save(data string) error {
	s.data = data
	return nil
}

func (s *CachingData) Load() (string, error) {
	return s.data, nil
}
