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
  wg *sync.WaitGroup
  funcs []reflect.Value
  Funcs []T
}

func NewWaitFuncs[T any]() *WaitFuncs[T] {
  ret := &WaitFuncs[T]{
    wg: &sync.WaitGroup{},
    funcs : []reflect.Value{},
  }
  return ret
}

func (s *WaitFuncs[T]) Add(f T) {
  s.wg.Add(1)
  newFuncRV := WrapWaitGroupRV(s.wg, reflect.ValueOf(f))
  s.funcs = append(s.funcs, newFuncRV)
  s.Funcs = append(s.Funcs, newFuncRV.Interface().(T))
}

func (s *WaitFuncs[T]) Wait() {
  s.wg.Wait()
}