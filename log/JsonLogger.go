package log

import (
  "encoding/json"
  "fmt"
  "time"

  "github.com/illublank/go-common/format"
)

type JsonFormatter struct {
  format.Formatter
}

func (s *JsonFormatter) Format(obj interface{}) string {
  bs, _ := json.Marshal(obj)
  return string(bs)
}

type JsonLogger struct {
  Logger
  name      string
  formatter format.Formatter
  level     Level
}

func NewJsonLogger(name string) *JsonLogger {
  return &JsonLogger{
    name:      name,
    formatter: &JsonFormatter{},
    level:     Info,
  }
}

func (s *JsonLogger) SetLevel(level Level) *JsonLogger {
  s.level = level
  return s
}

func (s *JsonLogger) Log(level Level, msg ...interface{}) {
  newMsg := fmt.Sprintln(msg...)
  fmt.Println(s.formatter.Format(map[string]interface{}{"time": time.Now().Format("2006-01-02 15:04:05"), "level": fmt.Sprintf("%5s", level), "name": s.name, "msg": newMsg[:len(newMsg)-1]}))
}

func (s *JsonLogger) Error(msg ...interface{}) {
  if s.level <= Error {
    s.Log(Error, msg...)
  }
}

func (s *JsonLogger) Warn(msg ...interface{}) {
  if s.level <= Warn {
    s.Log(Warn, msg...)
  }
}

func (s *JsonLogger) Info(msg ...interface{}) {
  if s.level <= Info {
    s.Log(Info, msg...)
  }
}

func (s *JsonLogger) Debug(msg ...interface{}) {
  if s.level <= Debug {
    s.Log(Debug, msg...)
  }
}
