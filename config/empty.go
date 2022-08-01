package config

type EmptyConfig struct {
}

// WithPrefix todo
func (s *EmptyConfig) WithPrefix(p string) Config {
  return &EmptyConfig{}
}

func (s *EmptyConfig) WithoutPrefix(p string) Config {
  return &EmptyConfig{}
}

// Get todo
func (s *EmptyConfig) Get(name string) (any, bool) {
  return nil, false
}

// GetString todo
func (s *EmptyConfig) GetString(name string, _default string) string {
  return _default
}

// GetInt todo
func (s *EmptyConfig) GetInt(name string, _default int) int {
  return _default
}

// GetInt64 todo
func (s *EmptyConfig) GetInt64(name string, _default int64) int64 {
  return _default
}

// GetBool todo
func (s *EmptyConfig) GetBool(name string, _default bool) bool {
  return _default
}

func (s *EmptyConfig) String() string {
  return "<EmptyConfig>"
}

var Empty = &EmptyConfig{}
