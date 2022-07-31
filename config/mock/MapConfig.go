package mock

import (
  "strings"

  "github.com/illublank/go-common/config"
  "github.com/illublank/go-common/typ/collection"
)

type MapConfig struct {
  config.Config
  prefix string
  m      collection.GoMap
}

func NewMapConfig(m collection.GoMap) *MapConfig {
  return &MapConfig{
    m: m,
  }
}

func (s *MapConfig) WithPrefix(p string) *MapConfig {

  return &MapConfig{
    prefix: s.prefix + p,
    m:      s.m,
  }
}

func (s *MapConfig) WithoutPrefix(p string) *MapConfig {
  return &MapConfig{
    m: s.m.Filter(
      func(k string, v interface{}) bool { return strings.HasPrefix(k, p) },
      func(k string) string { return strings.TrimPrefix(k, p) },
      nil,
    ),
  }
}

// Get todo
func (s *MapConfig) Get(name string) (interface{}, bool) {
  return s.m.Get(name)
}

// GetString todo
func (s *MapConfig) GetString(name string, _default string) string {
  return s.m.GetString(name, _default)
}

// GetInt todo
func (s *MapConfig) GetInt(name string, _default int) int {
  return s.m.GetInt(name, _default)
}

// GetInt64 todo
func (s *MapConfig) GetInt64(name string, _default int64) int64 {
  return s.m.GetInt64(name, _default)
}

// GetBool todo
func (s *MapConfig) GetBool(name string, _default bool) bool {
  return s.m.GetBool(name, _default)
}

// GetMap todo
func (s *MapConfig) GetMap(name string) collection.GoMap {
  return s.m.GetMap(name, collection.NewGoMap())
}

func (s *MapConfig) String() string {
  return s.m.String()
}
