package config

type emptyConfig struct {
}

// WithPrefix todo
func (s *emptyConfig) WithPrefix(p string) Config {
  return &emptyConfig{}
}

func (s *emptyConfig) WithoutPrefix(p string) Config {
  return &emptyConfig{}
}

// Get todo
func (s *emptyConfig) Get(name string) (any, bool) {
  return nil, false
}

// GetString todo
func (s *emptyConfig) GetString(name string, _default string) string {
  return _default
}

// GetInt todo
func (s *emptyConfig) GetInt(name string, _default int) int {
  return _default
}

// GetInt64 todo
func (s *emptyConfig) GetInt64(name string, _default int64) int64 {
  return _default
}

// GetBool todo
func (s *emptyConfig) GetBool(name string, _default bool) bool {
  return _default
}

func (s *emptyConfig) String() string {
  return "<EmptyConfig>"
}

var EmptyConfig = &emptyConfig{}
