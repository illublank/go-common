package ptr_test

import (
  "fmt"
  "testing"

  "github.com/illublank/go-common/typ/ptr"
)

func TestPtr(t *testing.T) {
  str := ptr.BoolPtr(true)

  fmt.Println(str)
}
