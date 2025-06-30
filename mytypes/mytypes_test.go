package mytypes_test

import (
	"mytypes"
	"testing"
)

func TestTwice(t *testing.T) {
	t.Parallel()
	input := mytypes.MyInt(9)
	want := mytypes.MyInt(18)
	got := input.Twice()
	if want != got {
		t.Errorf("%d.twice():want %d got %d", input, want, got)
	}
}

func TestStringLength(t *testing.T) {
	t.Parallel()
	input := mytypes.MyString("hello world")
	want := 11
	got := input.Len()
	if want != got {
		t.Errorf("%q.Len():want %d got %d", input, want, got)
	}
}
