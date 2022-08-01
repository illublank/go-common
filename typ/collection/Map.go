package collection

import (
  "encoding/json"
)

type GoMap map[string]any

func NewGoMap() GoMap {
  return GoMap(map[string]any{})
}

func (s GoMap) Get(key string) (any, bool) {
  val, exists := s[key]
  return val, exists
}

func (s GoMap) GetString(key string, def string) string {
  if val, exists := s[key]; exists {
    return val.(string)
  }
  return def
}

func (s GoMap) GetInt(key string, def int) int {
  if val, exists := s[key]; exists {
    return val.(int)
  }
  return def
}

func (s GoMap) GetInt64(key string, def int64) int64 {
  if val, exists := s[key]; exists {
    return val.(int64)
  }
  return def
}

func (s GoMap) GetFloat(key string, def float64) float64 {
  if val, exists := s[key]; exists {
    return val.(float64)
  }
  return def
}

func (s GoMap) GetBool(key string, def bool) bool {
  if val, exists := s[key]; exists {
    return val.(bool)
  }
  return def
}

func (s GoMap) GetMap(key string, def GoMap) GoMap {
  if val, exists := s[key]; exists {
    return val.(GoMap)
  }
  return def
}

func (s GoMap) Put(key string, val any) GoMap {
  s[key] = val
  return s
}

func (s GoMap) String() string {
  bs, _ := json.Marshal(s)
  return string(bs)
}

func (s GoMap) Filter(filterFunc func(k string, v any) bool, keyHandleFunc func(k string) string, valHandleFunc func(k any) any) GoMap {
  m := make(GoMap)
  for k, v := range s {
    if filterFunc(k, v) {
      if keyHandleFunc != nil {
        k = keyHandleFunc(k)
      }
      if valHandleFunc != nil {
        v = valHandleFunc(v)
      }
      m.Put(k, v)
    }
  }
  return m
}
