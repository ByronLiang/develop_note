package worker

import "testing"

func TestDeferSample(t *testing.T) {
	DeferSample()
}

func TestDeferHandle(t *testing.T) {
	i := 10
	res := DeferHandle(&i)
	t.Logf("res: %d, i: %d", res, i)
	j := 10
	res2 := DeferHandleWithName(&j)
	t.Logf("res2: %d, i: %d", res2, j)
}
