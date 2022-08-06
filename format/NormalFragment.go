package format

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
