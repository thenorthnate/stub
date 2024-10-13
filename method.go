package stub

import "testing"

type Method struct {
	tb testing.TB
}

func NewMethod(tb testing.TB) *Method {
	return &Method{
		tb: tb,
	}
}

func (m *Method) Call(args ...any) *Call {
	return &Call{}
}
