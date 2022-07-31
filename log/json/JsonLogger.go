package json

import (
  "encoding/json"
  "fmt"
  "time"

  "github.com/illublank/go-common/format"
  "github.com/illublank/go-common/log"
)

type JsonFormatter struct {
  log.Logger
  format.Formatter
}

func (s *JsonFormatter) Format(obj interface{}) string {
  bs, _ := json.Marshal(obj)
  return string(bs)
}

type JsonLogger struct {
  log.Logger
  name      string
  formatter format.Formatter
  level     log.Level
}

func NewJsonLogger(name string) *JsonLogger {
  return &JsonLogger{
    name:      name,
    formatter: &JsonFormatter{},
    level:     log.Info,
  }
}

func (s *JsonLogger) SetLevel(level log.Level) log.Logger {
  s.level = level
  return s
}

func (s *JsonLogger) Log(level log.Level, msg ...interface{}) {
  newMsg := fmt.Sprintln(msg...)
  fmt.Println(s.formatter.Format(map[string]interface{}{"time": time.Now().Format("2006-01-02 15:04:05"), "level": fmt.Sprintf("%5s", level), "name": s.name, "msg": newMsg[:len(newMsg)-1]}))
}

func (s *JsonLogger) Error(msg ...interface{}) {
  if s.level <= log.Error {
    s.Log(log.Error, msg...)
  }
}

func (s *JsonLogger) Warn(msg ...interface{}) {
  if s.level <= log.Warn {
    s.Log(log.Warn, msg...)
  }
}

func (s *JsonLogger) Info(msg ...interface{}) {
  if s.level <= log.Info {
    s.Log(log.Info, msg...)
  }
}

func (s *JsonLogger) Debug(msg ...interface{}) {
  if s.level <= log.Debug {
    s.Log(log.Debug, msg...)
  }
}
