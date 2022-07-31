package ptr

type String *string

func StringPtr(v string) String {
  return &v
}

type Int *int

func IntPtr(v int) Int {
  return &v
}

type Int8 *int8

func Int8Ptr(v int8) Int8 {
  return &v
}

type Int16 *int16

func Int16Ptr(v int16) Int16 {
  return &v
}

type Int32 *int32

func Int32Ptr(v int32) Int32 {
  return &v
}

type Int64 *int64

type Uint *uint

type Uint8 *uint8

type Uint16 *uint16

type Uint32 *uint32

type Uint64 *uint64

type Float32 *float32

type Float64 *float64

type Bool *bool

func BoolPtr(v bool) Bool {
  return &v
}
