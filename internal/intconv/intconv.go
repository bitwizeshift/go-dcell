/*
Package intconv provides functions for narrowing conversions between
signed and unsigned integers.
*/
package intconv

import (
	"errors"
	"fmt"
	"math"
	"reflect"

	"golang.org/x/exp/constraints"
)

var (
	// ErrOverflow indicates that a value is too large to fit in the target type.
	ErrOverflow = errors.New("integer overflow")
	// ErrUnderflow indicates that a value is too small to fit in the target type.
	ErrUnderflow = errors.New("integer underflow")
)

// ConvertError represents an error that occurs during type conversion.
// It contains information about the source and destination types and the
// value that caused the error.
//
// If this is converted with [errors.As], the underlying error kind will be
// resolvable as either [ErrOverflow] or [ErrUnderflow] with [errors.Is].
type ConvertError struct {
	From, To reflect.Type
	Value    any
	err      error
}

func overflow(from, to reflect.Type, value any) error {
	return &ConvertError{
		From:  from,
		To:    to,
		Value: value,
		err:   ErrOverflow,
	}
}

func underflow(from, to reflect.Type, value any) error {
	return &ConvertError{
		From:  from,
		To:    to,
		Value: value,
		err:   ErrUnderflow,
	}
}

func (e *ConvertError) Error() string {
	return fmt.Sprintf("%v: %v cannot be represented in %s", e.err, e.Value, e.To.Name())
}

func (e *ConvertError) Unwrap() error {
	return e.err
}

// Int8 converts a value of type T to int8. If the receiver integer overflows
// an [ErrOverflow] error is returned, and if it underflows an [ErrUnderflow]
// error is returned.
func Int8[T constraints.Integer](v T) (int8, error) {
	return convertInt[T, int8](v, math.MinInt8, math.MaxInt8)
}

// Int16 converts a value of type T to int16. If the receiver integer overflows
// an [ErrOverflow] error is returned, and if it underflows an [ErrUnderflow]
// error is returned.
func Int16[T constraints.Integer](v T) (int16, error) {
	return convertInt[T, int16](v, math.MinInt16, math.MaxInt16)
}

// Int32 converts a value of type T to int32. If the receiver integer overflows
// an [ErrOverflow] error is returned, and if it underflows an [ErrUnderflow]
// error is returned.
func Int32[T constraints.Integer](v T) (int32, error) {
	return convertInt[T, int32](v, math.MinInt32, math.MaxInt32)
}

// Int64 converts a value of type T to int64. If the receiver integer overflows
// an [ErrOverflow] error is returned, and if it underflows an [ErrUnderflow]
// error is returned.
func Int64[T constraints.Integer](v T) (int64, error) {
	return convertInt[T, int64](v, math.MinInt64, math.MaxInt64)
}

// Int converts a value of type T to uint8. If the receiver integer overflows
// an [ErrOverflow] error is returned, and if it underflows an [ErrUnderflow]
// error is returned.
func Int[T constraints.Integer](v T) (int, error) {
	return convertInt[T, int](v, math.MinInt, math.MaxInt)
}

// Uint8 converts a value of type T to uint8. If the receiver integer overflows
// an [ErrOverflow] error is returned, and if it underflows an [ErrUnderflow]
// error is returned.
func Uint8[T constraints.Integer](v T) (uint8, error) {
	return convertInt[T, uint8](v, 0, math.MaxUint8)
}

// Uint16 converts a value of type T to uint16. If the receiver integer overflows
// an [ErrOverflow] error is returned, and if it underflows an [ErrUnderflow]
// error is returned.
func Uint16[T constraints.Integer](v T) (uint16, error) {
	return convertInt[T, uint16](v, 0, math.MaxUint16)
}

// Uint32 converts a value of type T to uint32. If the receiver integer overflows
// an [ErrOverflow] error is returned, and if it underflows an [ErrUnderflow]
// error is returned.
func Uint32[T constraints.Integer](v T) (uint32, error) {
	return convertInt[T, uint32](v, 0, math.MaxUint32)
}

// Uint64 converts a value of type T to uint64. If the receiver integer overflows
// an [ErrOverflow] error is returned, and if it underflows an [ErrUnderflow]
// error is returned.
func Uint64[T constraints.Integer](v T) (uint64, error) {
	return convertInt[T, uint64](v, 0, math.MaxUint64)
}

// Uint converts a value of type T to uint. If the receiver integer overflows
// an [ErrOverflow] error is returned, and if it underflows an [ErrUnderflow]
// error is returned.
func Uint[T constraints.Integer](v T) (uint, error) {
	return convertInt[T, uint](v, 0, math.MaxUint)
}

func convertInt[From, To constraints.Integer](v From, low, high To) (To, error) {
	from, to := reflect.TypeOf(v), reflect.TypeFor[To]()
	fromSigned := isSigned(from)
	toSigned := isSigned(to)
	switch {
	case fromSigned && toSigned:
		if int64(v) < int64(low) {
			return 0, underflow(from, to, v)
		}
		if int64(v) > int64(high) {
			return 0, overflow(from, to, v)
		}
	case fromSigned && !toSigned:
		if v < 0 {
			return 0, underflow(from, to, v)
		}
		if uint64(v) > uint64(high) {
			return 0, overflow(from, to, v)
		}
	case !fromSigned && toSigned:
		if uint64(v) > uint64(high) {
			return 0, overflow(from, to, v)
		}
	case !fromSigned && !toSigned:
		if uint64(v) > uint64(high) {
			return 0, overflow(from, to, v)
		}
	}
	return To(v), nil
}

func isSigned(rt reflect.Type) bool {
	switch rt.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}
