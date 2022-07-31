package mock_test

import (
  "fmt"
  "testing"

  "github.com/illublank/go-common/config/mock"
  "github.com/illublank/go-common/typ/collection"
)

func TestMock(t *testing.T) {
  m := collection.GoMap(map[string]interface{}{
    "abc_bcd_cde": 1,
    "bcd_cde":     "abc",
  })
  config := mock.NewMapConfig(m)

  config = config.WithoutPrefix("abc_")

  fmt.Println(config)
}
