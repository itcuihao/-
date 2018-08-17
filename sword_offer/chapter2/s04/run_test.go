package s04

import "testing"

func TestRun(t *testing.T) {
	s := "We are happy."
	t.Log(run(s))
}
