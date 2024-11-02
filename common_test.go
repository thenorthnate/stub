package stub

import (
	"context"
	"testing"
)

// foo is a random interface that we want to create a stub for!
type foo interface {
	Do(ctx context.Context, name string) error
	Get(i int) (float64, error)
}

// execute is a pretend function that is going to take the foo interface and do some things
// with it.
func execute(f foo, name string) (float64, error) {
	ctx := context.Background()
	if err := f.Do(ctx, name); err != nil {
		return 0, err
	}
	v, err := f.Get(0)
	if err != nil {
		return 0, err
	}
	return v, nil
}

type fooStub struct {
	do  *Method
	get *Method
}

func newFooStub(t *testing.T) (*Stub, *fooStub) {
	s := New(t)
	fs := &fooStub{
		do:  s.NewMethod(),
		get: s.NewMethod(),
	}
	return s, fs
}

func (fs *fooStub) Do(ctx context.Context, name string) error {
	c := fs.do.Call(ctx, name)
	return GetReturn[error](c)
}

func (fs *fooStub) Get(i int) (float64, error) {
	c := fs.get.Call(i)
	return GetReturn[float64](c), GetReturn[error](c)
}
