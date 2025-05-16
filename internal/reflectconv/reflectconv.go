package reflectconv

import (
	"fmt"
	"math"
	"reflect"

	"golang.org/x/exp/constraints"
	"rodusek.dev/pkg/dcell/internal/intconv"
)

// IsNil checks if the given reflect.Value is nil or empty.
func IsNil(rv reflect.Value) bool {
	if !rv.IsValid() {
		return true
	}
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Map:
		return rv.IsNil()
	}
	return false
}

// IsInt checks if the given reflect.Type is an integer type.
func IsInt(rt reflect.Type) bool {
	switch rt.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	}
	return false
}

// IsSigned checks if the given reflect.Type is a signed integer type.
func IsSigned(rt reflect.Type) bool {
	switch rt.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

// IsUnsigned checks if the given reflect.Type is an unsigned integer type.
func IsUnsigned(rt reflect.Type) bool {
	switch rt.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return true
	}
	return false
}

// IsFloat checks if the given reflect.Type is a float type.
func IsFloat(rt reflect.Type) bool {
	switch rt.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

// IsString checks if the given reflect.Type is a string type.
func IsString(rt reflect.Type) bool {
	return rt.Kind() == reflect.String
}

// IsBool checks if the given reflect.Type is a boolean type.
func IsBool(rt reflect.Type) bool {
	return rt.Kind() == reflect.Bool
}

// IsTruthy checks if the given reflect.Value is either a boolean value, or
// something that can be implicitly converted to a boolean value.
// It returns true for non-zero numbers, non-empty strings, non-empty slices,
// non-empty maps, non-nil pointers, and non-nil channels. An object/struct
// is always considered true.
func IsTruthy(rv reflect.Value) bool {
	if !rv.IsValid() {
		return false
	}
	switch rv.Kind() {
	case reflect.Bool:
		return rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() != 0
	case reflect.String:
		return rv.String() != ""
	case reflect.Float32, reflect.Float64:
		return rv.Float() != 0
	case reflect.Slice, reflect.Array:
		return rv.Len() > 0
	case reflect.Map:
		return rv.Len() > 0
	case reflect.Pointer:
		return !rv.IsNil()
	case reflect.Struct:
		return true
	case reflect.Chan:
		return !rv.IsNil()
	case reflect.Func:
		return !rv.IsNil()
	}
	return !rv.IsZero()
}

// Bool converts the given [reflect.Value] into a boolean value.
func Bool(rv reflect.Value) (bool, error) {
	if !rv.IsValid() {
		return false, fmt.Errorf("invalid value")
	}
	rv = Deref(rv)
	if rv.Kind() == reflect.Bool {
		return rv.Bool(), nil
	}
	return false, fmt.Errorf("cannot convert %s to bool", rv.Type().Name())
}

// Int64s converts the given [reflect.Value] into a slice of int64 values.
func Int64s(rvs ...reflect.Value) ([]int64, error) {
	return sequence(Int64, rvs...)
}

// Int64 converts the given [reflect.Value] into an int64 value.
func Int64(rv reflect.Value) (int64, error) {
	return convertInt(rv, intconv.Int64, intconv.Int64)
}

// Int32 converts the given [reflect.Value] into an int32 value.
func Int32(rv reflect.Value) (int32, error) {
	return convertInt(rv, intconv.Int32, intconv.Int32)
}

// Int16 converts the given [reflect.Value] into an int16 value.
func Int16(rv reflect.Value) (int16, error) {
	return convertInt(rv, intconv.Int16, intconv.Int16)
}

// Int8 converts the given [reflect.Value] into an int8 value.
func Int8(rv reflect.Value) (int8, error) {
	return convertInt(rv, intconv.Int8, intconv.Int8)
}

// Int converts the given [reflect.Value] into an int value.
func Int(rv reflect.Value) (int, error) {
	return convertInt(rv, intconv.Int, intconv.Int)
}

// Uint64s converts the given [reflect.Value] into a slice of int64 values.
func Uint64s(rvs ...reflect.Value) ([]uint64, error) {
	return sequence(Uint64, rvs...)
}

// Uint64 converts the given [reflect.Value] into a uint64 value.
func Uint64(rv reflect.Value) (uint64, error) {
	return convertInt(rv, intconv.Uint64, intconv.Uint64)
}

// Uint32 converts the given [reflect.Value] into a uint32 value.
func Uint32(rv reflect.Value) (uint32, error) {
	return convertInt(rv, intconv.Uint32, intconv.Uint32)
}

// Uint16 converts the given [reflect.Value] into a uint16 value.
func Uint16(rv reflect.Value) (uint16, error) {
	return convertInt(rv, intconv.Uint16, intconv.Uint16)
}

// Uint8 converts the given [reflect.Value] into a uint8 value.
func Uint8(rv reflect.Value) (uint8, error) {
	return convertInt(rv, intconv.Uint8, intconv.Uint8)
}

// Uint converts the given [reflect.Value] into a uint value.
func Uint(rv reflect.Value) (uint, error) {
	return convertInt(rv, intconv.Uint, intconv.Uint)
}

type signedConv[T any] func(int64) (T, error)
type unsignedConv[T any] func(uint64) (T, error)

func convertInt[T constraints.Integer](rv reflect.Value, sconv signedConv[T], uconv unsignedConv[T]) (T, error) {
	if !rv.IsValid() {
		return 0, fmt.Errorf("invalid value")
	}
	rv = Deref(rv)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return sconv(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uconv(rv.Uint())
	}
	return 0, fmt.Errorf("cannot convert %s to int", rv.Type().Name())
}

// Float64s converts the given [reflect.Value] into a slice of float64 values.
func Float64s(rvs ...reflect.Value) ([]float64, error) {
	return sequence(Float64, rvs...)
}

// Float64 converts the given [reflect.Value] into a float64 value.
func Float64(rv reflect.Value) (float64, error) {
	if !rv.IsValid() {
		return 0, fmt.Errorf("invalid value")
	}
	rv = Deref(rv)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(rv.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(rv.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return rv.Float(), nil
	}
	return 0, fmt.Errorf("cannot convert %s to float64", rv.Type().Name())
}

// Float32 converts the given [reflect.Value] into a float32 value.
func Float32(rv reflect.Value) (float32, error) {
	f, err := Float64(rv)
	if err != nil {
		return 0, err
	}
	result := float32(f)
	if result > math.MaxFloat32 || result < -math.MaxFloat32 {
		return 0, fmt.Errorf("cannot convert %s to float32", rv.Type().Name())
	}
	return result, nil
}

// String converts the given [reflect.Value] into a string value.
func String(rv reflect.Value) (string, error) {
	if !rv.IsValid() {
		return "", fmt.Errorf("invalid value")
	}
	rv = Deref(rv)
	if rv.Kind() == reflect.String {
		return rv.String(), nil
	}
	return "", fmt.Errorf("cannot convert %s to string", rv.Type().Name())
}

func sequence[T any](convert func(reflect.Value) (T, error), rvs ...reflect.Value) ([]T, error) {
	result := make([]T, len(rvs))
	for i, rv := range rvs {
		got, err := convert(rv)
		if err != nil {
			return nil, err
		}
		result[i] = got
	}
	return result, nil
}

// Deref dereferences the given reflect.Value until it is no longer a pointer or
// interface.
func Deref(rv reflect.Value) reflect.Value {
	for (rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface) && !rv.IsNil() {
		rv = rv.Elem()
	}
	return rv
}
