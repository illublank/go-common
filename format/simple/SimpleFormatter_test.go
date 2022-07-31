package simple_test

import (
  "bytes"
  "fmt"
  "testing"

  "github.com/illublank/go-common/format/simple"
)

func TestXxx(t *testing.T) {
  formatter := simple.NewSimpleFormatter("abc${a${a}${a}}b}")

  fmt.Println(formatter.Format(map[string]string{"a": "123"}))
}

func TestBuf(t *testing.T) {
  buf := &bytes.Buffer{}

  buf.Write([]byte("abc"))

  fmt.Println(buf.Bytes())

  buf.Write([]byte("${a}"))
  // buf.Reset()

  fmt.Println(buf.Bytes())
}
