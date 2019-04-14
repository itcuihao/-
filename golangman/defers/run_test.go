package defers

import "testing"

func TestA(t *testing.T) {
	i := a()
	t.Log(i)
	j := b()
	t.Log(j)
	k := c()
	t.Log(k)
	m := e()
	t.Log(m)
}
