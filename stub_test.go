package stub

import (
	"context"
	"testing"
)

type foo interface {
	do(ctx context.Context) error
	get(i int) (float64, error)
}

func execute(f foo) (float64, error) {
	ctx := context.Background()
	if err := f.do(ctx); err != nil {
		return 0, err
	}
	v, err := f.get(0)
	if err != nil {
		return 0, err
	}
	return v, nil
}

type fooStub struct {
	methodDo  *Method
	methodGet *Method
}

func (fs *fooStub) do(ctx context.Context) error {
	c := fs.methodDo.Call(ctx)
	return c.Error()
}

func (fs *fooStub) get(i int) (float64, error) {
	c := fs.methodGet.Call(i)
	return c.Float64(), c.Error()
}

func TestExecute(t *testing.T) {
	fs := &fooStub{}

	v, err := execute(fs)
	if err != nil {
		t.Fatal("execute returned an error: ", err)
	}
	if v != 10 {
		t.Fatal("expected 10, but got: ", v)
	}
}
