package intconv_test

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/exp/constraints"
	"rodusek.dev/pkg/dcell/internal/intconv"
)

func convert[From, To constraints.Integer](v From, fn func(From) (To, error)) func() (To, error) {
	return func() (To, error) {
		return fn(v)
	}
}

func TestInt8(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		testcase func() (int8, error)
		want     int8
		wantErr  error
	}{
		{
			name:     "int8 input",
			testcase: convert(int8(42), intconv.Int8),
			want:     42,
		}, {
			name:     "int16 input within range",
			testcase: convert(int16(42), intconv.Int8),
			want:     42,
		}, {
			name:     "int16 input overflow",
			testcase: convert(int16(128), intconv.Int8),
			wantErr:  intconv.ErrOverflow,
		}, {
			name:     "int16 input underflow",
			testcase: convert(int16(math.MinInt8-1), intconv.Int8),
			wantErr:  intconv.ErrUnderflow,
		}, {
			name:     "uint8 input within range",
			testcase: convert(uint8(42), intconv.Int8),
			want:     42,
		}, {
			name:     "uint8 input overflow",
			testcase: convert(uint8(math.MaxInt8+1), intconv.Int8),
			wantErr:  intconv.ErrOverflow,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			conversion, err := tc.testcase()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Int8(...) = err %v, want %v", got, want)
			}
			if got, want := conversion, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Int8(...) = %v, want %v", got, want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		testcase func() (int16, error)
		want     int16
		wantErr  error
	}{
		{
			name:     "int16 input",
			testcase: convert(int16(42), intconv.Int16),
			want:     42,
		}, {
			name:     "int32 input within range",
			testcase: convert(int32(42), intconv.Int16),
			want:     42,
		}, {
			name:     "int32 input overflow",
			testcase: convert(int32(32768), intconv.Int16),
			wantErr:  intconv.ErrOverflow,
		}, {
			name:     "int32 input underflow",
			testcase: convert(int32(math.MinInt16-1), intconv.Int16),
			wantErr:  intconv.ErrUnderflow,
		}, {
			name:     "uint16 input within range",
			testcase: convert(uint16(42), intconv.Int16),
			want:     42,
		}, {
			name:     "uint16 input overflow",
			testcase: convert(uint16(math.MaxInt16+1), intconv.Int16),
			wantErr:  intconv.ErrOverflow,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			conversion, err := tc.testcase()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Int16(...) = err %v, want %v", got, want)
			}
			if got, want := conversion, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Int16(...) = %v, want %v", got, want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		testcase func() (int32, error)
		want     int32
		wantErr  error
	}{
		{
			name:     "int32 input",
			testcase: convert(int32(42), intconv.Int32),
			want:     42,
		}, {
			name:     "int64 input within range",
			testcase: convert(int64(42), intconv.Int32),
			want:     42,
		}, {
			name:     "int64 input overflow",
			testcase: convert(int64(2147483648), intconv.Int32),
			wantErr:  intconv.ErrOverflow,
		}, {
			name:     "int64 input underflow",
			testcase: convert(int64(math.MinInt32-1), intconv.Int32),
			wantErr:  intconv.ErrUnderflow,
		}, {
			name:     "uint32 input within range",
			testcase: convert(uint32(42), intconv.Int32),
			want:     42,
		}, {
			name:     "uint32 input overflow",
			testcase: convert(uint32(math.MaxInt32+1), intconv.Int32),
			wantErr:  intconv.ErrOverflow,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			conversion, err := tc.testcase()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Int32(...) = err %v, want %v", got, want)
			}
			if got, want := conversion, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Int32(...) = %v, want %v", got, want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		testcase func() (int64, error)
		want     int64
		wantErr  error
	}{
		{
			name:     "int64 input",
			testcase: convert(int64(42), intconv.Int64),
			want:     42,
		}, {
			name:     "int input within range",
			testcase: convert(int(42), intconv.Int64),
			want:     42,
		}, {
			name:     "uint64 input within range",
			testcase: convert(uint64(42), intconv.Int64),
			want:     42,
		}, {
			name:     "uint64 input overflow",
			testcase: convert(uint64(math.MaxInt64+1), intconv.Int64),
			wantErr:  intconv.ErrOverflow,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			conversion, err := tc.testcase()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Int64(...) = err %v, want %v", got, want)
			}
			if got, want := conversion, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Int64(...) = %v, want %v", got, want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		testcase func() (int, error)
		want     int
		wantErr  error
	}{
		{
			name:     "int input",
			testcase: convert(int(42), intconv.Int),
			want:     42,
		}, {
			name:     "int64 input within range",
			testcase: convert(int64(42), intconv.Int),
			want:     42,
		}, {
			name:     "uint64 input within range",
			testcase: convert(uint64(42), intconv.Int),
			want:     42,
		}, {
			name:     "uint64 input overflow",
			testcase: convert(uint64(math.MaxInt64+1), intconv.Int),
			wantErr:  intconv.ErrOverflow,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			conversion, err := tc.testcase()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Int(...) = err %v, want %v", got, want)
			}
			if got, want := conversion, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Int(...) = %v, want %v", got, want)
			}
		})
	}
}

