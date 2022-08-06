package format

import "io"

// Formatter todo
type Formatter interface {
  io.WriterTo
  Format(any) string
  Params(params any) Formatter
}

type Fragment interface {
  Bytes() []byte
}
