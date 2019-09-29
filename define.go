package xtype

type Map map[T]T
type MapInterface map[interface{}]T
type MapString map[string]T
type MapInt64 map[int64]T
type MapInt32 map[int32]T
type MapInt map[int64]T
type Hash map[string]interface{}
type MapStringInterface map[string]interface{}
type Slice []T
type SliceInterface []interface{}
type SliceString []string
type SliceInt64 []int64

type StringInterface interface {
	String() string
}

type ErrorInterface interface {
	Error() string
}