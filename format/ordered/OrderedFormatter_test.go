package ordered_test

import (
  "bytes"
  "fmt"
  "testing"

  "github.com/illublank/go-common/format/ordered"
)

func TestOrdered(t *testing.T) {
  formatter := ordered.NewOrderedFormatter("adfdsa{}bca{sdf{...}ad{%.2f}}")

  buf := &bytes.Buffer{}

  formatter.Params([]any{1, 23, 123.0, "Asb"})
  formatter.WriteTo(buf)

  fmt.Println(buf.String())
}
