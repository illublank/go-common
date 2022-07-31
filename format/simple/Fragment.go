package simple

import (
  "bytes"
  "fmt"
)

type Fragment interface {
  Bytes() []byte
}

type NormalFragment struct {
  Fragment
  orignal []byte
}

func (s *NormalFragment) Bytes() []byte {
  return s.orignal
}

func NewNormalFragment(orignal []byte) *NormalFragment {
  return &NormalFragment{
    orignal: orignal,
  }
}

type KeyFragment struct {
  Fragment
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

func (s KeyFragmentMap) Replace(m map[string]interface{}) {
  for k, v := range s {
    if val, exists := m[k]; exists {
      v.Replace([]byte(fmt.Sprintf("%v", val)))
    }

  }
}

type FragmentList struct {
  list        []Fragment
  keyFragment KeyFragmentMap
}

func (s *FragmentList) AddNormal(bs []byte) {
  nbs := make([]byte, len(bs))
  copy(nbs, bs)
  s.list = append(s.list, NewNormalFragment(nbs))
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

func NewFragmentList() *FragmentList {
  return &FragmentList{
    list:        []Fragment{},
    keyFragment: map[string]*KeyFragment{},
  }
}
