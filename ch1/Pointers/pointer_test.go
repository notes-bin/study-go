package pointers_test

import (
	"testing"
)

func zeroval(val int) {
	val = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func TestPointer(t *testing.T) {
	i := 1
	t.Logf("initial: %#v\n", i)

	zeroval(i)
	t.Logf("zeroval: %#v\n", i)

	zeroptr(&i)
	t.Logf("zeroptr: %#v\n", i)

	t.Logf("pointer: %#v\n", &i)
}
