package dcell_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell"
)

func TestCompile(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		expr    string
		wantErr error
	}{
		{
			name: "valid expression",
			expr: "1 + 2",
		},
		{
			name:    "complex",
			expr:    "%syntax error",
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := dcell.Compile(tc.expr)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Compile() error = %v, want %v", got, want)
			}
		})
	}
}

func TestMustCompile_Error(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustCompile() did not panic")
		}
	}()
	_ = dcell.MustCompile("%syntax error")
}

func TestWithFunc(t *testing.T) {
	t.Parallel()
	sut := dcell.MustCompile("test()", dcell.WithFunc("test", func() int {
		return 42
	}))

	result, err := sut.Eval(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("Eval() error = %v", err)
	}
	i, err := result.Int64()
	if err != nil {
		t.Fatalf("Eval() error = %v", err)
	}
	if got, want := i, int64(42); got != want {
		t.Errorf("Eval() = %v, want %v", got, want)
	}
}

func TestMustCompile_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("MustCompile() panicked: %v", r)
		}
	}()

	expr := dcell.MustCompile("1 + 2")
	if expr == nil {
		t.Errorf("MustCompile() returned nil")
	}
}

func TestExpr_Eval(t *testing.T) {
	t.Parallel()
	sut := dcell.MustCompile("1 + 2")

	result, err := sut.Eval(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("Eval() error = %v", err)
	}
	i, err := result.Int64()
	if err != nil {
		t.Fatalf("Eval() error = %v", err)
	}
	if got, want := i, int64(3); got != want {
		t.Errorf("Eval() = %v, want %v", got, want)
	}
}

func TestExpr_MustEval_Error(t *testing.T) {
	t.Parallel()
	sut := dcell.MustCompile("1 / 0")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustEval() did not panic")
		}
	}()
	_ = sut.MustEval(nil)
}

func TestExpr_MustEval_Success(t *testing.T) {
	t.Parallel()
	sut := dcell.MustCompile("1 + 2")

	result := sut.MustEval(nil)

	i, err := result.Int64()
	if err != nil {
		t.Fatalf("MustEval() error = %v", err)
	}
	if got, want := i, int64(3); got != want {
		t.Errorf("MustEval() = %v, want %v", got, want)
	}
}

func TestExpr_String(t *testing.T) {
	t.Parallel()
	input := "1 + 2"

	sut := dcell.MustCompile(input)

	result := sut.String()
	if got, want := result, input; got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

func TestExpr_UnmarshalText(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:  "valid expression",
			input: "1 + 2",
		},
		{
			name:    "invalid expression",
			input:   "%syntax error",
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var sut dcell.Expr

			err := sut.UnmarshalText([]byte(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("UnmarshalText() error = %v, want %v", got, want)
			}
		})
	}
}

func TestExpr_MarshalText(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input string
	}{
		{
			name:  "valid expression",
			input: "1 + 2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			sut := dcell.MustCompile(tc.input)

			result, err := sut.MarshalText()
			if err != nil {
				t.Fatalf("MarshalText() error = %v", err)
			}

			if got, want := string(result), tc.input; got != want {
				t.Errorf("MarshalText() = %v, want %v", got, want)
			}
		})
	}
}
