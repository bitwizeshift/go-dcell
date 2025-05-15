package reflectcmp

import (
	"reflect"
	"strings"

	"rodusek.dev/pkg/dcell/internal/intcmp"
	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// Equal checks if two [reflect.Values] are representationally equal.
func Equal(lhs, rhs reflect.Value) bool {
	return Compare(lhs, rhs) == 0
}

// Compare compares two [reflect.Values] and returns an integer indicating
// their relative order. It returns -1 if lhs < rhs, 0 if lhs == rhs, and
// 1 if lhs > rhs.
func Compare(lhs, rhs reflect.Value) int {
	lhs, rhs = reflectconv.Deref(lhs), reflectconv.Deref(rhs)
	ltype, rtype := classifyType(lhs), classifyType(rhs)
	if ltype < rtype {
		return -1
	}
	if ltype > rtype {
		return 1
	}
	switch ltype {
	case rtypeInt:
		return compareInt(lhs, rhs)
	case rtypeFloat:
		return compareFloat(lhs, rhs)
	case rtypeString:
		return strings.Compare(lhs.String(), rhs.String())
	case rtypeBool:
		return compareBool(lhs, rhs)
	case rtypeStruct:
		return compareStruct(lhs, rhs)
	case rtypeSlice:
		return compareSlice(lhs, rhs)
	case rtypeMap:
		return compareMap(lhs, rhs)
	case rtypeNil:
		return 0
	}
	return 1
}

func compareInt(lhs, rhs reflect.Value) int {
	if isSigned(lhs) {
		if isSigned(rhs) {
			return intcmp.Compare(lhs.Int(), rhs.Int())
		}
		return intcmp.Compare(lhs.Int(), rhs.Uint())
	}
	if isUnsigned(rhs) {
		return intcmp.Compare(lhs.Uint(), rhs.Uint())
	}
	return intcmp.Compare(lhs.Uint(), rhs.Int())
}

func compareFloat(lhs, rhs reflect.Value) int {
	lfloat, rfloat := lhs.Float(), rhs.Float()
	if lfloat < rfloat {
		return -1
	}
	if lfloat > rfloat {
		return 1
	}
	return 0
}

func compareBool(lhs, rhs reflect.Value) int {
	lbool, rbool := lhs.Bool(), rhs.Bool()
	if lbool && !rbool {
		return 1
	}
	if !lbool && rbool {
		return -1
	}
	return 0
}

func compareSlice(lhs, rhs reflect.Value) int {
	llen, rlen := lhs.Len(), rhs.Len()
	if llen != rlen {
		return intcmp.Compare(llen, rlen)
	}
	for i := range llen {
		if cmp := Compare(lhs.Index(i), rhs.Index(i)); cmp != 0 {
			return cmp
		}
	}
	return 0
}

func compareStruct(lhs, rhs reflect.Value) int {
	ltype, rtype := lhs.Type(), rhs.Type()
	if ltype != rtype {
		return strings.Compare(ltype.String(), rtype.String())
	}
	for i := range ltype.NumField() {
		lfield, rfield := lhs.Field(i), rhs.Field(i)
		if cmp := Compare(lfield, rfield); cmp != 0 {
			return cmp
		}
	}
	return 0
}

func compareMap(lhs, rhs reflect.Value) int {
	lkeys, rkeys := lhs.MapKeys(), rhs.MapKeys()
	llen, rlen := len(lkeys), len(rkeys)
	if llen != rlen {
		return intcmp.Compare(llen, rlen)
	}
	for i := range llen {
		lkey, rkey := lkeys[i], rkeys[i]
		if cmp := Compare(lkey, rkey); cmp != 0 {
			return cmp
		}
		if cmp := Compare(lhs.MapIndex(lkey), rhs.MapIndex(rkey)); cmp != 0 {
			return cmp
		}
	}
	return 0
}

type rtype int

const (
	rtypeInt rtype = iota
	rtypeFloat
	rtypeString
	rtypeBool
	rtypeStruct
	rtypeSlice
	rtypeMap
	rtypeUnknown
	rtypeNil
)

func classifyType(rv reflect.Value) rtype {
	if !rv.IsValid() {
		return rtypeNil
	}
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rtypeInt
	case reflect.Float32, reflect.Float64:
		return rtypeFloat
	case reflect.String:
		return rtypeString
	case reflect.Bool:
		return rtypeBool
	case reflect.Struct:
		return rtypeStruct
	case reflect.Slice, reflect.Array:
		return rtypeSlice
	case reflect.Map:
		return rtypeMap
	}
	return rtypeUnknown
}

func isSigned(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}
func isUnsigned(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	}
	return false
}
