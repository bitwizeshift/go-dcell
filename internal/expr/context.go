package expr

import "reflect"

// Context is used to keep track of the current evaluation context
// during expression evaluation. It holds the root value and the current
// value being evaluated. This allows for nested evaluations and
// maintaining the state of the evaluation context.
type Context struct {
	// Root is the value that calls the expression.
	Root reflect.Value

	// Current is the current value being evaluated; typically a sub-value of
	// Root.
	Current reflect.Value
}

// NewContext creates a new Context with the given root value.
func NewContext(root reflect.Value) *Context {
	return &Context{
		Root:    root,
		Current: root,
	}
}

// Next creates a new Context based on the current context,
func (c *Context) Next(v reflect.Value) *Context {
	result := c.clone()
	result.Current = v
	return result
}

func (c *Context) clone() *Context {
	return &Context{
		Root:    c.Root,
		Current: c.Current,
	}
}
