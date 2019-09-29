package xtype

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

var (
	emptyStringMap = map[string]struct{}{
		"":      struct{}{},
		"0":     struct{}{},
		"off":   struct{}{},
		"false": struct{}{},
		"False": struct{}{},
	}
)

// Convert converts the variable <i> to the type <t>, the type <t> is specified by string.
// The unnecessary parameter <params> is used for additional parameter passing.
func Convert(i interface{}, t string, params ...interface{}) interface{} {
	switch t {
	case "int":
		return Int(i)
	case "int8":
		return Int8(i)
	case "int16":
		return Int16(i)
	case "int32":
		return Int32(i)
	case "int64":
		return Int64(i)
	case "uint":
		return Uint(i)
	case "uint8":
		return Uint8(i)
	case "uint16":
		return Uint16(i)
	case "uint32":
		return Uint32(i)
	case "uint64":
		return Uint64(i)
	case "float32":
		return Float32(i)
	case "float64":
		return Float64(i)
	case "bool":
		return Bool(i)
	case "string":
		return String(i)
	case "[]byte":
		return Bytes(i)
	case "[]int":
		return Ints(i)
	case "[]string":
		return Strings(i)

	case "Time", "time.Time":
		if len(params) > 0 {
			return Time(i, String(params[0]))
		}
		return Time(i)

	case "gtime.Time":
		if len(params) > 0 {
			return GTime(i, String(params[0]))
		}
		return *GTime(i)

	case "GTime", "*gtime.Time":
		if len(params) > 0 {
			return GTime(i, String(params[0]))
		}
		return GTime(i)

	case "Duration", "time.Duration":
		return Duration(i)
	default:
		return i
	}
}

func String(v interface{}) string {
	if v == nil {
		return ""
	}
	switch value := v.(type) {
	case int:
		return strconv.FormatInt(int64(value), 10)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(int64(value), 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(uint64(value), 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case []rune:
		return string(value)
	default:
		if f, ok := value.(StringInterface); ok {
			return f.String()
		} else if f, ok := value.(ErrorInterface); ok {
			return f.Error()
		} else {
			if jsonContent, err := json.Marshal(value); err != nil {
				return fmt.Sprint(value)
			} else {
				return string(jsonContent)
			}
		}
	}
}

// Int converts <i> to int.
func Int(i interface{}) int {
	if i == nil {
		return 0
	}
	if v, ok := i.(int); ok {
		return v
	}
	return int(Int64(i))
}

// Int8 converts <i> to int8.
func Int8(i interface{}) int8 {
	if i == nil {
		return 0
	}
	if v, ok := i.(int8); ok {
		return v
	}
	return int8(Int64(i))
}

// Int16 converts <i> to int16.
func Int16(i interface{}) int16 {
	if i == nil {
		return 0
	}
	if v, ok := i.(int16); ok {
		return v
	}
	return int16(Int64(i))
}

// Int32 converts <i> to int32.
func Int32(i interface{}) int32 {
	if i == nil {
		return 0
	}
	if v, ok := i.(int32); ok {
		return v
	}
	return int32(Int64(i))
}

func Int64(v interface{}) int64 {
	if v == nil {
		return 0
	}
	switch value := v.(type) {
	case int:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case uint:
		return int64(value)
	case uint8:
		return int64(value)
	case uint16:
		return int64(value)
	case uint32:
		return int64(value)
	case uint64:
		return int64(value)
	case float32:
		return int64(value)
	case float64:
		return int64(value)
	case bool:
		if value {
			return 1
		}
		return 0
	case []byte:
		return DecodeToInt64(value)
	default:
		s := String(value)
		// Hexadecimal
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, e := strconv.ParseInt(s[2:], 16, 64); e == nil {
				return v
			}
		}
		// Octal
		if len(s) > 1 && s[0] == '0' {
			if v, e := strconv.ParseInt(s[1:], 8, 64); e == nil {
				return v
			}
		}
		// Decimal
		if v, e := strconv.ParseInt(s, 10, 64); e == nil {
			return v
		}
		// Float64
		return int64(Float64(value))
	}
}

// Uint converts <i> to uint.
func Uint(i interface{}) uint {
	if i == nil {
		return 0
	}
	if v, ok := i.(uint); ok {
		return v
	}
	return uint(Uint64(i))
}

// Uint8 converts <i> to uint8.
func Uint8(i interface{}) uint8 {
	if i == nil {
		return 0
	}
	if v, ok := i.(uint8); ok {
		return v
	}
	return uint8(Uint64(i))
}

// Uint16 converts <i> to uint16.
func Uint16(i interface{}) uint16 {
	if i == nil {
		return 0
	}
	if v, ok := i.(uint16); ok {
		return v
	}
	return uint16(Uint64(i))
}

// Uint32 converts <i> to uint32.
func Uint32(i interface{}) uint32 {
	if i == nil {
		return 0
	}
	if v, ok := i.(uint32); ok {
		return v
	}
	return uint32(Uint64(i))
}

// Uint64 converts <i> to uint64.
func Uint64(i interface{}) uint64 {
	if i == nil {
		return 0
	}
	switch value := i.(type) {
	case int:
		return uint64(value)
	case int8:
		return uint64(value)
	case int16:
		return uint64(value)
	case int32:
		return uint64(value)
	case int64:
		return uint64(value)
	case uint:
		return uint64(value)
	case uint8:
		return uint64(value)
	case uint16:
		return uint64(value)
	case uint32:
		return uint64(value)
	case uint64:
		return value
	case float32:
		return uint64(value)
	case float64:
		return uint64(value)
	case bool:
		if value {
			return 1
		}
		return 0
	case []byte:
		return DecodeToUint64(value)
	default:
		s := String(value)
		// Hexadecimal
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, e := strconv.ParseUint(s[2:], 16, 64); e == nil {
				return v
			}
		}
		// Octal
		if len(s) > 1 && s[0] == '0' {
			if v, e := strconv.ParseUint(s[1:], 8, 64); e == nil {
				return v
			}
		}
		// Decimal
		if v, e := strconv.ParseUint(s, 10, 64); e == nil {
			return v
		}
		// Float64
		return uint64(Float64(value))
	}
}

