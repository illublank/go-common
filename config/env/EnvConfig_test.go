package env_test

import (
  "fmt"
  "testing"

  "github.com/illublank/go-common/config/env"
)

func Test(t *testing.T) {
  conf := env.LoadAllWithoutPrefix("abc_")
  fmt.Println(conf)
}
