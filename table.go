package dcell

import (
	"sync"

	"rodusek.dev/pkg/dcell/internal/invocation"
)

// tableV1 creates the function table for the V1 version of this library.
var tableV1 = sync.OnceValue(func() *invocation.Table {
	table := invocation.NewTable()

	return table
})