func Float(i interface{}) float32  {
	return Float32(i)
}

// Float32 converts <i> to float32.
func Float32(i interface{}) float32 {
	if i == nil {
		return 0
	}
	switch value := i.(type) {
	case float32:
		return value
	case float64:
		return float32(value)
	case []byte:
		return DecodeToFloat32(value)
	default:
		v, _ := strconv.ParseFloat(String(i), 64)
		return float32(v)
	}
}

// Float64 converts <i> to float64.
func Float64(i interface{}) float64 {
	if i == nil {
		return 0
	}
	switch value := i.(type) {
	case float32:
		return float64(value)
	case float64:
		return value
	case []byte:
		return DecodeToFloat64(value)
	default:
		v, _ := strconv.ParseFloat(String(i), 64)
		return v
	}
}



// Byte converts <i> to byte.
func Byte(i interface{}) byte {
	if v, ok := i.(byte); ok {
		return v
	}
	return byte(Uint8(i))
}

// Bytes converts <i> to []byte.
func Bytes(i interface{}) []byte {
	if i == nil {
		return nil
	}
	switch value := i.(type) {
	case string:
		return []byte(value)
	case []byte:
		return value
	default:
		return Encode(i)
	}
}

// Rune converts <i> to rune.
func Rune(i interface{}) rune {
	if v, ok := i.(rune); ok {
		return v
	}
	return rune(Int32(i))
}

// Runes converts <i> to []rune.
func Runes(i interface{}) []rune {
	if v, ok := i.([]rune); ok {
		return v
	}
	return []rune(String(i))
}

func Bool(i interface{}) bool {
	if i == nil {
		return false
	}
	switch value := i.(type) {
	case bool:
		return value
	case string:
		if _, ok := emptyStringMap[value]; ok {
			return false
		}
		return true
	default:
		rv := reflect.ValueOf(i)
		switch rv.Kind() {
		case reflect.Ptr:
			return !rv.IsNil()
		case reflect.Map:
			fallthrough
		case reflect.Array:
			fallthrough
		case reflect.Slice:
			return rv.Len() != 0
		case reflect.Struct:
			return true
		default:
			s := String(i)
			if _, ok := emptyStringMap[s]; ok {
				return false
			}
			return true
		}
	}
}

// Ints converts <i> to []int.
func Ints(i interface{}) []int {
	if i == nil {
		return nil
	}
	if r, ok := i.([]int); ok {
		return r
	} else {
		array := make([]int, 0)
		switch value := i.(type) {
		case []string:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []int8:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []int16:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []int32:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []int64:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []uint:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []uint8:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []uint16:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []uint32:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []uint64:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []bool:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []float32:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []float64:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []interface{}:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case [][]byte:
			for _, v := range value {
				array = append(array, Int(v))
			}
		default:
			return []int{Int(i)}
		}
		return array
	}
}

// Strings converts <i> to []string.
func Strings(i interface{}) []string {
	if i == nil {
		return nil
	}
	if r, ok := i.([]string); ok {
		return r
	} else {
		array := make([]string, 0)
		switch value := i.(type) {
		case []int:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []int8:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []int16:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []int32:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []int64:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []uint:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []uint8:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []uint16:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []uint32:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []uint64:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []bool:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []float32:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []float64:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []interface{}:
			for _, v := range value {
				array = append(array, String(v))
			}
		case [][]byte:
			for _, v := range value {
				array = append(array, String(v))
			}
		default:
			return []string{String(i)}
		}
		return array
	}
}






