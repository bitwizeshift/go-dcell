package expr_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/expr"
	"rodusek.dev/pkg/dcell/internal/expr/exprtest"
	"rodusek.dev/pkg/dcell/internal/reflectcmp"
)

func TestLessThan(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name        string
		left, right expr.Expr
		want        reflect.Value
		wantErr     error
	}{
		{
			name:  "Is less than",
			left:  exprtest.Integer(1),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(true),
		}, {
			name:  "Is not less than",
			left:  exprtest.Integer(2),
			right: exprtest.Integer(1),
			want:  reflect.ValueOf(false),
		}, {
			name:    "left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Integer(1),
			wantErr: testErr,
		}, {
			name:    "right returns error",
			left:    exprtest.Integer(1),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.LessThan(tc.left, tc.right)

			result, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Eval() error = %v, want %v", got, want)
			}
			if got, want := result, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("Eval() = %v, want %v", got, want)
			}
		})
	}
}

func TestLessThanOrEqual(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name        string
		left, right expr.Expr
		want        reflect.Value
		wantErr     error
	}{
		{
			name:  "Is less than",
			left:  exprtest.Integer(1),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(true),
		}, {
			name:  "Is equal",
			left:  exprtest.Integer(2),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(true),
		}, {
			name:  "Is not less than or equal",
			left:  exprtest.Integer(3),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(false),
		}, {
			name:    "left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Integer(1),
			wantErr: testErr,
		}, {
			name:    "right returns error",
			left:    exprtest.Integer(1),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.LessThanOrEqual(tc.left, tc.right)

			result, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Eval() error = %v, want %v", got, want)
			}
			if got, want := result, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("Eval() = %v, want %v", got, want)
			}
		})
	}
}

func TestGreaterThan(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name        string
		left, right expr.Expr
		want        reflect.Value
		wantErr     error
	}{
		{
			name:  "Is greater than",
			left:  exprtest.Integer(3),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(true),
		}, {
			name:  "Is not greater than",
			left:  exprtest.Integer(1),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(false),
		}, {
			name:    "left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Integer(1),
			wantErr: testErr,
		}, {
			name:    "right returns error",
			left:    exprtest.Integer(1),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.GreaterThan(tc.left, tc.right)

			result, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Eval() error = %v, want %v", got, want)
			}
			if got, want := result, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("Eval() = %v, want %v", got, want)
			}
		})
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name        string
		left, right expr.Expr
		want        reflect.Value
		wantErr     error
	}{
		{
			name:  "Is greater than",
			left:  exprtest.Integer(3),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(true),
		}, {
			name:  "Is equal",
			left:  exprtest.Integer(2),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(true),
		}, {
			name:  "Is not greater than or equal",
			left:  exprtest.Integer(1),
			right: exprtest.Integer(2),
			want:  reflect.ValueOf(false),
		}, {
			name:    "left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Integer(1),
			wantErr: testErr,
		}, {
			name:    "right returns error",
			left:    exprtest.Integer(1),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.GreaterThanOrEqual(tc.left, tc.right)

			result, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Eval() error = %v, want %v", got, want)
			}
			if got, want := result, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("Eval() = %v, want %v", got, want)
			}
		})
	}
}
