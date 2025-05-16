package compile_test

import (
	"bufio"
	_ "embed"
	"os"
	"strings"
	"testing"

	"rodusek.dev/pkg/dcell/internal/compile"
	"rodusek.dev/pkg/dcell/internal/invocation"
)

func expressionsFromFile(t *testing.T, filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed to open testdata file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		expr := strings.TrimSpace(scanner.Text())
		if expr == "" || strings.HasPrefix(expr, "#") {
			continue
		}
		result = append(result, expr)
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("Failed to read testdata file: %v", err)
	}
	return result

}

func validExpressions(t *testing.T) []string {
	return expressionsFromFile(t, "testdata/valid-expressions.txt")
}

// This is a general acceptance test for the valid expressions to ensure that
// nothing is rejected.
func TestNewTree_ValidExpressions(t *testing.T) {
	t.Parallel()
	for _, expr := range validExpressions(t) {
		t.Run(expr, func(t *testing.T) {
			t.Parallel()
			table := invocation.NewTable()
			table.AddFunc("func", func(...any) (any, error) {
				return nil, nil
			})
			_, err := compile.NewTree(expr, &compile.Config{
				FuncTable: table,
			})

			if err != nil {
				t.Errorf("Failed to compile expression %q: %v", expr, err)
			}
		})
	}
}
