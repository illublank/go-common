package log

import (
  "fmt"
  "time"

  "github.com/illublank/go-common/format"
  "github.com/illublank/go-common/format/simple"
)

var CommonFormatter = simple.NewSimpleFormatter("${time} [${level}] ${name} : ${msg}")

type CommonLogger struct {
  Logger
  name      string
  formatter format.Formatter
  level     Level
}

func NewCommonLogger(name string) *CommonLogger {
  return &CommonLogger{
    name:      name,
    formatter: CommonFormatter,
    level:     Info,
  }
}

func (s *CommonLogger) SetLevel(level Level) Logger {
  s.level = level
  return s
}

func (s *CommonLogger) Log(level Level, msg ...interface{}) {
  newMsg := fmt.Sprintln(msg...)
  fmt.Println(s.formatter.Format(map[string]interface{}{"time": time.Now().Format("2006-01-02 15:04:05"), "level": fmt.Sprintf("%5s", level), "name": s.name, "msg": newMsg[:len(newMsg)-1]}))
}

func (s *CommonLogger) Error(msg ...interface{}) {
  if s.level <= Error {
    s.Log(Error, msg...)
  }
}

func (s *CommonLogger) Warn(msg ...interface{}) {
  if s.level <= Warn {
    s.Log(Warn, msg...)
  }
}

func (s *CommonLogger) Info(msg ...interface{}) {
  if s.level <= Info {
    s.Log(Info, msg...)
  }
}

func (s *CommonLogger) Debug(msg ...interface{}) {
  if s.level <= Debug {
    s.Log(Debug, msg...)
  }
}
