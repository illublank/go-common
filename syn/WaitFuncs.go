package syn

import (
	"reflect"
	"sync"
)

type funcWrapper struct {
  FuncRV reflect.Value
  wg     *sync.WaitGroup
}

func (s *funcWrapper) newFuncRV(args []reflect.Value) []reflect.Value {
  
  s.wg.Add(1)
  ret := s.FuncRV.Call(args)
  // wrap key action
  s.wg.Done()
  return ret
}

func WrapWaitGroupRV (wg *sync.WaitGroup, funcRV reflect.Value) reflect.Value {
  fw := &funcWrapper{
    FuncRV: funcRV,
    wg:     wg,
  }
  return reflect.MakeFunc(funcRV.Type(), fw.newFuncRV)
}

func WrapWaitGroup[T any](wg *sync.WaitGroup, f T) T {
  funcRV := reflect.ValueOf(f)
  return WrapWaitGroupRV(wg, funcRV).Interface().(T)
}


type WaitFuncs[T any] struct {
  WG *sync.WaitGroup
  funcRVs []reflect.Value
  Funcs []T
}

func NewWaitFuncs[T any]() *WaitFuncs[T] {
  ret := &WaitFuncs[T]{
    WG: &sync.WaitGroup{},
    funcRVs : []reflect.Value{},
    Funcs: []T{},
  }
  return ret
}

func (s *WaitFuncs[T]) Add(f T) {
  newFuncRV := WrapWaitGroupRV(s.WG, reflect.ValueOf(f))
  s.funcRVs = append(s.funcRVs, newFuncRV)
  s.Funcs = append(s.Funcs, newFuncRV.Interface().(T))
}

func (s *WaitFuncs[T]) ForEach(handler func(T)) {
  for _, f := range s.Funcs {
    s.WG.Add(1)
    handler(f)
  }
}

func (s *WaitFuncs[T]) Wait() {
  s.WG.Wait()
}