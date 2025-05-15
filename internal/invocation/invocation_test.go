package invocation_test

import (
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/dcell/internal/invocation"
	"rodusek.dev/pkg/dcell/internal/invocation/arity"
	"rodusek.dev/pkg/dcell/internal/invocation/invocationtest"
)

func TestTable_Lookup(t *testing.T) {
	sut := invocation.NewTable()
	sut.Add("exists", invocationtest.AlwaysReturn(42))

	testCases := []struct {
		name string
		key  string
		want bool
	}{
		{
			name: "exists",
			key:  "exists",
			want: true,
		}, {
			name: "not exists",
			key:  "not-exists",
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, ok := sut.Lookup(tc.key)

			if got, want := ok, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Lookup(%q) = %v, want %v", tc.key, got, want)
			}
		})
	}
}

func TestTable_New(t *testing.T) {
	table := invocation.NewTable()
	table.Add("parent", invocationtest.AlwaysReturn(42))
	child := table.New()
	child.Add("child", invocationtest.AlwaysReturn(43))
	child.Add("parent", invocationtest.AlwaysReturn(44))

	entry, ok := child.Lookup("parent")
	if !ok {
		t.Fatalf("failed to lookup parent")
	}
	result, err := entry.Invoke()
	if err != nil {
		t.Fatalf("failed to invoke parent: %v", err)
	}

	if got, want := result.Interface(), 44; !cmp.Equal(got, want) {
		t.Errorf("Invoke() = %v, want %v", got, want)
	}
}

func TestTable_FunctionNames(t *testing.T) {
	table := invocation.NewTable()
	table.Add("parent", invocationtest.AlwaysReturn(42))
	child := table.New()
	child.Add("child", invocationtest.AlwaysReturn(43))
	child.Add("parent", invocationtest.AlwaysReturn(44))
	want := []string{"child", "parent", "parent"}

	got := slices.Collect(child.FunctionNames())

	if !cmp.Equal(got, want, cmpopts.SortSlices(strings.Compare)) {
		t.Errorf("FunctionNames() = %v, want %v", got, want)
	}
}

func TestTable_AddFunc(t *testing.T) {
	sut := invocation.NewTable()

	testCases := []struct {
		name    string
		input   any
		wantErr error
	}{
		{
			name:    "Not a function",
			input:   42,
			wantErr: invocation.ErrBadFunc,
		}, {
			name:    "Function returns too many arguments",
			input:   func() (int, int, int) { return 1, 2, 3 },
			wantErr: invocation.ErrBadFunc,
		}, {
			name:    "Function returns too few arguments",
			input:   func() { return },
			wantErr: invocation.ErrBadFunc,
		}, {
			name:    "Function second return is not an error",
			input:   func() (int, int) { return 1, 2 },
			wantErr: invocation.ErrBadFunc,
		}, {
			name:    "Valid function",
			input:   func(int) (int, error) { return 1, nil },
			wantErr: nil,
		}, {
			name:    "Valid variadic function",
			input:   func(int, ...int) (int, error) { return 1, nil },
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := sut.AddFunc("example", tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("AddFunc(...) = %v, want %v", err, tc.wantErr)
			}
		})
	}
}

