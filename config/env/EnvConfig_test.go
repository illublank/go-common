package env_test

import (
  "fmt"
  "testing"

  "github.com/illublank/go-common/config"
  "github.com/illublank/go-common/config/env"
)

func Test(t *testing.T) {
  var conf config.Config
  conf = env.LoadAllWithoutPrefix("abc_")
  fmt.Println(conf)
}
