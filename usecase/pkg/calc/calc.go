package calc

import (
	"context"
	"github.com/turanukimaru/ca/domain/pkg/dummylogic"
)

type Adder struct {
	A int
	B int
}

// Add implements add.
func (c *Adder) Add(ctx context.Context) (res int, err error) {
	return dummylogic.Add(c.A, c.B), err
}
