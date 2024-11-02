package stub

import "testing"

type Method struct {
	id int
	tb testing.TB
	s  *Stub
}

func newMethod(id int, tb testing.TB, s *Stub) *Method {
	return &Method{
		id: id,
		tb: tb,
		s:  s,
	}
}

func (m *Method) Call(args ...any) *Call {
	return m.s.call(m, args)
}
