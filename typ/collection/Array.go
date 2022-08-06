package collection

type Array[T comparable] []T

func (s Array[T]) Len() int { return len(s) }

func (s Array[T]) IndexOf(item T) int {
  for i, v := range s {
    if v == item {
      return i
    }
  }
  return -1
}

func ParseArray[T comparable](arr []T) Array[T] {
  return Array[T](arr)
}