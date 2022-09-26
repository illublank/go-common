package algo

import (
  "sort"
  "testing"
)

func TestSkipTest(t *testing.T) {
  list := NewSkipList[int](3)
  list.Add(1)
}

type A interface {
  comparable
  sort.Interface
}
