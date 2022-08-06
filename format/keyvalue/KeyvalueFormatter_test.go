package keyvalue_test

import (
  "bytes"
  "fmt"
  "testing"

  "github.com/illublank/go-common/format/keyvalue"
)

func TestXxx(t *testing.T) {
  formatter := keyvalue.NewKeyvalueFormatter("abc${a${a}${a}}b}")

  buf := &bytes.Buffer{}
  formatter.Params(map[string]any{"a": "123"}).WriteTo(buf)

  fmt.Println(buf.String())
}

func TestBuf(t *testing.T) {
  buf := &bytes.Buffer{}

  buf.Write([]byte("abc"))

  fmt.Println(buf.Bytes())

  buf.Write([]byte("${a}"))
  // buf.Reset()

  fmt.Println(buf.Bytes())
}
