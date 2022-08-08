package syn_test

import (
  "fmt"
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

func TestWaitFuncs2(t *testing.T) {
  wf := syn.NewWaitFuncs[func(int, int) int]()

  wf.Add(add)
  wf.Add(add2)

  wf.ForEach(func(f func(int, int) int) {
    f(3, 2)
  })

  wf2 := syn.NewWaitFuncs[func(int) int]()
  wf2.Add(double)

  wf2.ForEach(func(f func(int) int) {
    f(3)
  })
}
