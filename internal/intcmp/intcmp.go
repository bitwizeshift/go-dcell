/*
Package intcmp provides integer comparison functions for comparing signed and
unsigned integers.
*/
package intcmp

import (
	"reflect"

	"golang.org/x/exp/constraints"
)

// Equal compares two values of any integer type and returns true if they are
// equal.
func Equal[L, R constraints.Integer](l L, r R) bool {
	return Compare(l, r) == 0
}

// Compare compares two values of any type and returns
func Compare[L, R constraints.Integer](l L, r R) int {
	lv, rv := reflect.ValueOf(l), reflect.ValueOf(r)
	if isSigned(lv) {
		if isSigned(rv) {
			return cmpInt(lv.Int(), rv.Int())
		}
		if isUnsigned(rv) {
			lint := lv.Int()
			if lint < 0 {
				return -1
			}
			return cmpInt(lint, int64(rv.Uint()))
		}
	}
	if isUnsigned(rv) {
		return cmpInt(lv.Uint(), rv.Uint())
	}
	rint := rv.Int()
	if rint < 0 {
		return 1
	}
	return cmpInt(lv.Uint(), uint64(rint))
}

func cmpInt[T constraints.Integer](l, r T) int {
	switch {
	case l < r:
		return -1
	case l > r:
		return 1
	default:
		return 0
	}
}

func isSigned(r reflect.Value) bool {
	switch r.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

func isUnsigned(r reflect.Value) bool {
	return !isSigned(r)
}
