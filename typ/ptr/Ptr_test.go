package ptr_test

import (
  "fmt"
  "reflect"
  "testing"

  "github.com/illublank/go-common/log"
  "github.com/illublank/go-common/typ/ptr"
)

func TestPtr(t *testing.T) {
  b := ptr.Ptr(true)

  fmt.Println(b)
  logger := log.NewCommonLogger("testLogger")

  s := ptr.Ptr("abc")

  typ := reflect.TypeOf(s)
  logger.Info(typ.PkgPath(), typ, ptr.StringTyp.PkgPath(), ptr.StringTyp)

  logger.Info(ptr.StringTyp, ptr.IntTyp)
}
