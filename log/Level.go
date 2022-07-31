package log

// Level todo
type Level int

const (
  _ Level = iota
  // Debug log level
  Debug
  // Info log level
  Info
  // Warn log level
  Warn
  // Error log level
  Error
)

func (s Level) String() string {
  switch s {
  case Error:
    return "ERROR"
  case Info:
    return "INFO"
  case Warn:
    return "WARN"
  case Debug:
    return "DEBUG"
  default:
    return ""
  }
}
