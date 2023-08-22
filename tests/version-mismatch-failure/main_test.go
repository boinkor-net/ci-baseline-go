package main

import "testing"

func TestSuccessButBuildFails(t *testing.T) {
	s := []string{"a", "b", "c"}
	a := [3]string(s)
	t.Log(a)
}
