package stub

import "testing"

type Call struct {
	tb         testing.TB
	retVals    []any
	returnsIDX int
}

func newCall(tb testing.TB, retVals []any) *Call {
	return &Call{
		tb:      tb,
		retVals: retVals,
	}
}

func GetReturn[T any](c *Call) T {
	if len(c.retVals) <= c.returnsIDX {
		c.tb.Fatal("asking for an unexpected number of return values")
	}
	vany := c.retVals[c.returnsIDX]
	v, ok := vany.(T)
	if !ok {
		c.tb.Fatalf("expected type %T type but got %T", v, vany)
	}
	c.returnsIDX++
	return v
}
