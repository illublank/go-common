package ordered

import (
  "bytes"
  "fmt"
  "io"

  "github.com/illublank/go-common/format"
  "github.com/illublank/go-common/maths"
)

/*
  token {abc}
  expr abc
*/
type OrderedFragment struct {
  format.Fragment
  token   string
  expr    string
  orignal []byte
  replace []byte
}

func (s *OrderedFragment) Bytes() []byte {
  return s.replace
}

func (s *OrderedFragment) Replace(replace []byte) {
  s.replace = replace
}

func (s *OrderedFragment) Reset() {
  s.replace = s.orignal
}

func NewOrderedFragment(orignal []byte) *OrderedFragment {
  token := string(orignal)
  return &OrderedFragment{
    token:   token,
    expr:    token[1 : len(token)-1],
    orignal: orignal,
    replace: orignal,
  }
}

type OrderedFragmentList []*OrderedFragment

func (s OrderedFragmentList) Reset() {
  for _, v := range s {
    v.Reset()
  }
}

func (s OrderedFragmentList) Replace(m ...any) {
  for i := 0; i < maths.Min(len(s), len(m)); i++ {
    item := s[i]
    token := item.token
    switch {
    case item.orignal[1] == '%':
      item.Replace([]byte(fmt.Sprintf(item.expr, m[i])))
    case token == "{...}":
      str := fmt.Sprintln(m[i:]...)
      item.Replace([]byte(str[:len(str)-1]))
    default:
      item.Replace([]byte(fmt.Sprint(m[i])))
    }
  }
}

type FragmentList struct {
  io.WriterTo
  list        []format.Fragment
  orderedlist OrderedFragmentList
}

func (s *FragmentList) AddNormal(bs []byte) {
  nbs := make([]byte, len(bs))
  copy(nbs, bs)
  s.list = append(s.list, format.NewNormalFragment(nbs))
}

func (s *FragmentList) AddOrdered(bs []byte) {
  nbs := make([]byte, len(bs))
  copy(nbs, bs)
  fragment := NewOrderedFragment(nbs)
  s.orderedlist = append(s.orderedlist, fragment)
  s.list = append(s.list, fragment)
}

func (s *FragmentList) String() string {
  buf := &bytes.Buffer{}
  for _, item := range s.list {
    bs := item.Bytes()
    buf.Write(bs)
  }
  return buf.String()
}

func (s *FragmentList) WriteTo(w io.Writer) (int64, error) {
  n := int64(0)
  for _, item := range s.list {
    n0, err := w.Write(item.Bytes())
    if err != nil {
      return n + int64(n0), err
    }
    n += int64(n0)
  }
  return n, nil
}

func NewFragmentList() *FragmentList {
  return &FragmentList{
    list:        []format.Fragment{},
    orderedlist: []*OrderedFragment{},
  }
}
