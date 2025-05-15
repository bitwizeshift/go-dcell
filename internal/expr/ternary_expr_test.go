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

func TestTernaryExpr(t *testing.T) {
	t.Parallel()

	testErr := errors.New("test error")
	testCases := []struct {
		name      string
		condition expr.Expr
		trueExpr  expr.Expr
		falseExpr expr.Expr
		want      reflect.Value
		wantErr   error
	}{
		{
			name:      "condition evaluates to true",
			condition: exprtest.Boolean(true),
			trueExpr:  exprtest.String("true"),
			falseExpr: exprtest.String("false"),
			want:      reflect.ValueOf("true"),
		}, {
			name:      "condition evaluates to false",
			condition: exprtest.Boolean(false),
			trueExpr:  exprtest.String("true"),
			falseExpr: exprtest.String("false"),
			want:      reflect.ValueOf("false"),
		}, {
			name:      "condition evaluates to error",
			condition: exprtest.Error(testErr),
			trueExpr:  exprtest.String("true"),
			falseExpr: exprtest.String("false"),
			wantErr:   testErr,
		}, {
			name:      "true expression evaluates to error",
			condition: exprtest.Boolean(true),
			trueExpr:  exprtest.Error(testErr),
			falseExpr: exprtest.String("false"),
			wantErr:   testErr,
		}, {
			name:      "false expression evaluates to error",
			condition: exprtest.Boolean(false),
			trueExpr:  exprtest.String("true"),
			falseExpr: exprtest.Error(testErr),
			wantErr:   testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.Ternary(tc.condition, tc.trueExpr, tc.falseExpr)

			result, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Ternary() error = %v, want %v", got, want)
			}
			if got, want := result, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("Ternary() = %v, want %v", got.Interface(), want.Interface())
			}
		})
	}
}

func TestElvisExpr(t *testing.T) {
	testErr := errors.New("test error")
	t.Parallel()
	testCases := []struct {
		name      string
		condition expr.Expr
		falseExpr expr.Expr
		want      reflect.Value
		wantErr   error
	}{
		{
			name:      "condition evaluates to true",
			condition: exprtest.Boolean(true),
			falseExpr: exprtest.String("false"),
			want:      reflect.ValueOf(true),
		}, {
			name:      "condition evaluates to false",
			condition: exprtest.Boolean(false),
			falseExpr: exprtest.String("false"),
			want:      reflect.ValueOf("false"),
		}, {
			name:      "condition evaluates to error",
			condition: exprtest.Error(testErr),
			falseExpr: exprtest.String("false"),
			wantErr:   testErr,
		}, {
			name:      "false expression evaluates to error",
			condition: exprtest.Boolean(false),
			falseExpr: exprtest.Error(testErr),
			wantErr:   testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sut := expr.Elvis(tc.condition, tc.falseExpr)

			result, err := sut.Eval(nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Elvis() error = %v, want %v", got, want)
			}
			if got, want := result, tc.want; !reflectcmp.Equal(got, want) {
				t.Errorf("Elvis() = %v, want %v", got.Interface(), want.Interface())
			}
		})
	}
}
