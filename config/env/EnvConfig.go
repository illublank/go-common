package env

import (
  "fmt"
  "os"
  "strings"

  "github.com/illublank/go-common/config"
  "github.com/illublank/go-common/typ/collection"
)

var (
// reg = regexp.MustCompile(`([a-zA-Z0-9][a-zA-Z0-9_]*?)_(\d+)`)
)

// EnvConfig todo
type EnvConfig struct {
  config.Config
  prefix string
  m      collection.GoMap
}

// LoadAll todo
func LoadAll() *EnvConfig {
  return LoadAllWithoutPrefix("")
}

// LoadAllWithoutPrefix todo
func LoadAllWithoutPrefix(prefix string) *EnvConfig {
  envs := os.Environ()
  config := collection.NewGoMap()
  for _, item := range envs {
    if strings.HasPrefix(item, prefix) {
      pair := strings.SplitN(strings.TrimPrefix(item, prefix), "=", 2)
      if len(pair) < 2 {
        fmt.Println("envConfig item error:" + item)
      } else {
        config.Put(pair[0], pair[1])
        // config.Set(pair[0], pair[1])
      }
    }
  }
  return &EnvConfig{
    m: config,
  }
}

// WithPrefix todo
func (s *EnvConfig) WithPrefix(p string) config.Config {
  return &EnvConfig{
    prefix: s.prefix + p,
    m:      s.m,
  }
}

func (s *EnvConfig) WithoutPrefix(p string) config.Config {
  return &EnvConfig{
    m: s.m.Filter(
      func(k string, v any) bool { return strings.HasPrefix(k, p) },
      func(k string) string { return strings.TrimPrefix(k, p) },
      nil,
    ),
  }
}

// Get todo
func (s *EnvConfig) Get(name string) (any, bool) {
  return s.m.Get(name)
}

// GetString todo
func (s *EnvConfig) GetString(name string, _default string) string {
  return s.m.GetString(name, _default)
}

// GetInt todo
func (s *EnvConfig) GetInt(name string, _default int) int {
  return s.m.GetInt(name, _default)
}

// GetInt64 todo
func (s *EnvConfig) GetInt64(name string, _default int64) int64 {
  return s.m.GetInt64(name, _default)
}

// GetBool todo
func (s *EnvConfig) GetBool(name string, _default bool) bool {
  return s.m.GetBool(name, _default)
}

// GetMap todo
func (s *EnvConfig) GetMap(name string) collection.GoMap {
  return s.m.GetMap(name, collection.NewGoMap())
}

func (s *EnvConfig) String() string {
  return s.m.String()
}
