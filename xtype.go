package xtype

import (
	"reflect"
)

type XType struct {
	val interface{}
}

var _ T = &XType{}

func New(o interface{}) XType {
	switch o.(type) {
	case T:
		return XType{o.(T).Interface()}
	default:
		return XType{val: o}
	}
}

func (t XType) Interface() interface{} {
	return t.val
}

func (t XType) String() string {
	return ParseString(t.val)
}

func (t XType) Float64() float64 {
	return ParseFloat64(t.val)
}
func (t XType) Float32() float32 {
	return ParseFloat32(t.val)
}

func (t XType) Int64() int64 {
	return ParseInt64(t.val)
}
func (t XType) Int() int {
	return ParseInt(t.val)
}
func (t XType) Int32() int32 {
	return int32(ParseInt64(t.val))
}
func (t XType) Int16() int16 {
	return int16(ParseInt64(t.val))
}
func (t XType) Int8() int8 {
	return int8(ParseInt64(t.val))
}

func (t XType) Uint64() uint64 {
	return ParseUint64(t.val)
}
func (t XType) Uint() uint {
	return ParseUint(t.val)
}
func (t XType) Uint32() uint32 {
	return uint32(ParseUint(t.val))
}
func (t XType) Uint16() uint16 {
	return uint16(ParseUint(t.val))
}
func (t XType) Uint8() uint8 {
	return uint8(ParseUint(t.val))
}

func (t XType) Bool() bool {
	return ParseBool(t.val)
}

func (t XType) Slice() Slice {
	ref := reflect.Indirect(reflect.ValueOf(t.val))
	l := ref.Len()
	v := ref.Slice(0, l)
	var res = Slice{}
	for i := 0; i < l; i++ {
		res = append(res, New(v.Index(i).Interface()))
	}
	return res
}

func (t XType) SliceInterface() []interface{} {
	s := t.Slice()
	var res = []interface{}{}
	for _, item := range s {
		res = append(res, item.Interface())
	}
	return res
}

func (t XType) SliceMapStringInterface() []map[string]interface{} {
	s := t.Slice()
	var res = []map[string]interface{}{}
	for _, item := range s {
		res = append(res, item.MapStringInterface())
	}
	return res
}

func (t XType) SliceHash() []Hash {
	s := t.Slice()
	var results []Hash
	for _, item := range s {
		results = append(results, item.Hash())
	}
	return results
}

func (t XType) SliceString() []string {
	s := t.Slice()
	var res = []string{}
	for _, item := range s {
		res = append(res, item.String())
	}
	return res
}

func (t XType) SliceInt64() []int64 {
	s := t.Slice()
	var res = []int64{}
	for _, item := range s {
		res = append(res, item.Int64())
	}
	return res
}

func (t XType) Map() Map {
	ref := reflect.Indirect(reflect.ValueOf(t.val))
	var res = make(Map)
	keys := ref.MapKeys()
	for _, item := range keys {
		res[New(item.Interface())] = New(ref.MapIndex(item).Interface())
	}
	return res
}

func (t XType) MapInterface() MapInterface {
	m := t.Map()
	var res = make(MapInterface)
	for k, v := range m {
		res[k.Interface()] = v
	}
	return res
}

func (t XType) MapString() MapString {
	m := t.Map()
	var res = make(MapString)
	for k, v := range m {
		res[k.String()] = v
	}
	return res
}

func (t XType) MapStringInterface() map[string]interface{} {
	m := t.Map()
	var res = make(map[string]interface{})
	for k, v := range m {
		res[k.String()] = v.Interface()
	}
	return res
}

func (t XType) Hash() Hash {
	m := t.Map()
	var results = make(Hash)
	for k, v := range m {
		results[k.String()] = v
	}
	return results
}

func (t XType) MapInt64() MapInt64 {
	m := t.Map()
	var res = make(MapInt64)
	for k, v := range m {
		res[k.Int64()] = v
	}
	return res
}
