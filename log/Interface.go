package log

// Logger todo
type Logger interface {
  Debug(...any)
  Debugf(string, ...any)
  Warn(...any)
  Warnf(string, ...any)
  Info(...any)
  Infof(string, ...any)
  Error(...any)
  Errorf(string, ...any)
  SetLevel(Level) Logger
  // SetFormatter(Formatter) Logger
}
