package stub

import "testing"

type Call struct {
	tb        testing.TB
	returns   []any
	inquiries int
}

func (c *Call) Error() error {
	c.tb.Cleanup(func() {
		if c.inquiries != len(c.returns) {
			c.tb.Fatal("expected to have %v inquiries but only got %v", len(c.returns), c.inquiries)
		}
	})
	return nil
}

func (c *Call) Float64() float64 {
	return 0
}
