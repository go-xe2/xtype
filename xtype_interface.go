package xtype

type T interface {
	Interface() interface{}
	String() string
	Float64() float64
	Float32() float32
	Int64() int64
	Int() int
	Int32() int32
	Int16() int16
	Int8() int8
	Uint64() uint64
	Uint() uint
	Uint32() uint32
	Uint16() uint16
	Uint8() uint8
	Bool() bool
	Slice() Slice
	SliceInterface() []interface{}
	SliceMapStringInterface() []map[string]interface{}
	SliceHash() []Hash
	SliceString() []string
	SliceInt64() []int64
	Map() Map
	MapInterface() MapInterface
	MapString() MapString
	MapInt64() MapInt64
	MapStringInterface() map[string]interface{}
	Hash() Hash
}

