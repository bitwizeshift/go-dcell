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

func TestIndexSliceEval(t *testing.T) {
	t.Parallel()
	testErr := errors.New("test error")
	testCases := []struct {
		name       string
		current    reflect.Value
		begin, end expr.Expr
		want       reflect.Value
		wantErr    error
	}{
		{
			name:    "basic slice [1:3]",
			current: reflect.ValueOf([]int{10, 20, 30, 40}),
			begin:   exprtest.Integer(1),
			end:     exprtest.Integer(3),
			want:    reflect.ValueOf([]int{20, 30}),
		}, {
			name:    "slice to end [2:]",
			current: reflect.ValueOf([]int{10, 20, 30, 40}),
			begin:   exprtest.Integer(2),
			end:     nil,
			want:    reflect.ValueOf([]int{30, 40}),
		}, {
			name:    "slice from start [:2]",
			current: reflect.ValueOf([]int{10, 20, 30, 40}),
			begin:   exprtest.Integer(0),
			end:     exprtest.Integer(2),
			want:    reflect.ValueOf([]int{10, 20}),
		}, {
			name:    "full slice [:]",
			current: reflect.ValueOf([]int{10, 20, 30}),
			begin:   exprtest.Integer(0),
			end:     nil,
			want:    reflect.ValueOf([]int{10, 20, 30}),
		}, {
			name:    "begin out of bounds",
			current: reflect.ValueOf([]int{1, 2}),
			begin:   exprtest.Integer(-1),
			end:     exprtest.Integer(1),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "end out of bounds",
			current: reflect.ValueOf([]int{1, 2}),
			begin:   exprtest.Integer(0),
			end:     exprtest.Integer(3),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "begin greater than end",
			current: reflect.ValueOf([]int{1, 2, 3}),
			begin:   exprtest.Integer(2),
			end:     exprtest.Integer(1),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "end negative index",
			current: reflect.ValueOf([]int{1, 2, 3, 4}),
			begin:   exprtest.Integer(1),
			end:     exprtest.Integer(-1),
			want:    reflect.ValueOf([]int{2, 3}),
		}, {
			name:    "end negative out of bounds",
			current: reflect.ValueOf([]int{1, 2, 3}),
			begin:   exprtest.Integer(0),
			end:     exprtest.Integer(-4),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "not a slice or array",
			current: reflect.ValueOf(42),
			begin:   exprtest.Integer(0),
			end:     exprtest.Integer(1),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "nil current",
			current: reflect.Value{},
			begin:   exprtest.Integer(0),
			end:     exprtest.Integer(1),
			want:    reflect.Value{},
		}, {
			name:    "begin returns error",
			current: reflect.ValueOf([]int{1, 2}),
			begin:   exprtest.Error(testErr),
			end:     exprtest.Integer(1),
			wantErr: testErr,
		}, {
			name:    "end returns error",
			current: reflect.ValueOf([]int{1, 2}),
			begin:   exprtest.Integer(0),
			end:     exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "begin is not integer",
			current: reflect.ValueOf([]int{1, 2}),
			begin:   exprtest.String("foo"),
			end:     exprtest.Integer(1),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "end is not integer",
			current: reflect.ValueOf([]int{1, 2}),
			begin:   exprtest.Integer(0),
			end:     exprtest.String("foo"),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.IndexSlice(tc.begin, tc.end)
			ctx := &expr.Context{Current: tc.current}

			got, err := sut.Eval(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("Eval() = %v, want %v", got, want)
			}
		})
	}
}
