package json_test

import (
  "fmt"
  "testing"

  "github.com/illublank/go-common/log"
  "github.com/illublank/go-common/log/json"
)

func TestJsonLogger(t *testing.T) {
  var logger log.Logger = json.NewJsonLogger("testlogger")
  logger.SetLevel(log.Info)

  logger.Debug("debug", "haha")
  logger.Info("info", "abc", 3)
  logger.Warn("warn", "abc", 3)
  logger.Error("error", "abc", 3)
  fmt.Println(fmt.Sprintln("testing", "abc", 3))
}
