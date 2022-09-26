package algo

import "container/list"

type LevelList struct {
  level int
  list  *list.List
}

func NewLevelList(level int) *LevelList {

  return &LevelList{
    level: level,
    list:  list.New(),
  }
}

type SkipList[T comparable] struct {
  list []*LevelList
}

func NewSkipList[T comparable](level int) *SkipList[T] {
  list := make([]*LevelList, level)
  for i := range list {
    list[i] = NewLevelList(i)
  }
  return &SkipList[T]{
    list: list,
  }
}

func (s *SkipList[T]) Add(obj T) {
  for _, l := range s.list {
    l.list.PushBack(obj)
  }
}