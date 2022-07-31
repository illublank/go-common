package log_test

import (
  "fmt"
  "testing"

  "github.com/illublank/go-common/log"
)

func TestCommonLogger(t *testing.T) {
  logger := log.NewCommonLogger("testlogger")
  logger.SetLevel(log.Info)

  logger.Debug("debug", "haha")
  logger.Info("info", "abc", 3)
  logger.Warn("warn", "abc", 3)
  logger.Error("error", "abc", 3)
  fmt.Println(fmt.Sprintln("testing", "abc", 3))
}
