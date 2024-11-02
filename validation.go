package stub

import "testing"

func validate(tb testing.TB, expect, got []any) {
	if len(expect) != len(got) {
		tb.Fatalf("expected %v input values, but got %v: expect=%v, got=%v", len(expect), len(got), expect, got)
	}
	for i := range expect {
		switch v := expect[i].(type) {
		case Input:
			switch v {
			case Any:
				// We don't care what the value is... it is allowed to be anything
				continue
			default:
				tb.Fatalf("%v is an unexpected Input type", v)
			}
		}
	}
}
