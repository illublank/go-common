package ptr

import (
	"reflect"
)

type String *string

var StringTyp = reflect.TypeOf(String(nil))

type Int *int

var IntTyp = reflect.TypeOf(Int(nil))

type Int8 *int8

var Int8Typ = reflect.TypeOf(Int8(nil))

type Int16 *int16

var Int16Typ = reflect.TypeOf(Int16(nil))

type Int32 *int32

var Int32Typ = reflect.TypeOf(Int32(nil))

type Int64 *int64

var Int64Typ = reflect.TypeOf(Int64(nil))

type Uint *uint

var UintTyp = reflect.TypeOf(Uint(nil))

type Uint8 *uint8

var Uint8Typ = reflect.TypeOf(Uint8(nil))

type Uint16 *uint16

var Uint16Typ = reflect.TypeOf(Uint16(nil))

type Uint32 *uint32

var Uint32Typ = reflect.TypeOf(Uint32(nil))

type Uint64 *uint64

var Uint64Typ = reflect.TypeOf(Uint64(nil))

type Float32 *float32

var Float32Typ = reflect.TypeOf(Float32(nil))

type Float64 *float64

var Float64Typ = reflect.TypeOf(Float64(nil))

type Bool *bool

var BoolTyp = reflect.TypeOf(Bool(nil))

func PkgPathAndName(t reflect.Type) string {
  return t.PkgPath() + "." + t.Name()
}

type TypArray []reflect.Type

func (s TypArray) Len() int { return len(s) }

func (s TypArray) IndexOf(item reflect.Type) int {
  for i, v := range s {
    if v == item {
      return i
    }
  }
  return -1
}

func (s TypArray) CanConvertFrom(t reflect.Type) bool {
  for _, v := range s {
    if t.ConvertibleTo(v) {
      return true
    }
  }
  return false
}

var PtrTypList = TypArray{StringTyp, IntTyp, Int8Typ, Int16Typ, Int32Typ, Int64Typ, UintTyp, Uint8Typ, Uint16Typ, Uint32Typ, Uint64Typ, Float32Typ, Float64Typ, BoolTyp}

func Ptr[T any](v T) *T {
  return &v
}

func AntiPtr(v any) any {
  rv := reflect.ValueOf(v)
  if PtrTypList.CanConvertFrom(rv.Type()) {
    return rv.Elem()
  }
  return v
}