func TestUint8(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		testcase func() (uint8, error)
		want     uint8
		wantErr  error
	}{
		{
			name:     "uint8 input",
			testcase: convert(uint8(42), intconv.Uint8),
			want:     42,
		}, {
			name:     "int16 input within range",
			testcase: convert(int16(42), intconv.Uint8),
			want:     42,
		}, {
			name:     "int16 input overflow",
			testcase: convert(int16(256), intconv.Uint8),
			wantErr:  intconv.ErrOverflow,
		}, {
			name:     "int16 input underflow",
			testcase: convert(int16(-1), intconv.Uint8),
			wantErr:  intconv.ErrUnderflow,
		}, {
			name:     "uint16 input within range",
			testcase: convert(uint16(42), intconv.Uint8),
			want:     42,
		}, {
			name:     "uint16 input overflow",
			testcase: convert(uint16(math.MaxUint8+1), intconv.Uint8),
			wantErr:  intconv.ErrOverflow,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			conversion, err := tc.testcase()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Uint8(...) = err %v, want %v", got, want)
			}
			if got, want := conversion, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Uint8(...) = %v, want %v", got, want)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		testcase func() (uint16, error)
		want     uint16
		wantErr  error
	}{
		{
			name:     "uint16 input",
			testcase: convert(uint16(42), intconv.Uint16),
			want:     42,
		}, {
			name:     "int32 input within range",
			testcase: convert(int32(42), intconv.Uint16),
			want:     42,
		}, {
			name:     "int32 input overflow",
			testcase: convert(int32(65536), intconv.Uint16),
			wantErr:  intconv.ErrOverflow,
		}, {
			name:     "int32 input underflow",
			testcase: convert(int32(-1), intconv.Uint16),
			wantErr:  intconv.ErrUnderflow,
		}, {
			name:     "uint32 input within range",
			testcase: convert(uint32(42), intconv.Uint16),
			want:     42,
		}, {
			name:     "uint32 input overflow",
			testcase: convert(uint32(math.MaxUint16+1), intconv.Uint16),
			wantErr:  intconv.ErrOverflow,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			conversion, err := tc.testcase()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Uint16(...) = err %v, want %v", got, want)
			}
			if got, want := conversion, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Uint16(...) = %v, want %v", got, want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		testcase func() (uint32, error)
		want     uint32
		wantErr  error
	}{
		{
			name:     "uint32 input",
			testcase: convert(uint32(42), intconv.Uint32),
			want:     42,
		}, {
			name:     "int64 input within range",
			testcase: convert(int64(42), intconv.Uint32),
			want:     42,
		}, {
			name:     "int64 input overflow",
			testcase: convert(int64(4294967296), intconv.Uint32),
			wantErr:  intconv.ErrOverflow,
		}, {
			name:     "int64 input underflow",
			testcase: convert(int64(-1), intconv.Uint32),
			wantErr:  intconv.ErrUnderflow,
		}, {
			name:     "uint64 input within range",
			testcase: convert(uint64(42), intconv.Uint32),
			want:     42,
		}, {
			name:     "uint64 input overflow",
			testcase: convert(uint64(math.MaxUint32+1), intconv.Uint32),
			wantErr:  intconv.ErrOverflow,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			conversion, err := tc.testcase()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Uint32(...) = err %v, want %v", got, want)
			}
			if got, want := conversion, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Uint32(...) = %v, want %v", got, want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		testcase func() (uint64, error)
		want     uint64
		wantErr  error
	}{
		{
			name:     "uint64 input",
			testcase: convert(uint64(42), intconv.Uint64),
			want:     42,
		}, {
			name:     "int input within range",
			testcase: convert(int(42), intconv.Uint64),
			want:     42,
		}, {
			name:     "int input underflow",
			testcase: convert(int(-1), intconv.Uint64),
			wantErr:  intconv.ErrUnderflow,
		}, {
			name:     "uint input within range",
			testcase: convert(uint(42), intconv.Uint64),
			want:     42,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			conversion, err := tc.testcase()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Uint64(...) = err %v, want %v", got, want)
			}
			if got, want := conversion, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Uint64(...) = %v, want %v", got, want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		testcase func() (uint, error)
		want     uint
		wantErr  error
	}{
		{
			name:     "uint input",
			testcase: convert(uint(42), intconv.Uint),
			want:     42,
		}, {
			name:     "int64 input within range",
			testcase: convert(int64(42), intconv.Uint),
			want:     42,
		}, {
			name:     "int64 input underflow",
			testcase: convert(int64(-1), intconv.Uint),
			wantErr:  intconv.ErrUnderflow,
		}, {
			name:     "uint64 input within range",
			testcase: convert(uint64(42), intconv.Uint),
			want:     42,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			conversion, err := tc.testcase()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Uint(...) = err %v, want %v", got, want)
			}
			if got, want := conversion, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Uint(...) = %v, want %v", got, want)
			}
		})
	}
}
