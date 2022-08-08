package syn

import (
	"sync"
)

type WaitFuncs[T any] struct {
  funcs []T
}

func NewWaitFuncs[T any]() *WaitFuncs[T] {
  ret := &WaitFuncs[T]{
    funcs: []T{},
  }
  return ret
}

func (s *WaitFuncs[T]) Add(f T) {
  s.funcs = append(s.funcs, f)
}

func (s *WaitFuncs[T]) ForEach(handler func(T)) {
  wg := &sync.WaitGroup{}
  wg.Add(len(s.funcs))
  for _, item := range s.funcs {
    go func(f T) {
      handler(f)
      wg.Done()
    }(item)
  }
  wg.Wait()
}