package stub

import (
	"testing"
)

func TestExecute(t *testing.T) {
	s, fs := newFooStub(t)
	s.Expect(fs.do, Any, "bar").Return(nil).NTimes(1)
	s.Expect(fs.get, 0).Return(12.1, nil)

	v, err := execute(fs, "bar")
	if err != nil {
		t.Fatal("execute returned an error: ", err)
	}
	if v != 12.1 {
		t.Fatal("expected 10, but got: ", v)
	}
}
