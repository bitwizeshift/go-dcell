package reflectconv_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/intconv"
	"rodusek.dev/pkg/dcell/internal/reflectconv"
)

func ptr[T any](v T) *T {
	return &v
}

func TestIsNil(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		input reflect.Value
		want  bool
	}{
		{
			name:  "nil pointer",
			input: reflect.ValueOf((*int)(nil)),
			want:  true,
		}, {
			name:  "nil map",
			input: reflect.ValueOf((map[string]int)(nil)),
			want:  true,
		}, {
			name:  "non-nil pointer",
			input: reflect.ValueOf(ptr(42)),
			want:  false,
		}, {
			name:  "non-nil value",
			input: reflect.ValueOf(42),
			want:  false,
		}, {
			name:  "invalid value",
			input: reflect.Value{},
			want:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			isNil := reflectconv.IsNil(tc.input)

			if got, want := isNil, tc.want; !cmp.Equal(got, want) {
				t.Errorf("IsNil(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestIsInt(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		val  any
		want bool
	}{
		{
			name: "int",
			val:  int(1),
			want: true,
		}, {
			name: "int8",
			val:  int8(1),
			want: true,
		}, {
			name: "uint",
			val:  uint(1),
			want: true,
		}, {
			name: "float64",
			val:  float64(1),
			want: false,
		}, {
			name: "string",
			val:  "foo",
			want: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rt := reflect.TypeOf(tc.val)

			isInt := reflectconv.IsInt(rt)

			if got, want := isInt, tc.want; !cmp.Equal(got, want) {
				t.Errorf("IsInt(%#v) = %v, want %v", tc.val, got, want)
			}
		})
	}
}

func TestIsSignedUnsigned(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name   string
		val    any
		signed bool
	}{
		{
			name:   "int",
			val:    int(1),
			signed: true,
		}, {
			name:   "int8",
			val:    int8(1),
			signed: true,
		}, {
			name:   "uint",
			val:    uint(1),
			signed: false,
		}, {
			name:   "uint8",
			val:    uint8(1),
			signed: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rt := reflect.TypeOf(tc.val)

			t.Run("IsSigned", func(t *testing.T) {
				t.Parallel()
				isSigned := reflectconv.IsSigned(rt)

				if got, want := isSigned, tc.signed; !cmp.Equal(got, want) {
					t.Errorf("IsSigned(%#v) = %v, want %v", tc.val, got, want)
				}
			})
			t.Run("IsUnsigned", func(t *testing.T) {
				t.Parallel()
				isUnsigned := reflectconv.IsUnsigned(rt)

				if got, want := isUnsigned, !tc.signed; !cmp.Equal(got, want) {
					t.Errorf("IsUnsigned(%#v) = %v, want %v", tc.val, got, want)
				}
			})
		})
	}
}

func TestIsFloat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		val  any
		want bool
	}{
		{
			name: "float64",
			val:  float64(1),
			want: true,
		}, {
			name: "float32",
			val:  float32(1),
			want: true,
		}, {
			name: "not float",
			val:  "foo",
			want: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rt := reflect.TypeOf(tc.val)

			isFloat := reflectconv.IsFloat(rt)

			if got, want := isFloat, tc.want; !cmp.Equal(got, want) {
				t.Errorf("IsFloat(%#v) = %v, want %v", tc.val, got, tc.want)
			}
		})
	}
}

func TestIsString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		val  any
		want bool
	}{
		{
			name: "not string",
			val:  float32(1),
			want: false,
		}, {
			name: "string",
			val:  "foo",
			want: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rt := reflect.TypeOf(tc.val)

			isString := reflectconv.IsString(rt)

			if got, want := isString, tc.want; !cmp.Equal(got, want) {
				t.Errorf("IsString(%#v) = %v, want %v", tc.val, got, tc.want)
			}
		})
	}
}

func TestIsBool(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		val  any
		want bool
	}{
		{
			name: "bool",
			val:  true,
			want: true,
		}, {
			name: "not bool",
			val:  int(1),
			want: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rt := reflect.TypeOf(tc.val)

			isBool := reflectconv.IsBool(rt)

			if got, want := isBool, tc.want; !cmp.Equal(got, want) {
				t.Errorf("IsBool(%#v) = %v, want %v", tc.val, got, tc.want)
			}
		})
	}
}

