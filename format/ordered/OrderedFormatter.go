package ordered

import (
  "bytes"
  "io"

  "github.com/illublank/go-common/format"
)

type OrderedFormatter struct {
  format.Formatter

  formatStr string

  list *FragmentList
}

func NewOrderedFormatter(formatStr string) *OrderedFormatter {
  fbs := []byte(formatStr)

  list := NewFragmentList()

  buf := &bytes.Buffer{}
  token := &bytes.Buffer{}
  for i := 0; i < len(fbs); i++ {
    b := fbs[i]
    if b == '{' {
      if token.Len() == 1 && token.Bytes()[0] == '{' {
        buf.WriteByte('{')
        token.Reset()
      } else {
        if token.Len() > 0 {
          buf.Write(token.Bytes())
        }
        if buf.Len() > 0 {
          list.AddNormal(buf.Bytes())
        }
        token.Reset()
        buf.Reset()
        token.WriteByte(b)
      }
    } else if b == '}' {
      token.WriteByte(b)
      bs := token.Bytes()
      if len(bs) > 1 && bs[0] == '{' {
        if buf.Len() > 0 {
          list.AddNormal(buf.Bytes())
        }
        list.AddOrdered(token.Bytes())
        token.Reset()
        buf.Reset()
      }
    } else {
      token.WriteByte(b)
    }
  }

  if token.Len() > 0 {
    buf.Write(token.Bytes())
  }
  if buf.Len() > 0 {
    list.AddNormal(buf.Bytes())
  }
  return &OrderedFormatter{
    formatStr: formatStr,
    list:      list,
  }
}

func (s *OrderedFormatter) Format(params any) string {
  s.list.orderedlist.Replace(params.([]any)...)
  result := s.list.String()
  s.list.orderedlist.Reset()
  return result
}

func (s *OrderedFormatter) Params(params any) format.Formatter {
  s.list.orderedlist.Replace(params.([]any)...)
  return s
}

func (s *OrderedFormatter) Reset() *OrderedFormatter {
  s.list.orderedlist.Reset()
  return s
}

// WriteTo(w Writer) (n int64, err error)
func (s *OrderedFormatter) WriteTo(w io.Writer) (int64, error) {
  n, err := s.list.WriteTo(w)
  s.list.orderedlist.Reset()
  return n, err
}
