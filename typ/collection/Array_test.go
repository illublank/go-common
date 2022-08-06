package collection_test

import (
  "fmt"
  "reflect"
  "testing"
)

func TestArray(t *testing.T) {

  // a := reflect.TypeOf(1)
  b := reflect.TypeOf("a")
  // c := reflect.TypeOf(0.1)
  d := reflect.TypeOf("c")
  // var arr Array = collection.Array{a, b, c}

  fmt.Println(reflect.TypeOf(b) == reflect.TypeOf(d))
}
