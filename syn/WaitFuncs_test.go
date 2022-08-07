package syn_test

import (
  "fmt"
  "sync"
  "testing"
  "time"

  "github.com/illublank/go-common/syn"
)

func add(a, b int) int {
  ret := a + b
  time.Sleep(time.Duration(ret) * time.Second)
  fmt.Println(ret)
  return ret
}

func add2(a, b int) int {
  ret := b
  time.Sleep(time.Duration(ret) * time.Second)
  fmt.Println(ret)
  return ret
}

func double(a int) int {
  ret := a * 2
  time.Sleep(time.Duration(ret) * time.Second)
  fmt.Println(ret)
  return ret
}

func TestWaitFuncs(t *testing.T) {
  wg := &sync.WaitGroup{}
  newAdd := syn.WrapWaitGroup(wg, add)
  newDouble := syn.WrapWaitGroup(wg, double)
  go newAdd(1, 9)

  go newDouble(2)

  wg.Wait()

}

func TestWaitFuncs2(t *testing.T) {
  wf := syn.NewWaitFuncs[func(int, int) int]()

  wf.Add(add)
  wf.Add(add2)

  for _, item := range wf.Funcs {
    go item(5, 5)
  }

  wf.Wait()
}
