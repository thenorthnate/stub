package stub

import "testing"

type Expectation struct {
	tb            testing.TB
	m             *Method
	inputVals     []any
	retVals       []any
	expectedCount int
	callCount     int
}

func newExpectation(tb testing.TB, m *Method, inputs []any) *Expectation {
	return &Expectation{
		tb:            tb,
		m:             m,
		inputVals:     inputs,
		expectedCount: 1,
	}
}

func (e *Expectation) Return(args ...any) *Expectation {
	e.retVals = args
	return e
}

func (e *Expectation) NTimes(n uint) *Expectation {
	e.expectedCount = int(n)
	return e
}

func (e *Expectation) AnyTimes() *Expectation {
	e.expectedCount = -1
	return e
}

func (e *Expectation) matches(m *Method) bool {
	return m == e.m
}

func (e *Expectation) call(m *Method, args []any) (*Call, bool) {
	e.callCount++
	if e.callCount > e.expectedCount && e.expectedCount >= 0 {
		e.tb.Fatal("called too many times")
	}
	if m != e.m {
		// the method being called does not match this expectation
		// TODO : get the current stack trace and get the actual name of the method that was called
		e.tb.Fatal("the method called does not match the expected method")
	}
	validate(e.tb, e.inputVals, args)
	c := newCall(e.tb, e.retVals)
	return c, e.callCount == e.expectedCount
}
