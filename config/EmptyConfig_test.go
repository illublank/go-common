package config_test

import (
  "fmt"
  "testing"

  "github.com/illublank/go-common/config"
)

func TestEmptyConfig(t *testing.T) {
  emptyConfig := config.EmptyConfig
  fmt.Println(emptyConfig)
}
