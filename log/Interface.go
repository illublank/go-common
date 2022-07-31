package log

// Logger todo
type Logger interface {
  Debug(...interface{})
  Debugf(string, ...interface{})
  Warn(...interface{})
  Warnf(string, ...interface{})
  Info(...interface{})
  Infof(string, ...interface{})
  Error(...interface{})
  Errorf(string, ...interface{})
  SetLevel(Level) Logger
  // SetFormatter(Formatter) Logger
}
