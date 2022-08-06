package keyvalue

import (
  "bytes"
  "fmt"
  "io"

  "github.com/illublank/go-common/format"
)

type KeyFragment struct {
  format.Fragment
  orignal []byte
  replace []byte
}

func (s *KeyFragment) Bytes() []byte {
  return s.replace
}

func (s *KeyFragment) Replace(replace []byte) {
  s.replace = replace
}

func (s *KeyFragment) Reset() {
  s.replace = s.orignal
}

func NewKeyFragment(orignal []byte) *KeyFragment {
  return &KeyFragment{
    orignal: orignal,
    replace: orignal,
  }
}

type KeyFragmentMap map[string]*KeyFragment

func (s KeyFragmentMap) Reset() {
  for _, v := range s {
    v.Reset()
  }
}

func (s KeyFragmentMap) Replace(m map[string]any) {
  for k, v := range s {
    if val, exists := m[k]; exists {
      v.Replace([]byte(fmt.Sprintf("%v", val)))
    }

  }
}

type FragmentList struct {
  io.WriterTo
  list        []format.Fragment
  keyFragment KeyFragmentMap
}

func (s *FragmentList) AddNormal(bs []byte) {
  nbs := make([]byte, len(bs))
  copy(nbs, bs)
  s.list = append(s.list, format.NewNormalFragment(nbs))
}

func (s *FragmentList) AddKey(bs []byte) {
  key := string(bs[2 : len(bs)-1])
  if fragment, exists := s.keyFragment[key]; exists {
    s.list = append(s.list, fragment)
  } else {
    nbs := make([]byte, len(bs))
    copy(nbs, bs)
    fragment := NewKeyFragment(nbs)
    s.keyFragment[key] = fragment
    s.list = append(s.list, fragment)
  }
}

func (s *FragmentList) GetReplaceMap() KeyFragmentMap {
  return s.keyFragment
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
    keyFragment: map[string]*KeyFragment{},
  }
}
