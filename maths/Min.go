package maths

import "github.com/illublank/go-common/typ"

func Min[T typ.Ordered] (a T, b T) T {
  if a > b {
    return b
  }
  return a
}