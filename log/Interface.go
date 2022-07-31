package log

// Logger todo
type Logger interface {
  Debug(...interface{})
  Warn(...interface{})
  Info(...interface{})
  Error(...interface{})
  SetLevel(Level) Logger
  // SetFormatter(Formatter) Logger
}