func TestTable_AddFunc_Invoke(t *testing.T) {
	testCases := []struct {
		name    string
		fn      any
		params  []reflect.Value
		want    reflect.Value
		wantErr error
	}{
		{
			name:    "function accepting no params returning two values",
			fn:      func() (int, error) { return 42, nil },
			params:  []reflect.Value{},
			want:    reflect.ValueOf(42),
			wantErr: nil,
		}, {
			name:    "function accepting no params returning one value",
			fn:      func() int { return 42 },
			params:  []reflect.Value{},
			want:    reflect.ValueOf(42),
			wantErr: nil,
		}, {
			name:    "function accepting no params called with params",
			fn:      func() int { return 42 },
			params:  []reflect.Value{reflect.ValueOf(1)},
			want:    reflect.Value{},
			wantErr: arity.ErrBadArity,
		}, {
			name:    "function accepting one param returning two values",
			fn:      func(i int) (int, error) { return i, nil },
			params:  []reflect.Value{reflect.ValueOf(42)},
			want:    reflect.ValueOf(42),
			wantErr: nil,
		}, {
			name:    "function accepting one param returning one value",
			fn:      func(i int) int { return i },
			params:  []reflect.Value{reflect.ValueOf(42)},
			want:    reflect.ValueOf(42),
			wantErr: nil,
		}, {
			name:    "function accepting one param called with wrong type",
			fn:      func(i int) int { return i },
			params:  []reflect.Value{reflect.ValueOf("42")},
			want:    reflect.Value{},
			wantErr: invocation.ErrBadArgument,
		}, {
			name:    "function accepting variadic param returning two values",
			fn:      func(i int, _ ...int) (int, error) { return i, nil },
			params:  []reflect.Value{reflect.ValueOf(42)},
			want:    reflect.ValueOf(42),
			wantErr: nil,
		}, {
			name:    "function accepting variadic param returning one value",
			fn:      func(i int, _ ...int) int { return i },
			params:  []reflect.Value{reflect.ValueOf(42)},
			want:    reflect.ValueOf(42),
			wantErr: nil,
		}, {
			name:    "function accepting variadic param called with wrong type",
			fn:      func(i int, _ ...int) int { return i },
			params:  []reflect.Value{reflect.ValueOf("42")},
			want:    reflect.Value{},
			wantErr: invocation.ErrBadArgument,
		}, {
			name:    "function accepting variadic param called with multiple params",
			fn:      func(i int, _ ...int) int { return i },
			params:  []reflect.Value{reflect.ValueOf(42), reflect.ValueOf(1), reflect.ValueOf(2)},
			want:    reflect.ValueOf(42),
			wantErr: nil,
		}, {
			name:    "function accepting variadic param called with multiple wrong types",
			fn:      func(i int, _ ...int) int { return i },
			params:  []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2), reflect.ValueOf("42")},
			want:    reflect.Value{},
			wantErr: invocation.ErrBadArgument,
		}, {
			name:    "function returns an error",
			fn:      func() (int, error) { return 0, arity.ErrBadArity },
			params:  []reflect.Value{},
			want:    reflect.Value{},
			wantErr: arity.ErrBadArity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := invocation.NewTable()
			err := sut.AddFunc("example", tc.fn)
			if err != nil {
				t.Fatalf("failed to add function: %v", err)
			}
			entry, ok := sut.Lookup("example")
			if !ok {
				t.Fatalf("failed to lookup entry")
			}

			result, err := entry.Invoke(tc.params...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Table.AddFunc(...).Invoke(...) = %v, want %v", got, want)
			}
			if got, want := result, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Table.AddFunc(...).Invoke(...) = %v, want %v", got, want)
			}
		})
	}
}

func TestEntry_Invoke(t *testing.T) {
	sut := invocation.NewTable()
	sut.Add("example", invocationtest.AlwaysReturn(42)).SetArity(arity.None())

	testCases := []struct {
		name    string
		params  []reflect.Value
		want    reflect.Value
		wantErr error
	}{
		{
			name:    "no params",
			params:  []reflect.Value{},
			want:    reflect.ValueOf(42),
			wantErr: nil,
		}, {
			name:    "wrong number of params",
			params:  []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)},
			want:    reflect.Value{},
			wantErr: arity.ErrBadArity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			entry, ok := sut.Lookup("example")
			if !ok {
				t.Fatalf("failed to lookup entry")
			}

			result, err := entry.Invoke(tc.params...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Invoke(%v) = %v, want %v", tc.params, got, want)
			}
			if got, want := result, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Invoke(%v) = %v, want %v", tc.params, got, want)
			}
		})
	}

}
