package dcell

import (
	"reflect"

	"rodusek.dev/pkg/dcell/internal/reflectcmp"
	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

// Result is the result of evaluating a dcell expression.
// This may be a
type Result struct {
	inner reflect.Value
}

// String returns the result as a string.
func (r *Result) String() (string, error) {
	return reflectconv.String(r.inner)
}

// Int64 attempts to convert the result to a 64-bit signed integer value.
// If the value is not convertible to a string, or would lose information in the
// conversion, an error is returned.
func (r *Result) Int64() (int64, error) {
	return reflectconv.Int64(r.inner)
}

// Int32 attempts to convert the result to a 32-bit signed integer value.
// If the value is not convertible to a string, or would lose information in the
// conversion, an error is returned.
func (r *Result) Int32() (int32, error) {
	return reflectconv.Int32(r.inner)
}

// Int16 attempts to convert the result to a 16-bit signed integer value.
// If the value is not convertible to a string, or would lose information in the
// conversion, an error is returned.
func (r *Result) Int16() (int16, error) {
	return reflectconv.Int16(r.inner)
}

// Int8 attempts to convert the result to an 8-bit signed integer value.
// If the value is not convertible to a string, or would lose information in the
// conversion, an error is returned.
func (r *Result) Int8() (int8, error) {
	return reflectconv.Int8(r.inner)
}

// Int attempts to convert the result to a signed integer value.
// If the value is not convertible to a string, or would lose information in the
// conversion, an error is returned.
func (r *Result) Int() (int, error) {
	return reflectconv.Int(r.inner)
}

// Uint64 attempts to convert the result to a 64-bit unsigned integer value.
// If the value is not convertible to a string, or would lose information in the
// conversion, an error is returned.
func (r *Result) Uint64() (uint64, error) {
	return reflectconv.Uint64(r.inner)
}

// Uint32 attempts to convert the result to a 32-bit unsigned integer value.
// If the value is not convertible to a string, or would lose information in the
// conversion, an error is returned.
func (r *Result) Uint32() (uint32, error) {
	return reflectconv.Uint32(r.inner)
}

// Uint16 attempts to convert the result to a 16-bit unsigned integer value.
// If the value is not convertible to a string, or would lose information in the
// conversion, an error is returned.
func (r *Result) Uint16() (uint16, error) {
	return reflectconv.Uint16(r.inner)
}

// Uint8 attempts to convert the result to an 8-bit unsigned integer value.
// If the value is not convertible to a string, or would lose information in the
// conversion, an error is returned.
func (r *Result) Uint8() (uint8, error) {
	return reflectconv.Uint8(r.inner)
}

// Uint attempts to convert the result to an unsigned integer value.
// If the value is not convertible to a string, or would lose information in the
// conversion, an error is returned.
func (r *Result) Uint() (uint, error) {
	return reflectconv.Uint(r.inner)
}

// Float64 attempts to convert the result to a 64-bit floating point value.
// If the value is not convertible to a string, or would lose information in the
// conversion, an error is returned.
func (r *Result) Float64() (float64, error) {
	return reflectconv.Float64(r.inner)
}

// Float32 attempts to convert the result to a 32-bit floating point value.
// If the value is not convertible to a string, or would lose information in the
// conversion, an error is returned.
func (r *Result) Float32() (float32, error) {
	return reflectconv.Float32(r.inner)
}

// Bool attempts to convert the result to a boolean value.
func (r *Result) Bool() (bool, error) {
	return reflectconv.Bool(r.inner)
}

// IsTruthy returns true if the result is a truthy value.
// Truthiness is determined by the following rules:
//
//   - Numeric values are truthy if they are not zero.
//   - Strings are truthy if they are not empty.
//   - Slices and maps are truthy if they are not nil and contain at least one
//     element.
//   - Pointers are truthy if they are not nil.
//   - Structs are always truthy
//   - All else is considered truthy only if [reflect.Value.IsZero] returns false.
func (r *Result) IsTruthy() bool {
	return reflectconv.IsTruthy(r.inner)
}

// IsNil returns true if there is no value in the result, or if the value that
// would be returned is a nil value -- such as a pointer, slice, or map.
func (r *Result) IsNil() bool {
	return reflectconv.IsNil(r.inner)
}

// IsZero returns true if the result is a zero value, or nil.
func (r *Result) IsZero() bool {
	return !r.inner.IsValid() || r.inner.IsZero()
}

// Interface returns the underlying value of the result. If the IsNil would
// return true, this will return nil.
func (r *Result) Interface() any {
	if r.inner.IsValid() {
		return reflectconv.Deref(r.inner).Interface()
	}
	return nil
}

// Value returns the underlying value of the result. If the IsNil would
// return true, this will return a zero value of the type of the result.
func (r *Result) Value() reflect.Value {
	return r.inner
}

// Equal returns true if the two results are semantically equal.
// This is a deep comparison of the underlying values, but not the underlying
// storage or types -- meaning that two results carrying different types with
// the same value are considered equal, such as `int(42)` and `int8(42)`.
//
// This check is done recursively for structures, maps, and slices.
func (r *Result) Equal(other *Result) bool {
	if lhsNil, rhsNil := r.IsNil(), other.IsNil(); lhsNil || rhsNil {
		return lhsNil == rhsNil
	}
	return reflectcmp.Equal(r.inner, other.inner)
}
