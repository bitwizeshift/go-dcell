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

func TestFreeFuncExpr_Eval(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		fn      func(args ...reflect.Value) (reflect.Value, error)
		args    []expr.Expr
		want    reflect.Value
		wantErr error
	}{
		{
			name: "no args, returns constant",
			fn: func(_ ...reflect.Value) (reflect.Value, error) {
				return reflect.ValueOf(42), nil
			},
			args: nil,
			want: reflect.ValueOf(42),
		}, {
			name: "one arg, returns arg",
			fn: func(args ...reflect.Value) (reflect.Value, error) {
				return args[0], nil
			},
			args: []expr.Expr{exprtest.Integer(7)},
			want: reflect.ValueOf(7),
		}, {
			name: "two args, returns sum",
			fn: func(args ...reflect.Value) (reflect.Value, error) {
				return reflect.ValueOf(args[0].Int() + args[1].Int()), nil
			},
			args: []expr.Expr{exprtest.Integer(3), exprtest.Integer(4)},
			want: reflect.ValueOf(int64(7)),
		}, {
			name: "fn returns error",
			fn: func(_ ...reflect.Value) (reflect.Value, error) {
				return reflect.Value{}, testErr
			},
			args:    nil,
			wantErr: testErr,
		}, {
			name: "arg returns error",
			fn: func(_ ...reflect.Value) (reflect.Value, error) {
				return reflect.ValueOf(0), nil
			},
			args:    []expr.Expr{exprtest.Error(testErr)},
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.FreeFunc(tc.fn, tc.args...)

			got, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("FreeFuncExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("FreeFuncExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}

func TestMemberFuncExpr_Eval(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		current any
		fn      func(args ...reflect.Value) (reflect.Value, error)
		args    []expr.Expr
		want    reflect.Value
		wantErr error
	}{
		{
			name:    "current is nil",
			current: nil,
			fn: func(_ ...reflect.Value) (reflect.Value, error) {
				return reflect.ValueOf("should not be called"), nil
			},
			args: nil,
			want: reflect.Value{},
		}, {
			name:    "current and one arg, returns sum",
			current: 5,
			fn: func(args ...reflect.Value) (reflect.Value, error) {
				return reflect.ValueOf(args[0].Int() + args[1].Int()), nil
			},
			args: []expr.Expr{exprtest.Integer(7)},
			want: reflect.ValueOf(int64(12)),
		}, {
			name:    "current and two args, returns product",
			current: 2,
			fn: func(args ...reflect.Value) (reflect.Value, error) {
				return reflect.ValueOf(args[0].Int() * args[1].Int() * args[2].Int()), nil
			},
			args: []expr.Expr{exprtest.Integer(3), exprtest.Integer(4)},
			want: reflect.ValueOf(int64(24)),
		}, {
			name:    "fn returns error",
			current: 1,
			fn: func(_ ...reflect.Value) (reflect.Value, error) {
				return reflect.Value{}, testErr
			},
			args:    nil,
			wantErr: testErr,
		}, {
			name:    "arg returns error",
			current: 1,
			fn: func(_ ...reflect.Value) (reflect.Value, error) {
				return reflect.ValueOf(0), nil
			},
			args:    []expr.Expr{exprtest.Error(testErr)},
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := &expr.Context{Current: reflect.ValueOf(tc.current)}
			sut := expr.MemberFunc(tc.fn, tc.args...)

			got, err := sut.Eval(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("MemberFuncExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("MemberFuncExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
