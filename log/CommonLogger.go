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

func (s *CommonLogger) Log(level Level, fstr string, msg ...any) {
  var newMsg string
  if len(fstr) == 0 {
    newMsg = fmt.Sprintln(msg...)
    newMsg = newMsg[:len(newMsg)-1]
  } else {
    newMsg = fmt.Sprintf(fstr, msg...)
  }
  fmt.Println(s.formatter.Format(map[string]any{"time": time.Now().Format("2006-01-02 15:04:05"), "level": fmt.Sprintf("%5s", level), "name": s.name, "msg": newMsg}))
}

func (s *CommonLogger) Error(msg ...any) {
  if s.level <= Error {
    // s.Log(Error, msg...)
    s.Errorf("", msg...)
  }
}

func (s *CommonLogger) Warn(msg ...any) {
  if s.level <= Warn {
    // s.Log(Warn, msg...)
    s.Warnf("", msg...)
  }
}

func (s *CommonLogger) Info(msg ...any) {
  if s.level <= Info {
    // s.Log(Info, msg...)
    s.Infof("", msg...)
  }
}

func (s *CommonLogger) Debug(msg ...any) {
  if s.level <= Debug {
    // s.Log(Debug, msg...)
    s.Debugf("", msg...)
  }
}

func (s *CommonLogger) Errorf(fstr string, msg ...any) {
  if s.level <= Error {
    s.Log(Error, fstr, msg...)
  }
}

func (s *CommonLogger) Warnf(fstr string, msg ...any) {
  if s.level <= Warn {
    s.Log(Warn, fstr, msg...)
  }
}

func (s *CommonLogger) Infof(fstr string, msg ...any) {
  if s.level <= Info {
    s.Log(Info, fstr, msg...)
  }
}

func (s *CommonLogger) Debugf(fstr string, msg ...any) {
  if s.level <= Debug {
    s.Log(Debug, fstr, msg...)
  }
}
