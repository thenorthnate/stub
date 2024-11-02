package stub

import "testing"

type Input int

const (
	Any Input = iota + 1
)

type Stub struct {
	tb           testing.TB
	methodCount  int
	orderMatters bool
	// expect contains the calls that you expect
	// you can only have:
	// 	- exact match mode - every call is defined
	// 	- N match mode - calls must be in order, but they can be called N number of times
	// 	- any match mode - calls can be made in any order, N times
	expect        []*Expectation
	currentExpIdx int
}

func New(tb testing.TB) *Stub {
	return &Stub{
		tb:           tb,
		orderMatters: true,
	}
}

func (s *Stub) NewMethod() *Method {
	method := newMethod(s.methodCount, s.tb, s)
	s.methodCount++
	return method
}

func (s *Stub) OrderMatters(ok bool) {
	s.orderMatters = ok
}

func (s *Stub) Expect(m *Method, args ...any) *Expectation {
	e := newExpectation(s.tb, m)
	s.expect = append(s.expect, e)
	return e
}

func (s *Stub) call(m *Method, args ...any) *Call {
	if s.orderMatters {
		e := s.expect[s.currentExpIdx]
		call, moveOn := e.call(m, args) // may fail the test
		if moveOn {
			s.currentExpIdx++
		}
		return call
	}
	for _, e := range s.expect {
		if e.matches(m) {
			call, moveOn := e.call(m, args) // may fail the test
			if moveOn {
				s.currentExpIdx++
			}
			return call
		}
	}
	s.tb.Fatal("failed to match any expectations")
	return nil
}