func TestIsTruthy(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		input reflect.Value
		want  bool
	}{
		{
			name:  "true bool",
			input: reflect.ValueOf(true),
			want:  true,
		}, {
			name:  "false bool",
			input: reflect.ValueOf(false),
			want:  false,
		}, {
			name:  "nonzero uint",
			input: reflect.ValueOf(uint(1)),
			want:  true,
		}, {
			name:  "zero uint",
			input: reflect.ValueOf(uint(0)),
			want:  false,
		}, {
			name:  "nonzero int",
			input: reflect.ValueOf(1),
			want:  true,
		}, {
			name:  "zero int",
			input: reflect.ValueOf(0),
			want:  false,
		}, {
			name:  "nonzero float",
			input: reflect.ValueOf(1.1),
			want:  true,
		}, {
			name:  "zero float",
			input: reflect.ValueOf(0.0),
			want:  false,
		}, {
			name:  "non-empty string",
			input: reflect.ValueOf("foo"),
			want:  true,
		}, {
			name:  "empty string",
			input: reflect.ValueOf(""),
			want:  false,
		}, {
			name:  "non-empty slice",
			input: reflect.ValueOf([]int{1}),
			want:  true,
		}, {
			name:  "empty slice",
			input: reflect.ValueOf([]int{}),
			want:  false,
		}, {
			name:  "non-empty map",
			input: reflect.ValueOf(map[string]int{"a": 1}),
			want:  true,
		}, {
			name:  "empty map",
			input: reflect.ValueOf(map[string]int{}),
			want:  false,
		}, {
			name:  "nil pointer",
			input: reflect.ValueOf((*int)(nil)),
			want:  false,
		}, {
			name:  "nil map",
			input: reflect.ValueOf((map[string]int)(nil)),
			want:  false,
		}, {
			name:  "struct",
			input: reflect.ValueOf(struct{}{}),
			want:  true,
		}, {
			name:  "invalid value",
			input: reflect.Value{},
			want:  false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			isTruthy := reflectconv.IsTruthy(tc.input)

			if got, want := isTruthy, tc.want; !cmp.Equal(got, want) {
				t.Errorf("IsTruthy(%#v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestBool(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name    string
		input   reflect.Value
		want    bool
		wantErr error
	}{
		{
			name:  "boolean true value",
			input: reflect.ValueOf(true),
			want:  true,
		}, {
			name:    "string value",
			input:   reflect.ValueOf("hello"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "invalid value",
			input:   reflect.Value{},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			b, err := reflectconv.Bool(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Bool(%v) = %v, want %v", tc.input, got, want)
			}
			if got, want := b, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Bool(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   reflect.Value
		want    int64
		wantErr error
	}{
		{
			name:  "int64 value",
			input: reflect.ValueOf(int64(42)),
			want:  42,
		}, {
			name:    "string value",
			input:   reflect.ValueOf("hello"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "invalid value",
			input:   reflect.Value{},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			i, err := reflectconv.Int64(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Int64(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := i, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Int64(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   reflect.Value
		want    uint64
		wantErr error
	}{
		{
			name:  "uint64 value",
			input: reflect.ValueOf(uint64(42)),
			want:  42,
		}, {
			name:    "string value",
			input:   reflect.ValueOf("hello"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "invalid value",
			input:   reflect.Value{},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			u, err := reflectconv.Uint64(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Uint64(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := u, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Uint64(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   reflect.Value
		want    float64
		wantErr error
	}{
		{
			name:  "float64 value",
			input: reflect.ValueOf(float64(3.14)),
			want:  3.14,
		}, {
			name:  "signed int value",
			input: reflect.ValueOf(int64(42)),
			want:  42.0,
		}, {
			name:  "unsigned int value",
			input: reflect.ValueOf(uint64(42)),
			want:  42.0,
		}, {
			name:    "string value",
			input:   reflect.ValueOf("hello"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "invalid value",
			input:   reflect.Value{},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			u, err := reflectconv.Float64(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Float64(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := u, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Float64(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   reflect.Value
		want    float32
		wantErr error
	}{
		{
			name:  "float64 value",
			input: reflect.ValueOf(float64(3.14)),
			want:  3.14,
		}, {
			name:  "signed int value",
			input: reflect.ValueOf(int64(42)),
			want:  42.0,
		}, {
			name:  "unsigned int value",
			input: reflect.ValueOf(uint64(42)),
			want:  42.0,
		}, {
			name:    "string value",
			input:   reflect.ValueOf("hello"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "invalid value",
			input:   reflect.Value{},
			wantErr: cmpopts.AnyError,
		}, {
			name:    "float out of range",
			input:   reflect.ValueOf(float64(math.MaxFloat64)),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			u, err := reflectconv.Float32(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Float32(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := u, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Float32(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   reflect.Value
		want    string
		wantErr error
	}{
		{
			name:  "string value",
			input: reflect.ValueOf("foo"),
			want:  "foo",
		}, {
			name:    "int value",
			input:   reflect.ValueOf(42),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "invalid value",
			input:   reflect.Value{},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			s, err := reflectconv.String(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("String(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := s, tc.want; !cmp.Equal(got, want) {
				t.Errorf("String(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestInt64s(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   []reflect.Value
		want    []int64
		wantErr error
	}{
		{
			name:  "valid int64 slice",
			input: []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)},
			want:  []int64{1, 2},
		}, {
			name:    "contains non-int value",
			input:   []reflect.Value{reflect.ValueOf(1), reflect.ValueOf("foo")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			i, err := reflectconv.Int64s(tc.input...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Int64s(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := i, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Int64s(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestUint64s(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   []reflect.Value
		want    []uint64
		wantErr error
	}{
		{
			name:  "valid uint64 slice",
			input: []reflect.Value{reflect.ValueOf(uint64(1)), reflect.ValueOf(uint64(2))},
			want:  []uint64{1, 2},
		}, {
			name:    "contains non-uint value",
			input:   []reflect.Value{reflect.ValueOf(uint64(1)), reflect.ValueOf("foo")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			u, err := reflectconv.Uint64s(tc.input...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Uint64s(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := u, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Uint64s(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestFloat64s(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   []reflect.Value
		want    []float64
		wantErr error
	}{
		{
			name:  "valid float64 slice",
			input: []reflect.Value{reflect.ValueOf(1.1), reflect.ValueOf(2.2)},
			want:  []float64{1.1, 2.2},
		}, {
			name:    "contains non-float value",
			input:   []reflect.Value{reflect.ValueOf(1.1), reflect.ValueOf("foo")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f, err := reflectconv.Float64s(tc.input...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Float64s(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := f, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Float64s(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestDeref(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		val  any
		want any
	}{
		{
			name: "plain value",
			val:  42,
			want: 42,
		}, {
			name: "pointer",
			val:  ptr(42),
			want: 42,
		}, {
			name: "pointer to pointer",
			val:  ptr(ptr(42)),
			want: 42,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rv := reflect.ValueOf(tc.val)

			got := reflectconv.Deref(rv)

			if !got.IsValid() || got.Interface() != tc.want {
				t.Errorf("Deref(%#v) = %v, want %v", tc.val, got.Interface(), tc.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   reflect.Value
		want    int32
		wantErr error
	}{
		{
			name:  "int32 value",
			input: reflect.ValueOf(int32(42)),
			want:  42,
		}, {
			name:    "overflow",
			input:   reflect.ValueOf(int64(2147483648)),
			wantErr: intconv.ErrOverflow,
		}, {
			name:    "underflow",
			input:   reflect.ValueOf(int64(-2147483649)),
			wantErr: intconv.ErrUnderflow,
		}, {
			name:    "string value",
			input:   reflect.ValueOf("hello"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "invalid value",
			input:   reflect.Value{},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			i, err := reflectconv.Int32(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Int32(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := i, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Int32(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   reflect.Value
		want    int16
		wantErr error
	}{
		{
			name:  "int16 value",
			input: reflect.ValueOf(int16(42)),
			want:  42,
		}, {
			name:    "overflow",
			input:   reflect.ValueOf(int32(32768)),
			wantErr: intconv.ErrOverflow,
		}, {
			name:    "underflow",
			input:   reflect.ValueOf(int32(-32769)),
			wantErr: intconv.ErrUnderflow,
		}, {
			name:    "string value",
			input:   reflect.ValueOf("hello"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "invalid value",
			input:   reflect.Value{},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			i, err := reflectconv.Int16(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Int16(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := i, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Int16(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   reflect.Value
		want    int8
		wantErr error
	}{
		{
			name:  "int8 value",
			input: reflect.ValueOf(int8(42)),
			want:  42,
		}, {
			name:    "overflow",
			input:   reflect.ValueOf(int16(128)),
			wantErr: intconv.ErrOverflow,
		}, {
			name:    "underflow",
			input:   reflect.ValueOf(int16(-129)),
			wantErr: intconv.ErrUnderflow,
		}, {
			name:    "string value",
			input:   reflect.ValueOf("hello"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "invalid value",
			input:   reflect.Value{},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			i, err := reflectconv.Int8(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Int8(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := i, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Int8(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   reflect.Value
		want    int
		wantErr error
	}{
		{
			name:  "int value",
			input: reflect.ValueOf(int(42)),
			want:  42,
		}, {
			name:  "overflow",
			input: reflect.ValueOf(int64(9223372036854775807)), // math.MaxInt64 for 64-bit
			want:  9223372036854775807,
		}, {
			name:  "underflow",
			input: reflect.ValueOf(int64(-9223372036854775808)), // math.MinInt64 for 64-bit
			want:  -9223372036854775808,
		}, {
			name:    "string value",
			input:   reflect.ValueOf("hello"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "invalid value",
			input:   reflect.Value{},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			i, err := reflectconv.Int(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Int(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := i, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Int(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   reflect.Value
		want    uint32
		wantErr error
	}{
		{
			name:  "uint32 value",
			input: reflect.ValueOf(uint32(42)),
			want:  42,
		}, {
			name:    "overflow",
			input:   reflect.ValueOf(uint64(math.MaxUint64)),
			wantErr: intconv.ErrOverflow,
		}, {
			name:    "underflow",
			input:   reflect.ValueOf(int64(-1)),
			wantErr: intconv.ErrUnderflow,
		}, {
			name:    "string value",
			input:   reflect.ValueOf("hello"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "invalid value",
			input:   reflect.Value{},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			u, err := reflectconv.Uint32(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Uint32(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := u, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Uint32(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   reflect.Value
		want    uint16
		wantErr error
	}{
		{
			name:  "uint16 value",
			input: reflect.ValueOf(uint16(42)),
			want:  42,
		}, {
			name:    "overflow",
			input:   reflect.ValueOf(uint32(65536)),
			wantErr: intconv.ErrOverflow,
		}, {
			name:    "underflow",
			input:   reflect.ValueOf(int32(-1)),
			wantErr: intconv.ErrUnderflow,
		}, {
			name:    "string value",
			input:   reflect.ValueOf("hello"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "invalid value",
			input:   reflect.Value{},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			u, err := reflectconv.Uint16(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Uint16(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := u, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Uint16(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestUint8(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   reflect.Value
		want    uint8
		wantErr error
	}{
		{
			name:  "uint8 value",
			input: reflect.ValueOf(uint8(42)),
			want:  42,
		}, {
			name:    "overflow",
			input:   reflect.ValueOf(uint16(256)),
			wantErr: intconv.ErrOverflow,
		}, {
			name:    "underflow",
			input:   reflect.ValueOf(int16(-1)),
			wantErr: intconv.ErrUnderflow,
		}, {
			name:    "string value",
			input:   reflect.ValueOf("hello"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "invalid value",
			input:   reflect.Value{},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			u, err := reflectconv.Uint8(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Uint8(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := u, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Uint8(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   reflect.Value
		want    uint
		wantErr error
	}{
		{
			name:  "uint value",
			input: reflect.ValueOf(uint(42)),
			want:  42,
		}, {
			name:  "overflow",
			input: reflect.ValueOf(uint64(^uint(0))), // math.MaxUint
			want:  ^uint(0),
		}, {
			name:    "underflow",
			input:   reflect.ValueOf(int64(-1)),
			wantErr: intconv.ErrUnderflow,
		}, {
			name:    "string value",
			input:   reflect.ValueOf("hello"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "invalid value",
			input:   reflect.Value{},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			u, err := reflectconv.Uint(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Uint(%v) error = %v, want %v", tc.input, got, want)
			}
			if got, want := u, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Uint(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}
