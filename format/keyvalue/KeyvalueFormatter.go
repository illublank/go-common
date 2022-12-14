package keyvalue

import (
  "bytes"
  "io"

  "github.com/illublank/go-common/format"
)

type KeyvalueFormatter struct {
  format.Formatter

  formatStr string

  list *FragmentList
}

func NewKeyvalueFormatter(formatStr string) *KeyvalueFormatter {
  fbs := []byte(formatStr)

  list := NewFragmentList()

  buf := &bytes.Buffer{}
  token := &bytes.Buffer{}
  for i := 0; i < len(fbs); i++ {
    b := fbs[i]
    if b == '$' {
      if token.Len() > 0 {
        buf.Write(token.Bytes())
      }
      if buf.Len() > 0 {
        list.AddNormal(buf.Bytes())
      }
      token.Reset()
      buf.Reset()
      token.WriteByte(b)
    } else if b == '}' {
      token.WriteByte(b)
      bs := token.Bytes()
      if len(bs) > 3 && bs[0] == '$' && bs[1] == '{' {
        if buf.Len() > 0 {
          list.AddNormal(buf.Bytes())
        }
        list.AddKey(token.Bytes())
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
  return &KeyvalueFormatter{
    formatStr: formatStr,
    list:      list,
  }
}

func (s *KeyvalueFormatter) Format(params any) string {
  s.list.keyFragment.Replace(params.(map[string]any))
  result := s.list.String()
  s.list.keyFragment.Reset()
  return result
}

func (s *KeyvalueFormatter) Params(params any) format.Formatter {
  s.list.keyFragment.Replace(params.(map[string]any))
  return s
}

func (s *KeyvalueFormatter) Reset() *KeyvalueFormatter {
  s.list.keyFragment.Reset()
  return s
}

// WriteTo(w Writer) (n int64, err error)
func (s *KeyvalueFormatter) WriteTo(w io.Writer) (int64, error) {
  n, err := s.list.WriteTo(w)
  s.list.keyFragment.Reset()
  return n, err
}
