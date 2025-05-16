package expr_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/errs"
	"rodusek.dev/pkg/dcell/internal/expr"
	"rodusek.dev/pkg/dcell/internal/expr/exprtest"
	"rodusek.dev/pkg/dcell/internal/reflectcmp"
)

func TestAddExpr(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		left    expr.Expr
		right   expr.Expr
		want    reflect.Value
		wantErr error
	}{{
		name:  "int + int",
		left:  exprtest.Integer(2),
		right: exprtest.Integer(3),
		want:  reflect.ValueOf(int64(5)),
	}, {
		name:  "float + float",
		left:  exprtest.Float(2.5),
		right: exprtest.Float(3.5),
		want:  reflect.ValueOf(6.0),
	}, {
		name:  "int + float",
		left:  exprtest.Integer(2),
		right: exprtest.Float(3.5),
		want:  reflect.ValueOf(5.5),
	}, {
		name:  "float + int",
		left:  exprtest.Float(2.5),
		right: exprtest.Integer(3),
		want:  reflect.ValueOf(5.5),
	}, {
		name:  "string + string",
		left:  exprtest.String("foo"),
		right: exprtest.String("bar"),
		want:  reflect.ValueOf("foobar"),
	}, {
		name:    "int + string (incompatible)",
		left:    exprtest.Integer(1),
		right:   exprtest.String("bar"),
		wantErr: errs.ErrIncompatible,
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
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.Add(tc.left, tc.right)

			got, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) && (tc.wantErr != cmpopts.AnyError || got == nil) {
				t.Errorf("Eval() error = %v, want %v", got, want)
			}
			if tc.wantErr == nil {
				if got, want := got, tc.want; !reflectcmp.Equal(got, want) {
					t.Errorf("Eval() = %v, want %v", got, want)
				}
			}
		})
	}
}

func TestSubtractExpr(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		left    expr.Expr
		right   expr.Expr
		want    reflect.Value
		wantErr error
	}{{
		name:  "int - int",
		left:  exprtest.Integer(5),
		right: exprtest.Integer(3),
		want:  reflect.ValueOf(int64(2)),
	}, {
		name:  "float - float",
		left:  exprtest.Float(5.5),
		right: exprtest.Float(2.5),
		want:  reflect.ValueOf(3.0),
	}, {
		name:  "int - float",
		left:  exprtest.Integer(5),
		right: exprtest.Float(2.5),
		want:  reflect.ValueOf(2.5),
	}, {
		name:  "float - int",
		left:  exprtest.Float(5.5),
		right: exprtest.Integer(2),
		want:  reflect.ValueOf(3.5),
	}, {
		name:    "string - string (incompatible)",
		left:    exprtest.String("foo"),
		right:   exprtest.String("bar"),
		wantErr: errs.ErrIncompatible,
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
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.Subtract(tc.left, tc.right)

			got, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) && (tc.wantErr != cmpopts.AnyError || got == nil) {
				t.Errorf("Eval() error = %v, want %v", got, want)
			}
			if tc.wantErr == nil {
				if got, want := got, tc.want; !reflectcmp.Equal(got, want) {
					t.Errorf("Eval() = %v, want %v", got, want)
				}
			}
		})
	}
}
